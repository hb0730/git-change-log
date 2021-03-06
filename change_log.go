package git

import (
	"errors"
	"fmt"
	"github.com/apex/log"
	"regexp"
)

// GetChangeLogs get git previous tag to  current tag  change logs
func GetChangeLogs(prev, currentTag string, level log.Level) (string, error) {
	log.SetLevel(level)
	if currentTag == "" {
		return "", errors.New("current tag is null")
	}
	if prev == "" {
		previous, _ := GetPreviousTag(currentTag)
		if previous == "" {
			// get first commit
			result, err := Clean(Run("rev-list", "--max-parents=0", "HEAD"))
			if err != nil {
				return "", err
			}
			prev = result
		} else {
			prev = previous
		}
	}

	return Log(prev, currentTag)
}

var validSHA1 = regexp.MustCompile(`^[a-fA-F0-9]{40}$`)

// Log get git previous tag to  current tag  change logs
func Log(prev, current string) (string, error) {
	args := []string{"log", "--pretty=oneline", "--abbrev-commit", "--no-decorate", "--no-color"}
	if validSHA1.MatchString(prev) {
		args = append(args, prev, current)
	} else {
		args = append(args, fmt.Sprintf("tags/%s..tags/%s", prev, current))
	}
	return Run(args...)
}
