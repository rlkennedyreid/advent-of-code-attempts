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

func TestBinarySpacePartioner_initialRange(t *testing.T) {
	type fields struct {
		Power int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"1",
			fields{1},
			2,
		},
		{
			"10",
			fields{10},
			1024,
		},
		{
			"0",
			fields{0},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			partitioner := BinarySpacePartitioner{
				Power: tt.fields.Power,
			}
			if got := partitioner.initialRange(); got != tt.want {
				t.Errorf("BinarySpacePartitioner.initialRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySpacePartioner_indexFromTree(t *testing.T) {
	type fields struct {
		Power int
	}
	type args struct {
		tree string
		key  BinaryKey
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"0",
			fields{7},
			args{
				"FBFBBFF",
				BinaryKey{
					lower: "F",
					upper: "B",
				},
			},
			44,
		},
		{
			"0",
			fields{7},
			args{
				"FBFBBFB",
				BinaryKey{
					lower: "F",
					upper: "B",
				},
			},
			45,
		},
		{
			"0",
			fields{7},
			args{
				"FFFFFFF",
				BinaryKey{
					lower: "F",
					upper: "B",
				},
			},
			0,
		},
		{
			"0",
			fields{7},
			args{
				"BBBBBBB",
				BinaryKey{
					lower: "F",
					upper: "B",
				},
			},
			127,
		},
		{
			"0",
			fields{1},
			args{
				"F",
				BinaryKey{
					lower: "F",
					upper: "B",
				},
			},
			0,
		},
		{
			"0",
			fields{1},
			args{
				"B",
				BinaryKey{
					lower: "F",
					upper: "B",
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			partitioner := BinarySpacePartitioner{
				Power: tt.fields.Power,
			}
			if got := partitioner.indexFromTree(tt.args.tree, tt.args.key); got != tt.want {
				t.Errorf("BinarySpacePartitioner.indexFromTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySpacePartitioner_getPartitionIndexesFrom(t *testing.T) {
	type fields struct {
		Power int
	}
	type args struct {
		input []string
		key   BinaryKey
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			"0",
			fields{7},
			args{
				[]string{
					"FBFBBFF",
					"BBBBBBB",
					"FFFFFFF",
				},
				BinaryKey{
					lower: "F",
					upper: "B",
				},
			},
			[]int{
				44,
				127,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			partitioner := BinarySpacePartitioner{
				Power: tt.fields.Power,
			}
			if got := partitioner.getPartitionIndexesFrom(tt.args.input, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinarySpacePartitioner.getPartitionIndexesFrom() = %v, want %v", got, tt.want)
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
