package golog

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	ChannelSingle = iota
	ChannelDaily
)

type FileHandler struct {
	path    string
	fh      *os.File
	mu      sync.Mutex
	channel int
	cdate   int
}

func NewFileHandler(path string, channel int) *FileHandler {
	f := &FileHandler{
		path:    path,
		channel: channel,
	}
	return f
}

func (f *FileHandler) Write(p []byte) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if err := f.checkFh(); err != nil {
		return 0, err
	}
	return f.fh.Write(p)
}

func (f *FileHandler) Close() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.close()
}

func (f *FileHandler) close() error {
	if f.fh == nil {
		return nil
	}
	err := f.fh.Close()
	f.fh = nil
	return err
}

func (f *FileHandler) checkFh() (err error) {
	if f.fh == nil {
		err = f.openFile()
		return
	}
	if f.cdate != todayDate() {
		f.close()
		err = f.openFile()
		return
	}
	return nil
}

func (f *FileHandler) openFile() error {
	fh, err := os.OpenFile(f.logFile(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	f.fh = fh
	f.cdate = todayDate()
	return nil
}

func (f *FileHandler) logFile() (filePath string) {
	switch f.channel {
	case ChannelSingle:
		filePath = f.path
	case ChannelDaily:
		filePath = strings.Join([]string{f.path, ".", time.Now().Format("20060102")}, "")
	default:
		filePath = f.path
	}
	return
}

func todayDate() int {
	t, _ := strconv.Atoi(time.Now().Format("20060102"))
	return t
}
