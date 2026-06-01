# nconv
nconv is a golang library made to assist artists in their value studies, nconv breaks an image into its N - Value conterpart i.e it represents the image using only N greyscale colors

## Usage
Its a small library the nconv file only exports two functions!
```Golang
import (
	"fmt"
	"image"
	"github.com/SEZ3N/nconv"
	"os"
)

func main(){

	img,err := openImage("")// load image
	if err != nil {
		fmt.Println(err)
		return 
	}

	f,err := os.Create("out.jpg") //returns *os.file which implements the io.Writer interface

  if err != nil {
		fmt.Println(err)
		return 
	}

  //Convert returns an RGBA which implements the color.Color interface
  img := nconv.Convert(img,2)


  //ConvertAndWrite takes the image, the n value, an io.Writer and the jpg quality outputs the image to the file specified by the writer
	err = nconv.ConvertAndWrite(img,2,f,100)
	if err != nil {
		fmt.Println(err)
		return	
	}
}



func openImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, form, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}


```
