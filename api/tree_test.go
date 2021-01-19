package api

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFileIterate(t *testing.T) {
	// rootpath, _ := os.Getwd()
	var rootFile *File
	rootFile = FileIterate("rw-upload/temp")
	output, _ := json.MarshalIndent(rootFile, "", "     ")
	fmt.Println(string(output))
}
