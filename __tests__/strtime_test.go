package strtime

import (
	"fmt"
	"testing"
	"time"

	"github.com/afeiship/go-strtime"
)

func TestFormat(t *testing.T) {
	// create a time : 2019-01-01 00:00:00
	currentTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

	// fmt.Println("当前时间  : ", currentTime)
	// fmt.Println("当前时间字符串: ", currentTime.String())
	// fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
	// fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))

	// start test date
	format := "YYYY-MM-DD"
	res := strtime.Format(format, currentTime)
	if res != "2019-01-01" {
		t.Error("Format failed!")
	}

	// start test time
	format = "HH:mm:SS"
	res = strtime.Format(format, currentTime)
	if res != "00:00:00" {
		t.Error("Format failed!")
	}

	// start test datetime
	format = "YYYY-MM-DD HH:mm:SS"
	res = strtime.Format(format, currentTime)
	if res != "2019-01-01 00:00:00" {
		t.Error("Format failed!")
	}

	// test dateHook
	format = "date"
	res = strtime.Format(format, currentTime)
	if res != "2019-01-01" {
		t.Error("Format failed!")
	}

	// test timeHook
	// start test time
	format = "time"
	res = strtime.Format(format, currentTime)
	if res != "00:00:00" {
		t.Error("Format failed!")
	}

	// test datetimeHook
	// start test datetime
	format = "datetime"
	res = strtime.Format(format, currentTime)
	if res != "2019-01-01 00:00:00" {
		t.Error("Format failed!")
	}

	// test dbHook
	format = "dbdt"
	res = strtime.Format(format, currentTime)
	fmt.Println(res)
	if res != "20190101_000000" {
		t.Error("Format failed!")
	}
}

func TestNow(t *testing.T) {
	// create a time : 2019-01-01 00:00:00
	nowTs := strtime.Now()
	nowTime := time.Unix(nowTs, 0)
	t.Log("Now time: ", nowTime)
	t.Log("Now timestamp: ", strtime.Now())
}
