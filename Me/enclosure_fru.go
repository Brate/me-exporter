package Me

type EnclosureFru struct {
	ObjectName                string `json:"object-name"`
	Meta                      string `json:"meta"`
	Name                      string `json:"name"`
	Description               string `json:"description"`
	PartNumber                string `json:"part-number"`
	SerialNumber              string `json:"serial-number"`
	Revision                  string `json:"revision"`
	DashLevel                 string `json:"dash-level"`
	FruShortname              string `json:"fru-shortname"`
	MfgDate                   string `json:"mfg-date"`
	MfgDateNumeric            int    `json:"mfg-date-numeric"`
	MfgLocation               string `json:"mfg-location"`
	MfgVendorID               string `json:"mfg-vendor-id"`
	FruLocation               string `json:"fru-location"`
	ConfigurationSerialnumber string `json:"configuration-serialnumber"`
	FruStatus                 string `json:"fru-status"`
	FruStatusNumeric          int    `json:"fru-status-numeric"`
	OriginalSerialnumber      string `json:"original-serialnumber"`
	OriginalPartnumber        string `json:"original-partnumber"`
	OriginalRevision          string `json:"original-revision"`
	EnclosureID               int    `json:"enclosure-id"`
}
