package main

import (
	// "bufio"
	"fmt"
	"os"
	"path/filepath"
	// "regexp"
	"strconv"
	"strings"
	"time"
	"syscall"
)

func fetchSourceList(root string, fromDate int) []string {
	result := make([]string, 0, 1000)

	sub_begin := len(root)
	if root[sub_begin-1] == '\\' || root[sub_begin-1] == '/' {
		root = string([]byte(root)[:sub_begin-1])
	}

	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if err == nil && !f.IsDir() && !strings.Contains(path, ".svn") && !strings.Contains(path, ".db") {
			subPath := []byte(path)[sub_begin+1:]
			if sys, ok:= (f.Sys()).(*syscall.Win32FileAttributeData); ok{
				createTime := (time.Duration)(sys.CreationTime.Nanoseconds())
				modTime := f.ModTime()
				duration := time.Since(modTime)
				
				var pass_days = (int)(duration.Hours() / 24)
				var create_pass_days = (int)((time.Now().Unix() - (int64)(createTime.Seconds())) / (60 * 60 * 24))
				
				// createDate := modTime.Add(-createTime)
				// offset := (int64)(createTime.Seconds()) - modTime.Unix()
				if pass_days > create_pass_days{
					pass_days = create_pass_days
				}
				if fromDate < 0 || pass_days <= (fromDate + 1) {
					result = append(result, strings.Replace(string(subPath), "\\", "/", -1))
					// fmt.Printf("last modification date %+v%s%s", modTime, " duration in days", pass_days)
				}
				//fmt.Println("Unix()", offset, " pass days ", pass_days);
			}else{
				println("error");
			}
		}
		return err
	})
	return result
}

func main() {
	argCnt := len(os.Args)

	url_prefix := "http://zfswz.hcgame.cn:8080/"
	fromDate := 366 * 20
	root := "./"
	if argCnt > 1 {
		root = os.Args[1]
	}
	if argCnt > 2 {
		fromDate, _ = strconv.Atoi(os.Args[2])
	}

	fmt.Println("parse dir:", root)
	fmt.Println("parse fromDate:", fromDate, " days.")

	now := time.Now()
	result := fetchSourceList(root, fromDate)

	if f, err := os.OpenFile("resource.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm); err == nil {
		from_date := now.AddDate(0, 0, -fromDate)
		f.WriteString("From " + from_date.Format("2006-01-02 15:04") + "to " + now.Format("2006-01-02 15:04") + "\r\n")
		f.WriteString("Update count: " + strconv.Itoa(len(result)) + "\r\n")
		for i := 0; i < len(result); i++ {
			f.WriteString(url_prefix + result[i] + "\r\n")
		}
		f.Close()
	}

	fmt.Println("file count:", len(result))
	fmt.Println("cost :", time.Since(now))

}
