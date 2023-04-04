package log

type EventInfo struct {
	Event            string
	CustomDimensions map[string]string
}

func (info *EventInfo) ToString() string {
	return ""
}
