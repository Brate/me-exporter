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

type httpFrus struct {
	EnclosureSku []EnclosureSku `json:"enclosure-sku"`
	EnclosureFru []EnclosureFru `json:"enclosure-fru"`
	HttpStatus   []Status       `json:"status"`
}

func (f *httpFrus) GetAndDeserialize(url string) error {
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

func NewMe4Frus(url string) []EnclosureFru {
	f := &httpFrus{}
	err := f.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return f.EnclosureFru
}

func (sti *httpFrus) FromJson(body []byte) error {
	err := json.Unmarshal(body, sti)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4FrusFrom(body []byte) (sti []EnclosureFru, err error) {
	hst := &httpFrus{}
	err = json.Unmarshal(body, hst)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return
	}

	sti = hst.EnclosureFru
	return
}

func NewMe4FrusFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]EnclosureFru, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []EnclosureFru{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []EnclosureFru{}, err
	}

	return NewMe4FrusFrom(body)
}
