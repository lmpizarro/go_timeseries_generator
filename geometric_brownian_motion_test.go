package go_timeseries_generator

import (
	"github.com/lmpizarro/go_timeseries_generator"
	"testing"
)

func TestGeometricBrownianMotion(t *testing.T) {
	vals := go_timeseries_generator.GeometricBrownianMotionDefault(1024)
	err := go_timeseries_generator.Plt(vals, "img/geometric_brownian_motion.png")
	if err != nil {
		t.Error(err)
	}
}
