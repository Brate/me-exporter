package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

type SystemCacheSettings struct {
	ObjectName                string                      `json:"object-name"`
	Meta                      string                      `json:"meta"`
	DurableID                 string                      `json:"durable-id"`
	OperationMode             string                      `json:"operation-mode"`
	OperationModeNumeric      int                         `json:"operation-mode-numeric"`
	PiFormat                  string                      `json:"pi-format"`
	PiFormatNumeric           int                         `json:"pi-format-numeric"`
	CacheBlockSize            int                         `json:"cache-block-size"`
	ControllerCacheParameters []ControllerCacheParameters `json:"controller-cache-parameters"`
}
type httpCacheParameters struct {
	CacheSettings []SystemCacheSettings `json:"cache-settings"`
	HttpStatus    []Status              `json:"status"`
}

func (hcp *httpCacheParameters) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, hcp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return err
	}

	return nil
}

func NewMe4CacheParameters(url string) []SystemCacheSettings {
	cp := &httpCacheParameters{}
	err := cp.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return cp.CacheSettings
}

func (hcp *httpCacheParameters) FromJson(body []byte) error {
	err := json.Unmarshal(body, hcp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return err
	}

	return nil
}

func NewMe4CacheSettingsFrom(body []byte) (sti []SystemCacheSettings, err error) {
	hcp := &httpCacheParameters{}
	err = json.Unmarshal(body, hcp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return
	}

	sti = hcp.CacheSettings
	return
}

func NewMe4CacheSettingsFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]SystemCacheSettings, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []SystemCacheSettings{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []SystemCacheSettings{}, err
	}

	return NewMe4CacheSettingsFrom(body)
}

// We need to correct the JSON structure for this to work
// DELL ME Controller sends a broken JSON with repeated keys, like this:
//
//	"cache-settings":[{
//	   "object-name":"system-cache-parameters",
//	   "cache-block-size":512,
//	   "controller-cache-parameters":[{
//	         "object-name":"controller-a-cache-parameters",
//	         "controller-id":"A",
//	       } ],
//	   "controller-cache-parameters":[{
//	         "object-name":"controller-b-cache-parameters",
//	         "controller-id":"B",
//	       }]
//	 }]
func (scs *SystemCacheSettings) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(strings.NewReader(string(data)))
	correctedMap := make(map[string]interface{})

	_, err := dec.Token()
	if err != nil {
		return err
	}

	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			return err
		}
		var val interface{}
		err = dec.Decode(&val)
		if err != nil {
			return err
		}

		if key != "controller-cache-parameters" {
			correctedMap[key.(string)] = val
			continue
		}

		_, exists := correctedMap[key.(string)]
		if !exists {
			correctedMap[key.(string)] = []interface{}{}
		}

		arrVal := val.([]interface{})
		correctedMap[key.(string)] = append(correctedMap[key.(string)].([]interface{}), arrVal[0])
	}

	jsonBytes, err := json.Marshal(correctedMap)
	if err != nil {
		return err
	}

	type Alias SystemCacheSettings
	var a Alias
	if err := json.Unmarshal(jsonBytes, &a); err != nil {
		return err
	}

	*scs = SystemCacheSettings(a)

	return nil
}
