package visualiser

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFileMod_extractDataFromFile(t *testing.T) {
	convey.Convey("Given a FileMod instance with a file containing dependencies", t, func() {
		fileMod := &FileMod{filePath: "testdata/testmod.mod", deps: make(map[string][]string)}

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
					"v1.0.0":  {"github.com/jackpal/bencode-go v1.0.0"},
					"v1.15.0": {"github.com/spf13/viper v1.15.0"},
					"v1.6.1":  {"github.com/spf13/cobra v1.6.1"},
					"v1.8.1":  {"github.com/stretchr/testify v1.8.1"},
				}

				convey.So(fileMod.deps, convey.ShouldResemble, expectedDeps)
			})
		})
	})
}
