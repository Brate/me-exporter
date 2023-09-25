package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"io"
	"net/http"
	"regexp"
)

type Pools struct {
	ObjectName                 string  `json:"object-name"`
	Meta                       string  `json:"meta"`
	Name                       string  `json:"name"`
	SerialNumber               string  `json:"serial-number"`
	URL                        string  `json:"url"`
	StorageType                string  `json:"storage-type"`
	StorageTypeNumeric         int     `json:"storage-type-numeric"`
	Blocksize                  int     `json:"blocksize"`
	TotalSize                  string  `json:"total-size"`
	TotalSizeNumeric           int64   `json:"total-size-numeric"`
	TotalAvail                 string  `json:"total-avail"`
	TotalAvailNumeric          int64   `json:"total-avail-numeric"`
	SnapSize                   string  `json:"snap-size"`
	SnapSizeNumeric            int     `json:"snap-size-numeric"`
	AllocatedPages             int     `json:"allocated-pages"`
	AvailablePages             int     `json:"available-pages"`
	Overcommit                 string  `json:"overcommit"`
	OvercommitNumeric          int     `json:"overcommit-numeric"`
	OverCommitted              string  `json:"over-committed"`
	OverCommittedNumeric       int     `json:"over-committed-numeric"`
	DiskGroupsCount            int     `json:"disk-groups-counts"`
	DiskGroups                 []Disk  `json:"disk-groups"`
	Volumes                    int     `json:"volumes"`
	PageSize                   string  `json:"page-size"`
	PageSizeNumeric            int     `json:"page-size-numeric"`
	LowThreshold               string  `json:"low-threshold"`
	MiddleThreshold            string  `json:"middle-threshold"`
	HighThreshold              string  `json:"high-threshold"`
	UtilityRunning             string  `json:"utility-running"`
	UtilityRunningNumeric      int     `json:"utility-running-numeric"`
	PreferredOwner             string  `json:"preferred-owner"`
	PreferredOwnerNumeric      int     `json:"preferred-owner-numeric"`
	Owner                      string  `json:"owner"`
	OwnerNumeric               int     `json:"owner-numeric"`
	Rebalance                  string  `json:"rebalance"`
	RebalanceNumeric           int     `json:"rebalance-numeric"`
	Migration                  string  `json:"migration"`
	MigrationNumeric           int     `json:"migration-numeric"`
	ZeroScan                   string  `json:"zero-scan"`
	ZeroScanNumeric            int     `json:"zero-scan-numeric"`
	IdlePageCheck              string  `json:"idle-page-check"`
	IdlePageCheckNumeric       int     `json:"idle-page-check-numeric"`
	ReadFlashCache             string  `json:"read-flash-cache"`
	ReadFlashCacheNumeric      int     `json:"read-flash-cache-numeric"`
	MetadataVolSize            string  `json:"metadata-vol-size"`
	MetadataVolSizeNumeric     int     `json:"metadata-vol-size-numeric"`
	TotalRfcSize               string  `json:"total-rfc-size"`
	TotalRfcSizeNumeric        int64   `json:"total-rfc-size-numeric"`
	AvailableRfcSize           string  `json:"available-rfc-size"`
	AvailableRfcSizeNumeric    int     `json:"available-rfc-size-numeric"`
	ReservedSize               string  `json:"reserved-size"`
	ReservedSizeNumeric        int     `json:"reserved-size-numeric"`
	ReservedUnallocSize        string  `json:"reserved-unalloc-size"`
	ReservedUnallocSizeNumeric int     `json:"reserved-unalloc-size-numeric"`
	PoolSectorFormat           string  `json:"pool-sector-format"`
	PoolSectorFormatNumeric    int     `json:"pool-sector-format-numeric"`
	Health                     string  `json:"health"`
	HealthNumeric              int     `json:"health-numeric"`
	HealthReason               string  `json:"health-reason"`
	HealthRecommendation       string  `json:"health-recommendation"`
	Tiers                      []Tiers `json:"tiers"`
}

type httpPools struct {
	Pools      []Pools  `json:"pools"`
	HttpStatus []Status `json:"status"`
}

func (ps *httpPools) GetAndDeserialize(url string) error {
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
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4Pools(url string) []Pools {
	ps := &httpPools{}
	err := ps.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return ps.Pools
}

func (dk *httpPools) FromJson(body []byte) error {
	err := json.Unmarshal(body, dk)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4PoolsFrom(body []byte) (sti []Pools, err error) {
	diskGp := &httpPools{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = diskGp.Pools
	return
}

func NewMe4PoolsFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]Pools, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []Pools{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Pools{}, err
	}

	// Corrige JSON mal formado
	// duas ocorrencias de disk-groups

	regex := regexp.MustCompile(`"disk-groups":(\s*\d+)`)
	body = regex.ReplaceAll(body, []byte(`"disk-groups-count":$1`))
	return NewMe4PoolsFrom(body)
}
