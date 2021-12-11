package git

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/apex/log"
	"os/exec"
	"strings"
)

func Run(args ...string) (string, error) {
	// TODO: use exex.CommandContext here and refactor.
	extraArgs := []string{
		"-c", "log.showSignature=false",
	}
	args = append(extraArgs, args...)
	/* #nosec */
	cmd := exec.Command("git", args...)

	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	log.WithField("args", args).Debug("running git")
	err := cmd.Run()

	log.WithField("stdout", stdout.String()).
		WithField("stderr", stderr.String()).
		Debug("git result")

	if err != nil {
		return "", errors.New(stderr.String())
	}
	return stdout.String(), nil
}

// Clean the output.
func Clean(output string, err error) (string, error) {
	output = strings.ReplaceAll(strings.Split(output, "\n")[0], "'", "")
	if err != nil {
		err = errors.New(strings.TrimSuffix(err.Error(), "\n"))
	}
	return output, err
}
func GetTag() (string, error) {
	var tag string
	var err error
	for _, fn := range []func() (string, error){
		func() (string, error) {
			return Clean(Run("tag", "--points-at", "HEAD", "--sort", "-version:refname"))
		},
		func() (string, error) {
			return Clean(Run("describe", "--tags", "--abbrev=0"))
		},
	} {
		tag, err = fn()
		if tag != "" || err != nil {
			return tag, err
		}
	}

	return tag, err
}

// GetPreviousTag get previousTag
func GetPreviousTag(current string) (string, error) {
	if current == "" {
		return "", errors.New("current tags is null")
	}
	return Clean(Run("describe", "--tags", "--abbrev=0", fmt.Sprintf("tags/%s^", current)))
}
