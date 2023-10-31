package Me

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"

	//"github.com/go-kit/log"
	//"github.com/go-kit/log/level"
	"io"
	"net/http"
)

type httpPortStatistics struct {
	HostPortStatistics []HostPortStatistics `json:"host-port-statistics"`
	HttpStatus         []Status             `json:"status"`
}

func (ps *httpPortStatistics) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, ps)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4PortStatistics(url string) []HostPortStatistics {
	ps := &httpPortStatistics{}
	err := ps.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return ps.HostPortStatistics
}

func NewMe4PortStatisticsFrom(body []byte) (portS []HostPortStatistics, err error) {
	portSta := &httpPortStatistics{}
	err = json.Unmarshal(body, portSta)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v\n", err)
		err = errors.Errorf("Unmarshal error: %s", err)
		return
	}

	portS = portSta.HostPortStatistics
	return
}
