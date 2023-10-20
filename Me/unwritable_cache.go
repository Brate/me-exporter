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

type UnwritableCache struct {
	ObjectName            string `json:"object-name"`
	Meta                  string `json:"meta"`
	UnwritableAPercentage int    `json:"unwritable-a-percentage"`
	UnwritableBPercentage int    `json:"unwritable-b-percentage"`
}

type httpUnwritableCache struct {
	UnwritableCache []UnwritableCache `json:"unwritable-cache"`
	HttpStatus      []Status          `json:"status"`
}

func (uwc *httpUnwritableCache) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, uwc)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return err
	}

	return nil
}

func NewMe4UnwritableCache(url string) []UnwritableCache {
	uwc := &httpUnwritableCache{}
	err := uwc.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return nil
	}
	return uwc.UnwritableCache
}

func (uwc *httpUnwritableCache) FromJson(body []byte) error {
	err := json.Unmarshal(body, uwc)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return err
	}

	return nil
}

func NewMe4UnwritableCacheFrom(body []byte) (sti []UnwritableCache, err error) {
	diskGp := &httpUnwritableCache{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return
	}

	sti = diskGp.UnwritableCache
	return
}

func NewMe4UnwritableCacheFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]UnwritableCache, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []UnwritableCache{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []UnwritableCache{}, err
	}

	return NewMe4UnwritableCacheFrom(body)
}
