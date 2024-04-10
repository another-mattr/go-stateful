package state_machine

import "fmt"

// State represents the state of the state machine.
// Right now this is a simple string, but we'll need to persist more complex state data in the future.
// Things like the number of results, the results themselves, etc. These will feed into a larger statemachine.
type State string

const (
	Initial		 State = "Initial"
	BeginExperiment State = "BeginExperiment"
	NoResults       State = "NoResults"
	OneResult       State = "OneResult"
	TwoResults      State = "TwoResults"
	AnalyzeResults  State = "AnalyzeResults"
)

// Event represents the event that triggers a state transition in the state machine. Again this is a simple string,
// but we'll need to pass more complex data in the future. For example, the results of an experiment.
type Event string

const (
	BeginExperimentEvent Event = "BeginExperimentEvent"
	WaitForResults	   Event = "WaitForResults"
	ResultsEvent         Event = "ResultsEvent"
)

// Action represents an action to be executed when entering a state.
type Action func()

type ExperimentStateMachine struct {
	currentState State
	transitions  map[State]map[Event]State
	onEnter      map[State]Action
}

// NewExperimentStateMachine creates a new instance of the ExperimentStateMachine.
// It initializes the state machine with the initial state, defines transitions between states, and defines actions to be executed when entering a state.
func NewExperimentStateMachine() *ExperimentStateMachine {
	sm := &ExperimentStateMachine{
		currentState: Initial,
		transitions: map[State]map[Event]State{
			Initial: {
				BeginExperimentEvent: BeginExperiment,
			},
			BeginExperiment: {
				WaitForResults: NoResults,
			},
			NoResults: {
				ResultsEvent: OneResult,
			},
			OneResult: {
				ResultsEvent: TwoResults,
			},
			TwoResults: {
				ResultsEvent: AnalyzeResults,
			},
		},
		onEnter: make(map[State]Action),
	}

	sm.onEnter[BeginExperiment] = func() {
		fmt.Println("Experiment started")
		fmt.Println("Requested 3 Janky Builds")
		sm.HandleEvent(WaitForResults)
	}

	sm.onEnter[NoResults] = func() {	fmt.Println("No results yet") }
	sm.onEnter[OneResult] = func() {	fmt.Println("One result so far") }
	sm.onEnter[TwoResults] = func() {	fmt.Println("Two results so far") }
	sm.onEnter[AnalyzeResults] = func() {	fmt.Println("The results are in! Let's analyze.") }

	return sm
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
//   sm := NewExperimentStateMachine()
//   sm.HandleEvent(EventStart)
func (sm *ExperimentStateMachine) HandleEvent(event Event) {
	if transitions, ok := sm.transitions[sm.currentState]; ok {
		if nextState, ok := transitions[event]; ok {
			sm.currentState = nextState
			if action, ok := sm.onEnter[nextState]; ok {
				action()
			}
		}
	} else {
		fmt.Printf("No transitions found for the current state %s\n", sm.currentState)
	}
}

// GetCurrentState returns the current state of the state machine.
func (sm *ExperimentStateMachine) GetCurrentState() State {
	return sm.currentState
}
