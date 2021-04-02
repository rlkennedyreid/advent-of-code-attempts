package main

import (
	"reflect"
	"testing"
)

func Test_substringsAt(t *testing.T) {
	type args struct {
		fullString string
		pivot      int
	}
	tests := []struct {
		name  string
		args  args
		want1 string
		want2 string
	}{
		{
			"1",
			args{
				"0123456789",
				7,
			},
			"0123456",
			"789",
		},
		{
			"2",
			args{
				"0123456789",
				0,
			},
			"",
			"0123456789",
		},
		{
			"3",
			args{
				"0123456789",
				10,
			},
			"0123456789",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := substringsAt(tt.args.fullString, tt.args.pivot)
			if got != tt.want1 {
				t.Errorf("substringsAt() got = %v, want1 %v", got, tt.want1)
			}
			if got1 != tt.want2 {
				t.Errorf("substringsAt() got1 = %v, want1 %v", got1, tt.want2)
			}
		})
	}
}

func Test_splitSliceStrings(t *testing.T) {
	type args struct {
		input []string
		pivot int
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			"1",
			args{
				[]string{
					"FBFBBBBLRR",
					"BBFFBFBLRL",
					"FBBFBFFRLR",
					"FBBFFFBLRR",
					"BFBFBFFRRL",
				},
				7,
			},
			[]string{
				"FBFBBBB",
				"BBFFBFB",
				"FBBFBFF",
				"FBBFFFB",
				"BFBFBFF",
			},
			[]string{
				"LRR",
				"LRL",
				"RLR",
				"LRR",
				"RRL",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitSliceStrings(tt.args.input, tt.args.pivot)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitSliceStrings() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("splitSliceStrings() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func Test_getSeatIDsFrom(t *testing.T) {
	type args struct {
		rows []int
		cols []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"0",
			args{
				[]int{
					44,
					70,
					14,
					102,
				},
				[]int{
					5,
					7,
					7,
					4,
				},
			},
			[]int{
				357,
				567,
				119,
				820,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSeatIDsFrom(tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSeatIDsFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharSequenceToBinarySequence(t *testing.T) {
	type args struct {
		input string
		key   BinaryKey
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{
				"FBFBBFF",
				BinaryKey{
					lower: "F",
					upper: "B",
				},
			},
			want: "0101100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CharSequenceToBinarySequence(tt.args.input, tt.args.key); got != tt.want {
				t.Errorf("CharSequenceToBinarySequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryStringSliceToInts(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0",
			args: args{
				[]string{
					"0101100",
					"0000000",
				},
			},
			want: []int{
				44,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryStringSliceToInts(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryStringSliceToInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapToBinarySequenceSlice(t *testing.T) {
	type args struct {
		input []string
		key   BinaryKey
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "0",
			args: args{
				[]string{
					"FBFBBFF",
				},
				BinaryKey{
					lower: "F",
					upper: "B",
				},
			},
			want: []string{
				"0101100",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToBinarySequenceSlice(tt.args.input, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToBinarySequenceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapSequencesToDecimals(t *testing.T) {
	type args struct {
		input []string
		key   BinaryKey
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0",
			args: args{
				[]string{
					"FBFBBFF",
				},
				BinaryKey{
					lower: "F",
					upper: "B",
				},
			},
			want: []int{
				44,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapSequencesToDecimals(tt.args.input, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapSequencesToDecimals() = %v, want %v", got, tt.want)
			}
		})
	}
}
