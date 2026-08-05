package game

import "math/rand"

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
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

type Result int

const (
	Correct Result = iota
	TooHigh
	TooLow
	Finished
)

type Round struct {
	secret   int
	chances  int
	attempts int
	won      bool
}

func New(secret int, d Difficulty) *Round {
	return &Round{secret: secret, chances: d.Chances()}
}

func RandomSecret() int {
	return rand.Intn(100) + 1
}

func (r *Round) Guess(n int) Result {
	if r.Over() {
		return Finished
	}

	r.attempts++

	switch {
	case n == r.secret:
		r.won = true
		return Correct
	case n > r.secret:
		return TooHigh
	default:
		return TooLow
	}
}

func (r *Round) Won() bool      { return r.won }
func (r *Round) Over() bool     { return r.won || r.attempts >= r.chances }
func (r *Round) Attempts() int  { return r.attempts }
func (r *Round) Chances() int   { return r.chances }
func (r *Round) Remaining() int { return r.chances - r.attempts }
func (r *Round) Secret() int    { return r.secret }
