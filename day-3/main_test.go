package main

import "testing"

func Test_removeNumberFromStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "input",
			args: args{
				s: "mul(3,518)",
			},
			want:  3,
			want1: 518,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeNumberFromStr(tt.args.s)
			if got != tt.want {
				t.Errorf("removeNumberFromStr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("removeNumberFromStr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
