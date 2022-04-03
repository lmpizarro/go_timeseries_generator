package go_timeseries_generator

import (
	"fmt"
	"github.com/lmpizarro/go_timeseries_generator"
	"testing"
)

func TestUniformProcess(t *testing.T) {
	vals := go_timeseries_generator.UniformProcess(1024)

	filename := fmt.Sprintf("img/uniform_process.png")
	err := go_timeseries_generator.Plt(vals, filename)
	if err != nil {
		t.Error(err)
	}
}
