package golden_image

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
)

func saveSnapshot(testID, snapPath string, img image.Image) error {
	pathToFile, err := filepath.Abs(buildPath(snapPath, testID))

	f, _ := os.Create(pathToFile)
	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}

func getPrevSnapshot(testID, snapPath string) (image.Image, error) {
	pathToFile, err := filepath.Abs(buildPath(snapPath, testID))
	println(pathToFile)
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	img, _, err2 := image.Decode(file)
	if err2 != nil {
		return nil, err2
	}

	return img, nil
}

func cleanDiffPath(diffPath string) {
	err := os.Remove(diffPath)

	if err != nil {
		return
	}
}
