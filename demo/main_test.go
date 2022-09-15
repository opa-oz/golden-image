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

	t.Run("should compare golden to golden", func(t *testing.T) {
		goldenImage.ToGildImage(t, 0.02, golden)
	})
	t.Run("should compare copper to copper", func(t *testing.T) {
		goldenImage.ToGildImage(t, 0.02, copper)
	})
}
