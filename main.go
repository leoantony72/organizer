package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	file_ext := map[string]string{
		".txt": "text file",
		".pdf": "pdf",
	}

	files, err := os.Open(".")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer files.Close()

	fileinfo, err := files.Readdir(-1)
	if err != nil {
		fmt.Println("error reading directory:", err) //if directory is not read properly print error message
		return
	}
	for _, f := range fileinfo {

		if !f.IsDir() {
			ext := filepath.Ext(f.Name())
			fmt.Println(f.Name())

			if des, ok := file_ext[ext]; ok {
				os.Mkdir(des, 0755)
                

                fname := f.Name()
                newpath := filepath.Join(des,filepath.Base(fname))

                os.Rename(fname,newpath)
			}

		}
	}

}
