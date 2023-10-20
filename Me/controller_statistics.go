package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

type ControllerStatistics struct {
	ObjectName string `json:"object-name"`
	Meta       string `json:"meta"`
	DurableID  string `json:"durable-id"`

	CPULoad               int    `json:"cpu-load"`
	PowerOnTime           int    `json:"power-on-time"`
	WriteCacheUsed        int    `json:"write-cache-used"`
	BytesPerSecond        string `json:"bytes-per-second"`
	BytesPerSecondNumeric int    `json:"bytes-per-second-numeric"`
	Iops                  int    `json:"iops"`
	NumberOfReads         int64  `json:"number-of-reads"`
	ReadCacheHits         int64  `json:"read-cache-hits"`
	ReadCacheMisses       int64  `json:"read-cache-misses"`
	NumberOfWrites        int64  `json:"number-of-writes"`
	WriteCacheHits        int64  `json:"write-cache-hits"`
	WriteCacheMisses      int64  `json:"write-cache-misses"`
	DataRead              string `json:"data-read"`
	DataReadNumeric       int64  `json:"data-read-numeric"`
	DataWritten           string `json:"data-written"`
	DataWrittenNumeric    int64  `json:"data-written-numeric"`
	NumForwardedCmds      int    `json:"num-forwarded-cmds"`
	ResetTime             string `json:"reset-time"`
	ResetTimeNumeric      int    `json:"reset-time-numeric"`
	TimeSinceStatsReset   int64

	StartSampleTime          string `json:"start-sample-time"`
	StartSampleTimeNumeric   int    `json:"start-sample-time-numeric"`
	StopSampleTime           string `json:"stop-sample-time"`
	StopSampleTimeNumeric    int    `json:"stop-sample-time-numeric"`
	TotalPowerOnHoursString  string `json:"total-power-on-hours"`
	TotalPowerOnHoursNumeric float64
}

type httpControllerStatistics struct {
	ControllerStatistics []ControllerStatistics `json:"controller-statistics"`
	HttpStatus           []Status               `json:"status"`
}

func (scs *httpControllerStatistics) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, scs)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}
func (scs *httpControllerStatistics) FromJson(body []byte) error {
	err := json.Unmarshal(body, scs)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	for _, controller := range scs.ControllerStatistics {
		controller.TotalPowerOnHoursNumeric, _ =
			strconv.ParseFloat(controller.TotalPowerOnHoursString, 64)
		controller.TimeSinceStatsReset = time.Now().Unix() - int64(controller.ResetTimeNumeric)
	}
	return nil
}

func NewMe4ControllerStatistics(url string) []ControllerStatistics {
	scs := &httpControllerStatistics{}
	err := scs.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return scs.ControllerStatistics
}
func NewMe4ControllerStatisticsFrom(body []byte) (sti []ControllerStatistics, err error) {
	diskGp := &httpControllerStatistics{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return
	}

	sti = diskGp.ControllerStatistics
	return
}
func NewMe4ControllerStatisticsFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]ControllerStatistics, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []ControllerStatistics{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []ControllerStatistics{}, err
	}

	return NewMe4ControllerStatisticsFrom(body)
}
