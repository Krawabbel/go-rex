package rex

import (
	"reflect"
	"testing"
)

func Test_generate_failse(t *testing.T) {
	type args struct {
		entry     *state
		max_depth int
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 bool
	}{
		{"nil entry", args{entry: nil, max_depth: 0}, nil, false},
		{"max depth reached", args{entry: new(state), max_depth: -1}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := generate(tt.args.entry, tt.args.max_depth)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("generate() gotValid = %v, want %v", got1, tt.want1)
			}
		})
	}
}
