package Me

type PowerSupplies struct {
	ObjectName                string `json:"object-name"`
	Meta                      string `json:"meta"`
	DurableID                 string `json:"durable-id"`
	EnclosureID               int    `json:"enclosure-id"`
	DomID                     int    `json:"dom-id"`
	SerialNumber              string `json:"serial-number"`
	PartNumber                string `json:"part-number"`
	Description               string `json:"description"`
	Name                      string `json:"name"`
	FwRevision                string `json:"fw-revision"`
	Revision                  string `json:"revision"`
	Model                     string `json:"model"`
	Vendor                    string `json:"vendor"`
	Location                  string `json:"location"`
	Position                  string `json:"position"`
	PositionNumeric           int    `json:"position-numeric"`
	DashLevel                 string `json:"dash-level"`
	FruShortname              string `json:"fru-shortname"`
	MfgDate                   string `json:"mfg-date"`
	MfgDateNumeric            int    `json:"mfg-date-numeric"`
	MfgLocation               string `json:"mfg-location"`
	MfgVendorID               string `json:"mfg-vendor-id"`
	ConfigurationSerialnumber string `json:"configuration-serialnumber"`
	Dc12V                     int    `json:"dc12v"`
	Dc5V                      int    `json:"dc5v"`
	Dc33V                     int    `json:"dc33v"`
	Dc12I                     int    `json:"dc12i"`
	Dc5I                      int    `json:"dc5i"`
	Dctemp                    int    `json:"dctemp"`
	Health                    string `json:"health"`
	HealthNumeric             int    `json:"health-numeric"`
	HealthReason              string `json:"health-reason"`
	HealthRecommendation      string `json:"health-recommendation"`
	Status                    string `json:"status"`
	StatusNumeric             int    `json:"status-numeric"`
	Fan                       []Fans `json:"fan"`
}
