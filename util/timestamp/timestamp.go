package timestamp

import (
	"time"

	"github.com/dustin/go-humanize"
)

type Timestamp int64

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}

// FormatTimeRel format the timestamp for human consumption
func (t Timestamp) FormatTimeRel() string {
	return humanize.Time(t.Time())
}

func (t Timestamp) FormatTime() string {
	return t.Time().Format("Mon Jan 2 15:04:05 2006 +0200")
}
