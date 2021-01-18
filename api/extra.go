package api

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Unzip unzip file to destination directory
func Unzip(src, dst string) error {
	// src not exists
	zr, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	// handle close error
	defer func() {
		if err := zr.Close(); err != nil {
			panic(err)
		}
	}()

	// mkdir
	err = os.MkdirAll(dst, 0755)
	if err != nil {
		return err
	}

	// extra files
	extraFiles := func(f *zip.File) error {
		// open file
		fo, err := f.Open()
		if err != nil {
			return err
		}

		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()

		// concat file path
		path := filepath.Join(dst, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dst)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		// if dirctory then create else copy
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())

			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, fo)

			if err != nil {
				return err
			}
		}

		return nil
	}

	for _, f := range zr.File {
		// skip macosx dir
		switch f.Name[0:8] {
		case "__MACOSX":
		default:
			err := extraFiles(f)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// DelZip remove file from temp diretory
func DelZip(src string) error {
	name := path.Join("temp", src+".zip")
	if err := os.Remove(name); err != nil {
		return err
	}
	return nil
}
