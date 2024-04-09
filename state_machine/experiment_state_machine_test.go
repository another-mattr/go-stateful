package state_machine

import (
	"testing"
)

func TestNewExperimentStateMachine(t *testing.T) {
	sm := NewExperimentStateMachine()
	if sm.currentState != Initial {
		t.Errorf("Expected initial state to be %v, but got %v", Initial, sm.currentState)
	}
}

func TestTransitionFromInitial(t *testing.T) {
	sm := NewExperimentStateMachine()
	sm.HandleEvent(BeginExperimentEvent)
	if sm.currentState != NoResults {
		t.Errorf("Expected state to be %v after BeginExperimentEvent, but got %v", NoResults, sm.currentState)
	}
}

func TestResultsTransitions(t *testing.T) {
	sm := NewExperimentStateMachine()
	sm.HandleEvent(BeginExperimentEvent)

	sm.HandleEvent(ResultsEvent)
	if sm.currentState != OneResult {
		t.Errorf("Expected state to be %v after ResultsEvent, but got %v", OneResult, sm.currentState)
	}

	sm.HandleEvent(ResultsEvent)
	if sm.currentState != TwoResults {
		t.Errorf("Expected state to be %v after ResultsEvent, but got %v", TwoResults, sm.currentState)
	}
	
	sm.HandleEvent(ResultsEvent)
	if sm.currentState != AnalyzeResults {
		t.Errorf("Expected state to be %v after ResultsEvent, but got %v", AnalyzeResults, sm.currentState)
	}
}
