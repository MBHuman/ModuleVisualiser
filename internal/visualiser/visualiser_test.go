package visualiser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestVisualiser(t *testing.T) {
	// Create a new Visualiser instance
	v := Visualiser{}

	convey.Convey("Given a list of files", t, func() {
		files := []string{"file1.mod", "file2.mod", "file3.mod"}

		convey.Convey("When we call the ReadFiles method", func() {
			v.ReadFiles(files)

			convey.Convey("Then the Visualiser's files should be set to the given list of files", func() {
				convey.So(v.files, convey.ShouldResemble, files)
			})
		})

		convey.Convey("When we call the ReadFolder method", func() {
			// Create a temporary directory for testing
			tmpDir := t.TempDir()
			// Create some test files
			testFiles := []string{"test1.mod", "test2.mod", "test3.xml", "test4.xml"}
			for _, file := range testFiles {
				f, err := os.Create(filepath.Join(tmpDir, file))
				if err != nil {
					t.Fatalf("error creating file: %v", err)
				}
				f.Close()
			}

			// Set up a FilePattern for testing
			pattern, err := NewFilePattern(`\.mod$`)
			convey.So(err, convey.ShouldBeNil)

			err = v.ReadFolder(tmpDir, pattern)

			convey.Convey("Then the Visualiser's files should be set to the list of files in the folder matching the pattern", func() {
				convey.So(err, convey.ShouldBeNil)
				expectedFiles := []string{"test1.mod", "test2.mod"}
				convey.So(v.files, convey.ShouldResemble, expectedFiles)
			})
		})
	})
}
