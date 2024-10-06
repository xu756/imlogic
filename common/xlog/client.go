package xlog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"imlogic/common/config"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type logData map[string]interface{}

const waitTime = time.Second * 3

var xlog *XLog

func init() {
	if config.RunData().OpenTelemetry.UserName == "" {

	}
	xlog = setup()
}

type XLog struct {
	*sync.Mutex
	ctx            context.Context
	httpc          *http.Client
	addr           string
	c              chan logData
	waitList       []logData
	requestTimeout time.Duration
	stream         string
	waitTimer      *time.Timer
	fullSize       int
}

func setup() *XLog {
	l := &XLog{
		Mutex:          &sync.Mutex{},
		ctx:            context.Background(),
		addr:           config.RunData().OpenTelemetry.Addr,
		c:              make(chan logData, 10),
		requestTimeout: time.Second * 10,
		waitTimer:      time.NewTimer(waitTime),
		stream:         config.RunData().OpenTelemetry.Straem,
		fullSize:       config.RunData().OpenTelemetry.FullSize,
	}
	l.httpc = &http.Client{}
	l.httpc.Timeout = l.requestTimeout
	l.run()
	l.waitList = make([]logData, 0)
	return l
}
func (l *XLog) run() {
	go func() {
		for {
			select {
			case val := <-l.c:
				l.Lock()
				l.waitTimer.Reset(waitTime)
				l.waitList = append(l.waitList, val)
				if len(l.waitList) >= l.fullSize {
					data := l.waitList
					l.waitList = make([]logData, 0)
					l.request(data)
				}
				l.Unlock()
			case <-l.waitTimer.C:
				if len(l.waitList) > 0 {
					l.Lock()
					data := l.waitList
					l.waitList = make([]logData, 0)
					l.request(data)
					l.Unlock()
				}
			case <-l.ctx.Done():
				log.Println("XLog.run ctx.Done")
				l.Lock()
				data := l.waitList
				l.waitList = make([]logData, 0)
				l.request(data)
				l.Unlock()
				return
			}

		}
	}()
}

func (l *XLog) request(obj []logData) {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Println("XLog.request json.Marshal:", err)
		return
	}
	log.Printf("发送日志:%d条", len(obj))
	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/api/default/%s/_json", l.addr, l.stream),
		bytes.NewReader(data))
	if err != nil {
		log.Println("XLog.request NewRequest:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(config.RunData().OpenTelemetry.UserName, config.RunData().OpenTelemetry.PassWord)
	response, err := l.httpc.Do(req)
	if err != nil {
		log.Println("XLog.request DoRequest:", err)
		return
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		return
	}
	body, _ := io.ReadAll(response.Body)
	log.Println("XLog.request request error:", string(body))
	return
}
func (l *XLog) send(data logData) {
	if data == nil {
		return
	}
	l.request([]logData{data})
}

// SendSync 同步发送
func SendSync(data logData) {
	go xlog.request([]logData{data})
}

// Send 异步发送
func Send(data logData) {
	xlog.c <- data
}
