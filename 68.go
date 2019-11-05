package leetcode

import (
	"strings"
)

// 2019/11/06 0:47 by fzls
type LineOfWords struct {
	Begin                int
	End                  int
	LenOfWordsWithASpace int
}

func fullJustify(words []string, maxWidth int) []string {
	n := len(words)

	// find each line's start word and end word' idx
	var lineOfWords []LineOfWords
	start := 0
	for start < n {
		end := start
		currentLen := len(words[end])
		for end+1 < n {
			// len after add next word and a space
			if currentLen+len(words[end+1])+1 > maxWidth {
				break
			}

			currentLen += len(words[end+1]) + 1
			end++
		}

		lineOfWords = append(lineOfWords, LineOfWords{
			Begin:                start,
			End:                  end,
			LenOfWordsWithASpace: currentLen,
		})
		start = end + 1
	}
	//fmt.Println(lineOfWords)

	// format as requirements
	nLine := len(lineOfWords)
	var res []string
	// print line except last line
	for i := 0; i <= nLine-2; i++ {
		line := lineOfWords[i]
		remainingSpace := maxWidth - line.LenOfWordsWithASpace
		gap := line.End - line.Begin

		if gap == 0 {
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
	line := lineOfWords[nLine-1]
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
