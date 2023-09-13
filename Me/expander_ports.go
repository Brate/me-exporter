package Me

type ExpanderPorts struct {
	ObjectName           string `json:"object-name"`
	Meta                 string `json:"meta"`
	DurableID            string `json:"durable-id"`
	EnclosureID          int    `json:"enclosure-id"`
	Controller           string `json:"controller"`
	ControllerNumeric    int    `json:"controller-numeric"`
	SasPortType          string `json:"sas-port-type"`
	SasPortTypeNumeric   int    `json:"sas-port-type-numeric"`
	SasPortIndex         int    `json:"sas-port-index"`
	Name                 string `json:"name"`
	Status               string `json:"status"`
	StatusNumeric        int    `json:"status-numeric"`
	Health               string `json:"health"`
	HealthNumeric        int    `json:"health-numeric"`
	HealthReason         string `json:"health-reason"`
	HealthRecommendation string `json:"health-recommendation"`
}
