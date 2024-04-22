package chunkers_test

import (
	"testing"

	"github.com/nohe427/chunkers"
)

func TestChunker(t *testing.T) {
	text := "Hello World.\nThis is\n a test sentence! Have a good day? Haha. Haha"

	// t.Run("split to three chunks", func(t *testing.T) {
	// 	output := chunkers.Chunk(text, &chunkers.SplitOptions{MinLength: 7, MaxLength: 9})
	// 	if len(output) != 3 {
	// 		t.Errorf("Expected 3 chunks, got %d", len(output))
	// 	}
	// })
	//
	// t.Run("no splits", func(t *testing.T) {
	// 	output := chunkers.Chunk(text, nil)
	// 	if len(output) != 1 {
	// 		t.Errorf("Expected 1 chunk, got %d", len(output))
	// 	}
	// })

	t.Run("should be split to six chunks", func(t *testing.T) {
		output := chunkers.Chunk(text, &chunkers.SplitOptions{Splitter: chunkers.Sentence})
		if len(output) != 6 {
			t.Errorf("Expected 6 chunks, got %d", len(output))
		}
	})
	//
	// t.Run("should be split into 4 chunks", func(t *testing.T) {
	// 	output := chunkers.Chunk(text, &chunkers.SplitOptions{MinLength: 10, Splitter: chunkers.Sentence})
	// 	if len(output) != 4 {
	// 		t.Errorf("Expected 4 chunks, got %d", len(output))
	// 	}
	// })
}
