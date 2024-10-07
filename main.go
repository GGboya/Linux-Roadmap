package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	cnt := map[string]int{}

	filePath := "nginx-access.log.txt"

	// 使用os.Open()打开文件
	file, err := os.Open(filePath)
	if err != nil {
		// 处理错误
		fmt.Println("打开文件错误:", err)
		return
	}
	defer file.Close() // 确保文件在main函数结束前关闭

	// 创建Scanner对象，用于逐行读取文件
	scanner := bufio.NewScanner(file)

	// 使用scanner.Scan()逐行读取文件内容
	for scanner.Scan() {
		// 读取当前行
		line := scanner.Text()

		// 读取第三个 . ，遇到空格停止
		dotCnt := 0
		tmpStr := []byte{}
		for i := range line {
			if line[i] == '.' {
				dotCnt++
			} else if line[i] == ' ' {
				break
			}
			tmpStr = append(tmpStr, line[i])
		}
		if dotCnt == 3 {
			cnt[string(tmpStr)]++
		}
	}

	type pair struct {
		key   string
		value int
	}
	arr := []pair{}
	for k, v := range cnt {
		arr = append(arr, pair{k, v})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].value > arr[j].value
	})

	fmt.Println("Top 5 IP addresses with the most requests:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s - %d requests\n", arr[i].key, arr[i].value)
	}

	// 检查是否有读取错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时出现错误:", err)
	}
}
