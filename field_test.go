package wrigleyball

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkRunners(t *testing.T, f *Field, first, second, third bool) {
	assert.Equal(t, f.b1.on, first, fmt.Sprintf("First base incorrect - 1st:%t, 2nd:%t, 3rd:%t", f.b1.on, f.b2.on, f.b3.on))
	assert.Equal(t, f.b2.on, second, fmt.Sprintf("Second base incorrect - 1st:%t, 2nd:%t, 3rd:%t", f.b1.on, f.b2.on, f.b3.on))
	assert.Equal(t, f.b3.on, third, fmt.Sprintf("Third base incorrect - 1st:%t, 2nd:%t, 3rd:%t", f.b1.on, f.b2.on, f.b3.on))
}

// TestAdvRunnersSingleNoManOn checks that a single places a runner on first.
func TestAdvRunnersSingleNoManOn(t *testing.T) {
	f := NewField()

	f.AdvRunners(1, true)

	checkRunners(t, f, true, false, false)
	assert.Equal(t, f.score[f.atBat], 0)
}

// TestAdvRunnersDoubleNoManOn checks that a double places a runner on second.
func TestAdvRunnersDoubleNoManOn(t *testing.T) {
	f := NewField()

	f.AdvRunners(2, true)

	checkRunners(t, f, false, true, false)
	assert.Equal(t, f.score[f.atBat], 0)
}

// TestAdvRunnersTripleNoManOn checks that a triple places a runner on third.
func TestAdvRunnersTripleNoManOn(t *testing.T) {
	f := NewField()

	f.AdvRunners(3, true)

	checkRunners(t, f, false, false, true)
	assert.Equal(t, f.score[f.atBat], 0)
}

// TestAdvRunnersSingleFirstAndThird checks that a single places a runner on first and only advances runner on first.  Result should be bases loaded.
func TestAdvRunnersSingleFirstAndThird(t *testing.T) {
	f := NewField()

	f.b1.on = true
	f.b3.on = true

	f.AdvRunners(1, true)

	checkRunners(t, f, true, true, true)
	assert.Equal(t, f.score[f.atBat], 0)
}

// TestAdvRunnersDoubleFirstAndThird checks that a double places a runner on second and runner on third scores.  Result should be 2nd and 3rd loaded.
func TestAdvRunnersDoubleFirstAndThird(t *testing.T) {
	f := NewField()

	f.b1.on = true
	f.b3.on = true

	f.AdvRunners(2, true)

	checkRunners(t, f, false, true, true)
	assert.Equal(t, f.score[f.atBat], 1)
}

// TestAdvRunnersHomeRunFirstAndThird checks that a home run clears the bases.  Three runs should score.
func TestAdvRunnersHomeRunFirstAndThird(t *testing.T) {
	f := NewField()

	f.b1.on = true
	f.b3.on = true

	f.AdvRunners(4, true)

	checkRunners(t, f, false, false, false)
	assert.Equal(t, f.score[f.atBat], 3)
}
