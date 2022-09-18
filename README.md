# GoldenImage

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/opa-oz/golden-image) [![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/opa-oz/golden-image)

<p>Snapshot-testing tool with build-in image comparison</p>

<p align="center">
<img src="https://github.com/opa-oz/golden-image/raw/main/docs/GildYourImage.jpg" alt="Banner"/>
</p>

## Installation

To install `golden-image`, simply use `go get`:

```bash
$> go get github.com/opa-oz/golden-image
```

## Usage

Hungry for image comparison in your autotests? Simply gild your images for the first time and enjoy comparison

```go
import (
    goldenImage "github.com/opa-oz/golden-image"
    "testing"
)

func TestGet(t *testing.T) {
    firstImage, err := GetImage(1) // image.Image type
    if err != nil {
        t.Error(err)
        return
    }
    
    secondImage, err2 := GetImage(2)
    if err2 != nil {
        t.Error(err2)
        return
    }
    
    t.Run("should compare first to first", func (t *testing.T) {
        goldenImage.ToGildImage(t, 0.02, firstImage) // compare image to gilded one, or gild if it's not existing yet
    })
    
    t.Run("should compare second to second", func (t *testing.T) {
        goldenImage.ToMatchSnapshot(t, 0.02, secondImage) // you can also use alias `ToMatchSnapshot`
    })
}
```

You will find your baseline images (a.k.a. gilded images) inside `__baseimage__`

### But...

If your image has differences with the golden-set's image, you'll get conveniently printed `FAIL`:

```go
--- FAIL: TestGet (0.63s)
    --- FAIL: TestGet/should_compare_first_to_first (0.50s)
        main_test.go:22: • Found diff: 0.474205.
            Highlights are stored at .../__baseimages__/__diffs__/main_test.[TestGet-should_compare_first_to_first-0001].png.
```

And what is `highlights` you may ask?

Let's pretend this is our original image (the gilded one):
<p><img src="https://github.com/opa-oz/golden-image/raw/main/docs/golden.png" width="150" alt="Original"/></p>

And after our `image generation algorithm` it was slightly changed:
<p><img src="https://github.com/opa-oz/golden-image/raw/main/docs/changed.png" width="150" alt="Original"/></p>

In this case, your comparison will fail and we'll see the difference, highlighted like that:
<p><img src="https://github.com/opa-oz/golden-image/raw/main/docs/highlights.png" width="150" alt="Original"/></p>

## Update?

Of course, you can update gilded images, just casually run:

```bash
$> UPDATE_SNAPS=true go test
```

## Cheers 🥂

- _[Original Gophers' images repo](https://github.com/egonelbre/gophers)_
- Mostly inspired by _[go-snaps](https://github.com/gkampitakis/go-snaps)_