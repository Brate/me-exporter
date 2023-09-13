package Me

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ServiceTagInfo struct {
	ObjectName  string `json:"object-name"`
	Meta        string `json:"meta"`
	EnclosureID int    `json:"enclosure-id"`
	ServiceTag  string `json:"service-tag"`
}

type httpServiceTagInfo struct {
	ServiceTagInfo []ServiceTagInfo `json:"service-tag-info"`
	HttpStatus     []Status         `json:"status"`
}

func (sti *httpServiceTagInfo) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = sti.FromJson(body)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func (sti *httpServiceTagInfo) FromJson(body []byte) error {
	err := json.Unmarshal(body, sti)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4ServiceTagInfoFrom(body []byte) (sti []ServiceTagInfo, err error) {
	hst := &httpServiceTagInfo{}
	err = json.Unmarshal(body, hst)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = hst.ServiceTagInfo
	return
}

func NewMe4ServiceTagInfoFromRequest(client *http.Client, req *http.Request) ([]ServiceTagInfo, error) {
	resp, err := client.Do(req)
	if err != nil {
		return []ServiceTagInfo{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []ServiceTagInfo{}, err
	}

	return NewMe4ServiceTagInfoFrom(body)
}

func NewMe4ServiceTagInfo(url string) []ServiceTagInfo {
	sti := &httpServiceTagInfo{}
	err := sti.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return sti.ServiceTagInfo
}
