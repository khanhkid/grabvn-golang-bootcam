package readtxtfiles

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//ReadTxtFiles this is comment
func ReadTxtFiles(pathOfDir string) []string {
	var listFile []string
	err := filepath.Walk(pathOfDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(path, ".txt") == true {
				listFile = append(listFile, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return listFile
}
