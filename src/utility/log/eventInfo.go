package log

type EventInfo struct {
	Name             string
	CustomDimensions map[string]string
}

func (info *EventInfo) ToString() string {
	return ""
}
