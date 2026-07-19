package terminal

import (
	"time"
)

func Time(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
