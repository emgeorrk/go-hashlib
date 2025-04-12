package tests

import (
	"testing"
)

type Hasher interface {
	Hash(data []byte) []byte
}

func HashConsistency(t *testing.T, hasher Hasher) {
	type args struct {
		data [][]byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "HashConsistency #1",
			args: args{
				data: [][]byte{
					[]byte("hello world"),
					[]byte("hello world"),
					[]byte("hello world"),
					[]byte("hello world"),
					[]byte("hello world"),
				},
			},
		},
		{
			name: "HashConsistency #2",
			args: args{
				data: [][]byte{
					[]byte(""),
					[]byte(""),
					[]byte(""),
					[]byte(""),
					[]byte(""),
				},
			},
		},
		{
			name: "HashConsistency #3",
			args: args{
				data: [][]byte{
					[]byte("This is a very long string that should be tested against the hash function for consistency."),
					[]byte("This is a very long string that should be tested against the hash function for consistency."),
					[]byte("This is a very long string that should be tested against the hash function for consistency."),
					[]byte("This is a very long string that should be tested against the hash function for consistency."),
					[]byte("This is a very long string that should be tested against the hash function for consistency."),
					[]byte("This is a very long string that should be tested against the hash function for consistency."),
					[]byte("This is a very long string that should be tested against the hash function for consistency."),
				},
			},
		},
		{
			name: "HashConsistency #4",
			args: args{
				data: [][]byte{
					[]byte{0x01, 0x02, 0x03},
					[]byte{0x01, 0x02, 0x03},
					[]byte{0x01, 0x02, 0x03},
					[]byte{0x01, 0x02, 0x03},
					[]byte{0x01, 0x02, 0x03},
					[]byte{0x01, 0x02, 0x03},
				},
			},
		},
		{
			name: "HashConsistency #5",
			args: args{
				data: [][]byte{
					[]byte("a"),
					[]byte("a"),
					[]byte("a"),
				},
			},
		},
		{
			name: "HashCollision #6",
			args: args{
				data: [][]byte{
					[]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
					[]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
				},
			},
		},
		{
			name: "HashConsistency #7",
			args: args{
				data: [][]byte{
					[]byte("aaaaaaaaaaaaaaaa"),
					[]byte("aaaaaaaaaaaaaaaa"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHash := make(map[string]struct{})
			for _, data := range tt.args.data {
				got := hasher.Hash(data)

				gotHash[string(got)] = struct{}{}
			}

			if len(gotHash) != 1 {
				t.Errorf("HashConsistency() gotHash = %v, args %v", gotHash, tt.args.data)
			}
		})
	}
}

func HashCollisions(t *testing.T, hasher Hasher) {
	type args struct {
		data [][]byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "HashCollision #1",
			args: args{
				data: [][]byte{
					[]byte("hello world1"),
					[]byte("hello world2"),
					[]byte("hello world3"),
					[]byte("hello world4"),
					[]byte("hello world5"),
				},
			},
		},
		{
			name: "HashCollision #2",
			args: args{
				data: [][]byte{
					[]byte("a"),
					[]byte("b"),
					[]byte("c"),
					[]byte("d"),
					[]byte("e"),
				},
			},
		},
		{
			name: "HashCollision #3",
			args: args{
				data: [][]byte{
					[]byte("This is a very long string that should be tested against the hash function for collision."),
					[]byte("Another long string that is used to check for hash collision."),
					[]byte("Yet another long string to test hash collision over large inputs."),
					[]byte("Long strings help verify that the hash function performs well with big inputs."),
					[]byte("A fifth long string for final collision test."),
				},
			},
		},
		{
			name: "HashCollision #4",
			args: args{
				data: [][]byte{
					[]byte{0x01, 0x02, 0x03},
					[]byte{0x04, 0x05, 0x06},
					[]byte{0x07, 0x08, 0x09},
					[]byte{0x10, 0x11, 0x12},
					[]byte{0x13, 0x14, 0x15},
				},
			},
		},
		{
			name: "HashCollision #5",
			args: args{
				data: [][]byte{
					[]byte("aaaaaaaaaaaaaaaaaaaaaaaa"),
					[]byte("bbbbbbbbbbbbbbbbbbbbbbbb"),
					[]byte("cccccccccccccccccccccccc"),
					[]byte("dddddddddddddddddddddddd"),
					[]byte("eeeeeeeeeeeeeeeeeeeeeeee"),
				},
			},
		},
		{
			name: "HashCollision #6",
			args: args{
				data: [][]byte{
					[]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
					[]byte("baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHash := make(map[string]struct{})
			for _, data := range tt.args.data {
				got := hasher.Hash(data)

				gotHash[string(got)] = struct{}{}
			}

			if len(gotHash) != len(tt.args.data) {
				t.Errorf("HashCollisions() gotHash = %v, args %v", gotHash, tt.args.data)
			}
		})
	}
}
