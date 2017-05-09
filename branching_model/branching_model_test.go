package branching_model

import "testing"

func TestNew(t *testing.T) {
	model := New(2)
	model.ParseFile("../dataset/trace_1.txt")
	model.DisplayStatistics()
}
