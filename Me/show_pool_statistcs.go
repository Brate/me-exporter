package Me

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PoolStatistics struct {
	ObjectName                       string                 `json:"object-name"`
	Meta                             string                 `json:"meta"`
	SampleTime                       string                 `json:"sample-time"`
	SampleTimeNumeric                int                    `json:"sample-time-numeric"`
	SerialNumber                     string                 `json:"serial-number"`
	Pool                             string                 `json:"pool"`
	PagesAllocPerMinute              int                    `json:"pages-alloc-per-minute"`
	PagesAllocPerHour                int                    `json:"pages-alloc-per-hour"`
	PagesDeallocPerMinute            int                    `json:"pages-dealloc-per-minute"`
	PagesDeallocPerHour              int                    `json:"pages-dealloc-per-hour"`
	PagesUnmapPerMinute              int                    `json:"pages-unmap-per-minute"`
	PagesUnmapPerHour                int                    `json:"pages-unmap-per-hour"`
	NumBlockedSsdPromotionsPerMinute int                    `json:"num-blocked-ssd-promotions-per-minute"`
	NumBlockedSsdPromotionsPerHour   int                    `json:"num-blocked-ssd-promotions-per-hour"`
	NumPageAllocations               int                    `json:"num-page-allocations"`
	NumPageDeallocations             int                    `json:"num-page-deallocations"`
	NumPageUnmaps                    int                    `json:"num-page-unmaps"`
	NumPagePromotionsToSsdBlocked    int                    `json:"num-page-promotions-to-ssd-blocked"`
	NumHotPageMoves                  int                    `json:"num-hot-page-moves"`
	NumColdPageMoves                 int                    `json:"num-cold-page-moves"`
	ResettableStatistics             []ResettableStatistics `json:"resettable-statistics"`
	TierStatistics                   []TierStatistics       `json:"tier-statistics"`
}

type httpShowPoolStatistics struct {
	PoolStatistics []PoolStatistics `json:"pool-statistics"`
	Status         []Status         `json:"status"`
}

func (sps *httpShowPoolStatistics) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, sps)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4ShowPoolStatistics(url string) []PoolStatistics {
	sps := &httpShowPoolStatistics{}
	err := sps.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return sps.PoolStatistics
}

func (sps *httpShowPoolStatistics) FromJson(body []byte) error {
	err := json.Unmarshal(body, sps)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4ShowPoolStatisticsFrom(body []byte) (sti []PoolStatistics, err error) {
	diskGp := &httpShowPoolStatistics{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = diskGp.PoolStatistics
	return
}

func NewMe4ShowPoolStatisticsFromRequest(client *http.Client, req *http.Request) ([]PoolStatistics, error) {
	resp, err := client.Do(req)
	if err != nil {
		return []PoolStatistics{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []PoolStatistics{}, err
	}

	return NewMe4ShowPoolStatisticsFrom(body)
}
