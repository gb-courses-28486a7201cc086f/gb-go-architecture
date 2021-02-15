package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"gb-go-architecture/lesson-3/httptester/workerpool"
)

const (
	requestsDefault = 1000
	timeoutDefault  = 0
)

var (
	workers     int    = 100
	requests    int    = 0
	timeout     int    = 0
	url         string = ""
	method      string = "GET"
	payload     string = ""
	httpTimeOut int    = 5000

	errLimitFlags = errors.New("One of: \"-c\" or \"-t\" flags should be provided")
	errEmptyURL   = errors.New("URL flag \"-url\" should be provided")
)

func setUp() (*Config, error) {
	flag.Parse()

	missRequests := requests == 0
	missTimeout := timeout == 0
	// one of limit should be set
	if missRequests && missTimeout {
		return nil, errLimitFlags
	}
	// setup defauls if one of flags missed
	if missRequests {
		requests = requestsDefault
	}
	if missTimeout {
		timeout = timeoutDefault
	}
	timeOutStr := fmt.Sprintf("%dms", timeout)
	timeoutDuration, err := time.ParseDuration(timeOutStr)
	if err != nil {
		return nil, err
	}

	if url == "" {
		return nil, errEmptyURL
	}

	return &Config{
		Workers:     workers,
		Requests:    requests,
		Timeout:     timeoutDuration,
		URL:         url,
		Method:      method,
		HTTPTimeOut: time.Second * 5,
	}, nil
}

func init() {
	flag.IntVar(&workers, "w", 100, "Parallel workers which perform requests")
	flag.IntVar(&requests, "c", 0, "Total count of requests to send")
	flag.IntVar(&timeout, "t", 0, "Time limit (milliseconds) to perform test. 0 means no time limit")
	flag.StringVar(&url, "url", "", "URL to test")
	flag.StringVar(&method, "method", "GET", "HTTP method for test requests")
	flag.StringVar(&payload, "data", "", "Payload for test requests")
	flag.IntVar(&httpTimeOut, "reqtimeout", 5000, "Timeout for HTTP client (milliseconds)")
}

func main() {
	config, err := setUp()
	if err != nil {
		fmt.Println(err)
		os.Exit(255)
	}
	fmt.Printf("Starting test for %s %s\n\n", config.Method, config.URL)

	pool, _ := workerpool.NewPool(config.Workers)

	cancelChan := make(chan struct{})
	// should store 1 message w-o blocking
	// because timeout can be before and select-case will not proceed
	exceedChan := make(chan struct{}, 1)

	wg := &sync.WaitGroup{}
	wg.Add(2) // JobProducer + JobReporter
	go JobProducer(wg, config, pool.JobsChan, cancelChan)
	go JobReporter(wg, config.Requests, pool.ResultsChan, exceedChan)

	// timer to stop on duration
	var timerChan <-chan time.Time
	if config.Timeout.Milliseconds() > 0 {
		timerChan = time.NewTimer(config.Timeout).C
	}

	// stop producer on either timeout or count exceed
	select {
	case <-timerChan:
		cancelChan <- struct{}{}
	case <-exceedChan:
		cancelChan <- struct{}{}
	}

	pool.Join()
	wg.Wait()
}
