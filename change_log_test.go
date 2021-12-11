package git

import "testing"

func TestGetChangeLogs(t *testing.T) {
	type args struct {
		currentTag string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChangeLogs("", tt.args.currentTag)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChangeLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetChangeLogs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
