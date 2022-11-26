package lib

import (
	"log"
	"os"
)

// AppendFile takes two parameters once the path to the file and the content
// which should be appended to the file.
func AppendFile(path string, content string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(content)); err != nil {
		err := f.Close()
		if err != nil {
			return
		} // ignore error; Write error takes precedence
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
