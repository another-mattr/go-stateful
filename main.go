package main

import (
	"github.com/another-mattr/go-stateful/example"
	sm "github.com/another-mattr/go-stateful/state_machine"
)

func main() {

	// Create a new instance of the MyStateMachine.
	// Theoretically, we would want a way to pass in an initial state, such as something retrieved from a database.
	// This would allow us to completely restore the state of the state machine if the application crashes or needs to power
	// down for some reason.
	mySm := example.NewMyStateMachine()

	// These events would likely be from webhooks or some other external source. They might even come from
	// something monitoring janky for results on a regular basis.
	mySm.HandleEvent(sm.Event(example.BeginExperimentEvent))
	mySm.HandleEvent(sm.Event(example.ResultsEvent))
	mySm.HandleEvent(sm.Event(example.ResultsEvent))
	mySm.HandleEvent(sm.Event(example.ResultsEvent))
}
