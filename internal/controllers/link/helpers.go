package link

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"
)

func pickKey() (string, error) {
	f, err := os.OpenFile("./assets/keys.txt", os.O_RDWR, os.ModeAppend)
	if err != nil {
		return "", err
	}
	fi, err := f.Stat()
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(buf, f)
	if err != nil {
		return "", err
	}

	line, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}
	nw, err := io.Copy(f, buf)
	if err != nil {
		return "", err
	}
	err = f.Truncate(nw)
	if err != nil {
		return "", err
	}
	err = f.Sync()
	if err != nil {
		return "", err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(line)), nil
}

func isValidLink(link string) bool {
	pattern, err := regexp.Compile(`^(http|ftp|https)?(\:\/\/)?[\w-]+(\.[\w-]+)+([\w.,@?^!=%&amp;:\/~+#-]*[\w@?^=%&amp;\/~+#-])+$`)
	if err != nil {
		return false
	}

	return pattern.MatchString(link)

}
