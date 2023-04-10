package log

import (
	"MoneyManager/src/utility/utils"
	"fmt"
)

type ErrorInfo struct {
	Error            error
	TraceID          string
	UID              uint32
	CustomDimensions map[string]string
}

func (info *ErrorInfo) ToString() string {
	return fmt.Sprintf("[Error][%s][%d] %s args: \n%s", info.TraceID, info.UID, info.Error.Error(), utils.PrettyPrintJson(info.CustomDimensions))
}
