package Me

import (
	"encoding/json"
	"fmt"
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
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = hst.EnclosureFru
	return
}

func NewMe4FrusFromRequest(client *http.Client, req *http.Request) ([]EnclosureFru, error) {
	resp, err := client.Do(req)
	if err != nil {
		return []EnclosureFru{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []EnclosureFru{}, err
	}

	return NewMe4FrusFrom(body)
}
