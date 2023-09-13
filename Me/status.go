package Me

type Status struct {
	ObjectName          string `json:"object-name"`
	Meta                string `json:"meta"`
	ResponseType        string `json:"response-type"`
	ResponseTypeNumeric int    `json:"response-type-numeric"`
	Response            string `json:"response"`
	ReturnCode          int    `json:"return-code"`
	ComponentID         string `json:"component-id"`
	TimeStamp           string `json:"time-stamp"`
	TimeStampNumeric    int    `json:"time-stamp-numeric"`
}
