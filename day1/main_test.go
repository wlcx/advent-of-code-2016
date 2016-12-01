package main

import (
	"reflect"
	"testing"
)

func TestApplyDirection(t *testing.T) {
	tests := []struct {
		name       string
		s          state
		directions []string
		want       state
	}{
		{
			"Moving N",
			state{
				0, 0, west,
			},
			[]string{
				"R4",
			},
			state{
				0, 4, north,
			},
		},
		{
			"Moving E",
			state{
				-3, 69, north,
			},
			[]string{
				"R13",
			},
			state{
				10, 69, east,
			},
		},
		{
			"Moving S",
			state{
				-10, 3, west,
			},
			[]string{
				"L23",
			},
			state{
				-10, -20, south,
			},
		},
		{
			"Moving W",
			state{
				12, 42, north,
			},
			[]string{
				"L17",
			},
			state{
				-5, 42, west,
			},
		},
		{
			"Multiple",
			state{
				0, 0, north,
			},
			[]string{
				"R0",
				"R10",
				"L4",
				"R23",
			},
			state{
				4, -33, south,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, d := range tt.directions {
				tt.s.applyDirection(d)
			}
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("%s failed: got(%v) want(%v)", tt.name, tt.s, tt.want)
			}
		})
	}
}

func TestBlocksAway(t *testing.T) {
	tests := []struct {
		name       string
		s          state
		blocksAway int
	}{
		{
			"Given example 1",
			state{
				2, 3, north,
			},
			5,
		},
		{
			"Given example 2",
			state{
				10, 2, north,
			},
			12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.blocksAway != tt.s.blocksAway() {
				t.Errorf("%s failed: got(%d) want(%d)", tt.name, tt.s.blocksAway(), tt.blocksAway)
			}
		})
	}
}
