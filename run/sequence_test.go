package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
)

func TestSequenceInactive(t *testing.T) {
	// Construct the flow control object.
	flow := floc.NewFlowControl()
	defer flow.Release()

	flow.Complete(nil)

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	job := Sequence(jobIncrement)

	// Run the job.
	floc.Run(flow, state, updateCounter, job)

	if getCounter(state) != 0 {
		t.Fatalf("%s expects counter to be zero", t.Name())
	}
}
