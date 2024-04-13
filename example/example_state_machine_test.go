package example

import (
	"testing"

	sm "github.com/another-mattr/go-stateful/state_machine"
)

func TestNewMyStateMachine(t *testing.T) {
	mySm := NewMyStateMachine()
	if mySm.CurrentState != sm.State(Initial) {
		t.Errorf("Expected initial state to be %v, but got %v", Initial, mySm.CurrentState)
	}
}

func TestTransitionFromInitial(t *testing.T) {
	mySm := NewMyStateMachine()
	mySm.HandleEvent(sm.Event(BeginExperiment))
	if mySm.CurrentState != sm.State(NoResults) {
		t.Errorf("Expected state to be %v after BeginExperimentEvent, but got %v", NoResults, mySm.CurrentState)
	}
}

func TestResultsTransitions(t *testing.T) {
	mySm := NewMyStateMachine()
	mySm.HandleEvent(sm.Event(BeginExperiment))

	mySm.HandleEvent(sm.Event(ResultsEvent))
	if mySm.CurrentState != sm.State(OneResult) {
		t.Errorf("Expected state to be %v after ResultsEvent, but got %v", OneResult, mySm.CurrentState)
	}

	mySm.HandleEvent(sm.Event(ResultsEvent))
	if mySm.CurrentState != sm.State(TwoResults) {
		t.Errorf("Expected state to be %v after ResultsEvent, but got %v", TwoResults, mySm.CurrentState)
	}
	
	mySm.HandleEvent(sm.Event(ResultsEvent))
	if mySm.CurrentState != sm.State(AnalyzeResults) {
		t.Errorf("Expected state to be %v after ResultsEvent, but got %v", AnalyzeResults, mySm.CurrentState)
	}
}
