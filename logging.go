package observability

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
)

const cAppName = ""

// GetLogHdr function
func GetLogHdr() string {
	return "[" + cAppName + "] "
	//return ""
}

// LogMemory prints memory usage to the trace
func LogMemory(errorType string) {
	logN(errorType, fmt.Sprintf("%s", getMemUsageStr()), 4)
}

// Logger is externalised for first level caller that doesnt care
func Logger(errorType string, logString string) {
	log(errorType, logString, 3)
}

// logN is not externalised to cater for internal callers from observability
func logN(errorType string, logString string, n int) {
	log(errorType, logString, n)
}

// Log wraps glog
func log(errorType string, logString string, n int) {

	caller := Caller{}
	t := time.Now()

	if errorType == "Exit" {
		glog.Exitf("%s %s\t%s", GetLogHdr(), caller.get(n), logString)
	} else if errorType == "Fatal" {
		glog.Fatalf("%s %s\t%s", GetLogHdr(), caller.get(n), logString)
	} else if errorType == "Debug" {
		glog.Infof("%s %s\t%s", GetLogHdr(), caller.get(n), logString)
		fmt.Fprintf(os.Stdout, "%s %s %s\t%s\n", t.Format("2006-01-02 15:04:05.0000"), GetLogHdr(), caller.get(n), logString)
	} else if errorType == "Info" {
		glog.Infof("%s %s\t%s", GetLogHdr(), caller.get(n), logString)
		fmt.Fprintf(os.Stdout, "%s %s %s\t%s\n", t.Format("2006-01-02 15:04:05.0000"), GetLogHdr(), caller.get(n), logString)
	} else if errorType == "Warn" {
		glog.Warningf("%s %s\t%s", GetLogHdr(), caller.get(n), logString)
	} else if errorType == "Error" {
		glog.Errorf("%s %s\t%s", GetLogHdr(), caller.get(n), logString)
	} else if errorType == "" {
		glog.Infof("%s %s\t%s", GetLogHdr(), caller.get(n), logString)
	}

}
