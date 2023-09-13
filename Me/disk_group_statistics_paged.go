package Me

type DiskGroupStatisticsPaged struct {
	ObjectName             string `json:"object-name"`
	Meta                   string `json:"meta"`
	SerialNumber           string `json:"serial-number"`
	PagesAllocPerMinute    int    `json:"pages-alloc-per-minute"`
	PagesDeallocPerMinute  int    `json:"pages-dealloc-per-minute"`
	PagesReclaimed         int    `json:"pages-reclaimed"`
	NumPagesUnmapPerMinute int    `json:"num-pages-unmap-per-minute"`
}
