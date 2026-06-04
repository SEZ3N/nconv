# nconv
Nconv is a Golang library/CLI made to assist artists in their value studies by representing any given image using specified number of greyscale colors.
## CLI Installation
```terminal
go install github.com/SEZ3N/nconv/cmds/nconv@latest
```
## Library Installation
```terminal
go get -u github.com/SEZ3N/nconv
```

### Ouput Examples
Input images with their 2 and 3 value decompositions are shown below.

<img width="30%" alt="noelle-xHDN-da46r8-unsplash" src="https://github.com/user-attachments/assets/42b93759-5495-439d-94d1-572a752e5253" /> 
<img width="30%" alt="nconv-2b0341db51317e66" src="https://github.com/user-attachments/assets/57691457-dda0-48c7-adac-e42c94473cf8" />
<img width="30%" alt="nconv-dbe19e7a04238bac" src="https://github.com/user-attachments/assets/0a85b428-b05f-4283-b2c3-e3f5a8011a87" />
<br><br/>
<img width="30%" alt="maximilian-bungart-8_oQ-qNsFnw-unsplash(2)" src="https://github.com/user-attachments/assets/5d3dffa9-2b34-4f64-b9f7-05963e62a45c" />
<img width="30%" alt="nconv-a9dea30348de7ab3" src="https://github.com/user-attachments/assets/4151836f-d21f-4571-bb16-e84a9e9df66d" />
<img width="30%" alt="nconv-acd45670df53cbb4" src="https://github.com/user-attachments/assets/f180419e-6850-4abb-bdcb-d2cd8d39daee" />

## CLI Usage
```terminal
nconv [Options] <input_image_URL> <N_Value>
```
* input_image_URL could be a local image path or a web URL pointing to an image
* N_value is an unsigned int greater than 1
* All Options can be viewed using `nconv -h`

### Examples
Below snippet converts img.jpg into its 2 value counterpart and saves it with 50% quality in ./out directory with a random name, when its done it opens the image in users default image viewer.
```terminal
nconv -dir ./out -quality 50 ./image.jpg 2
```
Below snippet converts the image pointed to by the web URL into its 3 value counterpart and saves it in the working directory with the name output.jpg, with the quiet flag set to true it doesn't open the image in the default viewer.
```terminal
nconv -out-name output -quiet=true https://example.com/image.jpg 3
```

## Library Usage
The nconv library is really small. It only exports two functions!
```Golang

func main() {
	// load your image
	img, err := openImage("exampleImage.jpg")
	if err != nil {
		return
	}

	f, err := os.Create("out.jpg")
	if err != nil {
		return
	}

	// Convert returns an RGBA which implements the color.Color interface
	outImg := nconv.Convert(img, 2)
	// img now stores the 2 value representation of the input

	// ConvertAndWrite takes the image, the n value, an io.Writer and the jpg quality
	err = nconv.ConvertAndWrite(img, 3, f, 50)
	if err != nil {
		return
	}
	// The image would now be Outputted to the file specified by the writer
}

// helper func to read an image into an image.Image interface
func openImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	defer f.Close()
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

## Roadmap
I plan on adding support for .png and .webp images later, adding some convinient functions to the library and making the processing parallel is on my mind too.
