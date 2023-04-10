package visualiser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileMod struct {
	filePath    string
	deps        map[string][]string
	requirement *Requirement
}

// Create new fileMod object, automaticaly extracts data from file
// and build requirement graph
func NewFileMod(filePath string) (*FileMod, error) {
	fileMod := &FileMod{
		filePath:    filePath,
		deps:        make(map[string][]string),
		requirement: NewRequirement(),
	}
	if err := fileMod.extractDataFromFile(); err != nil {
		return nil, err
	}
	if err := fileMod.buildRequirements(); err != nil {
		return nil, err
	}
	return fileMod, nil
}

// Extract data from file
func (fileMod *FileMod) extractDataFromFile() error {
	file, err := os.Open(fileMod.filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	isDirect := true

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
					fileMod.deps[pkg] = nil
				} else {
					// The line contains a dependency of a package
					dep := strings.TrimSpace(line)
					lastPkg := "indirect"
					if isDirect {
						lastPkg = "direct"
					}
					fileMod.deps[lastPkg] = append(fileMod.deps[lastPkg], dep)
				}
			}
			isDirect = !isDirect
		}
	}

	return nil
}

// get All elements of fileMode.requirement
func (fileMode *FileMod) getChilds() (map[string][]string, error) {
	elements := make([]*RequirementNode, 0, CAP_SIZE)
	childs := make(map[string][]string, CAP_SIZE)
	elements = append(elements, fileMode.requirement.root)

	// classical BFS
	for len(elements) > 0 {
		cur := elements[0]
		elements = elements[1:]
		if len(cur.childs) > 0 {
			childs[cur.url] = make([]string, 0, CAP_SIZE)
		}
		for _, child := range fileMode.requirement.elementsSet[cur.url].childs {
			childs[cur.url] = append(childs[cur.url], child.url)
			elements = append(elements, child)
		}
	}

	return childs, nil
}

// build Requirements tree from requirements in file
// :TODO remove pkg != "indirect" and add indirect deps dependency on direct deps
func (fileMod *FileMod) buildRequirements() error {
	root := fileMod.requirement
	for pkg, deps := range fileMod.deps {
		if pkg != "indirect" {
			for _, dep := range deps {
				if len(strings.Split(dep, " ")) == 0 {
					return fmt.Errorf("DEP STRING HAS LENGTH 0")
				}
				root.AddSingleRequirement(strings.TrimSpace(strings.Split(dep, " ")[0]))
			}
		}
	}
	return nil
}
