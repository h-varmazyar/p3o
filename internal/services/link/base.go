package link

import (
	"bytes"
	"context"
	"github.com/h-varmazyar/p3o/internal/entities"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"regexp"
	"strings"
)

type linkRepository interface {
	Create(ctx context.Context, link entities.Link) (entities.Link, error)
	ReturnByKey(ctx context.Context, key string) (entities.Link, error)
	List(ctx context.Context, userId uint) ([]entities.Link, error)
	Update(ctx context.Context, link entities.Link) error
	TotalLinkCount(ctx context.Context, userId uint) (int64, error)
	Delete(ctx context.Context, key string) error
	Visits(ctx context.Context, userId uint) (int64, error)
}

type Service struct {
	log      *log.Logger
	linkRepo linkRepository
}

func New(log *log.Logger, linkRepo linkRepository) Service {
	return Service{
		log:      log,
		linkRepo: linkRepo,
	}
}

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
