package implementation04

import (
	"reflect"
	"testing"
)

func TestDelElem(t *testing.T) {
	type args[T any] struct {
		data []T
		idx  int
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    []T
		wantErr bool
	}
	tests := []testCase[string]{
		{
			name: "T type is string",
			args: args[string]{
				data: []string{"1", "2", "3"},
				idx:  0,
			},
			want:    []string{"2", "3"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DelElem(tt.args.data, tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DelElem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelElem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shrinkSlice(t *testing.T) {
	type args[T any] struct {
		data   []T
		oldCap int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	data := make([]int, 3, 100)
	data[0] = 1
	data[1] = 2
	data[2] = 3

	want := make([]int, 3, 50)
	want[0] = 1
	want[1] = 2
	want[2] = 3

	tests := []testCase[int]{
		{
			name: "shrink 0.5",
			args: args[int]{
				data:   data,
				oldCap: cap(data),
			},
			want: want,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shrinkSlice(tt.args.data, tt.args.oldCap)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shrinkSlice() = %v, want %v", got, tt.want)
			}

			if cap(got) != cap(tt.want) {
				t.Errorf("wrong cap : %v, want %v", cap(got), cap(tt.want))
			}
		})
	}
}

func Test_shrink(t *testing.T) {
	type args struct {
		newLen int
		oldCap int
	}
	tests := []struct {
		name       string
		args       args
		wantNewCap int
	}{
		{
			name: "no change",
			args: args{
				newLen: 100,
				oldCap: 200,
			},
			wantNewCap: 200,
		},
		{
			name: "shrink 0.5",
			args: args{
				newLen: 20,
				oldCap: 200,
			},
			wantNewCap: 100,
		},
		{
			name: "shrink between[0.5,1]",
			args: args{
				newLen: 170,
				oldCap: 300,
			},
			wantNewCap: 176,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewCap := shrink(tt.args.newLen, tt.args.oldCap); gotNewCap != tt.wantNewCap {
				t.Errorf("shrink() = %v, want %v", gotNewCap, tt.wantNewCap)
			}
		})
	}
}
