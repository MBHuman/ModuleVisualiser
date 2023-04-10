package visualiser

import (
	"os"
	"path/filepath"
)

type IVisualiser interface {
	Visualise() error
	ReadFiles(files []string) error
	ReadFolder(path string) error
}

type Visualiser struct {
	files       []string
	fileModules map[string]*FileModules
}

func (v Visualiser) Visualise() error {
	return nil
}

func (v *Visualiser) ReadFiles(files []string) {
	v.files = files
}

func (v *Visualiser) ReadFolder(path string, pattern *FilePattern) error {
	v.files = make([]string, 0, 10)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && pattern.Regex.MatchString(info.Name()) {
			v.files = append(v.files, info.Name())
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

type VisualiserConosle struct {
	Visualiser Visualiser
}

func (v VisualiserConosle) Visualise() error {
	return nil
}

type VisualiserPUML struct {
	Visualiser Visualiser
}

func (v VisualiserPUML) Visualise() error {
	return nil
}
