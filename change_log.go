package git

import (
	"errors"
	"fmt"
	"regexp"
)

func GetChangeLogs(prev, currentTag string) (string, error) {
	if currentTag == "" {
		return "", errors.New("current tag is null")
	}
	if prev == "" {
		prev, err := GetPreviousTag(currentTag)
		if err != nil {
			return "", err
		}
		if prev == "" {
			// get first commit
			result, err := Clean(Run("rev-list", "--max-parents=0", "HEAD"))
			if err != nil {
				return "", err
			}
			prev = result
		}
	}

	return Log(prev, currentTag)
}

var validSHA1 = regexp.MustCompile(`^[a-fA-F0-9]{40}$`)

func Log(prev, current string) (string, error) {
	args := []string{"log", "--pretty=oneline", "--abbrev-commit", "--no-decorate", "--no-color"}
	if validSHA1.MatchString(prev) {
		args = append(args, prev, current)
	} else {
		args = append(args, fmt.Sprintf("tags/%s..tags/%s", prev, current))
	}
	return Run(args...)
}
