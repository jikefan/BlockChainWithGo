package main

import (
	"goblockchain/utils"
	"reflect"
	"testing"
)

func TestToHexInt(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "Test 1", args: args{num: 97}, want: []byte("a")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ToHexInt(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHexInt() = %x, want %x", got, tt.want)
			}
		})
	}
}
