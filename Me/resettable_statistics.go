package Me

type ResettableStatistics struct {
	ObjectName            string `json:"object-name"`
	Meta                  string `json:"meta"`
	SerialNumber          string `json:"serial-number"`
	TimeSinceReset        int    `json:"time-since-reset"`
	TimeSinceSample       int    `json:"time-since-sample"`
	NumberOfReads         int64  `json:"number-of-reads"`
	NumberOfWrites        int64  `json:"number-of-writes"`
	DataRead              string `json:"data-read"`
	DataReadNumeric       int64  `json:"data-read-numeric"`
	DataWritten           string `json:"data-written"`
	DataWrittenNumeric    int64  `json:"data-written-numeric"`
	BytesPerSecond        string `json:"bytes-per-second"`
	BytesPerSecondNumeric int    `json:"bytes-per-second-numeric"`
	Iops                  int    `json:"iops"`
	AvgRspTime            int    `json:"avg-rsp-time"`
	AvgReadRspTime        int    `json:"avg-read-rsp-time"`
	AvgWriteRspTime       int    `json:"avg-write-rsp-time"`
}
