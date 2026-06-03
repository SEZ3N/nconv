package internals

import (
	"errors"
	"fmt"
	"image"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetImage(imageURL string) (image.Image, error) {
	if st, err := os.Stat(imageURL); err == nil && !st.IsDir() {
		return resolveLocalUrl(imageURL)
	} else if errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("File Specified doesn't Exist")
	}

	url, err := url.ParseRequestURI(imageURL)
	if err == nil && (url.Scheme == "http" || url.Scheme == "https") {
		return resolveHttpUrl(imageURL)
	}
	return nil, errors.New("Unsupported URL")
}

func resolveLocalUrl(imageURL string) (image.Image, error) {
	f, err := os.Open(imageURL)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, form, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	if form != "jpeg" {
		return nil, errors.New("Image not a Jpeg")
	}

	return img, nil
}

func resolveHttpUrl(imageURL string) (image.Image, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", imageURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	if err != nil {
		return nil, err
	}
	r, err := client.Do(req)
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Couldn't get Image from Source http response %v", r.Status)
	}
	defer r.Body.Close()
	img, form, err := image.Decode(r.Body)
	if form != "jpeg" {
		return nil, errors.New("Image Provided Not a Jpeg")
	}
	return img, nil
}
