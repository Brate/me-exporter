package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type Fans struct {
	ObjectName           string `json:"object-name"`
	Meta                 string `json:"meta"`
	DurableID            string `json:"durable-id"`
	Name                 string `json:"name"`
	Location             string `json:"location"`
	StatusSes            string `json:"status-ses"`
	StatusSesNumeric     int    `json:"status-ses-numeric"`
	ExtendedStatus       string `json:"extended-status"`
	Status               string `json:"status"`
	StatusNumeric        int    `json:"status-numeric"`
	Speed                int    `json:"speed"`
	Position             string `json:"position"`
	PositionNumeric      int    `json:"position-numeric"`
	SerialNumber         string `json:"serial-number"`
	PartNumber           string `json:"part-number"`
	FwRevision           string `json:"fw-revision"`
	HwRevision           string `json:"hw-revision"`
	LocatorLed           string `json:"locator-led"`
	LocatorLedNumeric    int    `json:"locator-led-numeric"`
	Health               string `json:"health"`
	HealthNumeric        int    `json:"health-numeric"`
	HealthReason         string `json:"health-reason"`
	HealthRecommendation string `json:"health-recommendation"`
}

type httpFans struct {
	Fans       []Fans   `json:"fan"`
	HttpStatus []Status `json:"Status"`
}

func (f *httpFans) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, f)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4Fans(url string) []Fans {
	f := &httpFans{}
	err := f.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return f.Fans
}

func (f *httpFans) FromJson(body []byte) error {
	err := json.Unmarshal(body, f)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4FansFrom(body []byte) (sti []Fans, err error) {
	hst := &httpFans{}
	err = json.Unmarshal(body, hst)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return
	}

	sti = hst.Fans
	return
}

func NewMe4FansFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]Fans, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []Fans{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Fans{}, err
	}

	return NewMe4FansFrom(body)
}
