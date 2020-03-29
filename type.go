package main

import (
	"net/http"
)

type  Task struct {
	Url string `json:"url"`
	Name string `json:"name"`
	Dir string `json:"dir"`
	Id int64 `json:"id"`
	Size uint64 `json:"size"`
	Chunk int `json:"chunk"`
	blockList map[int]*FileBlock
	header http.Header
}

type FileBlock struct {
	Begin uint64 `json:"begin"`
	End uint64 `json:"end"`
	Current uint64 `json:"current"`
	filePath string
	tempPath string
}

