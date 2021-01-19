package api

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// File info
type File struct {
	ModifiedTime time.Time `json:"modified_time"`
	IsLink       bool      `json:"is_link"`
	IsDir        bool      `json:"is_dir"`
	LinksTo      string    `json:"links_to"`
	Size         int64     `json:"size"`
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Ext          string    `json:"ext"`
	Children     []*File   `json:"children"`
}

// FileIterate iterate file to struct
func FileIterate(path string) *File {
	rootOSFile, _ := os.Stat(path)
	rootFile := toFile(rootOSFile, path) //start with root file
	stack := []*File{rootFile}

	for len(stack) > 0 { //until stack is empty,
		file := stack[len(stack)-1] //pop entry from stack
		stack = stack[:len(stack)-1]
		children, _ := ioutil.ReadDir(file.Path) //get the children of entry
		for _, chld := range children {          //for each child
			child := toFile(chld, filepath.Join(file.Path, chld.Name())) //turn it into a File object
			file.Children = append(file.Children, child)                 //append it to the children of the current file popped
			stack = append(stack, child)                                 //append the child to the stack, so the same process can be run again
		}
	}

	return rootFile
}

func toFile(file os.FileInfo, path string) *File {
	JSONFile := File{ModifiedTime: file.ModTime(),
		IsDir:    file.IsDir(),
		Size:     file.Size(),
		Name:     file.Name(),
		Path:     path,
		Children: []*File{},
	}
	if file.Mode()&os.ModeSymlink == os.ModeSymlink {
		JSONFile.IsLink = true
		JSONFile.LinksTo, _ = filepath.EvalSymlinks(filepath.Join(path, file.Name()))
	} // Else case is the zero values of the fields

	if !file.IsDir() {
		JSONFile.Ext = filepath.Ext(filepath.Join(path, file.Name()))[1:]
	}
	return &JSONFile
}
