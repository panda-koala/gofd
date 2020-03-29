package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func (t *Task) Download (block int) {
	fmt.Println(block)
	client := &http.Client{}
	req, err := http.NewRequest("GET", t.Url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b := t.blockList[block]
	fmt.Println("download", b.Begin, b.End)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", b.Begin, b.End))

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(b.filePath + strconv.Itoa(block))
	blockFile, err := os.Create(b.filePath + strconv.Itoa(block))
	defer blockFile.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = io.Copy(blockFile, resp.Body)
	group.Done()
}


func (t *Task) getUrlInfo() {
	client := &http.Client{}
	req, err := http.NewRequest("HEAD", t.Url, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	req.Header.Set("Transfer-Encoding", "chunked")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	t.header = resp.Header
	length, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	t.Size = uint64(length)
}


