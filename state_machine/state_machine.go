package state_machine

import "fmt"

// State represents the state of the state machine.
// Right now this is a simple string, but we'll need to persist more complex state data in the future.
// Things like the number of results, the results themselves, etc. These will feed into a larger statemachine.
type State string

// Event represents the event that triggers a state transition in the state machine. Again this is a simple string,
// but we'll need to pass more complex data in the future. For example, the results of an experiment.
type Event string

// Action represents an action to be executed when entering a state.
type Action func()

type StateMachine struct {
	CurrentState State
	Transitions  map[State]map[Event]State
	OnEnter      map[State]Action
}

// HandleEvent handles the given event and transitions the state machine to the next state if a valid transition exists.
// It checks if there are any transitions defined for the current state, and if a valid transition exists for the given event.
// If a valid transition is found, the state machine's current state is updated to the next state, and the corresponding onEnter action is executed, if defined.
// If no transitions are found for the current state, it prints a message indicating that no transitions are available.
//
// Parameters:
// - event: The event to be handled by the state machine.
//
// Example usage:
//
//	sm := NewMyStateMachine()
//	sm.HandleEvent(EventStart)
func (sm *StateMachine) HandleEvent(event Event) {
	if transitions, ok := sm.Transitions[sm.CurrentState]; ok {
		if nextState, ok := transitions[event]; ok {
			sm.CurrentState = nextState
			if action, ok := sm.OnEnter[nextState]; ok {
				action()
			}
		}
	} else {
		fmt.Printf("No transitions found for the current state %s\n", sm.CurrentState)
	}
}

// GetCurrentState returns the current state of the state machine.
func (sm *StateMachine) GetCurrentState() State {
	return sm.CurrentState
}
