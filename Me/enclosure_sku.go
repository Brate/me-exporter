package Me

type EnclosureSku struct {
	ObjectName      string `json:"object-name"`
	Meta            string `json:"meta"`
	SkuPartnumber   string `json:"sku-partnumber"`
	SkuSerialnumber string `json:"sku_serialnumber"`
	SkuRevision     string `json:"sku-revision"`
	EnclosureID     int    `json:"enclosure-id"`
}
