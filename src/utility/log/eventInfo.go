package log

import (
	"MoneyManager/src/utility/utils"
	"fmt"
)

type EventInfo struct {
	Name             string
	TraceID          string
	UID              uint32
	CustomDimensions map[string]string
}

func (info *EventInfo) ToString() string {
	return fmt.Sprintf("[Event][%s][%d] %s args: \n%s", info.TraceID, info.UID, info.Name, utils.PrettyPrintJson(info.CustomDimensions))
}
