package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_deals(t *testing.T) {
	type args struct {
		d        deck
		handSize int
	}
	tests := []struct {
		name  string
		args  args
		want  deck
		want1 deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := deals(tt.args.d, tt.args.handSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("deals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func Test_deck_print(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.print()
		})
	}
}

func Test_deck_saveToFile(t *testing.T) {
	os.Remove("_decktesting")

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		d       deck
		args    args
		wantErr bool
	}{
		{
			name: "Expected 16 cards in deck, got ",
			d:    newDeck(),
			args: args{"_decktesting"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.saveToFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("saveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_deck_shuffle(t *testing.T) {
	tests := []struct {
		name string
		d    deck
		want deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.shuffle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deck_toString(t *testing.T) {
	tests := []struct {
		name string
		d    deck
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.toString(); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newDeck(t *testing.T) {
	lengthTests := []struct {
		name string
		want int
	}{
		{
			name: "16個カードがあるか",
			want: 16,
		},
	}
	for _, tt := range lengthTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeck(); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("newDeck() = %v, want %v", got, tt.want)
			}
		})
	}

	valuesTests := []struct {
		name  string
		index int
		want  string
	}{
		{
			name:  "1マイ目想定したものか",
			index: 0,
			want:  "Ace of Spades",
		},
		{
			name:  "最終マイ目想定したものか",
			index: len([]string(newDeck())) - 1,
			want:  "Four of Clubs",
		},
	}
	for _, tt := range valuesTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeck(); !reflect.DeepEqual(got[tt.index], tt.want) {
				t.Errorf("newDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	newDeck := newDeck()
	newDeck.saveToFile("_decktesting")

	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ロードした結果16個分帰ってくるか",
			args: args{"_decktesting"},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeckFromFile(tt.args.filename); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("newDeckFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
