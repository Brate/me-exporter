package Me

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CacheSettings struct {
	ObjectName                string                      `json:"object-name"`
	Meta                      string                      `json:"meta"`
	OperationMode             string                      `json:"operation-mode"`
	OperationModeNumeric      int                         `json:"operation-mode-numeric"`
	PiFormat                  string                      `json:"pi-format"`
	PiFormatNumeric           int                         `json:"pi-format-numeric"`
	CacheBlockSize            int                         `json:"cache-block-size"`
	ControllerCacheParameters []ControllerCacheParameters `json:"controller-cache-parameters"`
}
type httpCacheParameters struct {
	CacheSettings []CacheSettings `json:"cache-settings"`
	HttpStatus    []Status        `json:"status"`
}

func (cp *httpCacheParameters) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, cp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4CacheParameters(url string) []CacheSettings {
	cp := &httpCacheParameters{}
	err := cp.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return cp.CacheSettings
}

func (cp *httpCacheParameters) FromJson(body []byte) error {
	err := json.Unmarshal(body, cp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4CacheSettings(body []byte) (sti []CacheSettings, err error) {
	hcp := &httpCacheParameters{}
	err = json.Unmarshal(body, hcp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = hcp.CacheSettings
	return
}

func NewMe4CacheSettingsFromRequest(client *http.Client, req *http.Request) ([]CacheSettings, error) {
	resp, err := client.Do(req)
	if err != nil {
		return []CacheSettings{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []CacheSettings{}, err
	}

	return NewMe4CacheSettings(body)
}
