package Me

type IscsiPort struct {
	ObjectName                   string `json:"object-name"`
	Meta                         string `json:"meta"`
	IPVersion                    string `json:"ip-version"`
	IPAddress                    string `json:"ip-address"`
	Gateway                      string `json:"gateway"`
	Netmask                      string `json:"netmask"`
	DefaultRouter                string `json:"default-router"`
	LinkLocalAddress             string `json:"link-local-address"`
	MacAddress                   string `json:"mac-address"`
	SfpStatus                    string `json:"sfp-status"`
	SfpStatusNumeric             int    `json:"sfp-status-numeric"`
	SfpPresent                   string `json:"sfp-present"`
	SfpPresentNumeric            int    `json:"sfp-present-numeric"`
	SfpVendor                    string `json:"sfp-vendor"`
	SfpPartNumber                string `json:"sfp-part-number"`
	SfpRevision                  string `json:"sfp-revision"`
	Sfp10GCompliance             string `json:"sfp-10G-compliance"`
	Sfp10GComplianceNumeric      int    `json:"sfp-10G-compliance-numeric"`
	SfpEthernetCompliance        string `json:"sfp-ethernet-compliance"`
	SfpEthernetComplianceNumeric int    `json:"sfp-ethernet-compliance-numeric"`
	SfpCableTechnology           string `json:"sfp-cable-technology"`
	SfpCableTechnologyNumeric    int    `json:"sfp-cable-technology-numeric"`
	SfpCableLength               int    `json:"sfp-cable-length"`
}
