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

// Zip zip srouce file or directory to zip format
func Zip(src, dst, zipName string) error {
	// check zipNmae exists in destination directory
	zipFile := path.Join(dst, zipName+".zip")
	if _, err := os.Stat(zipFile); os.IsExist(err) {
		return fmt.Errorf("filename exists")
	}

	// mkdir directory
	err := os.MkdirAll(dst, 0755)
	if err != nil {
		return fmt.Errorf("making folder for destination: %v", err)
	}

	// create zip file in temp diretory
	out, err := os.Create(zipFile)
	if err != nil {
		return fmt.Errorf("creating %s: %v", dst, err)
	}

	defer out.Close()

	//new a zip.Write
	zw := zip.NewWriter(out)

	defer func() {
		if err := zw.Close(); err != nil {
			panic(err)
		}
	}()

	// walk from path.Join(src, zipName)
	return filepath.Walk(path.Join(src, zipName), func(path string, fi os.FileInfo, errWalk error) error {
		if errWalk != nil {
			return errWalk
		}

		// create info header(zip header)
		zh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}
		zh.Name = strings.TrimPrefix(path, string(filepath.Separator))

		fmt.Println("zip name:", zh.Name)

		// if os.FileInfo is dir
		if fi.IsDir() {
			zh.Name += "/"
		}

		// write file info to zip file,return a writer struct
		w, err := zw.CreateHeader(zh)
		if err != nil {
			return err
		}

		// if not a file , no data need to write
		if !zh.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
			return err
		}

		// io copy
		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}

		return nil
	})
}
