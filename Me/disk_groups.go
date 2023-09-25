package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"io"
	"net/http"
)

type Disk struct {
	ObjectName                       string `json:"object-name"`
	Meta                             string `json:"meta"`
	Name                             string `json:"name"`
	Blocksize                        int    `json:"blocksize"`
	Size                             string `json:"size"`
	SizeNumeric                      int64  `json:"size-numeric"`
	Freespace                        string `json:"freespace"`
	FreespaceNumeric                 int64  `json:"freespace-numeric"`
	RawSize                          string `json:"raw-size"`
	RawSizeNumeric                   int64  `json:"raw-size-numeric"`
	StorageType                      string `json:"storage-type"`
	StorageTypeNumeric               int    `json:"storage-type-numeric"`
	Pool                             string `json:"pool"`
	PoolSerialNumber                 string `json:"pool-serial-number"`
	StorageTier                      string `json:"storage-tier"`
	StorageTierNumeric               int    `json:"storage-tier-numeric"`
	TotalPages                       int    `json:"total-pages"`
	AllocatedPages                   int    `json:"allocated-pages"`
	AvailablePages                   int    `json:"available-pages"`
	PoolPercentage                   int    `json:"pool-percentage"`
	PerformanceRank                  int    `json:"performance-rank"`
	Owner                            string `json:"owner"`
	OwnerNumeric                     int    `json:"owner-numeric"`
	PreferredOwner                   string `json:"preferred-owner"`
	PreferredOwnerNumeric            int    `json:"preferred-owner-numeric"`
	Raidtype                         string `json:"raidtype"`
	RaidtypeNumeric                  int    `json:"raidtype-numeric"`
	Diskcount                        int    `json:"diskcount"`
	Sparecount                       int    `json:"sparecount"`
	Chunksize                        string `json:"chunksize"`
	Status                           string `json:"status"`
	StatusNumeric                    int    `json:"status-numeric"`
	Lun                              int64  `json:"lun"`
	MinDriveSize                     string `json:"min-drive-size"`
	MinDriveSizeNumeric              int64  `json:"min-drive-size-numeric"`
	CreateDate                       string `json:"create-date"`
	CreateDateNumeric                int    `json:"create-date-numeric"`
	CacheReadAhead                   string `json:"cache-read-ahead"`
	CacheReadAheadNumeric            int    `json:"cache-read-ahead-numeric"`
	CacheFlushPeriod                 int    `json:"cache-flush-period"`
	ReadAheadEnabled                 string `json:"read-ahead-enabled"`
	ReadAheadEnabledNumeric          int    `json:"read-ahead-enabled-numeric"`
	WriteBackEnabled                 string `json:"write-back-enabled"`
	WriteBackEnabledNumeric          int    `json:"write-back-enabled-numeric"`
	JobRunning                       string `json:"job-running"`
	CurrentJob                       string `json:"current-job"`
	CurrentJobNumeric                int    `json:"current-job-numeric"`
	CurrentJobCompletion             string `json:"current-job-completion"`
	NumArrayPartitions               int    `json:"num-array-partitions"`
	LargestFreePartitionSpace        string `json:"largest-free-partition-space"`
	LargestFreePartitionSpaceNumeric int    `json:"largest-free-partition-space-numeric"`
	NumDrivesPerLowLevelArray        int    `json:"num-drives-per-low-level-array"`
	NumExpansionPartitions           int    `json:"num-expansion-partitions"`
	NumPartitionSegments             int    `json:"num-partition-segments"`
	NewPartitionLba                  string `json:"new-partition-lba"`
	NewPartitionLbaNumeric           int    `json:"new-partition-lba-numeric"`
	ArrayDriveType                   string `json:"array-drive-type"`
	ArrayDriveTypeNumeric            int    `json:"array-drive-type-numeric"`
	DiskDescription                  string `json:"disk-description"`
	DiskDescriptionNumeric           int    `json:"disk-description-numeric"`
	IsJobAutoAbortable               string `json:"is-job-auto-abortable"`
	IsJobAutoAbortableNumeric        int    `json:"is-job-auto-abortable-numeric"`
	SerialNumber                     string `json:"serial-number"`
	Blocks                           int64  `json:"blocks"`
	DiskDsdEnableVdisk               string `json:"disk-dsd-enable-vdisk"`
	DiskDsdEnableVdiskNumeric        int    `json:"disk-dsd-enable-vdisk-numeric"`
	DiskDsdDelayVdisk                int    `json:"disk-dsd-delay-vdisk"`
	ScrubDurationGoal                int    `json:"scrub-duration-goal"`
	AdaptTargetSpareCapacity         string `json:"adapt-target-spare-capacity"`
	AdaptTargetSpareCapacityNumeric  int    `json:"adapt-target-spare-capacity-numeric"`
	AdaptActualSpareCapacity         string `json:"adapt-actual-spare-capacity"`
	AdaptActualSpareCapacityNumeric  int    `json:"adapt-actual-spare-capacity-numeric"`
	AdaptCriticalCapacity            string `json:"adapt-critical-capacity"`
	AdaptCriticalCapacityNumeric     int    `json:"adapt-critical-capacity-numeric"`
	AdaptDegradedCapacity            string `json:"adapt-degraded-capacity"`
	AdaptDegradedCapacityNumeric     int    `json:"adapt-degraded-capacity-numeric"`
	AdaptLinearVolumeBoundary        int    `json:"adapt-linear-volume-boundary"`
	PoolSectorFormat                 string `json:"pool-sector-format"`
	PoolSectorFormatNumeric          int    `json:"pool-sector-format-numeric"`
	Health                           string `json:"health"`
	HealthNumeric                    int    `json:"health-numeric"`
	HealthReason                     string `json:"health-reason"`
	HealthRecommendation             string `json:"health-recommendation"`
}

type httpDiskGroups struct {
	DiskGroups []Disk   `json:"disk-groups"`
	HttpStatus []Status `json:"status"`
}

func (sti *httpDiskGroups) GetAndDeserialize(url string) error {
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

func NewMe4DiskGroups(url string) []Disk {
	sti := &httpDiskGroups{}
	err := sti.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return sti.DiskGroups
}

func (sti *httpDiskGroups) FromJson(body []byte) error {
	err := json.Unmarshal(body, sti)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4SDiskGroupsFrom(body []byte) (sti []Disk, err error) {
	hst := &httpDiskGroups{}
	err = json.Unmarshal(body, hst)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = hst.DiskGroups
	return
}

func NewMe4DiskGroupsFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]Disk, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []Disk{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Disk{}, err
	}

	return NewMe4SDiskGroupsFrom(body)
}
