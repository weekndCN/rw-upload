package api

import (
	"fmt"
	"testing"
)

func TestUnzip(t *testing.T) {
	err := Unzip("../temp/fic_data_screen.zip", "../uploads")
	if err != nil {
		fmt.Println(err)
	}
}
