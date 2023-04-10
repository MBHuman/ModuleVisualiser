package visualiser

import "regexp"

type FilePattern struct {
	Pattern string
	Regex   *regexp.Regexp
}

func NewFilePattern(pattern string) (*FilePattern, error) {
	regexp, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &FilePattern{Pattern: pattern, Regex: regexp}, nil
}
