# Observability

## Logging

### Options

- "Exit" - Exits
- "Fatal" - Panics
- "Debug"
- "Info"
- "Warn"
- "Error"

### Usage

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