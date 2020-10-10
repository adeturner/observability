package observability

import (
	"fmt"
	"regexp"
	"runtime"
)

// Caller is externalised
type Caller struct {
	fn   string
	line string
}

func (caller *Caller) init() {
	// _, caller.fn, caller.line, _ := runtime.Caller(1)
}

func (caller *Caller) get(n int) string {
	_, fn, line, _ := runtime.Caller(n)
	reStr := ".*/"
	re := regexp.MustCompile(reStr)
	fn = re.ReplaceAllLiteralString(fn, "")
	str := fmt.Sprintf("%s:%d", fn, line)
	if len(fn) > 12 {
		fn = fn[:12]
	}
	return str
}

// Get calls internal function get
func (caller *Caller) Get(n int) string {
	return caller.get(n)
}
