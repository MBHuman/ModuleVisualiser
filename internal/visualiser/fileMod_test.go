package visualiser

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFileMod_extractDataFromFile(t *testing.T) {
	convey.Convey("Given a FileMod instance with a file containing dependencies", t, func() {
		fileMod := &FileMod{
			filePath:    "testdata/testmod.mod",
			deps:        make(map[string][]string),
			requirement: NewRequirement(),
		}

		convey.Convey("When extractDataFromFile is called", func() {
			err := fileMod.extractDataFromFile()

			convey.Convey("It should not return an error", func() {
				convey.So(err, convey.ShouldBeNil)
			})

			convey.Convey("It should extract the dependencies correctly", func() {
				expectedDeps := map[string][]string{
					"indirect": {
						"github.com/davecgh/go-spew v1.1.1 // indirect",
						"github.com/fsnotify/fsnotify v1.6.0 // indirect",
						"github.com/hashicorp/hcl v1.0.0 // indirect",
						"github.com/inconshreveable/mousetrap v1.0.1 // indirect",
						"github.com/magiconair/properties v1.8.7 // indirect",
						"github.com/mitchellh/mapstructure v1.5.0 // indirect",
						"github.com/pelletier/go-toml/v2 v2.0.6 // indirect",
						"github.com/pmezard/go-difflib v1.0.0 // indirect",
						"github.com/spf13/afero v1.9.3 // indirect",
						"github.com/spf13/cast v1.5.0 // indirect",
						"github.com/spf13/jwalterweatherman v1.1.0 // indirect",
						"github.com/spf13/pflag v1.0.5 // indirect",
						"github.com/subosito/gotenv v1.4.2 // indirect",
						"golang.org/x/sys v0.3.0 // indirect",
						"golang.org/x/text v0.5.0 // indirect",
						"gopkg.in/ini.v1 v1.67.0 // indirect",
						"gopkg.in/yaml.v3 v3.0.1 // indirect",
					},
					"direct": {
						"github.com/jackpal/bencode-go v1.0.0",
						"github.com/spf13/viper v1.15.0",
						"github.com/spf13/cobra v1.6.1",
						"github.com/stretchr/testify v1.8.1",
					},
				}
				if _, ok := expectedDeps["indirect"]; ok {
					sort.Strings(expectedDeps["indirect"])
				}
				if _, ok := expectedDeps["direct"]; ok {
					sort.Strings(expectedDeps["direct"])
				}

				if _, ok := fileMod.deps["indirect"]; ok {
					sort.Strings(fileMod.deps["indirect"])
				}
				if _, ok := expectedDeps["direct"]; ok {
					sort.Strings(expectedDeps["direct"])
				}

				convey.So(fileMod.deps, convey.ShouldResemble, expectedDeps)
			})
		})
	})
}

func TestFileMod_getRequirements(t *testing.T) {
	convey.Convey("Given a FileMod instance with dependencies", t, func() {
		fileMod := &FileMod{
			filePath:    "testdata/testmod.mod",
			deps:        make(map[string][]string),
			requirement: NewRequirement(),
		}
		fileMod.extractDataFromFile()

		convey.Convey("When getRequirements() is called", func() {
			err := fileMod.buildRequirements()
			convey.Convey("it should not return an error", func() {
				convey.So(err, convey.ShouldBeNil)
			})

		})

		convey.Convey("It should add all non-internal dependecies to the requirement root", func() {
			expectedRequirements := map[string][]string{
				"root": {
					"github.com/jackpal/bencode-go",
					"github.com/spf13/viper",
					"github.com/spf13/cobra",
					"github.com/stretchr/testify",
				},
			}
			if _, ok := expectedRequirements["root"]; ok {
				sort.Strings(expectedRequirements["root"])
			}

			err := fileMod.buildRequirements()
			convey.Convey("it should not return an error", func() {
				convey.So(err, convey.ShouldBeNil)
			})

			actualRequirements, err := fileMod.getChilds()
			if _, ok := actualRequirements["root"]; ok {
				sort.Strings(actualRequirements["root"])
			}

			convey.Convey("It should add all non-internal dependencie to the requirement root", func() {
				convey.So(err, convey.ShouldBeNil)
				convey.So(actualRequirements, convey.ShouldResemble, expectedRequirements)
			})

		})
	})
}

func TestFileModWithZeroHash(t *testing.T) {
	convey.Convey("FileMod struct with zero hash should always return false", t, func() {
		// Create a temporary file for testing
		tmpfile, err := ioutil.TempFile("", "testfile")
		convey.So(err, convey.ShouldBeNil)
		defer os.Remove(tmpfile.Name())

		// Write some content to the file
		content := []byte("This is some test content")
		_, err = tmpfile.Write(content)
		convey.So(err, convey.ShouldBeNil)

		// Create a FileMod struct with hash value 0
		// Create a hash object
		hash := sha256.New()

		// Copy the contents of the file to the hash object
		_, err = io.Copy(hash, tmpfile)
		convey.So(err, convey.ShouldBeNil)

		// Compute the final hash value as a hexadecimal string
		hashValue := hex.EncodeToString(hash.Sum(nil))

		// Close the file
		err = tmpfile.Close()
		convey.So(err, convey.ShouldBeNil)

		fileMod := FileMod{hash: hashValue}

		// Compare the file's hash value
		hashMatch, err := fileMod.Compare(tmpfile.Name())

		convey.Convey("Should always return false", func() {
			convey.So(err, convey.ShouldBeNil)
			convey.So(hashMatch, convey.ShouldBeFalse)
		})
	})
}

func TestFileModUpdateHash(t *testing.T) {
	// Initialize the Convey test suite
	convey.Convey("Given a FileMod object", t, func() {
		fileMod := &FileMod{
			filePath: "test.txt",
		}

		// Create a test file with some contents
		file, err := os.Create(fileMod.filePath)
		convey.So(err, convey.ShouldBeNil)
		defer file.Close()
		file.WriteString("This is a test file")

		convey.Convey("When updateHash is called", func() {
			err := fileMod.updateHash()

			convey.Convey("No error should be returned", func() {
				convey.So(err, convey.ShouldBeNil)
			})

			convey.Convey("The hash should be computed correctly", func() {
				hash := sha256.New()
				file, err := os.Open(fileMod.filePath)
				convey.So(err, convey.ShouldBeNil)
				defer file.Close()
				if _, err := io.Copy(hash, file); err != nil {
					t.Fatal(err)
				}
				expectedHash := hex.EncodeToString(hash.Sum(nil))
				convey.So(fileMod.hash, convey.ShouldEqual, expectedHash)
			})
		})

		// Clean up the test file
		os.Remove(fileMod.filePath)
	})
}
