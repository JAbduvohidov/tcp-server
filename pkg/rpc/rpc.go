package rpc

import (
	"bufio"
	"io/ioutil"
	"log"
)

func ReadLine(reader *bufio.Reader) (line string, err error) {
	return reader.ReadString('\n')
}

func WriteLine(line string, writer *bufio.Writer) (err error) {
	_, err = writer.WriteString(line + "\n")
	if err != nil {
		return
	}
	err = writer.Flush()
	if err != nil {
		return
	}
	return
}

func GetFileList(path string) (fileList string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("can't read dir: %v", err)
	}
	for _, file := range files {
		if fileList == "" {
			fileList = fileList + file.Name()
		} else {
			fileList = fileList + " " + file.Name()
		}
	}
	fileList = fileList + "\n"
	return fileList
}

const (
	Download  = "download"
	Upload    = "upload"
	FileList  = "list"
	ResultErr = "result: error"
	ResultOK  = "result: ok"
	Tcp       = "tcp"
	Suffix    = "\n"

	ServerFilesPath = "files/"
	ClientFilesPath = "files/"
)
