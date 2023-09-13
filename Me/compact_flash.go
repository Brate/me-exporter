package Me

type CompactFlash struct {
	ObjectName           string `json:"object-name"`
	Meta                 string `json:"meta"`
	DurableID            string `json:"durable-id"`
	ControllerID         string `json:"controller-id"`
	ControllerIDNumeric  int    `json:"controller-id-numeric"`
	Name                 string `json:"name"`
	Status               string `json:"status"`
	StatusNumeric        int    `json:"status-numeric"`
	CacheFlush           string `json:"cache-flush"`
	CacheFlushNumeric    int    `json:"cache-flush-numeric"`
	Health               string `json:"health"`
	HealthNumeric        int    `json:"health-numeric"`
	HealthReason         string `json:"health-reason"`
	HealthRecommendation string `json:"health-recommendation"`
}
