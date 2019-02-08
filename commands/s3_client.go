package commands

import (
	"errors"
	"fmt"
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/s3"
	_ "github.com/graymeta/stow/s3"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//go:generate counterfeiter -o ./fakes/config_service.go --fake-name Config . Config
type Config interface {
	Config(name string) (string, bool)
	Set(name, value string)
}

type Stower interface {
	Dial(kind string, config Config) (stow.Location, error)
	Walk(container stow.Container, prefix string, pageSize int, fn stow.WalkFunc) error
}

type S3Configuration struct {
	Bucket          string `yaml:"bucket" validate:"required"`
	AccessKeyID     string `yaml:"access-key-id" validate:"required"`
	SecretAccessKey string `yaml:"secret-access-key" validate:"required"`
	RegionName      string `yaml:"region-name" validate:"required"`
	Endpoint        string `yaml:"endpoint" validate:"required"`
	DisableSSL      bool `yaml:"disable-ssl" `
}

type S3Client struct {
	stower Stower
	bucket string
	Config stow.Config
}

func NewS3Client(stower Stower, config S3Configuration) (*S3Client, error) {
	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		return nil, err
	}

	disableSSL := strconv.FormatBool(config.DisableSSL)
	stowConfig := stow.ConfigMap{
		s3.ConfigAccessKeyID: config.AccessKeyID,
		s3.ConfigSecretKey:   config.SecretAccessKey,
		s3.ConfigRegion:      config.RegionName,
		s3.ConfigEndpoint:    config.Endpoint,
		s3.ConfigDisableSSL:  disableSSL,
	}

	return &S3Client{
		stower: stower,
		Config: stowConfig,
		bucket: config.Bucket,
	}, nil
}

func (s3 S3Client) GetAllProductVersions(slug string) ([]string, error) {
	files, err := s3.ListFiles()
	if err != nil {
		return nil, err
	}

	validFile := regexp.MustCompile(fmt.Sprintf(`^%s-(.*?)_`, slug))

	var versions []string
	versionFound := make(map[string]bool)
	for _, f := range files {
		x := validFile.FindStringSubmatch(f)
		if len(x) == 2 {
			version := x[1]
			if !versionFound[version] {
				versions = append(versions, version)
				versionFound[version] = true
			}
		}

	}
	return versions, nil

}

func (s3 S3Client) GetLatestProductFile(slug, version, glob string) (*FileArtifact, error) {
	files, err := s3.ListFiles()
	if err != nil {
		return nil, err
	}

	validFile := regexp.MustCompile(fmt.Sprintf(`^%s-%s`, slug, version))
	var artifacts []string

	for _, f := range files {
		if validFile.MatchString(f) {
			matched, _ := filepath.Match(glob, f)
			if matched {
				artifacts = append(artifacts, f)
			}
		}
	}

	if len(artifacts) > 1 {
		return nil, fmt.Errorf("the glob '%s' matches multiple files. Write your glob to match exactly one of the following:\n  %s", glob, strings.Join(artifacts, "\n  "))
	}

	if len(artifacts) == 0 {
		return nil, fmt.Errorf("the glob '%s' matches no file", glob)
	}

	return &FileArtifact{Name: artifacts[0]}, nil
}

func (s3 S3Client) DownloadProductToFile(fa *FileArtifact, file *os.File) error {
	f, err := s3.DownloadFile(fa.Name)
	if err != nil {
		return err
	}

	if _, err := io.Copy(file, f); err != nil {
		return err
	}

	return nil
}

func (s3 S3Client) DownloadProductStemcell(fa *FileArtifact) (*stemcell, error) {
	return nil, errors.New("downloading stemcells for s3 is not supported at this time")
}

func (s *S3Client) ListFiles() ([]string, error) {
	location, err := s.stower.Dial("s3", s.Config)
	if err != nil {
		return nil, err
	}
	container, err := location.Container(s.bucket)
	if err != nil {
		return nil, err
	}

	var paths []string
	err = s.stower.Walk(container, stow.NoPrefix, 100, func(item stow.Item, err error) error {
		if err != nil {
			return err
		}
		paths = append(paths, item.ID())
		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (s *S3Client) DownloadFile(filename string) (io.ReadCloser, error) {
	location, err := s.stower.Dial("s3", s.Config)
	if err != nil {
		return nil, err
	}
	container, err := location.Container(s.bucket)
	if err != nil {
		return nil, err
	}
	item, err := container.Item(filename)
	if err != nil {
		return nil, err
	}

	return item.Open()
}
