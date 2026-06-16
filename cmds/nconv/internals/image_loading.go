package internals

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"
)

func GetImage(imageURL string) (image.Image, error) {
	url, err := url.Parse(imageURL)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	st,stErr := os.Stat(imageURL)
	if stErr == nil && !st.IsDir() {
		return resolveLocalUrl(imageURL)
	}
	if url.Scheme == "http" || url.Scheme == "https" {
		return resolveHttpUrl(imageURL)
	} 
	match,_ := regexp.MatchString(`^(?:[a-zA-Z]:[\\/]|\\\\|/)`,imageURL)
	if !match {
		err = errors.New("Unsupported Protocol")
	} else if errors.Is(stErr, os.ErrNotExist) {
		err = errors.New("File Specified doesn't Exist")
	}
	return nil, err
}

func resolveLocalUrl(imageURL string) (image.Image, error) {
	f, err := os.Open(imageURL)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	img, form, err := image.Decode(buf)
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
		return nil, fmt.Errorf("Couldn't get Image from Source, http response: %v", r.Status)
	}
	defer r.Body.Close()
	buf := bufio.NewReaderSize(r.Body,256*1024)
	img, form, err := image.Decode(buf)
	if form != "jpeg" {
		return nil, errors.New("Image Provided Not a Jpeg")
	}
	return img, nil
}
