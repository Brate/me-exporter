package Me

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SasStatusControllerA struct {
	ObjectName          string `json:"object-name"`
	Meta                string `json:"meta"`
	EnclosureID         int    `json:"enclosure-id"`
	DrawerID            int    `json:"drawer-id"`
	ExpanderID          int    `json:"expander-id"`
	Controller          string `json:"controller"`
	ControllerNumeric   int    `json:"controller-numeric"`
	WidePortIndex       int    `json:"wide-port-index"`
	PhyIndex            int    `json:"phy-index"`
	WidePortRole        string `json:"wide-port-role"`
	WidePortRoleNumeric int    `json:"wide-port-role-numeric"`
	WidePortNum         int    `json:"wide-port-num"`
	Type                string `json:"type"`
	Status              string `json:"status"`
	StatusNumeric       int    `json:"status-numeric"`
	ElemStatus          string `json:"elem-status"`
	ElemStatusNumeric   int    `json:"elem-status-numeric"`
	ElemDisabled        string `json:"elem-disabled"`
	ElemDisabledNumeric int    `json:"elem-disabled-numeric"`
	ElemReason          string `json:"elem-reason"`
	ElemReasonNumeric   int    `json:"elem-reason-numeric"`
	ChangeCounter       string `json:"change-counter"`
	CodeViolations      string `json:"code-violations"`
	DisparityErrors     string `json:"disparity-errors"`
	CrcErrors           string `json:"crc-errors"`
	ConnCrcErrors       string `json:"conn-crc-errors"`
	LostDwords          string `json:"lost-dwords"`
	InvalidDwords       string `json:"invalid-dwords"`
	ResetErrorCounter   string `json:"reset-error-counter"`
	FlagBits            string `json:"flag-bits"`
}

type httpExpanderStatus struct {
	SasStatusControllerA []SasStatusControllerA `json:"sas-status-controller-a"`
	HttpStatus           []Status               `json:"status"`
}

func (es *httpExpanderStatus) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, es)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4ExpanderStatus(url string) []SasStatusControllerA {
	es := &httpExpanderStatus{}
	err := es.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return es.SasStatusControllerA
}

func (es *httpExpanderStatus) FromJson(body []byte) error {
	err := json.Unmarshal(body, es)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4ExpanderStatusFrom(body []byte) (sti []SasStatusControllerA, err error) {
	hst := &httpExpanderStatus{}
	err = json.Unmarshal(body, hst)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = hst.SasStatusControllerA
	return
}

func NewMe4ExpanderStatusFromRequest(client *http.Client, req *http.Request) ([]SasStatusControllerA, error) {
	resp, err := client.Do(req)
	if err != nil {
		return []SasStatusControllerA{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []SasStatusControllerA{}, err
	}

	return NewMe4ExpanderStatusFrom(body)
}
