package log

import (
	"MoneyManager/src/utility/utils"
	"fmt"
)

type TraceInfo struct {
	Message          string
	TraceID          string
	UID              uint32
	CustomDimensions map[string]string
}

func (trace *TraceInfo) ToString() string {
	return fmt.Sprintf("[Trace][%s][%d] %s args: \n%s", trace.TraceID, trace.UID, trace.Message, utils.PrettyPrintJson(trace.CustomDimensions))
}
