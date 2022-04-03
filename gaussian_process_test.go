package go_timeseries_generator

import (
	"fmt"
	"github.com/lmpizarro/go_timeseries_generator"
	"testing"
)

func TestGaussianProcess(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)

	filename := fmt.Sprintf("img/gaussian_process.png")
	err := go_timeseries_generator.Plt(vals, filename)
	if err != nil {
		t.Error(err)
	}
}
