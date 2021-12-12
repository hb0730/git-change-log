package git

import (
	"github.com/apex/log"
	"testing"
)

func TestGetChangeLogs(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test", args: "v1.0.0.03-beta", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChangeLogs("", tt.args, log.DebugLevel)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChangeLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
