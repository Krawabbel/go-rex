package rex_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/Krawabbel/go-rex/rex"
)

func TestMatch(t *testing.T) {

	patterns := []string{
		"",
		"?",
		"+",
		"*",
		"\\+",
		"\\*",
		"\\?",
		"\\\\",
		"\\(",
		"\\)",
		"()",
		"a",
		"a?",
		"a+",
		"a*",
		"a|",
		"|a",
		"(a)",
		"aa",
		"(a",
		"a)",
		"a||",
		"a|+",
		"a|+?",
		"a|+?*",
		"(a?)(a?)(a?)(a?)a",
	}

	texts := []string{"", "a", "b", "aa", "bb", "aaaaa", "+", "*", "?", "\\", "(", ")", string([]byte{0x00})}

	for _, pattern := range patterns {
		for _, text := range texts {
			msg := []byte(text)
			t.Run(pattern+" ~ "+string(msg), func(t *testing.T) {
				got, err := rex.Match(pattern, msg)
				want, wantErr := regexp.Match(pattern, msg)
				if (err != nil) != (wantErr != nil) {
					t.Errorf("Match() error = %v, wantErr %v", err, wantErr)
					return
				}
				if got != want {
					t.Errorf("Match() = %v, want %v", got, want)
				}
			})
		}
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		args    args
		want    []string
		wantErr bool
	}{
		{args{"|a"}, []string{"", "a"}, false},
		{args{"a|"}, []string{"", "a"}, false},
		{args{"a"}, []string{"a"}, false},
		{args{"a|b"}, []string{"a", "b"}, false},
		{args{"a|b|c"}, []string{"a", "b", "c"}, false},
		{args{"a|(ab)"}, []string{"a", "ab"}, false},
		{args{"a?"}, []string{"", "a"}, false},
		{args{"(ab)c?"}, []string{"ab", "abc"}, false},
		{args{"*a"}, nil, true},
		{args{"(a"}, nil, true},
		{args{"a)"}, nil, true},
		{args{"(((a)(((|b)))))"}, []string{"a", "ab"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.args.pattern, func(t *testing.T) {
			got, err := rex.Generate(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegexpMatch(t *testing.T) {
	type args struct {
		pattern string
		msg     []byte
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"infinite recursion?", args{"(a?)*aaa", []byte("aaa")}, true, false},
		{"empty?", args{"(?)", []byte("")}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := regexp.Match(tt.args.pattern, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ref() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ref() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchSpecial(t *testing.T) {
	type args struct {
		pattern string
		msg     []byte
	}
	tests := []struct {
		args    args
		want    bool
		wantErr bool
	}{
		{args{"\\a", []byte("a")}, false, true},
		{args{"(?)", []byte("")}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.args.pattern+" ~ "+string(tt.args.msg), func(t *testing.T) {
			got, err := rex.Match(tt.args.pattern, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Match() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
