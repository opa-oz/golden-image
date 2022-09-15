package demo

import (
	goldenImage "github.com/opa-oz/golden-image"
	"testing"
)

func TestGet(t *testing.T) {
	golden, err := Get()
	if err != nil {
		t.Error(err)
		return
	}

	copper, err2 := GetCopper2()
	if err2 != nil {
		t.Error(err2)
		return
	}

	goldenImage.ToGildImage(t, 0.02, golden)
	goldenImage.ToGildImage(t, 0.02, copper)
}
