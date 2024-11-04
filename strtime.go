package strtime

import (
	"fmt"
	"strings"
	"time"
)

// YYYY-MM-DD HH:MM:SS
// define a pre-defined format hash
var formatHash = map[string]string{
	"YYYY": "2006",
	"MM":   "01",
	"DD":   "02",
	"HH":   "15",
	"mm":   "04",
	"SS":   "05",
}

var defaultHooks = map[string]string{
	"month":    "YYYY-MM",
	"date":     "YYYY-MM-DD",
	"time":     "HH:mm:SS",
	"datetime": "YYYY-MM-DD HH:mm:SS",
	"dbdt":     "YYYYMMDD_HHmmSS",
	"unix":     "unix",
}

func Format(format string, t time.Time) string {
	// 1. check if the format is in the default hooks
	if defaultFormat, ok := defaultHooks[format]; ok {
		format = defaultFormat
	}

	if format == "unix" {
		timestamp := t.Unix()
		return fmt.Sprintf("%d", timestamp)
	}

	for k, v := range formatHash {
		// replace all
		format = strings.Replace(format, k, v, -1)
	}
	return t.Format(format)
}

func Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
