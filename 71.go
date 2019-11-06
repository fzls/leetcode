package leetcode

import (
	"strings"
)

// 2019/11/06 21:27 by fzls
func simplifyPath(path string) string {
	var absoluteDirs []string

	// get canonical dir list
	canonicalDirs := strings.Split(path, "/")
	// turn it into absolute dir list
	for _, dir := range canonicalDirs {
		if len(dir) == 0 || dir == "." {
			continue
		}

		if dir == ".." {
			if len(absoluteDirs) != 0 {
				absoluteDirs = absoluteDirs[:len(absoluteDirs)-1]
			}
			continue
		}

		absoluteDirs = append(absoluteDirs, dir)
	}

	// root
	absPath := "/"
	if len(absoluteDirs) != 0 {
		// format absolute path
		var sb strings.Builder
		for _, dir := range absoluteDirs {
			sb.WriteByte('/')
			sb.WriteString(dir)
		}
		absPath = sb.String()
	}

	return absPath
}
