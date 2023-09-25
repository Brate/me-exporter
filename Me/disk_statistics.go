package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"io"
	"net/http"
)

type DiskStatistic struct {
	ObjectName                 string `json:"object-name"`
	Meta                       string `json:"meta"`
	DurableID                  string `json:"durable-id"`
	Location                   string `json:"location"`
	SerialNumber               string `json:"serial-number"`
	PowerOnHours               int    `json:"power-on-hours"`
	BytesPerSecond             string `json:"bytes-per-second"`
	BytesPerSecondNumeric      int    `json:"bytes-per-second-numeric"`
	Iops                       int    `json:"iops"`
	NumberOfReads              int64  `json:"number-of-reads"`
	NumberOfWrites             int    `json:"number-of-writes"`
	DataRead                   string `json:"data-read"`
	DataReadNumeric            int64  `json:"data-read-numeric"`
	DataWritten                string `json:"data-written"`
	DataWrittenNumeric         int64  `json:"data-written-numeric"`
	LifetimeDataRead           string `json:"lifetime-data-read"`
	LifetimeDataReadNumeric    int    `json:"lifetime-data-read-numeric"`
	LifetimeDataWritten        string `json:"lifetime-data-written"`
	LifetimeDataWrittenNumeric int    `json:"lifetime-data-written-numeric"`
	QueueDepth                 int    `json:"queue-depth"`
	ResetTime                  string `json:"reset-time"`
	ResetTimeNumeric           int    `json:"reset-time-numeric"`
	StartSampleTime            string `json:"start-sample-time"`
	StartSampleTimeNumeric     int    `json:"start-sample-time-numeric"`
	StopSampleTime             string `json:"stop-sample-time"`
	StopSampleTimeNumeric      int    `json:"stop-sample-time-numeric"`
	SmartCount1                int    `json:"smart-count-1"`
	SmartCount2                int    `json:"smart-count-2"`
	IoTimeoutCount1            int    `json:"io-timeout-count-1"`
	IoTimeoutCount2            int    `json:"io-timeout-count-2"`
	NoResponseCount1           int    `json:"no-response-count-1"`
	NoResponseCount2           int    `json:"no-response-count-2"`
	SpinupRetryCount1          int    `json:"spinup-retry-count-1"`
	SpinupRetryCount2          int    `json:"spinup-retry-count-2"`
	NumberOfMediaErrors1       int    `json:"number-of-media-errors-1"`
	NumberOfMediaErrors2       int    `json:"number-of-media-errors-2"`
	NumberOfNonmediaErrors1    int    `json:"number-of-nonmedia-errors-1"`
	NumberOfNonmediaErrors2    int    `json:"number-of-nonmedia-errors-2"`
	NumberOfBlockReassigns1    int    `json:"number-of-block-reassigns-1"`
	NumberOfBlockReassigns2    int    `json:"number-of-block-reassigns-2"`
	NumberOfBadBlocks1         int    `json:"number-of-bad-blocks-1"`
	NumberOfBadBlocks2         int    `json:"number-of-bad-blocks-2"`
}

type httpDiskStatistics struct {
	DiskStatistics []DiskStatistic `json:"disk-statistics"`
	HttpStatus     []Status        `json:"status"`
}

func (sti *httpDiskStatistics) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, sti)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4DiskStatistics(url string) []DiskStatistic {
	sti := &httpDiskStatistics{}
	err := sti.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return sti.DiskStatistics
}

func (sti *httpDiskStatistics) FromJson(body []byte) error {
	err := json.Unmarshal(body, sti)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4DiskStatisticFrom(body []byte) (sti []DiskStatistic, err error) {
	hst := &httpDiskStatistics{}
	err = json.Unmarshal(body, hst)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = hst.DiskStatistics
	return
}

func NewMe4DiskStatisticFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]DiskStatistic, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []DiskStatistic{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []DiskStatistic{}, err
	}

	return NewMe4DiskStatisticFrom(body)
}
