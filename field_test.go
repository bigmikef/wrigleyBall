package wrigleyball

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkRunners(t *testing.T, f Field, first, second, third bool) {
	assert.Equal(t, f.b1, first, fmt.Sprintf("First base incorrect - %t, %t, %t", f.b1, f.b2, f.b3))
	assert.Equal(t, f.b2, second, fmt.Sprintf("Second base incorrect - %t, %t, %t", f.b1, f.b2, f.b3))
	assert.Equal(t, f.b3, third, fmt.Sprintf("Third base incorrect - %t, %t, %t", f.b1, f.b2, f.b3))
}

// TestAdvRunnersSingleNoManOn checks that a single places a runner on first.
func TestAdvRunnersSingleNoManOn(t *testing.T) {
	var f Field

	f.advRunners(1, true)

	checkRunners(t, f, true, false, false)
	assert.Equal(t, f.score[f.atBat], 0)
}

// TestAdvRunnersDoubleNoManOn checks that a single places a runner on first.
func TestAdvRunnersDoubleNoManOn(t *testing.T) {
	var f Field

	f.advRunners(2, true)

	checkRunners(t, f, false, true, false)
	assert.Equal(t, f.score[f.atBat], 0)
}

// TestAdvRunnersTripleNoManOn checks that a single places a runner on first.
func TestAdvRunnersTripleNoManOn(t *testing.T) {
	var f Field

	f.advRunners(3, true)

	checkRunners(t, f, false, false, true)
	assert.Equal(t, f.score[f.atBat], 0)
}

// Does not work with current architecture.
// // TestAdvRunnersDoubleFirstAndThird checks that a single places a runner on first.
// func TestAdvRunnersDoubleFirstAndThird(t *testing.T) {
// 	var f Field

// 	f.advRunners(3, true)

// 	checkRunners(t, f, false, false, true)
// 	assert.Equal(t, f.score[f.atBat], 0)
// }
