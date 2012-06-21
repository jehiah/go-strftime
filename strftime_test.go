package strftime

import (
	"time"
	"testing"
)

func TestFormat(t *testing.T) {
	tt := time.Unix(1340244776, 0)
	f := Format("%Y-%m-%d %H:%M:%S", tt)
	if f != "2012-06-20 22:12:56" {
		t.Errorf("got %s", f)
	}
}