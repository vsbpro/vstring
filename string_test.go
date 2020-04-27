package vstring

import (
	"reflect"
	"testing"
)

func TestSplitByMultipleDelimiters(t *testing.T) {
	type args struct {
		s          string
		delimiters string
	}
	tests := []struct {
		name string
		args args
		want []*Token
	}{
		{
			name: "No delimiter encountered in the string",
			args: args{
				s:          "Some string",
				delimiters: ",:",
			},
			want: nil,
		},
		{
			name: "One delimiter encountered in the string",
			args: args{
				s:          "Some, string",
				delimiters: ",:",
			},
			want: []*Token{
				{
				StartDelimiter: Nothing,
				Value:          "Some",
				EndDelimiter:   ',',
				},
				{
					StartDelimiter: ',',
					Value:          " string",
					EndDelimiter:   Nothing,
				},
			},
		},
		{
			name: "Two delimiter encountered in the string",
			args: args{
				s:          "{ 'Key1' : 'data1', 'Key2' : 'data2',}",
				delimiters: ",:",
			},
			want: []*Token{
				{
					StartDelimiter: Nothing,
					Value:          "{ 'Key1' ",
					EndDelimiter:   ':',
				},
				{
					StartDelimiter: ':',
					Value:          " 'data1'",
					EndDelimiter:   ',',
				},
				{
					StartDelimiter: ',',
					Value:          " 'Key2' ",
					EndDelimiter:   ':',
				},
				{
					StartDelimiter: ':',
					Value:          " 'data2'",
					EndDelimiter:   ',',
				},
				{
					StartDelimiter: ',',
					Value:          "}",
					EndDelimiter:   Nothing,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitByMultipleDelimiters(tt.args.s, tt.args.delimiters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}