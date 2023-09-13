package Me

type Expanders struct {
	ObjectName           string `json:"object-name"`
	Meta                 string `json:"meta"`
	DurableID            string `json:"durable-id"`
	EnclosureID          int    `json:"enclosure-id"`
	DrawerID             int    `json:"drawer-id"`
	DomID                int    `json:"dom-id"`
	PathID               string `json:"path-id"`
	PathIDNumeric        int    `json:"path-id-numeric"`
	Name                 string `json:"name"`
	Location             string `json:"location"`
	Status               string `json:"status"`
	StatusNumeric        int    `json:"status-numeric"`
	ExtendedStatus       string `json:"extended-status"`
	FwRevision           string `json:"fw-revision"`
	Health               string `json:"health"`
	HealthNumeric        int    `json:"health-numeric"`
	HealthReason         string `json:"health-reason"`
	HealthRecommendation string `json:"health-recommendation"`
}
