package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// UploadService .

const limitSize = -1     // limit single file size
const maxBody = 32 << 20 // 32M limit body
const readBuff = 512     // read buffer size

// HandleUpload handler upload func
func HandleUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//vars := mux.Vars(r)

		// dirname, _ := vars["dirname"]

		// limit body
		r.Body = http.MaxBytesReader(w, r.Body, maxBody+512)

		// limit buffer size 32M
		// if overhead then store in disk the rest of data
		r.ParseMultipartForm(maxBody)
		// if no file receive
		if r.MultipartForm == nil {
			log.Println("MultipartForm is null")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("file not found in body"))
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
					w.WriteHeader(http.StatusPreconditionFailed)
					w.Write([]byte(fmt.Sprintf("%s 文件大小超过限制", file.Filename)))
					return
				}
			}
			// check name
			if file.Filename == "" {
				log.Printf("%s 文件名允许为空\n", file.Filename)
				w.WriteHeader(http.StatusPreconditionFailed)
				w.Write([]byte(fmt.Sprintf("%s 文件名允许为空", file.Filename)))
				return
			}

			f, err := file.Open()

			if err != nil {
				log.Printf("%s 文件打开失败\n", file.Filename)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			defer f.Close()
			// detech file type

			buff := make([]byte, readBuff)
			_, err = f.Read(buff)

			if err != nil {
				log.Printf("%s 文件读取失败\n", file.Filename)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			// limit file type
			filetype := http.DetectContentType(buff)
			switch filetype {
			case "image/jpeg",
				"image/jpg",
				"image/gif",
				"image/png",
				"application/pdf",
				"text/plain; charset=utf-8",
				"application/x-gzip",
				"application/octet-stream":
			default:
				log.Printf("%s 文件的%s 格式不支持\n", file.Filename, filetype)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("%s 文件的%s 格式不支持", file.Filename, filetype)))
				return
			}

			_, err = f.Seek(0, io.SeekStart)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			err = os.MkdirAll("./uploads", os.ModePerm)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			fd, err := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			defer fd.Close()

			// add progress bar
			pr := &Progress{
				TotalSize: file.Size,
				Name:      file.Filename,
			}

			_, err = io.Copy(fd, io.TeeReader(f, pr))
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("文件上传成功"))
	}
}
