package util

import (
	"fmt"
	"time"
)

func FormatCurrentDay(format string) string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf(format, year, int(month), day)
}
