package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"io"
	"net/http"
)

type VolumeStatistics struct {
	ObjectName              string `json:"object-name"`
	Meta                    string `json:"meta"`
	VolumeName              string `json:"volume-name"`
	SerialNumber            string `json:"serial-number"`
	BytesPerSecond          string `json:"bytes-per-second"`
	BytesPerSecondNumeric   int    `json:"bytes-per-second-numeric"`
	Iops                    int    `json:"iops"`
	NumberOfReads           int    `json:"number-of-reads"`
	NumberOfWrites          int    `json:"number-of-writes"`
	DataRead                string `json:"data-read"`
	DataReadNumeric         int64  `json:"data-read-numeric"`
	DataWritten             string `json:"data-written"`
	DataWrittenNumeric      int64  `json:"data-written-numeric"`
	AllocatedPages          int    `json:"allocated-pages"`
	PercentTierSsd          int    `json:"percent-tier-ssd"`
	PercentTierSas          int    `json:"percent-tier-sas"`
	PercentTierSata         int    `json:"percent-tier-sata"`
	PercentAllocatedRfc     int    `json:"percent-allocated-rfc"`
	PagesAllocPerMinute     int    `json:"pages-alloc-per-minute"`
	PagesDeallocPerMinute   int    `json:"pages-dealloc-per-minute"`
	SharedPages             int    `json:"shared-pages"`
	WriteCacheHits          int64  `json:"write-cache-hits"`
	WriteCacheMisses        int    `json:"write-cache-misses"`
	ReadCacheHits           int64  `json:"read-cache-hits"`
	ReadCacheMisses         int64  `json:"read-cache-misses"`
	SmallDestages           int    `json:"small-destages"`
	FullStripeWriteDestages int    `json:"full-stripe-write-destages"`
	ReadAheadOperations     int    `json:"read-ahead-operations"`
	WriteCacheSpace         int    `json:"write-cache-space"`
	WriteCachePercent       int    `json:"write-cache-percent"`
	ResetTime               string `json:"reset-time"`
	ResetTimeNumeric        int    `json:"reset-time-numeric"`
	StartSampleTime         string `json:"start-sample-time"`
	StartSampleTimeNumeric  int    `json:"start-sample-time-numeric"`
	StopSampleTime          string `json:"stop-sample-time"`
	StopSampleTimeNumeric   int    `json:"stop-sample-time-numeric"`
}

type htpShowVolumeStatics struct {
	VolumeStatistics []VolumeStatistics `json:"volume-statistics"`
	HttpStatus       []Status           `json:"status"`
}

func (svs *htpShowVolumeStatics) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, svs)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4ShowVolumeStatics(url string) []VolumeStatistics {
	svs := &htpShowVolumeStatics{}
	err := svs.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return svs.VolumeStatistics
}

func (svs *htpShowVolumeStatics) FromJson(body []byte) error {
	err := json.Unmarshal(body, svs)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4VolumeStatisticsFrom(body []byte) (sti []VolumeStatistics, err error) {
	diskGp := &htpShowVolumeStatics{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = diskGp.VolumeStatistics
	return
}

func NewMe4VolumeStatisticsFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]VolumeStatistics, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []VolumeStatistics{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []VolumeStatistics{}, err
	}

	return NewMe4VolumeStatisticsFrom(body)
}
