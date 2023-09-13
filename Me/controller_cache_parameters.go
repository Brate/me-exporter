package Me

type ControllerCacheParameters struct {
	ObjectName                string `json:"object-name"`
	Meta                      string `json:"meta"`
	DurableID                 string `json:"durable-id"`
	ControllerID              string `json:"controller-id"`
	ControllerIDNumeric       int    `json:"controller-id-numeric"`
	Name                      string `json:"name"`
	WriteBackStatus           string `json:"write-back-status"`
	WriteBackStatusNumeric    int    `json:"write-back-status-numeric"`
	CompactFlashStatus        string `json:"compact-flash-status"`
	CompactFlashStatusNumeric int    `json:"compact-flash-status-numeric"`
	CompactFlashHealth        string `json:"compact-flash-health"`
	CompactFlashHealthNumeric int    `json:"compact-flash-health-numeric"`
	CacheFlush                string `json:"cache-flush"`
	CacheFlushNumeric         int    `json:"cache-flush-numeric"`
}
