package observability

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/google/uuid"
)

var appName string
var corrId string
var loggingOn, loggingLevel string

// SetAppName -
func SetAppName(s string) {
	appName = s
}

// SetCorrId -
func SetCorrId(s string) {
	corrId = s
}

// GenCorrId -
func GenCorrId() {
	corrId = uuid.New().String()
}

// ClearCorrId -
func ClearCorrId() {
	corrId = ""
}

// GetCorrId -
func GetCorrId() string {
	return corrId
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

func LogEnvVars() {
	for _, pair := range os.Environ() {
		log("Info", pair, 3)
	}
}

// Log wraps glog
func log(errorType string, logString string, n int) {

	if loggingOn == "" {
		loggingOn = os.Getenv("LOG_ENABLED")
		loggingLevel = os.Getenv("LOG_LEVEL")

		if loggingOn == "" {
			loggingOn = "true"
		}
		if loggingLevel == "" {
			loggingLevel = "DEBUG"
		}
	}

	if loggingOn == "true" {

		caller := Caller{}
		t := time.Now()

		// In a single-threaded process, the thread ID is equal to the process ID
		p := fmt.Sprintf("%d:%d", syscall.Getpid(), syscall.Gettid())
		proc := p
		if corrId != "" {
			proc = fmt.Sprintf("%s %s", p, corrId)
		}

		// format message
		msg := fmt.Sprintf("%s [%s] %s %s\t%s\n", t.Format("2006-01-02 15:04:05.0000"), proc, getLogHdr(), caller.get(n), logString)

		if errorType == "Exit" {
			fmt.Fprintf(os.Stdout, "Q %s", msg)
			os.Exit(0)
		} else if errorType == "Fatal" {
			fmt.Fprintf(os.Stdout, "F %s", msg)
			os.Exit(3)
		} else if errorType == "Debug" && loggingLevel == "DEBUG" {
			fmt.Fprintf(os.Stdout, "D %s", msg)
		} else if errorType == "Info" && (loggingLevel == "INFO" || loggingLevel == "DEBUG") {
			fmt.Fprintf(os.Stdout, "I %s", msg)
		} else if errorType == "Warn" && (loggingLevel == "WARN" || loggingLevel == "INFO" || loggingLevel == "DEBUG") {
			fmt.Fprintf(os.Stdout, "W %s", msg)
		} else if errorType == "Error" && (loggingLevel == "ERROR" || loggingLevel == "WARN" || loggingLevel == "INFO" || loggingLevel == "DEBUG") {
			fmt.Fprintf(os.Stdout, "E %s", msg)
		} else if errorType == "" {
			fmt.Fprintf(os.Stdout, "? %s", msg)
		}
	}

}
