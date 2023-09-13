package Me

type NetworkParameters struct {
	ObjectName             string `json:"object-name"`
	Meta                   string `json:"meta"`
	DurableID              string `json:"durable-id"`
	ActiveVersion          int    `json:"active-version"`
	IPAddress              string `json:"ip-address"`
	Gateway                string `json:"gateway"`
	SubnetMask             string `json:"subnet-mask"`
	MacAddress             string `json:"mac-address"`
	AddressingMode         string `json:"addressing-mode"`
	AddressingModeNumeric  int    `json:"addressing-mode-numeric"`
	LinkSpeed              string `json:"link-speed"`
	LinkSpeedNumeric       int    `json:"link-speed-numeric"`
	DuplexMode             string `json:"duplex-mode"`
	DuplexModeNumeric      int    `json:"duplex-mode-numeric"`
	AutoNegotiation        string `json:"auto-negotiation"`
	AutoNegotiationNumeric int    `json:"auto-negotiation-numeric"`
	Health                 string `json:"health"`
	HealthNumeric          int    `json:"health-numeric"`
	HealthReason           string `json:"health-reason"`
	HealthRecommendation   string `json:"health-recommendation"`
	PingBroadcast          string `json:"ping-broadcast"`
	PingBroadcastNumeric   int    `json:"ping-broadcast-numeric"`
}
