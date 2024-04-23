package chunkers_test

import (
	"strings"
	"testing"

	"github.com/nohe427/chunkers"
)

func TestChunker(t *testing.T) {
	text := "Hello World.\nThis is\n a test sentence! Have a good day? Haha. Haha"

	t.Run("split to three chunks", func(t *testing.T) {
		output := chunkers.Chunk(text, &chunkers.SplitOptions{MinLength: 7, MaxLength: 9})
		if len(output) != 3 {
			t.Errorf("Expected 3 chunks, got %d", len(output))
		}
	})

	t.Run("no splits", func(t *testing.T) {
		output := chunkers.Chunk(text, nil)
		if len(output) != 1 {
			t.Errorf("Expected 1 chunk, got %d", len(output))
		}
	})

	t.Run("should be split to six chunks", func(t *testing.T) {
		output := chunkers.Chunk(text, &chunkers.SplitOptions{Splitter: chunkers.Sentence})
		if len(output) != 6 {
			t.Errorf("Expected 6 chunks, got %d", len(output))
		}
	})

	t.Run("should be split into 4 chunks", func(t *testing.T) {
		output := chunkers.Chunk(text, &chunkers.SplitOptions{MinLength: 10, Splitter: chunkers.Sentence})
		if len(output) != 4 {
			t.Errorf("Expected 4 chunks, got %d", len(output))
		}
	})

	t.Run("should be overlapped 3 characters minimum", func(t *testing.T) {
		chunks := chunkers.Chunk(text, &chunkers.SplitOptions{Overlap: 3, Splitter: chunkers.Sentence})
		firstWord := strings.Split(chunks[0], " ")
		secondWord := strings.Split(chunks[1], " ")
		thirdWord := strings.Split(chunks[2], " ")
		if len(chunks) != 7 {
			t.Errorf("Expected 7 chunks, got %d", len(chunks))
		}
		if firstWord[len(firstWord)-1] != secondWord[1] {
			t.Errorf("Expected %s, got %s", secondWord[1], firstWord[len(firstWord)-1])
		}
		// comparing the second index since the first starts with a space
		if secondWord[len(secondWord)-1] != thirdWord[2] {
			t.Errorf("Expected `%s`, got `%s`", thirdWord[2], secondWord[len(secondWord)-1])
		}
	})
}
