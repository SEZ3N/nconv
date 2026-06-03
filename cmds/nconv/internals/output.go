package internals

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func genRandomFileName(ext string) (string, error) {
	buff := make([]byte, 8)
	_, err := rand.Read(buff)
	if err != nil {
		return "", err
	}

	return "nconv-" + hex.EncodeToString(buff) + ext, err
}

func CreateOutputFile(dir string, name string) (*os.File, error) {
	absDirPath, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	if stat, err := os.Stat(absDirPath); err != nil || !stat.IsDir() {
		err := os.Mkdir(dir, os.ModeDir)
		if err != nil {
			return nil, err
		}
	}
	fname, err := genRandomFileName(".jpg")
	if err != nil {
		return nil, err
	}
	if name != "" {
		fname = name + ".jpg"
	}
	fpath := filepath.Join(absDirPath, fname)
	f, err := os.Create(fpath)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func OpenImageInViewer(imgPath string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer.exe", imgPath)
	case "darwin":
		cmd = exec.Command("open", imgPath)
	case "linux":
		cmd = exec.Command("xdg-open", imgPath)
	default:
		return errors.New("Unsupported Operating System")
	}
	return cmd.Start()
}
