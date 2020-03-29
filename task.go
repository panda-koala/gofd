package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const FileConfExt string = ".conf"
const ChuckExt string = ".stream"
const ChuckTmpExt string = ".tmp"

func (t *Task) CreateTask() {
	blockSize := t.Size / uint64(t.Chunk)
	t.blockList = make(map[int]*FileBlock)
	for i := 0; i < t.Chunk; i++ {
		begin := uint64(i) * blockSize
		end := begin + blockSize - 1
		if i+1 == t.Chunk {
			end = t.Size
		}
		t.blockList[i] = &FileBlock{
			Begin:begin,
			End:end,
			Current:0,
			filePath:path.Join(t.Dir, t.Name) + ChuckExt,
			tempPath:path.Join(t.Dir, t.Name) + ChuckTmpExt,
		}

		group.Add(1)
		go t.Download(i)
	}
}


func (t *Task) Config() error {
	if t.Dir == "" {
		err := t.setDir()
		if err != nil {return err}
	}

	err := t.getConf()
	if err != nil {return err}
	return nil
}

func (t *Task) getConf() error {

	_, err := os.Stat(path.Join(t.Dir, t.Name) + FileConfExt)
	if err != nil {
		return t.createConf()
	}
	return nil
}

func (t *Task) createConf() error {
	cf, err := os.Create(path.Join(t.Dir, t.Name) + FileConfExt)
	if err != nil {return err}
	defer cf.Close()

	c, err := json.Marshal(t)
	if err != nil {return err}

	_, err = cf.Write(c)
	if err != nil {return err}

	return nil
}

func (t *Task) setDir() error {
	t.Name = path.Base(t.Url)

	currentPath, err := os.Executable()
	if err != nil {return err}

	dir := filepath.Dir(currentPath)

	fileExt := path.Ext(t.Name)
	fileDir := path.Join(dir, strings.TrimSuffix(t.Name, fileExt))
	t.Dir = filepath.FromSlash(fileDir)
	err = os.Mkdir(t.Dir, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
