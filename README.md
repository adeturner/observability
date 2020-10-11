# Observability

## Usage

### Set via Command line options

```txt

import (
    "flag"
)

usage: goexecuteable -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]
```

### Set in code

```go

import (
    "flag"
)

flag.PrintDefaults()
flag.Set("logtostderr", true)
flag.Set("stderrthreshold", "INFO")  //  INFO|WARN|FATAL logs at or above this threshold go to stderr
flag.Parse()

// others...
flag.Set("alsologtostderr", false)
flag.Set("v", "log level for V logs")
flag.Set("vmodule", "comma-separated list of pattern=N settings for file-filtered logging")
flag.Set("log_backtrace_at", "when logging hits line file:N, emit a stack trace")
```

## Logging

### Code level logging Options

- "Exit" - Exits
- "Fatal" - Panics
- "Debug"
- "Info"
- "Warn"
- "Error"

### Logging Usage

```go
observability.Logger("Info", fmt.Sprintf("a=%v\n", a))
```

## Timing

```go
func (timer *Timer) Start(timing bool, str string) {
func (timer *Timer) EndAndPrint(timing bool) {
func (timer *Timer) EndAndPrintStderr(timing bool) {
```

```go
timingOn := true // toggle, eg set via environment variable
t1 := observability.Timer{}
t1.Start(timingOn, fmt.Sprintf("loadImageFromFile=%s", path))
// do stuff
t1.EndAndPrintStderr(timingOn)
```

## Write memory consumption to log output

```go
observability.LogMemory("Info")
```

## Metrics

```go
func (ms *Metrics) setKeyValue(key string, m Metric)
func (ms *Metrics) SetDuration(key string, d time.Duration)
func (ms *Metrics) SetInteger(key string, i int)
func (ms *Metrics) SetFloat(key string, f float64)
func (ms *Metrics) Dump()
```

```go
type a {
    metrics     observability.Metrics
}

a.metrics.Init()

t1 = time.Now()
// Do load
count++
t2 = time.Now()
ela = t2.Sub(t1)
a.metrics.SetInteger("LOAD_COUNT", count)
a.metrics.SetDuration("LOAD_TOTAL_TIME_S", ela)
```