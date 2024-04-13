package example

import (
	"fmt"

	sm "github.com/another-mattr/go-stateful/state_machine"
)

type MyState sm.State

const (
	Initial         MyState = "Initial"
	BeginExperiment MyState = "BeginExperiment"
	NoResults       MyState = "NoResults"
	OneResult       MyState = "OneResult"
	TwoResults      MyState = "TwoResults"
	AnalyzeResults  MyState = "AnalyzeResults"
)

type MyEvent sm.Event

const (
	BeginExperimentEvent MyEvent = "BeginExperimentEvent"
	WaitForResults       MyEvent = "WaitForResults"
	ResultsEvent         MyEvent = "ResultsEvent"
)

type MyStateMachine struct {
	sm.StateMachine
}

// NewMyStateMachine creates a new instance of the MyStateMachine.
// It initializes the state machine with the initial state, defines transitions between states, and defines actions to be executed when entering a state.
func NewMyStateMachine() *MyStateMachine {
	mySm := &MyStateMachine{
		StateMachine: sm.StateMachine{
			CurrentState: sm.State(Initial),
			Transitions: map[sm.State]map[sm.Event]sm.State{
				sm.State(Initial): {
					sm.Event(BeginExperiment): sm.State(BeginExperiment),
				},
				sm.State(BeginExperiment): {
					sm.Event(WaitForResults): sm.State(NoResults),
				},
				sm.State(NoResults): {
					sm.Event(ResultsEvent): sm.State(OneResult),
				},
				sm.State(OneResult): {
					sm.Event(ResultsEvent): sm.State(TwoResults),
				},
				sm.State(TwoResults): {
					sm.Event(ResultsEvent): sm.State(AnalyzeResults),
				},
			},
			OnEnter: make(map[sm.State]sm.Action),
		},
	}

	mySm.OnEnter[sm.State(BeginExperiment)] = func() {
		fmt.Println("Experiment started")
		fmt.Println("Requested 3 Janky Builds")
		mySm.HandleEvent(sm.Event(WaitForResults))
	}

	mySm.OnEnter[sm.State(NoResults)] = func() { fmt.Println("No results yet") }
	mySm.OnEnter[sm.State(OneResult)] = func() { fmt.Println("One result so far") }
	mySm.OnEnter[sm.State(TwoResults)] = func() { fmt.Println("Two results so far") }
	mySm.OnEnter[sm.State(AnalyzeResults)] = func() { fmt.Println("The results are in! Let's analyze.") }
	return mySm
}
