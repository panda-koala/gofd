package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

var group sync.WaitGroup

func main() {
	arg := new(Arg)
	arg.initArg()
	task := Task{Url:*arg.url, Dir:*arg.dir, Id:time.Now().Unix(), Chunk:*arg.chuck}
	//task.Url = " "
	task.getUrlInfo()
	err := task.Config()
	//if err != nil {
	//	fmt.Println(err)
	//}
	task.CreateTask()

	group.Wait()

	file, err := os.Create(path.Join(task.Dir, task.Name))
	if err != nil {
		return
	}


	for i := 0; i < len(task.blockList); i++ {
		n, _ := file.Seek(0, io.SeekEnd)
		src, err := os.Open(task.blockList[i].filePath + strconv.Itoa(i))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		content, err := ioutil.ReadAll(src)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		_, err = file.WriteAt(content, n)
	}
}
