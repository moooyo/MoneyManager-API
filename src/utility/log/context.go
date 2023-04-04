package log

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TraceID = "trace_id"

func RegisterTraceID(context *gin.Context) {
	traceID := context.GetHeader(TraceID)
	if traceID == "" {
		traceID = GenerateTraceID()
	}

	context.Set(TraceID, traceID)
}

func GenerateTraceID() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("GenerateTraceID error. err: ", err)
		return "INVALID"
	}

	uuidStr := uuid.String()
	md5Sum := md5.Sum([]byte(uuidStr))
	md5Str := fmt.Sprintf("%x", md5Sum)
	md5Str = strings.ToUpper(md5Str)
	return md5Str
}
