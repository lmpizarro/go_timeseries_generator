package go_timeseries_generator

import (
	"fmt"
	"github.com/lmpizarro/go_timeseries_generator"
	"testing"
)

func TestSineWave(t *testing.T) {
	sine := go_timeseries_generator.SineWave(1024)

	filename := fmt.Sprintf("img/sinewave.png")
	err := go_timeseries_generator.Plt(sine, filename)
	if err != nil {
		t.Error(err)
	}
}
