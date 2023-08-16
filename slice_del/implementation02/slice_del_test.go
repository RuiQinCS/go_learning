package implementation02

import (
	"reflect"
	"testing"
)

func TestDelElem(t *testing.T) {
	type args struct {
		data []int
		idx  int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "wrong index exp 1",
			args: args{
				data: []int{1, 2, 3},
				idx:  -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong index exp 2",
			args: args{
				data: []int{1, 2, 3},
				idx:  3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "right index exp 1",
			args: args{
				data: []int{1, 2, 3},
				idx:  0,
			},
			want:    []int{2, 3},
			wantErr: false,
		},
		{
			name: "right index exp 2",
			args: args{
				data: []int{1, 2, 3},
				idx:  2,
			},
			want:    []int{1, 2},
			wantErr: false,
		},
		{
			name: "right index exp 3",
			args: args{
				data: []int{1, 2, 3},
				idx:  1,
			},
			want:    []int{1, 3},
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

func BenchmarkDelElem(b *testing.B) {
	capacity := 100000
	data := make([]int, capacity)
	for i := 0; i < capacity; i++ {
		data[i] = i
	}
	idx := 2
	for i := 0; i < b.N; i++ {
		_, _ = DelElem(data, idx)
	}
}
