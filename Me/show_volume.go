package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"io"
	"net/http"
)

type Volumes struct {
	ObjectName                        string `json:"object-name"`
	Meta                              string `json:"meta"`
	DurableID                         string `json:"durable-id"`
	URL                               string `json:"url"`
	VirtualDiskName                   string `json:"virtual-disk-name"`
	StoragePoolName                   string `json:"storage-pool-name"`
	StoragePoolURL                    string `json:"storage-pool-url"`
	VolumeName                        string `json:"volume-name"`
	Size                              string `json:"size"`
	SizeNumeric                       int64  `json:"size-numeric"`
	TotalSize                         string `json:"total-size"`
	TotalSizeNumeric                  int64  `json:"total-size-numeric"`
	AllocatedSize                     string `json:"allocated-size"`
	AllocatedSizeNumeric              int64  `json:"allocated-size-numeric"`
	StorageType                       string `json:"storage-type"`
	StorageTypeNumeric                int    `json:"storage-type-numeric"`
	PreferredOwner                    string `json:"preferred-owner"`
	PreferredOwnerNumeric             int    `json:"preferred-owner-numeric"`
	Owner                             string `json:"owner"`
	OwnerNumeric                      int    `json:"owner-numeric"`
	SerialNumber                      string `json:"serial-number"`
	WritePolicy                       string `json:"write-policy"`
	WritePolicyNumeric                int    `json:"write-policy-numeric"`
	CacheOptimization                 string `json:"cache-optimization"`
	CacheOptimizationNumeric          int    `json:"cache-optimization-numeric"`
	ReadAheadSize                     string `json:"read-ahead-size"`
	ReadAheadSizeNumeric              int    `json:"read-ahead-size-numeric"`
	VolumeType                        string `json:"volume-type"`
	VolumeTypeNumeric                 int    `json:"volume-type-numeric"`
	VolumeClass                       string `json:"volume-class"`
	VolumeClassNumeric                int    `json:"volume-class-numeric"`
	TierAffinity                      string `json:"tier-affinity"`
	TierAffinityNumeric               int    `json:"tier-affinity-numeric"`
	Snapshot                          string `json:"snapshot"`
	SnapshotRetentionPriority         string `json:"snapshot-retention-priority"`
	SnapshotRetentionPriorityNumeric  int    `json:"snapshot-retention-priority-numeric"`
	VolumeQualifier                   string `json:"volume-qualifier"`
	VolumeQualifierNumeric            int    `json:"volume-qualifier-numeric"`
	Blocksize                         int    `json:"blocksize"`
	Blocks                            int64  `json:"blocks"`
	Capabilities                      string `json:"capabilities"`
	VolumeParent                      string `json:"volume-parent"`
	SnapPool                          string `json:"snap-pool"`
	ReplicationSet                    string `json:"replication-set"`
	Attributes                        string `json:"attributes"`
	VirtualDiskSerial                 string `json:"virtual-disk-serial"`
	VolumeDescription                 string `json:"volume-description"`
	Wwn                               string `json:"wwn"`
	Progress                          string `json:"progress"`
	ProgressNumeric                   int    `json:"progress-numeric"`
	ContainerName                     string `json:"container-name"`
	ContainerSerial                   string `json:"container-serial"`
	AllowedStorageTiers               string `json:"allowed-storage-tiers"`
	AllowedStorageTiersNumeric        int    `json:"allowed-storage-tiers-numeric"`
	ThresholdPercentOfPool            string `json:"threshold-percent-of-pool"`
	ReservedSizeInPages               int    `json:"reserved-size-in-pages"`
	AllocateReservedPagesFirst        string `json:"allocate-reserved-pages-first"`
	AllocateReservedPagesFirstNumeric int    `json:"allocate-reserved-pages-first-numeric"`
	ZeroInitPageOnAllocation          string `json:"zero-init-page-on-allocation"`
	ZeroInitPageOnAllocationNumeric   int    `json:"zero-init-page-on-allocation-numeric"`
	LargeVirtualExtents               string `json:"large-virtual-extents"`
	LargeVirtualExtentsNumeric        int    `json:"large-virtual-extents-numeric"`
	Raidtype                          string `json:"raidtype"`
	RaidtypeNumeric                   int    `json:"raidtype-numeric"`
	PiFormat                          string `json:"pi-format"`
	PiFormatNumeric                   int    `json:"pi-format-numeric"`
	CsReplicationRole                 string `json:"cs-replication-role"`
	CsCopyDest                        string `json:"cs-copy-dest"`
	CsCopyDestNumeric                 int    `json:"cs-copy-dest-numeric"`
	CsCopySrc                         string `json:"cs-copy-src"`
	CsCopySrcNumeric                  int    `json:"cs-copy-src-numeric"`
	CsPrimary                         string `json:"cs-primary"`
	CsPrimaryNumeric                  int    `json:"cs-primary-numeric"`
	CsSecondary                       string `json:"cs-secondary"`
	CsSecondaryNumeric                int    `json:"cs-secondary-numeric"`
	Health                            string `json:"health"`
	HealthNumeric                     int    `json:"health-numeric"`
	HealthReason                      string `json:"health-reason"`
	HealthRecommendation              string `json:"health-recommendation"`
	VolumeGroup                       string `json:"volume-group"`
	GroupKey                          string `json:"group-key"`
}

type httpShowVolume struct {
	Volumes    []Volumes `json:"volumes"`
	HttpStatus []Status  `json:"status"`
}

func (sv *httpShowVolume) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, sv)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4ShowVolume(url string) []Volumes {
	sv := &httpShowVolume{}
	err := sv.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return sv.Volumes
}

func (sv *httpShowVolume) FromJson(body []byte) error {
	err := json.Unmarshal(body, sv)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4VolumesFrom(body []byte) (sti []Volumes, err error) {
	diskGp := &httpShowVolume{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = diskGp.Volumes
	return
}

func NewMe4VolumesFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]Volumes, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []Volumes{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Volumes{}, err
	}

	return NewMe4VolumesFrom(body)
}
