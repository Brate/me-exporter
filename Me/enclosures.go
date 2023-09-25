package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"io"
	"net/http"
)

type Enclosures struct {
	ObjectName                  string          `json:"object-name"`
	Meta                        string          `json:"meta"`
	DurableID                   string          `json:"durable-id"`
	EnclosureID                 int             `json:"enclosure-id"`
	EnclosureWwn                string          `json:"enclosure-wwn"`
	Name                        string          `json:"name"`
	Type                        string          `json:"type"`
	TypeNumeric                 int             `json:"type-numeric"`
	IomType                     string          `json:"iom-type"`
	IomTypeNumeric              int             `json:"iom-type-numeric"`
	PlatformType                string          `json:"platform-type"`
	PlatformTypeNumeric         int             `json:"platform-type-numeric"`
	BoardModel                  string          `json:"board-model"`
	BoardModelNumeric           int             `json:"board-model-numeric"`
	Location                    string          `json:"location"`
	RackNumber                  int             `json:"rack-number"`
	RackPosition                int             `json:"rack-position"`
	NumberOfCoolingsElements    int             `json:"number-of-coolings-elements"`
	NumberOfDisks               int             `json:"number-of-disks"`
	NumberOfPowerSupplies       int             `json:"number-of-power-supplies"`
	Status                      string          `json:"status"`
	StatusNumeric               int             `json:"status-numeric"`
	ExtendedStatus              string          `json:"extended-status"`
	MidplaneSerialNumber        string          `json:"midplane-serial-number"`
	Vendor                      string          `json:"vendor"`
	Model                       string          `json:"model"`
	FruTlapn                    string          `json:"fru-tlapn"`
	FruShortname                string          `json:"fru-shortname"`
	FruLocation                 string          `json:"fru-location"`
	PartNumber                  string          `json:"part-number"`
	MfgDate                     string          `json:"mfg-date"`
	MfgDateNumeric              int             `json:"mfg-date-numeric"`
	MfgLocation                 string          `json:"mfg-location"`
	Description                 string          `json:"description"`
	Revision                    string          `json:"revision"`
	DashLevel                   string          `json:"dash-level"`
	EmpARev                     string          `json:"emp-a-rev"`
	EmpBRev                     string          `json:"emp-b-rev"`
	Rows                        int             `json:"rows"`
	Columns                     int             `json:"columns"`
	Slots                       int             `json:"slots"`
	LocatorLed                  string          `json:"locator-led"`
	LocatorLedNumeric           int             `json:"locator-led-numeric"`
	DriveOrientation            string          `json:"drive-orientation"`
	DriveOrientationNumeric     int             `json:"drive-orientation-numeric"`
	EnclosureArrangement        string          `json:"enclosure-arrangement"`
	EnclosureArrangementNumeric int             `json:"enclosure-arrangement-numeric"`
	EmpABusid                   string          `json:"emp-a-busid"`
	EmpATargetid                string          `json:"emp-a-targetid"`
	EmpBBusid                   string          `json:"emp-b-busid"`
	EmpBTargetid                string          `json:"emp-b-targetid"`
	EmpA                        string          `json:"emp-a"`
	EmpAChIDRev                 string          `json:"emp-a-ch-id-rev"`
	EmpB                        string          `json:"emp-b"`
	EmpBChIDRev                 string          `json:"emp-b-ch-id-rev"`
	MidplaneType                string          `json:"midplane-type"`
	MidplaneTypeNumeric         int             `json:"midplane-type-numeric"`
	MidplaneRev                 int             `json:"midplane-rev"`
	EnclosurePower              string          `json:"enclosure-power"`
	Pcie2Capable                string          `json:"pcie2-capable"`
	Pcie2CapableNumeric         int             `json:"pcie2-capable-numeric"`
	Health                      string          `json:"health"`
	HealthNumeric               int             `json:"health-numeric"`
	HealthReason                string          `json:"health-reason"`
	HealthRecommendation        string          `json:"health-recommendation"`
	Controllers                 []Controllers   `json:"controllers"`
	PowerSupplies               []PowerSupplies `json:"power-supplies"`
}

type httpEnclosures struct {
	Enclosures []Enclosures `json:"enclosures"`
	HttpStatus []Status     `json:"status"`
}

func (en *httpEnclosures) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, en)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4Enclosures(url string) []Enclosures {
	en := &httpEnclosures{}
	err := en.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return en.Enclosures
}

func (sti *httpEnclosures) FromJson(body []byte) error {
	err := json.Unmarshal(body, sti)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4EnclosuresFrom(body []byte) (sti []Enclosures, err error) {
	hst := &httpEnclosures{}
	err = json.Unmarshal(body, hst)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = hst.Enclosures
	return
}

func NewMe4EnclosuresFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]Enclosures, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []Enclosures{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Enclosures{}, err
	}

	return NewMe4EnclosuresFrom(body)
}
