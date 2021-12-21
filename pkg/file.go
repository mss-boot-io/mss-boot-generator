/*
 * @Author: lwnmengjing
 * @Date: 2021/12/16 7:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/12/16 7:39 下午
 */

package pkg

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// PathCreate create path
func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// PathExist path exist
func PathExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// FileCreate create file
func FileCreate(content bytes.Buffer, name string) error {
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	changeStr := strings.ReplaceAll(content.String(), `\$`, `$`)
	changeStr = strings.ReplaceAll(changeStr, `\}`, "}")
	changeStr = strings.ReplaceAll(changeStr, `\{`, "{")
	_, err = file.WriteString(changeStr)
	if err != nil {
		log.Println(err)
	}
	return err
}

// FileCopy copy file
func FileCopy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}