package Me

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"io"
	"net/http"
)

type Tiers struct {
	ObjectName           string `json:"object-name"`
	Meta                 string `json:"meta"`
	SerialNumber         string `json:"serial-number"`
	Pool                 string `json:"pool"`
	Tier                 string `json:"tier"`
	TierNumeric          int    `json:"tier-numeric"`
	PoolPercentage       int    `json:"pool-percentage"`
	Diskcount            int    `json:"diskcount"`
	RawSize              string `json:"raw-size"`
	RawSizeNumeric       int64  `json:"raw-size-numeric"`
	TotalSize            string `json:"total-size"`
	TotalSizeNumeric     int64  `json:"total-size-numeric"`
	AllocatedSize        string `json:"allocated-size"`
	AllocatedSizeNumeric int64  `json:"allocated-size-numeric"`
	AvailableSize        string `json:"available-size"`
	AvailableSizeNumeric int64  `json:"available-size-numeric"`
	AffinitySize         string `json:"affinity-size"`
	AffinitySizeNumeric  int    `json:"affinity-size-numeric"`
}

type httpTiers struct {
	Tiers      []Tiers  `json:"tiers"`
	HttpStatus []Status `json:"status"`
}

func (t *httpTiers) GetAndDeserialize(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, t)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}
func NewMe4TiersInfo(url string) []Tiers {
	sti := &httpTiers{}
	err := sti.GetAndDeserialize(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %v", err)
		return nil
	}
	return sti.Tiers
}
func (t *httpTiers) FromJson(body []byte) error {
	err := json.Unmarshal(body, t)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	return nil
}

func NewMe4TiersFrom(body []byte) (sti []Tiers, err error) {
	diskGp := &httpTiers{}
	err = json.Unmarshal(body, diskGp)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return
	}

	sti = diskGp.Tiers
	return
}

func NewMe4TiersFromRequest(client *http.Client, req *http.Request, log log.Logger) ([]Tiers, error) {
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(log).Log("msg", "request error", "error", err)

		return []Tiers{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Tiers{}, err
	}

	return NewMe4TiersFrom(body)
}
