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

type Ports struct {
	ObjectName             string      `json:"object-name"`
	Meta                   string      `json:"meta"`
	DurableID              string      `json:"durable-id"`
	Controller             string      `json:"controller"`
	ControllerNumeric      int         `json:"controller-numeric"`
	Port                   string      `json:"port"`
	PortType               string      `json:"port-type"`
	PortTypeNumeric        int         `json:"port-type-numeric"`
	Media                  string      `json:"media"`
	TargetID               string      `json:"target-id"`
	Status                 string      `json:"status"`
	StatusNumeric          int         `json:"status-numeric"`
	ActualSpeed            string      `json:"actual-speed"`
	ActualSpeedNumeric     int         `json:"actual-speed-numeric"`
	ConfiguredSpeed        string      `json:"configured-speed"`
	ConfiguredSpeedNumeric int         `json:"configured-speed-numeric"`
	FanOut                 int         `json:"fan-out"`
	Health                 string      `json:"health"`
	HealthNumeric          int         `json:"health-numeric"`
	HealthReason           string      `json:"health-reason"`
	HealthRecommendation   string      `json:"health-recommendation"`
	IscsiPort              []IscsiPort `json:"iscsi-port"`
}

type httpPorts struct {
	Port       []Ports  `json:"port"`
	HttpStatus []Status `json:"status"`
}

func (p *httpPorts) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4Ports(url string) []Ports {
	p := &httpPorts{}
	err := p.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return p.Port
}

func (dk *httpPorts) FromJson(body []byte) error {
	err := json.Unmarshal(body, dk)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return err
	}

	return nil
}

func NewMe4PortsFrom(body []byte) (sti []Ports, err error) {
	diskGp := &httpPorts{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return
	}
	sti = diskGp.Port
	return
}

func NewMe4PortsFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]Ports, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []Ports{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Ports{}, err
	}
	return NewMe4PortsFrom(body)
}
