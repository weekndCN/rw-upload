package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/weekndCN/rw-upload/api"
	"github.com/weekndCN/rw-upload/render"
)

// UploadService .
const (
	limitSize = -1       // limit single file size
	maxBody   = 32 << 20 // 32M limit body
	readBuff  = 512      // read buffer size
)

// HandleUpload handler upload func
func HandleUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//vars := mux.Vars(r)
		// dirname, _ := vars["dirname"]
		var uploadPath string

		baseDir, staticPath, tempPath := api.Basedir(staticDir, tmpDir)
		if baseDir == "" {
			render.BadRequest(w, render.ErrNotFound)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, maxBody+512)

		// limit buffer size 32M
		// if overhead then store in disk the rest of data
		r.ParseMultipartForm(maxBody)
		// limit body
		// if no file receive
		if r.MultipartForm == nil {
			log.Println("MultipartForm is null")
			render.BadRequest(w, render.ErrNotFound)
			return
		}

		// if custom directory
		dir := r.MultipartForm.Value["dir"][0]

		uploadPath = path.Join(staticPath, dir)

		err := os.MkdirAll(uploadPath, os.ModePerm)

		if err != nil {
			log.Println(err.Error())
			render.InternalError(w, err)
			return
		}

		// k is filename, v is fileheader pointer
		// receiver data
		// frontend must append file with name “file” ,avoid server loop map
		files := r.MultipartForm.File["file"]

		for _, file := range files {
			// check file size. -1 is no limit
			if limitSize > 0 {
				if file.Size > limitSize {
					log.Printf("%s 文件大小超过限制\n", file.Filename)
					render.BadRequest(w, render.ErrOverLimit)
					return
				}
			}
			// check name
			if file.Filename == "" {
				log.Printf("%s 文件名允许为空\n", file.Filename)
				render.BadRequest(w, render.ErrEmptyName)
				return
			}

			f, err := file.Open()

			if err != nil {
				log.Printf("%s 文件打开失败\n", file.Filename)
				render.InternalError(w, err)
				return
			}

			defer f.Close()
			// detech file type

			buff := make([]byte, readBuff)
			_, err = f.Read(buff)

			if err != nil {
				log.Printf("%s 文件读取失败\n", file.Filename)
				render.InternalError(w, err)
				return
			}

			// limit file type
			filetype := http.DetectContentType(buff)
			/*
				switch filetype {
				case "image/jpeg",
					"image/jpg",
					"image/gif",
					"image/png",
					"application/pdf",
					"text/plain; charset=utf-8",
					"application/x-gzip",
					"application/zip",
					"application/octet-stream":
				default:
					log.Printf("%s 文件的%s 格式不支持\n", file.Filename, filetype)
					render.NotImplemented(w, render.ErrFormatSupport)
					return
				}
			*/

			_, err = f.Seek(0, io.SeekStart)
			if err != nil {
				render.InternalError(w, err)
				return
			}

			var fd *os.File
			if filetype == "application/zip" {
				fd, err = os.Create(path.Join(tempPath, file.Filename))
			} else {
				fd, err = os.Create(path.Join(uploadPath, file.Filename))
			}

			if err != nil {
				log.Println(err.Error())
				render.BadRequest(w, err)
				return
			}

			defer fd.Close()

			// add progress bar
			pr := &api.Progress{
				TotalSize: file.Size,
				Name:      file.Filename,
			}

			_, err = io.Copy(fd, io.TeeReader(f, pr))
			if err != nil {
				log.Println(err.Error())
				render.InternalError(w, err)
				return
			}

			if filetype == "application/zip" {
				fileName := path.Join(tempPath, file.Filename)
				if err := api.Unzip(fileName, staticPath); err != nil {
					log.Println(err.Error())
				}
				api.DelZip(fileName)
			}
		}

		render.JSON(w, "finished", 200)
	}
}
