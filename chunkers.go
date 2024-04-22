package chunkers

import (
	"fmt"
	"regexp"
	"strings"
)

type Splitter int

const (
	Sentence  Splitter = 1
	Paragraph Splitter = 0
)

type SplitOptions struct {
	MinLength int
	MaxLength int
	Overlap   int
	Delimiter string
	Splitter  Splitter
}

var defaultSplitOptions = SplitOptions{
	MinLength: 0,
	MaxLength: 1000,
	Overlap:   0,
	Delimiter: "",
	Splitter:  Paragraph,
}

func splitChunk(currChunks []string, maxLength int, overlap int) (string, string, string) {
	chunkString := strings.Join(currChunks[:], " ")
	fmt.Println(chunkString)
	if len(chunkString) < maxLength {
		maxLength = len(chunkString)
	}
	fmt.Printf("Max Length: %v\n", maxLength)
	fmt.Printf("Chunk String: %v\n", chunkString)
	subChunk := chunkString[0:maxLength]
	restChunk := chunkString[maxLength:]

	blankPosition := strings.Index(restChunk, " ")
	if blankPosition == -1 {
		blankPosition = strings.Index(restChunk, "\n")
	}

	fmt.Printf("Sub String: %v\n", subChunk)

	if blankPosition != -1 {
		fmt.Println("Blank Position: ", blankPosition)
		subChunk = subChunk + restChunk[0:blankPosition]
		restChunk = restChunk[blankPosition+1:]
	}

	fmt.Printf("Sub String: %v\n", subChunk)

	overlapText := ""
	if overlap > 0 {
		blankPosition = strings.LastIndex(subChunk[0:len(subChunk)-overlap], " ")
		if blankPosition == -1 {
			blankPosition = strings.LastIndex(subChunk[0:len(subChunk)-overlap], "\n")
		}
		if blankPosition != -1 {
			overlapText = subChunk[blankPosition:]
		}
	}

	return subChunk, restChunk, overlapText
}

func Chunk(text string, opts *SplitOptions) []string {
	if opts == nil {
		opts = &defaultSplitOptions
	}
	if opts.MinLength == 0 {
		opts.MinLength = defaultSplitOptions.MinLength
	}
	if opts.MaxLength == 0 {
		opts.MaxLength = defaultSplitOptions.MaxLength
	}
	if opts.Overlap == 0 {
		opts.Overlap = defaultSplitOptions.Overlap
	}
	if opts.Delimiter == "" {
		opts.Delimiter = defaultSplitOptions.Delimiter
	}

	var regex = opts.Delimiter

	if regex == "" {
		splitStyle := opts.Splitter
		switch splitStyle {
		case Sentence:
			regex = "([.!?\\n])\\s*"
		case Paragraph:
			regex = "\\n{2,}"
		}
	}

	re := regexp.MustCompile(regex)
	baseChunk := re.Split(text, -1)

	fmt.Printf("Base Chunk: %s\n", strings.Join(baseChunk, "|||"))

	chunks := make([]string, 0)

	currChunks := make([]string, 0)
	currChunkLength := 0
	for i := 0; i < len(baseChunk); i = i + 2 {

		subChunk := baseChunk[i]
		if i+1 < len(baseChunk)-1 {
			subChunk = subChunk + " " + baseChunk[i+1]
		}

		currChunks = append(currChunks, subChunk)
		currChunkLength += len(subChunk)

		if currChunkLength >= opts.MinLength {
			fmt.Println(currChunks)
			fmt.Println(i)
			subChunk, restChunk, overlapText := splitChunk(
				currChunks, opts.MaxLength, opts.Overlap,
			)

			fmt.Printf("Returned : %s, %s, %s\n", subChunk, restChunk, overlapText)

			chunks = append(chunks, subChunk)

			currChunks = make([]string, 0)
			currChunkLength = len(overlapText) + len(restChunk)

			if overlapText != "" && i < len(baseChunk) {
				currChunks = append(currChunks, overlapText)
			}
			if restChunk != "" {
				currChunks = append(currChunks, restChunk)
			}
		}

		if len(currChunks) > 0 {
			subChunk, restChunk, _ := splitChunk(currChunks, opts.MaxLength, opts.Overlap)

			if subChunk != "" {
				chunks = append(chunks, subChunk)
			}
			if restChunk != "" {
				chunks = append(chunks, restChunk)
			}
		}

	}
	return chunks
}
