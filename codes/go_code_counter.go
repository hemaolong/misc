package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func fetchSourceList(root string, filter string) []string {
	result := make([]string, 0, 1000)

	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if err == nil && !f.IsDir() {
			ok, _ := regexp.MatchString(filter, path)
			if ok {
				result = append(result, path)
			}
		}
		return err
	})
	return result
}

func calcSourceLineCnt(path string) int {
	f, err := os.Open(path)
	if err != nil {
		return 0
	}

	defer f.Close()

	result := 0
	reader := bufio.NewReader(f)
	blockCommentBegin := false
	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		s := string(l)
		s = strings.Trim(s, " \r\n")

		// Skip empty line
		if len(s) < 1 {
			continue
		}
		//fmt.Println(s)

		// Skip line comment
		if strings.HasPrefix(s, "//") {
			continue
		}

		// /*xxxx*/
		if m, _ := regexp.MatchString("^(/\\*.*\\*/)$", s); m {
			continue
		}

		// xxx/*yyy*/zzz
		if m, _ := regexp.MatchString("(/\\*.*\\*/)", s); !m {
			if blockCommentBegin {
				if m, _ := regexp.MatchString("(\\*/)", s); m {
					blockCommentBegin = false
				}
				continue
			} else {
				if m, _ := regexp.MatchString("(/\\*)", s); m {
					blockCommentBegin = true
					continue
				}
			}
		}

		result++

	}
	return result
}

func main() {
	argCnt := len(os.Args)

	var filter string = "*"
	root := "./"
	if argCnt > 1 {
		root = os.Args[1]
	}
	if argCnt > 2 {
		filter = os.Args[2]
	}
	fmt.Println("parse dir:", root)
	fmt.Println("parse filter:", filter)

	now := time.Now()
	result := fetchSourceList(root, filter)
	totalCnt := 0
	for _, i := range result {
		totalCnt += calcSourceLineCnt(i)
	}

	fmt.Println("file count:", len(result))
	fmt.Println("total line:", totalCnt)
	fmt.Println("cost :", time.Since(now))
}
