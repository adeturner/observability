package observability

import (
	"fmt"
	"os"
	"time"
)

var appName string
var loggingOn, loggingLevel string

// SetAppName -
func SetAppName(s string) {
	appName = s
}

// getLogHdr function
func getLogHdr() string {
	if appName != "" {
		return "[" + appName + "] "
	}
	return ""
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

	if loggingOn == "" {
		loggingOn = os.Getenv("LOG_ENABLED")
		loggingLevel = os.Getenv("LOG_LEVEL")

		fmt.Println(fmt.Sprintf("1 loggingLevel=%s loggingOn=%s", loggingLevel, loggingOn))

		if loggingOn == "" {
			loggingOn = "true"
		}
		if loggingLevel == "" {
			loggingLevel = "INFO"
		}
	}

	if loggingOn == "true" {

		caller := Caller{}
		t := time.Now()

		if errorType == "Exit" {
			fmt.Fprintf(os.Stdout, "Q %s %s %s\t%s\n", t.Format("2006-01-02 15:04:05.0000"), getLogHdr(), caller.get(n), logString)
			os.Exit(0)
		} else if errorType == "Fatal" {
			fmt.Fprintf(os.Stdout, "F %s %s %s\t%s\n", t.Format("2006-01-02 15:04:05.0000"), getLogHdr(), caller.get(n), logString)
			os.Exit(3)
		} else if errorType == "Debug" && loggingLevel == "DEBUG" {
			fmt.Fprintf(os.Stdout, "D %s %s %s\t%s\n", t.Format("2006-01-02 15:04:05.0000"), getLogHdr(), caller.get(n), logString)
		} else if errorType == "Info" && (loggingLevel == "INFO" || loggingLevel == "DEBUG") {
			fmt.Fprintf(os.Stdout, "I %s %s %s\t%s\n", t.Format("2006-01-02 15:04:05.0000"), getLogHdr(), caller.get(n), logString)
		} else if errorType == "Warn" && (loggingLevel == "WARN" || loggingLevel == "INFO" || loggingLevel == "DEBUG") {
			fmt.Fprintf(os.Stdout, "W %s %s %s\t%s\n", t.Format("2006-01-02 15:04:05.0000"), getLogHdr(), caller.get(n), logString)
		} else if errorType == "Error" && (loggingLevel == "ERROR" || loggingLevel == "WARN" || loggingLevel == "INFO" || loggingLevel == "DEBUG") {
			fmt.Fprintf(os.Stdout, "E %s %s %s\t%s\n", t.Format("2006-01-02 15:04:05.0000"), getLogHdr(), caller.get(n), logString)
		} else if errorType == "" {
			fmt.Fprintf(os.Stdout, "? %s %s %s\t%s\n", t.Format("2006-01-02 15:04:05.0000"), getLogHdr(), caller.get(n), logString)
		}
	}

}
