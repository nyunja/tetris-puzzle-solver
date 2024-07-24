package lib

import "testing"

func Test_hasSuffix(t *testing.T) {
	type args struct {
		s      string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "hasSuffix", args: args{"hello.txt", ".txt"}, want: true,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSuffix(tt.args.s, tt.args.suffix); got != tt.want {
				t.Errorf("hasSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidFile(t *testing.T) {
	tests := []struct {
		name string
		fileName string
		want bool
	}{
		{name: "hasSuffix", fileName: "hello.txt", want: true,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidFile(tt.fileName); got != tt.want {
				t.Errorf("hasSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

