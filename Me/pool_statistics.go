package Me

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"

	//"github.com/go-kit/log"
	//"github.com/go-kit/log/level"
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
	ResettableStatistics             []ResetTableStatistics `json:"resettable-statistics"`
	TierStatistics                   []TierStatistics       `json:"tier-statistics"`
}

type httpPoolStatistics struct {
	PoolStatistics []PoolStatistics `json:"pool-statistics"`
	HttpStatus     []Status         `json:"status"`
}

func (ps *httpPoolStatistics) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, ps)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return err
	}

	return nil
}

func NewMe4PoolStatisticsFrom(body []byte) (ps []PoolStatistics, err error) {
	diskGp := &httpPoolStatistics{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)

		return
	}

	ps = diskGp.PoolStatistics
	return
}

func NewMe4PoolStatistics(url string) []PoolStatistics {
	ps := &httpPoolStatistics{}
	err := ps.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v\n", err)
		err = errors.Errorf("Request error: %s", err)
		return nil
	}
	return ps.PoolStatistics
}
