package Me

type HostPortStatistics struct {
	ObjectName             string `json:"object-name"`
	Meta                   string `json:"meta"`
	DurableID              string `json:"durable-id"`
	BytesPerSecond         string `json:"bytes-per-second"`
	BytesPerSecondNumeric  int    `json:"bytes-per-second-numeric"`
	Iops                   int    `json:"iops"`
	NumberOfReads          int64  `json:"number-of-reads"`
	NumberOfWrites         int64  `json:"number-of-writes"`
	DataRead               string `json:"data-read"`
	DataReadNumeric        int64  `json:"data-read-numeric"`
	DataWritten            string `json:"data-written"`
	DataWrittenNumeric     int64  `json:"data-written-numeric"`
	QueueDepth             int    `json:"queue-depth"`
	AvgRspTime             int    `json:"avg-rsp-time"`
	AvgReadRspTime         int    `json:"avg-read-rsp-time"`
	AvgWriteRspTime        int    `json:"avg-write-rsp-time"`
	ResetTime              string `json:"reset-time"`
	ResetTimeNumeric       int    `json:"reset-time-numeric"`
	StartSampleTime        string `json:"start-sample-time"`
	StartSampleTimeNumeric int    `json:"start-sample-time-numeric"`
	StopSampleTime         string `json:"stop-sample-time"`
	StopSampleTimeNumeric  int    `json:"stop-sample-time-numeric"`
}
