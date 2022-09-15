package demo

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

func Open(path string) (img image.Image, err error) {
	pathToFile, err := filepath.Abs(path)
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	img, _, err = image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, err
}

func Get() (image.Image, error) {
	img, err := Open("./subjects/golden.jpg")
	if err != nil {
		print(err.Error())
		return nil, err
	}

	return img, nil
}

func GetCopper() (image.Image, error) {
	img, err := Open("./subjects/copper.jpg")
	if err != nil {
		print(err.Error())
		return nil, err
	}

	return img, nil
}

func GetCopper2() (image.Image, error) {
	img, err := Open("./subjects/copper2.jpg")
	if err != nil {
		print(err.Error())
		return nil, err
	}

	return img, nil
}
