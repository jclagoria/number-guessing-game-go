package game

import "testing"

func TestDifficultyChances(t *testing.T) {
	cases := []struct {
		d    Difficulty
		want int
	}{
		{Easy, 10},
		{Medium, 5},
		{Hard, 3},
	}
	for _, c := range cases {
		if got := c.d.Chances(); got != c.want {
			t.Errorf("%v Chances() = %d, want %d", c.d, got, c.want)
		}
	}
}

func TestDifficultyString(t *testing.T) {
	cases := []struct {
		d    Difficulty
		want string
	}{
		{Easy, "Easy"},
		{Medium, "Medium"},
		{Hard, "Hard"},
	}
	for _, c := range cases {
		if got := c.d.String(); got != c.want {
			t.Errorf("%v String() = %q, want %q", c.d, got, c.want)
		}
	}
}

func TestGuessCorrect(t *testing.T) {
	r := New(50, Easy)
	if got := r.Guess(50); got != Correct {
		t.Errorf("Guess(50) = %v, want Correct", got)
	}
	if !r.Won() {
		t.Error("round should be won after a correct guess")
	}
	if !r.Over() {
		t.Error("round should be over after a correct guess")
	}
	if r.Attempts() != 1 {
		t.Errorf("Attempts() = %d, want 1", r.Attempts())
	}
}

func TestGuessFeedback(t *testing.T) {
	r := New(50, Easy)
	if got := r.Guess(75); got != TooHigh {
		t.Errorf("Guess(75) = %v, want TooHigh", got)
	}
	if got := r.Guess(25); got != TooLow {
		t.Errorf("Guess(25) = %v, want TooLow", got)
	}
	if r.Won() {
		t.Error("round should not be won after wrong guesses")
	}
	if r.Attempts() != 2 {
		t.Errorf("Attempts() = %d, want 2", r.Attempts())
	}
}

func TestChanceExhaustion(t *testing.T) {
	r := New(50, Hard) // 3 chances
	r.Guess(1)
	r.Guess(1)
	if r.Over() {
		t.Error("round should not be over after 2 of 3 chances")
	}
	r.Guess(1)
	if !r.Over() {
		t.Error("round should be over after exhausting all chances")
	}
	if r.Won() {
		t.Error("round should not be won on chance exhaustion")
	}
	if r.Attempts() != 3 {
		t.Errorf("Attempts() = %d, want 3", r.Attempts())
	}
}

func TestChances(t *testing.T) {
	r := New(50, Medium)
	if got := r.Chances(); got != 5 {
		t.Errorf("Chances() = %d, want 5", got)
	}
	if got := r.Remaining(); got != 5 {
		t.Errorf("Remaining() = %d, want 5", got)
	}
	r.Guess(25)
	if got := r.Remaining(); got != 4 {
		t.Errorf("Remaining() after one guess = %d, want 4", got)
	}
}

func TestGuessAfterOverIsFinished(t *testing.T) {
	r := New(50, Hard)
	r.Guess(25)
	r.Guess(25)
	r.Guess(25)
	if !r.Over() {
		t.Fatal("round should be over after 3 chances")
	}
	if got := r.Guess(50); got != Finished {
		t.Errorf("Guess() after over = %v, want Finished", got)
	}
	if r.Attempts() != 3 {
		t.Errorf("Attempts() = %d, want 3 (no extra attempts after over)", r.Attempts())
	}
}

func TestRandomSecretInRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := RandomSecret(); n < 1 || n > 100 {
			t.Fatalf("RandomSecret() = %d, want in [1,100]", n)
		}
	}
}
