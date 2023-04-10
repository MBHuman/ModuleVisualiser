package visualiser

import (
	"os"
	"path/filepath"
)

type Visualiser struct {
	files       []string
	fileModules map[string]*FileMod
}

func NewVisualiser() *Visualiser {
	visualiser := &Visualiser{
		files:       make([]string, 0),
		fileModules: make(map[string]*FileMod),
	}
	visualiser.preprocess()
	return visualiser
}

func (visualiser *Visualiser) preprocess() error {

	return nil
}

func (visualiser *Visualiser) Visualise() error {
	return nil
}

func (visualiser *Visualiser) ReadFiles(files []string) {
	visualiser.files = files
}

func (visualiser *Visualiser) ReadFolder(path string, pattern *FilePattern) error {
	visualiser.files = make([]string, 0, 10)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && pattern.Regex.MatchString(info.Name()) {
			visualiser.files = append(visualiser.files, info.Name())
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

func (visualiser *VisualiserConosle) Visualise() error {
	return nil
}

type VisualiserPUML struct {
	Visualiser Visualiser
}

func (visualiser *VisualiserPUML) Visualise() error {
	return nil
}
