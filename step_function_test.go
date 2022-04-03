package go_timeseries_generator

import (
	"fmt"
	"github.com/lmpizarro/go_timeseries_generator"
	"testing"
)

func TestStepFunction(t *testing.T) {
	vals := go_timeseries_generator.StepFunction(1024, 500, 100)
	filename := fmt.Sprintf("img/step_function.png")
	err := go_timeseries_generator.Plt(vals, filename)

	if err != nil {
		t.Error(err)
	}
}
