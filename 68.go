package leetcode

import (
	"strings"
)

// 2019/11/06 0:47 by fzls

func fullJustify(words []string, maxWidth int) []string {
	lineOfWords := getLineOfWords(words, maxWidth)
	return formatLines(lineOfWords, words, maxWidth)
}

type LineOfWords struct {
	Begin                int
	End                  int
	LenOfWordsWithASpace int
}

func getLineOfWords(words []string, maxWidth int) []*LineOfWords {
	// find each line's start word and end word' idx
	var lineOfWords []*LineOfWords
	start := 0
	for start < len(words) {
		end := start
		currentLen := len(words[end])
		for end+1 < len(words) && currentLen+len(words[end+1])+1 <= maxWidth {
			currentLen += len(words[end+1]) + 1
			end++
		}

		lineOfWords = append(lineOfWords, &LineOfWords{
			Begin:                start,
			End:                  end,
			LenOfWordsWithASpace: currentLen,
		})

		start = end + 1
	}

	return lineOfWords
}

func formatLines(lineOfWords []*LineOfWords, words []string, maxWidth int) []string {
	// format as requirements
	var res []string
	// print line except last line
	for i := 0; i <= len(lineOfWords)-2; i++ {
		line := lineOfWords[i]
		remainingSpace := maxWidth - line.LenOfWordsWithASpace
		gap := line.End - line.Begin

		if gap == 0 {
			// single word line
			var sb strings.Builder
			sb.WriteString(words[line.Begin])
			nLeft := maxWidth - sb.Len()
			for j := 0; j < nLeft; j++ {
				sb.WriteByte(' ')
			}
			res = append(res, sb.String())
			continue
		}

		// space for each gap
		nEvenSpace := remainingSpace/gap + 1
		// space that cannot allocate for each gap
		nUnevenSpace := remainingSpace % gap

		var sb strings.Builder
		sb.WriteString(words[line.Begin])
		for j := line.Begin + 1; j <= line.End; j++ {
			for k := 0; k < nEvenSpace; k++ {
				sb.WriteByte(' ')
			}
			if j-line.Begin <= nUnevenSpace {
				sb.WriteByte(' ')
			}
			sb.WriteString(words[j])
		}
		res = append(res, sb.String())
	}

	// print last line
	line := lineOfWords[len(lineOfWords)-1]
	var sb strings.Builder
	sb.WriteString(words[line.Begin])
	for j := line.Begin + 1; j <= line.End; j++ {
		sb.WriteByte(' ')
		sb.WriteString(words[j])
	}
	nLeft := maxWidth - sb.Len()
	for j := 0; j < nLeft; j++ {
		sb.WriteByte(' ')
	}
	res = append(res, sb.String())
	return res
}
