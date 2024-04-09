package main

import "github.com/another-mattr/go-stateful/state_machine"

func main() {

	// Create a new instance of the ExperimentStateMachine.
	// Theoretically, we would want a way to pass in an initial state, such as something retrieved from a database.
	// This would allow us to completely restore the state of the state machine if the application crashes or needs to power
	// down for some reason.
	sm := state_machine.NewExperimentStateMachine()

	// These events would likely be from webhooks or some other external source. They might even come from
	// something monitoring janky for results on a regular basis.
	sm.HandleEvent(state_machine.BeginExperimentEvent)
	sm.HandleEvent(state_machine.ResultsEvent)
	sm.HandleEvent(state_machine.ResultsEvent)
	sm.HandleEvent(state_machine.ResultsEvent)
}
