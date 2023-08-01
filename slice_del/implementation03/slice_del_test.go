package implementation03

import (
	"reflect"
	"testing"
)

func TestDelElemTypeString(t *testing.T) {
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
	tests := []testCase[string /* TODO: Insert concrete types here */]{
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

func TestDelElemTypeFloat(t *testing.T) {
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
	tests := []testCase[float64 /* TODO: Insert concrete types here */]{
		{
			name: "T type is float64",
			args: args[float64]{
				data: []float64{1.23, 2.345, 3.4567},
				idx:  0,
			},
			want:    []float64{2.345, 3.4567},
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
