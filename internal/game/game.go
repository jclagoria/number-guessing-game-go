package game

import (
	"math/rand"
	"time"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

const (
	Min = 1   // smallest possible secret/guess
	Max = 100 // largest possible secret/guess
)

func (d Difficulty) Chances() int {
	switch d {
	case Easy:
		return 10
	case Medium:
		return 5
	case Hard:
		return 3
	default:
		return 0
	}
}

func (d Difficulty) String() string {
	switch d {
	case Easy:
		return "Easy"
	case Medium:
		return "Medium"
	case Hard:
		return "Hard"
	default:
		return "Unknwon"
	}
}

type Result int

const (
	Correct Result = iota
	TooHigh
	TooLow
	Finished
)

type Round struct {
	secret      int
	chances     int
	attempts    int
	won         bool
	wrongStreak int
	lo, hi      int
	start       time.Time
}

func New(secret int, d Difficulty) *Round {
	return &Round{secret: secret, chances: d.Chances(), lo: Min, hi: Max, start: time.Now()}
}

func RandomSecret() int {
	return rand.Intn(Max-Min+1) + Min
}

func (r *Round) Guess(n int) Result {
	if r.Over() {
		return Finished
	}

	r.attempts++
	r.wrongStreak++

	switch {
	case n == r.secret:
		r.won = true
		return Correct
	case n > r.secret:
		if n-1 < r.hi {
			r.hi = n - 1
		}
		return TooHigh
	default:
		if n+1 > r.lo {
			r.lo = n + 1
		}
		return TooLow
	}
}

// Hint returns the narrowed possible range for the secret, derived only from
// guesses made so far, once the player has had 3 consecutive wrong guesses.
func (r *Round) Hint() (lo, hi int, ok bool) {
	if r.wrongStreak < 3 {
		return 0, 0, false
	}

	return r.lo, r.hi, true
}

func (r *Round) Won() bool              { return r.won }
func (r *Round) Over() bool             { return r.won || r.attempts >= r.chances }
func (r *Round) Attempts() int          { return r.attempts }
func (r *Round) Chances() int           { return r.chances }
func (r *Round) Remaining() int         { return r.chances - r.attempts }
func (r *Round) Secret() int            { return r.secret }
func (r *Round) Elapsed() time.Duration { return time.Since(r.start) }
