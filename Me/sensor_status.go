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

type SensorStatus struct {
	ObjectName          string `json:"object-name"`
	Meta                string `json:"meta"`
	DurableID           string `json:"durable-id"`
	EnclosureID         int    `json:"enclosure-id"`
	DrawerID            string `json:"drawer-id"`
	DrawerIDNumeric     int    `json:"drawer-id-numeric"`
	ControllerID        string `json:"controller-id"`
	ControllerIDNumeric int    `json:"controller-id-numeric"`
	SensorName          string `json:"sensor-name"`
	Value               string `json:"value"`
	Status              string `json:"status"`
	StatusNumeric       int    `json:"status-numeric"`
	Container           string `json:"container"`
	ContainerNumeric    int    `json:"container-numeric"`
	SensorType          string `json:"sensor-type"`
	SensorTypeNumeric   int    `json:"sensor-type-numeric"`
}

type httpSensors struct {
	SensorStatus []SensorStatus `json:"sensors"`
	HttpStatus   []Status       `json:"status"`
}

func (ss *httpSensors) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, ss)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return err
	}

	return nil
}

func NewMe4Sensors(url string) []SensorStatus {
	ss := &httpSensors{}
	err := ss.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return ss.SensorStatus
}

func (ss *httpSensors) FromJson(body []byte) error {
	err := json.Unmarshal(body, ss)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return err
	}

	return nil
}

func NewMe4SensorStatusFrom(body []byte) (sti []SensorStatus, err error) {
	diskGp := &httpSensors{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return
	}

	sti = diskGp.SensorStatus
	return
}

func NewMe4SensorStatusFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]SensorStatus, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []SensorStatus{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []SensorStatus{}, err
	}

	return NewMe4SensorStatusFrom(body)
}
