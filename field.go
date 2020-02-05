package wrigleyball

import "fmt"

// Team is type def to unsigned int 8
type Team int8

const (
	home Team = iota
	away
)

// Field is the game state
type Field struct {
	b1 bool
	b2 bool
	b3 bool

	balls   int8
	strikes int8

	inning int16
	atBat  Team

	score [2]int
}

func (f *Field) advRunners(nBases uint8, runner bool) error {
	if nBases > 4 || nBases < 1 {
		return fmt.Errorf("nBases must be 1 - 4 inclusive - %d", nBases)
	}

	var ret = 0
	for n := 1; n <= int(nBases); n++ {
		if f.b3 {
			ret++
		}
		f.b3 = f.b2
		f.b2 = f.b1
		if n == 1 {
			f.b1 = runner
		} else {
			f.b1 = false
		}
	}

	return nil
}
