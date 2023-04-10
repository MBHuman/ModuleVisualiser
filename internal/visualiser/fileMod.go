package visualiser

import (
	"bufio"
	"os"
	"strings"
)

type FileMod struct {
	filePath        string
	deps            map[string][]string
	requirementRoot *Requirement
}

func NewFileMod(filePath string) (*FileMod, error) {
	fileMod := &FileMod{
		filePath:        filePath,
		deps:            make(map[string][]string),
		requirementRoot: NewRequirement(),
	}
	if err := fileMod.extractDataFromFile(); err != nil {
		return nil, err
	}
	if err := fileMod.getRequirements(); err != nil {
		return nil, err
	}
	return fileMod, nil
}

func (f *FileMod) extractDataFromFile() error {
	file, err := os.Open(f.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "require (") {
			// We've found the dependencies section
			for scanner.Scan() {
				line := scanner.Text()
				if line == ")" {
					// We've reached the end of the dependencies section
					break
				}
				if !strings.HasPrefix(line, "\t") {
					// The line contains the package name and version
					pkg := strings.Split(line, " ")[0]
					f.deps[pkg] = nil
				} else {
					// The line contains a dependency of a package
					dep := strings.TrimSpace(line)
					lastPkg := strings.Split(dep, " ")[len(strings.Split(dep, " "))-1]
					f.deps[lastPkg] = append(f.deps[lastPkg], dep)
				}
			}
		}
	}

	return nil
}

func (f *FileMod) getRequirements() error {

	return nil
}
