package log

type ErrorInfo struct {
	Error            error
	CustomDimensions map[string]string
}

func (info *ErrorInfo) ToString() string {
	return ""
}
