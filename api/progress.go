package api

import (
	"log"
)

// Progress progress bar
type Progress struct {
	Name      string
	TotalSize int64
	BytesRead int64
}

// Write is used to satisfy the io.Writer interface.
// Instead of writing somewhere, it simply aggregates
// the total bytes on each read
func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

// Print displays the current progress of the file upload
// each time Write is called
func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		log.Printf("%s upload DONE!\n", pr.Name)
		return
	}

	log.Printf("%s upload in progress: %d\n", pr.Name, pr.BytesRead)
}
