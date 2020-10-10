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
timingOn := true // toggle, eg set via environment variable
t1 := observability.Timer{}
t1.Start(timingOn, fmt.Sprintf("loadImageFromFile=%s", path))
// do stuff
t1.EndAndPrintStderr(timingOn)
```

## log memory consumption

```go
observability.LogMemory("Info")
```

