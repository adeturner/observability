# Observability

## Logging

```
observability.Logger("Info", fmt.Sprintf("a=%v\n", a))
```

## Timing

```
t1 := observability.Timer{}
t1.Start(extendedTiming, fmt.Sprintf("%s:1 loadImageFromFile=%s", e.corrID, path))
// do stuff
t1.EndAndPrintStderr(extendedTiming)
```

## log memory consumption

```
observability.LogMemory("Info")
```
