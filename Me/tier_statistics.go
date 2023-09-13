package Me

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TierStatistics struct {
	ObjectName             string                 `json:"object-name"`
	Meta                   string                 `json:"meta"`
	SerialNumber           string                 `json:"serial-number"`
	Pool                   string                 `json:"pool"`
	Tier                   string                 `json:"tier"`
	TierNumeric            int                    `json:"tier-numeric"`
	PagesAllocPerMinute    int                    `json:"pages-alloc-per-minute"`
	PagesDeallocPerMinute  int                    `json:"pages-dealloc-per-minute"`
	PagesReclaimed         int                    `json:"pages-reclaimed"`
	NumPagesUnmapPerMinute int                    `json:"num-pages-unmap-per-minute"`
	ResettableStatistics   []ResettableStatistics `json:"resettable-statistics"`
}

type httpTierStatistics struct {
	TierStatistics []TierStatistics `json:"tier-statistics"`
	HttpStatus     []Status         `json:"status"`
}

func (ts *httpTierStatistics) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, ts)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4TierStatistics(url string) []TierStatistics {
	ts := &httpTierStatistics{}
	err := ts.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return ts.TierStatistics
}

func (dk *httpTierStatistics) FromJson(body []byte) error {
	err := json.Unmarshal(body, dk)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4TierStatisticsFrom(body []byte) (sti []TierStatistics, err error) {
	diskGp := &httpTierStatistics{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = diskGp.TierStatistics
	return
}

func NewMe4TierStatisticsFromRequest(client *http.Client, req *http.Request) ([]TierStatistics, error) {
	resp, err := client.Do(req)
	if err != nil {
		return []TierStatistics{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []TierStatistics{}, err
	}

	return NewMe4TierStatisticsFrom(body)
}
