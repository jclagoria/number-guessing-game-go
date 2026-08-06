package game

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRecordFirstScoreIsNewBest(t *testing.T) {
	h := &HighScores{}
	if !h.Record(Easy, 7) {
		t.Fatal("first score should be a new best")
	}
	if got := h.Best(Easy); got != 7 {
		t.Fatalf("Best(Easy) = %d, want 7", got)
	}
}

func TestRecordBetterScoreReplaces(t *testing.T) {
	h := &HighScores{}
	h.Record(Easy, 7)
	if !h.Record(Easy, 4) {
		t.Fatal("fewer attempts should be a new best")
	}
	if got := h.Best(Easy); got != 4 {
		t.Fatalf("Best(Easy) = %d, want 4", got)
	}
}

func TestRecordWorseScoreKeepsBest(t *testing.T) {
	h := &HighScores{}
	h.Record(Easy, 4)
	if h.Record(Easy, 9) {
		t.Fatal("more attempts should not replace the best")
	}
	if got := h.Best(Easy); got != 4 {
		t.Fatalf("Best(Easy) = %d, want 4", got)
	}
}

func TestRecordKeepsDifficultiesIndependent(t *testing.T) {
	h := &HighScores{}
	h.Record(Easy, 7)
	h.Record(Hard, 2)
	if got := h.Best(Easy); got != 7 {
		t.Fatalf("Best(Easy) = %d, want 7", got)
	}
	if got := h.Best(Hard); got != 2 {
		t.Fatalf("Best(Hard) = %d, want 2", got)
	}
	if got := h.Best(Medium); got != 0 {
		t.Fatalf("Best(Medium) = %d, want 0 (unset)", got)
	}
}

func TestSaveAndLoadRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "highscores.json")
	h := &HighScores{}
	h.Record(Easy, 5)
	h.Record(Medium, 3)
	if err := h.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}
	got, err := LoadHighScores(path)
	if err != nil {
		t.Fatalf("LoadHighScores: %v", err)
	}
	if got.Best(Easy) != 5 || got.Best(Medium) != 3 || got.Best(Hard) != 0 {
		t.Fatalf("loaded scores = %v, want [5 3 0]", *got)
	}
}

func TestLoadMissingFileIsEmpty(t *testing.T) {
	h, err := LoadHighScores(filepath.Join(t.TempDir(), "nope.json"))
	if err != nil {
		t.Fatalf("LoadHighScores(missing): %v", err)
	}
	if h.Best(Easy) != 0 || h.Best(Medium) != 0 || h.Best(Hard) != 0 {
		t.Fatalf("missing file should load empty, got %v", *h)
	}
}

func TestLoadCorruptFileReportsError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.json")
	if err := os.WriteFile(path, []byte("{not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	h, err := LoadHighScores(path)
	if err == nil {
		t.Fatal("corrupt file should report an error so the caller can warn")
	}
	if h.Best(Easy) != 0 || h.Best(Medium) != 0 || h.Best(Hard) != 0 {
		t.Fatalf("corrupt file should still yield empty scores, got %v", *h)
	}
}

func TestRecordTreatsNonPositiveStoredAsUnset(t *testing.T) {
	// A hand-edited file could store a negative or zero value. It must not
	// lock the difficulty into an unbeatable score.
	for _, stored := range []int{-5, 0} {
		h := &HighScores{}
		h[Easy] = stored
		if !h.Record(Easy, 3) {
			t.Fatalf("stored %d should be treated as unset and beaten by 3", stored)
		}
		if got := h.Best(Easy); got != 3 {
			t.Fatalf("Best(Easy) = %d, want 3", got)
		}
	}
}

func TestRecordOutOfRangeDifficultyIgnored(t *testing.T) {
	h := &HighScores{}
	h.Record(Easy, 4)
	if h.Record(Difficulty(7), 1) {
		t.Fatal("out-of-range difficulty should not record a score")
	}
	if got := h.Best(Easy); got != 4 {
		t.Fatalf("Best(Easy) = %d, want 4", got)
	}
	if got := h.Best(Difficulty(7)); got != 0 {
		t.Fatalf("Best(out-of-range) = %d, want 0", got)
	}
}
