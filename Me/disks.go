package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"io"
	"net/http"
)

type Drives struct {
	ObjectName                  string `json:"object-name"`
	Meta                        string `json:"meta"`
	DurableID                   string `json:"durable-id"`
	EnclosureID                 int    `json:"enclosure-id"`
	DrawerID                    int    `json:"drawer-id"`
	Slot                        int    `json:"slot"`
	Location                    string `json:"location"`
	URL                         string `json:"url"`
	Port                        int    `json:"port"`
	ScsiID                      int    `json:"scsi-id"`
	Blocksize                   int    `json:"blocksize"`
	Blocks                      int64  `json:"blocks"`
	SerialNumber                string `json:"serial-number"`
	Vendor                      string `json:"vendor"`
	Model                       string `json:"model"`
	Revision                    string `json:"revision"`
	SecondaryChannel            int    `json:"secondary-channel"`
	ContainerIndex              int    `json:"container-index"`
	MemberIndex                 int    `json:"member-index"`
	Description                 string `json:"description"`
	DescriptionNumeric          int    `json:"description-numeric"`
	Architecture                string `json:"architecture"`
	ArchitectureNumeric         int    `json:"architecture-numeric"`
	Interface                   string `json:"interface"`
	InterfaceNumeric            int    `json:"interface-numeric"`
	SinglePorted                string `json:"single-ported"`
	SinglePortedNumeric         int    `json:"single-ported-numeric"`
	Type                        string `json:"type"`
	TypeNumeric                 int    `json:"type-numeric"`
	Usage                       string `json:"usage"`
	UsageNumeric                int    `json:"usage-numeric"`
	JobRunning                  string `json:"job-running"`
	JobRunningNumeric           int    `json:"job-running-numeric"`
	State                       string `json:"state"`
	CurrentJobCompletion        string `json:"current-job-completion"`
	Blink                       int    `json:"blink"`
	LocatorLed                  string `json:"locator-led"`
	LocatorLedNumeric           int    `json:"locator-led-numeric"`
	Speed                       int    `json:"speed"`
	Smart                       string `json:"smart"`
	SmartNumeric                int    `json:"smart-numeric"`
	DualPort                    int    `json:"dual-port"`
	Error                       int    `json:"error"`
	FcP1Channel                 int    `json:"fc-p1-channel"`
	FcP1DeviceID                int    `json:"fc-p1-device-id"`
	FcP1NodeWwn                 string `json:"fc-p1-node-wwn"`
	FcP1PortWwn                 string `json:"fc-p1-port-wwn"`
	FcP1UnitNumber              int    `json:"fc-p1-unit-number"`
	FcP2Channel                 int    `json:"fc-p2-channel"`
	FcP2DeviceID                int    `json:"fc-p2-device-id"`
	FcP2NodeWwn                 string `json:"fc-p2-node-wwn"`
	FcP2PortWwn                 string `json:"fc-p2-port-wwn"`
	FcP2UnitNumber              int    `json:"fc-p2-unit-number"`
	DriveDownCode               int    `json:"drive-down-code"`
	Owner                       string `json:"owner"`
	OwnerNumeric                int    `json:"owner-numeric"`
	Index                       int    `json:"index"`
	Rpm                         int    `json:"rpm"`
	Size                        string `json:"size"`
	SizeNumeric                 int64  `json:"size-numeric"`
	SectorFormat                string `json:"sector-format"`
	SectorFormatNumeric         int    `json:"sector-format-numeric"`
	TransferRate                string `json:"transfer-rate"`
	TransferRateNumeric         int    `json:"transfer-rate-numeric"`
	Attributes                  string `json:"attributes"`
	AttributesNumeric           int    `json:"attributes-numeric"`
	EnclosureWwn                string `json:"enclosure-wwn"`
	EnclosureURL                string `json:"enclosure-url"`
	Status                      string `json:"status"`
	ReconState                  string `json:"recon-state"`
	ReconStateNumeric           int    `json:"recon-state-numeric"`
	CopybackState               string `json:"copyback-state"`
	CopybackStateNumeric        int    `json:"copyback-state-numeric"`
	VirtualDiskSerial           string `json:"virtual-disk-serial"`
	DiskGroup                   string `json:"disk-group"`
	StoragePoolName             string `json:"storage-pool-name"`
	StorageTier                 string `json:"storage-tier"`
	StorageTierNumeric          int    `json:"storage-tier-numeric"`
	SsdLifeLeft                 string `json:"ssd-life-left"`
	SsdLifeLeftNumeric          int    `json:"ssd-life-left-numeric"`
	LedStatus                   string `json:"led-status"`
	LedStatusNumeric            int    `json:"led-status-numeric"`
	DiskDsdCount                int    `json:"disk-dsd-count"`
	SpunDown                    int    `json:"spun-down"`
	NumberOfIos                 int    `json:"number-of-ios"`
	TotalDataTransferred        string `json:"total-data-transferred"`
	TotalDataTransferredNumeric int    `json:"total-data-transferred-numeric"`
	AvgRspTime                  int    `json:"avg-rsp-time"`
	FdeState                    string `json:"fde-state"`
	FdeStateNumeric             int    `json:"fde-state-numeric"`
	LockKeyID                   string `json:"lock-key-id"`
	ImportLockKeyID             string `json:"import-lock-key-id"`
	FdeConfigTime               string `json:"fde-config-time"`
	FdeConfigTimeNumeric        int    `json:"fde-config-time-numeric"`
	Temperature                 string `json:"temperature"`
	TemperatureNumeric          int    `json:"temperature-numeric"`
	TemperatureStatus           string `json:"temperature-status"`
	TemperatureStatusNumeric    int    `json:"temperature-status-numeric"`
	PiFormatted                 string `json:"pi-formatted"`
	PiFormattedNumeric          int    `json:"pi-formatted-numeric"`
	PowerOnHours                int    `json:"power-on-hours"`
	ExtendedStatus              int    `json:"extended-status"`
	Health                      string `json:"health"`
	HealthNumeric               int    `json:"health-numeric"`
	HealthReason                string `json:"health-reason"`
	HealthRecommendation        string `json:"health-recommendation"`
}

type httpDisks struct {
	Drives     []Drives `json:"drives"`
	HttpStatus []Status `json:"status"`
}

func (sti *httpDisks) GetAndDeserialize(url string) error {
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

func NewMe4Disks(url string) []Drives {
	sti := &httpDisks{}
	err := sti.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return sti.Drives
}

func (sti *httpDisks) FromJson(body []byte) error {
	err := json.Unmarshal(body, sti)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4SDisksFrom(body []byte) (sti []Drives, err error) {
	hst := &httpDisks{}
	err = json.Unmarshal(body, hst)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = hst.Drives
	return
}

func NewMe4DisksFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]Drives, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []Drives{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Drives{}, err
	}

	return NewMe4SDisksFrom(body)
}
