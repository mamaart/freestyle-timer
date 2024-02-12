package image

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"time"
)

type ImageState struct {
	images []image.Image
	ch     <-chan int
}

func NewImageState(name string, ch <-chan int) *ImageState {
	imgs, err := images(name)
	if err != nil {
		log.Fatal(err)
	}
	return &ImageState{
		images: imgs,
		ch:     ch,
	}
}

func (s *ImageState) Run(name string) {
	s.writeImg(name, 119)
	for n := range s.ch {
		s.writeImg(name, n)
	}
}

func (s *ImageState) writeImg(name string, n int) {
	file, err := os.Create(fmt.Sprintf("./svgs/%s.png", name))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img := s.images[n]
	if err := png.Encode(file, img); err != nil {
		log.Fatal(err)
	}
}

func images(name string) ([]image.Image, error) {
	images := make([]image.Image, 120)
	for i := 0; i < 120; i++ {
		filename := fmt.Sprintf("./svgs/%s-%03d.png", name, i)
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			return nil, err
		}

		images[i] = img
	}
	return images, nil
}

func setCurrent(time.Duration) error {
	return nil
}
