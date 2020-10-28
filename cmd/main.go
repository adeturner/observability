package main

import (
	"github.com/adeturner/observability"
)

func main() {

	observability.Logger("Debug", "Debug message")
	observability.Logger("Info", "Info message")
	observability.Logger("Warn", "Warn message")
	observability.Logger("Error", "Error message")

	observability.GenCorrId()
	observability.Logger("Info", "Test message with CorrId")
	observability.SetCausationId(observability.GetCorrId())
	observability.GenCorrId()
	observability.Logger("Info", "Test message with CorrId and CausationId")

	t := observability.Timer{}
	t.Start(true, "Timing Stage 1 - should print")
	t.EndAndPrint(true)
	t.Start(false, "Timing Stage 2 - should not print")
	t.EndAndPrint(false)

	observability.GenCorrId()
	observability.LogMemory("Info")

	observability.GenCorrId()
	m := observability.Metrics{}
	m.Init()
	m.SetInteger("count", 1)
	m.SetFloat("decimal", 0.1)
	m.Dump()

	observability.GenCorrId()
	observability.Logger("Fatal", "Fatal message")
	// wont get here:
	observability.Logger("Exit", "Exit message")

}
