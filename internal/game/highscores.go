package game

import (
	"encoding/json"
	"os"
)

// HighScores records the fewest attempts per difficulty level, indexed by
// Difficulty (Easy=0, Medium=1, Hard=2). Zero means no score stored yet.
type HighScores [3]int

// Best returns the fewest attempts for a difficulty, or 0 if unset.
// An out-of-range difficulty yields 0 rather than panicking.
func (h *HighScores) Best(d Difficulty) int {
	if d < Easy || d > Hard {
		return 0
	}

	return h[d]
}

// Record stores attempts as the best for d if it beats the current score.
// It reports whether the stored score was beaten (true on a new best).
// Non-positive stored values count as unset, so a hand-edited file can never
// lock a difficulty into an unbeatable score.
func (h *HighScores) Record(d Difficulty, attempts int) bool {
	if d < Easy || d > Hard {
		return false
	}

	if h[d] <= 0 || attempts < h[d] {
		h[d] = attempts
		return true
	}

	return false
}

func (h *HighScores) Save(path string) error {
	data, err := json.MarshalIndent(h, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

// LoadHighScores reads scores from path. A missing file yields an empty score
// set without error, so a first run starts clean. Any other read or decode
// failure (corrupt file, permissions) returns the error alongside an empty
// set, so the caller can warn and keep the game running.
func LoadHighScores(path string) (*HighScores, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &HighScores{}, nil
		}
		return &HighScores{}, err
	}

	h := &HighScores{}
	if err := json.Unmarshal(data, h); err != nil {
		return &HighScores{}, err
	}

	return h, nil
}
