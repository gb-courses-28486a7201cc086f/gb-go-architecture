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
	URL     string
	Method  string
	Payload *bytes.Buffer
}

func (tj *TestURLJob) Run() *workerpool.JobResult {
	req, _ := http.NewRequest(tj.Method, tj.URL, tj.Payload)
	start := time.Now()
	resp, err := tj.Client.Do(req)
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

type JobReport struct {
	reqTotal            int
	reqSuccess          int
	reqServerErrors     map[int]int
	reqFailed           int
	reqSuccessTotalTime time.Duration
	jobFirstStart       *time.Time
	jobLastEnd          *time.Time
}

func (jr *JobReport) Update(data *workerpool.JobResult) {
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
	if jr.reqSuccess > 0 {
		avgSuccessTime := jr.reqSuccessTotalTime.Seconds() / float64(jr.reqSuccess)
		log.Printf("Avg response time, sec: %.3f", avgSuccessTime)
	}
}

func (jr *JobReport) PrintRPS() {
	sentReq := float64(jr.reqTotal - jr.reqFailed)
	testExecTime := jr.jobLastEnd.Sub(*jr.jobFirstStart).Seconds()
	rpsSuccess := float64(jr.reqSuccess) / testExecTime

	var avgSuccessTime float64
	if jr.reqSuccess > 0 {
		avgSuccessTime = jr.reqSuccessTotalTime.Seconds() / float64(jr.reqSuccess)
	}

	fmt.Printf("\nResults:\nSent %.0f requests, %d success", sentReq, jr.reqSuccess)
	fmt.Printf("\nServer errors by code: %v\n", jr.reqServerErrors)
	fmt.Printf("\nRPS via success: %.0f\n", rpsSuccess)
	fmt.Printf("\nAvg response time: %.3f sec\n", avgSuccessTime)
}

func JobReporter(wg *sync.WaitGroup, resultsChan <-chan *workerpool.JobResult) {
	defer wg.Done()

	report := &JobReport{
		reqServerErrors: make(map[int]int),
	}

	// setup ticker to print sub results
	ticker := time.NewTicker(time.Second)

	running := true
	for running {
		select {
		case v, ok := <-resultsChan:
			// resultsChan closed => stop report
			if !ok {
				running = false
				break
			}
			report.Update(v)
		case <-ticker.C:
			report.LogAvgTime()
		}
	}

	// all results done -> ticker can stop
	ticker.Stop()

	// print final result
	report.PrintRPS()
}
