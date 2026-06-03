package internals

import (
	"errors"
	"flag"
	"strconv"
)

func ValidateArgs(imageURL, nValue string) error {
	if imageURL == "" {
		return errors.New("No Image URL Provided")
	}
	if nValue == "" {
		return errors.New("nValue not Provided")
	}
	N, err := strconv.Atoi(nValue)
	if err != nil {
		return errors.New("N Value Provided not an int")
	}
	if N < 1 {
		return errors.New("N value Cannot be less than 1")
	}
	return nil
}

type Flags struct {
	SaveAtDir   string
	JpegQuality uint64
	OutputName  string
	Quiet       bool
}

type Args struct {
	ImageURL string
	N        string
}

func ParseFlagsAndArgs() (*Flags, *Args) {
	var flags Flags
	flag.StringVar(&flags.SaveAtDir, "dir", "./", "Directory to Save the file at defaults to the current dir")
	flag.Uint64Var(&flags.JpegQuality, "quality", 100, "Quality of Resultant Image in the range [0,100]")
	flag.StringVar(&flags.OutputName, "out-name", "", "Optional Name of the Output File (Defaults to a Randomly Generated Name)")
	flag.BoolVar(&flags.Quiet, "quiet", false, "If set to True Image isn't Opened by your defualt image viewer")
	flag.Parse()
	return &flags, &Args{flag.Arg(0), flag.Arg(1)}
}
