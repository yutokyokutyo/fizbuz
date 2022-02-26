package main

import (
	"testing"
)

func TestFizBuz(t *testing.T) {
	tests := map[string]struct {
		in      string // 入力
		want    string // 期待する返り値
		wantErr bool   // エラーになることを期待するかどうか
	}{
		"3の倍数ならFizz":  {"3", "Fizz", false},
		"5の倍数ならBuzz":  {"5", "Buzz", false},
		"15の倍数ならBuzz": {"15", "FizzBuzz", false},
	}
	for _, tt := range tests {
		got, err := FizBuz(tt.in)
		if err != nil {
			if !tt.wantErr {
				t.Error(err)
			}
			continue
		}
		if tt.wantErr {
			t.Error("should not be error")
			continue
		}
		if got != tt.want {
			t.Errorf("got %v\nwant %v", got, tt.want)
		}
	}
}
