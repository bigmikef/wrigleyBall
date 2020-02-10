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
	b1 *base
	b2 *base
	b3 *base

	balls   int8
	strikes int8

	inning int16
	atBat  Team

	score [2]int
}

// NewField creates a field
func NewField() *Field {
	f := Field{}
	f.b3 = newBase(nil)
	f.b2 = newBase(f.b3)
	f.b1 = newBase(f.b2)
	return &f
}

// Base is a place where a runner can land.
type base struct {
	on bool
	n  *base
}

func (b *base) adv(r bool) int {
	var ret int
	if b.on {
		if b.n != nil {
			ret = b.n.adv(b.on)
		} else {
			ret = 1
		}
	}
	b.on = r
	return ret
}

// newBase creates a base for the field
func newBase(nextBase *base) *base {
	return &base{on: false, n: nextBase}
}

// AdvRunners force advances the runners, adding a runner to the base path if runner is true, incrementing the score of the team that is at bat.
func (f *Field) AdvRunners(nBases uint8, runner bool) error {
	if nBases > 4 || nBases < 1 {
		return fmt.Errorf("nBases must be 1 - 4 inclusive - %d", nBases)
	}

	var score int
	if nBases == 1 {
		score = f.b1.adv(true)
	} else if nBases == 2 {
		score = f.b1.adv(true)
		score += f.b1.adv(false)
	} else if nBases == 3 {
		score = f.b1.adv(true)
		score += f.b1.adv(false)
		score += f.b2.adv(false)
	} else if nBases == 4 {
		score = f.b1.adv(true)
		score += f.b1.adv(false)
		score += f.b2.adv(false)
		score += f.b3.adv(false)
	}

	f.score[f.atBat] += score

	return nil
}
