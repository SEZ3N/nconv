package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "image/jpeg"
	"os"
	"strconv"
	"github.com/SEZ3N/nconv"
	"github.com/SEZ3N/nconv/cmds/nconv/internals"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Printf("\nsyntax: nconv [options] <input_image_URL> <n_value>\n")
		fmt.Printf("\nAvailable Options:\n\n")
		flag.PrintDefaults()
	}

	flags, args, err  := internals.ParseFlagsAndArgs()

	if err != nil {
		internals.DisplayFomattedMessageFromErr(err)
		flag.CommandLine.Usage()
		return
	}
	if err := internals.ValidateArgs(args.ImageURL, args.N); err != nil {
		internals.DisplayFomattedMessageFromErr(err)
		return
	}
	n_value, _ := strconv.Atoi(args.N)
	img, err := internals.GetImage(args.ImageURL)
	if err != nil {
		internals.DisplayFomattedMessageFromErr(err)
		return
	}

	f, err := internals.CreateOutputFile(flags.SaveAtDir, flags.OutputName)
	if err != nil {
		internals.DisplayFomattedMessageFromErr(err)
		return
	}
	defer f.Close()
	wr := bufio.NewWriterSize(f, 512*1024)
	defer wr.Flush()
	err = nconv.ConvertAndWrite(img, uint(n_value), wr, int(flags.JpegQuality))
	if err != nil {
		internals.DisplayFomattedMessageFromErr(err)
		return
	}
	internals.DisplayMessageOnSuccess("Output saved at", f.Name())

	if !flags.Quiet {
		err := internals.OpenImageInViewer(f.Name())
		if err != nil {
			internals.DisplayFomattedMessageFromErr(err)
			return
		}
	}
}
