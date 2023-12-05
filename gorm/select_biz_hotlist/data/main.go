package main

import (
	"bufio"
	"fmt"
	"os"
)

func genArticleData() {
	filePath := "article_data.sql"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 100000; i++ {
		cur := fmt.Sprintf("insert into interactives(biz_id,biz,read_cnt,like_cnt,ctime,utime)values(%d,'article',%d,%d,1699853275614,1699931894538);\n", i+200000, i+100, i)
		write.WriteString(cur)

	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func genVideoData() {
	filePath := "video_data.sql"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 100000; i++ {
		cur := fmt.Sprintf("insert into interactives(biz_id,biz,read_cnt,like_cnt,ctime,utime)values(%d,'video',%d,%d,1699853275614,1699931894538);\n", i+200000, i+100, i)
		write.WriteString(cur)

	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func main() {
	genArticleData()
	genVideoData()
}
