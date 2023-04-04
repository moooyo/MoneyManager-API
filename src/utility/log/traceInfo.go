package log

type TraceInfo struct {
	Message          string
	CustomDimensions map[string]string
}

func (trace *TraceInfo) ToString() string {
	return ""
}
