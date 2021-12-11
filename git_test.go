package git

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "run git branch -l", args: []string{"branch", "-l"}, wantErr: false},
		{name: "get git init commit hash", args: []string{"rev-list", "--max-parents=0", "HEAD"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Run(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("print: %s \n", got)
		})
	}
}
