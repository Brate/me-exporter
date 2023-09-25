package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"io"
	"net/http"
)

type DiskGroupStatistics struct {
	ObjectName               string                     `json:"object-name"`
	Meta                     string                     `json:"meta"`
	SerialNumber             string                     `json:"serial-number"`
	Name                     string                     `json:"name"`
	TimeSinceReset           int                        `json:"time-since-reset"`
	TimeSinceSample          int                        `json:"time-since-sample"`
	NumberOfReads            int                        `json:"number-of-reads"`
	NumberOfWrites           int                        `json:"number-of-writes"`
	DataRead                 string                     `json:"data-read"`
	DataReadNumeric          int64                      `json:"data-read-numeric"`
	DataWritten              string                     `json:"data-written"`
	DataWrittenNumeric       int64                      `json:"data-written-numeric"`
	BytesPerSecond           string                     `json:"bytes-per-second"`
	BytesPerSecondNumeric    int                        `json:"bytes-per-second-numeric"`
	Iops                     int                        `json:"iops"`
	AvgRspTime               int                        `json:"avg-rsp-time"`
	AvgReadRspTime           int                        `json:"avg-read-rsp-time"`
	AvgWriteRspTime          int                        `json:"avg-write-rsp-time"`
	DiskGroupStatisticsPaged []DiskGroupStatisticsPaged `json:"disk-group-statistics-paged"`
}

type httpDiskGroupStatistics struct {
	DiskGroup  []DiskGroupStatistics `json:"disk-group-statistics"`
	HttpStatus []Status              `json:"status"`
}

func (dk *httpDiskGroupStatistics) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, dk)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4DiskGroupStatistics(url string) []DiskGroupStatistics {
	dk := &httpDiskGroupStatistics{}
	err := dk.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return dk.DiskGroup
}

func (dk *httpDiskGroupStatistics) FromJson(body []byte) error {
	err := json.Unmarshal(body, dk)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4DiskGroupStatisticsFrom(body []byte) (sti []DiskGroupStatistics, err error) {
	diskGp := &httpDiskGroupStatistics{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = diskGp.DiskGroup
	return
}

func NewMe4DiskGroupStatisticsFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]DiskGroupStatistics, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []DiskGroupStatistics{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []DiskGroupStatistics{}, err
	}

	return NewMe4DiskGroupStatisticsFrom(body)
}
