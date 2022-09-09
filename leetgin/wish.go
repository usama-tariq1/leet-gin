package leetgin

import "fmt"

type Gin struct {
}

func (gin Gin) wish(FileType string, Name string) {
	fmt.Println(FileType, Name)
}
