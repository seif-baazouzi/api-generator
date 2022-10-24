package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateDirectory(dirName string, doNotLog ...bool) string {
	if len(doNotLog) == 0 || !doNotLog[0] {
		fmt.Printf("[+] Create Directory %s\n", dirName)
	}

	path := filepath.Join(".", dirName)
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create `%s` directory\n", path)
		os.Exit(1)
	}

	return path
}

func RemoveDirectory(dirName string, doNotLog ...bool) string {
	if len(doNotLog) == 0 || !doNotLog[0] {
		fmt.Printf("[+] Remove Directory %s\n", dirName)
	}

	path := filepath.Join(".", dirName)
	err := os.RemoveAll(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not remove `%s` directory\n", path)
		os.Exit(1)
	}

	return path
}

func ClearDirectory(dirName string, doNotLog ...bool) string {
	RemoveDirectory(dirName)
	path := CreateDirectory(dirName)

	return path
}

func CreateFile(filePath string, content string, doNotLog ...bool) {
	if len(doNotLog) == 0 || !doNotLog[0] {
		fmt.Printf("[+] Create File %s\n", filePath)
	}

	err := os.WriteFile(filePath, []byte(content), os.ModePerm)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create `%s` file\n", filePath)
		os.Exit(1)
	}
}

func CopyFile(src string, dst string, doNotLog ...bool) {
	if len(doNotLog) == 0 || !doNotLog[0] {
		fmt.Printf("[+] Copy File %s to %s\n", src, dst)
	}

	in, err := os.Open(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read file %s\n", src)
		os.Exit(1)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create file %s\n", dst)
		os.Exit(1)
	}
	defer func() {
		err := out.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not close file %s\n", dst)
			os.Exit(1)
		}
	}()

	if _, err = io.Copy(out, in); err != nil {
		fmt.Fprintf(os.Stderr, "Could not copy %s into %s\n", src, dst)
		os.Exit(1)
	}

	err = out.Sync()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not write file %s\n", dst)
		os.Exit(1)
	}
}
