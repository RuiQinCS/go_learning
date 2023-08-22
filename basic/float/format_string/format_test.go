package format_string

import (
	"testing"
)

func TestGetFormatString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "21.34"},
			want: "%.2f",
		},
		{
			args: args{str: "21"},
			want: "%.0f",
		},
		{
			args: args{str: "21.34232"},
			want: "%.5f",
		},
		{
			args: args{str: "-"},
			want: "%v",
		},
		{
			args: args{str: "empty"},
			want: "%v",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFormatString(tt.args.str); got != tt.want {
				t.Errorf("GetFormatString() = %v, want %v", got, tt.want)
			}
		})
	}
}
