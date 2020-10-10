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
t1 := observability.Timer{}
t1.Start(extendedTiming, fmt.Sprintf("loadImageFromFile=%s", path))
// do stuff
t1.EndAndPrintStderr(extendedTiming)
```

## log memory consumption

```go
observability.LogMemory("Info")
```

