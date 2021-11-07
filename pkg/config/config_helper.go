package config

import (
	"fmt"
	"github.com/ncraft-io/ncraft-go/pkg/config/reader"
	"github.com/ncraft-io/ncraft-go/pkg/config/source"
	"github.com/ncraft-io/ncraft-go/pkg/config/source/env"
	"github.com/ncraft-io/ncraft-go/pkg/config/source/file"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var supportedFileSuffixes = map[string]bool{
	"yaml": true,
	"json": true,
	"toml": true,
}

const (
	DevDeployEnvironment   = "dev"
	TestDeployEnvironment  = "test"
	StageDeployEnvironment = "stage"
	ProdDeployEnvironment  = "prod"
)

func init() {
	workPath, _ := os.Getwd()
	sources := readSource(workPath + "/conf", os.Getenv("deploy_env"))

	sources = append(sources, env.NewSource())
	// flag.NewSource(),
	err := Load(sources...)
	if err != nil {
		// panic(err)
	}
}

func readSource(dir string, env string) []source.Source {
	var sources []source.Source

	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() {
			ss := readSource(f.Name(), env)
			sources = append(sources, ss...)
		} else {
			segments := strings.Split(f.Name(), ".")
			suffix := ""
			if len(segments) >= 2 {
				suffix = segments[len(segments)-1]
			}
			if !supportedFileSuffixes[suffix] {
				continue
			}
			p := path.Join(dir, f.Name())
			if len(env) > 0 {
				name := strings.Join(segments[:len(segments)-1], ".")
				if strings.HasSuffix(name, env) {
					sources = append(sources, file.NewSource(file.WithPath(p)))
				}
			} else {
				sources = append(sources, file.NewSource(file.WithPath(p)))
			}
		}
	}

	return sources
}

// Scan values to a go type
func ScanKey(key string, v interface{}) error {
	return Get(key).Scan(v)
}

// Get a value from the config
func GetValue(path ...string) reader.Value {
	if len(path) == 1 {
		segments := strings.Split(path[0], ".")
		return Get(segments...)
	}

	return Get(path...)
}

type watchCloser struct {
	exit chan struct{}
}

func (w watchCloser) Close() error {
	fmt.Println("close")
	w.exit <- struct{}{}
	return nil
}

func WatchFunc(handle func(reader.Value), paths ...string) (io.Closer, error) {
	path := make([]string, 0, len(paths))
	for _, v := range paths {
		path = append(path, strings.Split(v, ".")...)
	}

	exit := make(chan struct{})
	w, err := Watch(path...)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			v, err := w.Next()
			//if err == err_code.WatchStoppedError {
			//	return
			//}
			if err != nil {
				continue
			}

			//if v.Empty() {
			//	continue
			//}

			if handle != nil {
				handle(v)
			}
		}
	}()

	go func() {
		select {
		case <-exit:
			_ = w.Stop()
		}
	}()

	return watchCloser{exit: exit}, nil
}
