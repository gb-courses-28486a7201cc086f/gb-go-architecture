package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"gb-go-architecture/lesson-3/httptester/workerpool"
)

type Config struct {
	Workers     int
	Requests    int
	Timeout     time.Duration
	URL         string
	Method      string
	Payload     []byte
	HTTPTimeOut time.Duration
}

type TestURLJob struct {
	ID      int
	Client  *http.Client
	Request *http.Request
}

func (tj *TestURLJob) Run() *workerpool.JobResult {
	start := time.Now()
	resp, err := tj.Client.Do(tj.Request)
	end := time.Now()
	if err != nil {
		return &workerpool.JobResult{
			Code:      -1,
			Message:   err.Error(),
			Payload:   []byte{},
			StartTime: start,
			EndTime:   end,
		}
	}
	defer resp.Body.Close()

	return &workerpool.JobResult{
		Code:      resp.StatusCode,
		Message:   resp.Status,
		Payload:   []byte{},
		StartTime: start,
		EndTime:   end,
	}
}

type TestJobPool struct {
	jobs []*TestURLJob
	size int
	next int
	mx   *sync.Mutex
}

func (jp *TestJobPool) Next() (job *TestURLJob) {
	jp.mx.Lock()
	if jp.next == jp.size {
		job = jp.jobs[0]
		jp.next = 1
	} else {
		job = jp.jobs[jp.next]
		jp.next = jp.next + 1
	}
	jp.mx.Unlock()
	return job
}

func NewTestJobPool(conf *Config) (*TestJobPool, error) {
	jobs := make([]*TestURLJob, conf.Workers)
	for i := 0; i < conf.Workers; i++ {
		client := &http.Client{
			Timeout: conf.HTTPTimeOut,
		}
		req, err := http.NewRequest(conf.Method, conf.URL, bytes.NewBuffer(conf.Payload))
		if err != nil {
			return nil, err
		}
		jobs[i] = &TestURLJob{
			ID:      i + 1,
			Client:  client,
			Request: req,
		}
	}
	return &TestJobPool{
		mx:   &sync.Mutex{},
		jobs: jobs,
		size: conf.Workers,
	}, nil
}

type JobReport struct {
	mx                  *sync.RWMutex
	maxRequests         int
	doneChan            chan<- struct{}
	isDone              bool
	reqTotal            int
	reqSuccess          int
	reqServerErrors     map[int]int
	reqFailed           int
	reqSuccessTotalTime time.Duration
	jobFirstStart       *time.Time
	jobLastEnd          *time.Time
}

func (jr *JobReport) Update(data *workerpool.JobResult) {
	jr.mx.Lock()
	defer jr.mx.Unlock()

	// skip if report exceeded expected results count
	if jr.reqSuccess >= jr.maxRequests {
		// raise done signal only once
		if !jr.isDone {
			jr.isDone = true
			jr.doneChan <- struct{}{}
		}
		return
	}

	jr.reqTotal++
	// collect failed requests (client/network issues)
	if data.Code < 0 {
		jr.reqFailed++
		return
	}

	// collect success jobs count and exec time
	if data.Code >= 200 && data.Code < 500 {
		// non server errors
		jr.reqSuccess++
		jr.reqSuccessTotalTime = jr.reqSuccessTotalTime + data.EndTime.Sub(data.StartTime)
	} else {
		jr.reqServerErrors[data.Code] = jr.reqServerErrors[data.Code] + 1
	}

	// collect report start and end time
	if jr.jobFirstStart == nil || data.StartTime.Before(*jr.jobFirstStart) {
		jr.jobFirstStart = &data.StartTime
	}
	if jr.jobLastEnd == nil || data.EndTime.After(*jr.jobLastEnd) {
		jr.jobLastEnd = &data.EndTime
	}
}

func (jr *JobReport) LogAvgTime() {
	var avgSuccessTime float64
	var print bool
	jr.mx.RLock()
	if jr.reqSuccess > 0 {
		avgSuccessTime = jr.reqSuccessTotalTime.Seconds() / float64(jr.reqSuccess)
		print = true
	}
	jr.mx.RUnlock()

	if print {
		log.Printf("Avg response time, sec: %.3f", avgSuccessTime)
	}
}

func (jr *JobReport) PrintRPS() {
	var rpsSuccess, rpsSent, avgSuccessTime float64
	jr.mx.RLock()
	sentReq := float64(jr.reqTotal - jr.reqFailed)
	testExecTime := jr.jobLastEnd.Sub(*jr.jobFirstStart).Seconds()
	rpsSent = sentReq / testExecTime
	rpsSuccess = float64(jr.reqSuccess) / testExecTime
	if jr.reqSuccess > 0 {
		avgSuccessTime = jr.reqSuccessTotalTime.Seconds() / float64(jr.reqSuccess)
	}
	jr.mx.RUnlock()

	fmt.Printf("\nResults:\nSent %.0f requests, %d success", sentReq, jr.reqSuccess)
	fmt.Printf("\nServer errors by code: %v\n", jr.reqServerErrors)
	fmt.Printf("\nRPS via success: %.0f\nRPS via sent: %.0f\n", rpsSuccess, rpsSent)
	fmt.Printf("\nAvg response time: %.3f sec\n", avgSuccessTime)
}

func JobReporter(wg *sync.WaitGroup, maxResults int, resultsChan <-chan *workerpool.JobResult, doneChan chan<- struct{}) {
	defer wg.Done()

	report := &JobReport{
		mx:              &sync.RWMutex{},
		maxRequests:     maxResults,
		doneChan:        doneChan,
		reqServerErrors: make(map[int]int),
	}

	// setup ticker to print sub results
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			report.LogAvgTime()
		}
	}()

	// collect results
	for result := range resultsChan {
		report.Update(result)
	}

	// all results done -> ticker can stop
	ticker.Stop()

	// print final result
	report.PrintRPS()
}

func JobProducer(wg *sync.WaitGroup, conf *Config, jobChan chan<- workerpool.Job, cancelChan <-chan struct{}) {
	defer wg.Done()
	defer close(jobChan)

	jobPool, _ := NewTestJobPool(conf)

	for {
		job := jobPool.Next()
		select {
		case <-cancelChan:
			return
		default:
			jobChan <- job
		}
	}
}
