// Copyright 2018 tinystack Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package golog

import (
    "os"
    "fmt"
)

type FileHandler struct {
    file    *os.File
    path    string
}

func NewFileHandler(path string) *FileHandler {
    file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
    if err != nil {
        panic(fmt.Sprintf("log file '%s' create failed, err output: %s", path, err.Error()))
    }
    return &FileHandler{
        file: file,
        path: path,
    }
}

func (f *FileHandler) Write(p []byte) (n int, err error) {
    return f.file.Write(p)
}

func (f *FileHandler) Close() {
    f.file.Close()
}

