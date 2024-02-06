package main

// set GOOS=linux
// set GOOS=windows
// set GOARCH=amd64 go build
// ./parseconf reposyncDirr insert_position_bblayer.conf insert_position_local.conf
// ./sep reposyncDir/tmp_dir/deploy/imgDir/spdx/
import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

// INHERIT += "create-spdx-2.2"
// SPDX_PRETTY = "1"
// SPDX_ARCHIVE_PACKAGED = "1"
// SPDX_INCLUDE_SOURCES = "1"
// SPDX_ARCHIVE_SOURCES = "1"
func main() {
	fmt.Println(strings.Repeat("=", 10), "Start of file", path.Base(os.Args[0]), strings.Repeat("=", 10))
	var wg sync.WaitGroup
	wg.Add(1)

	inheritText := "INHERIT += \"create-spdx-2.2\"" + "\n" + "SPDX_PRETTY = \"1\"" + "\n" + "SPDX_ARCHIVE_PACKAGED = \"1\"" + "\n" + "SPDX_INCLUDE_SOURCES = \"1\"" + "\n" + "SPDX_ARCHIVE_SOURCES = \"1\""

	workDir := "workdir"
	confDir := workDir + string(os.PathSeparator) + "conf"
	localConfFile := confDir + string(os.PathSeparator) + "local.conf"

	var folderPath string
	var posToInsertLocalConfFile int
	if len(os.Args) == 3 {
		folderPath = os.Args[1]
		posToInsertLocalConfFile, _ = strconv.Atoi(os.Args[2])
	} else {
		fmt.Println("Please provide reposync folder to process that contains workdir , insert position (starts from 0) to local conf file as an argument")
		fmt.Println("check https://git.yoctoproject.org/poky/tree/meta/classes/create-spdx-2.2.bbclass")
		os.Exit(1)
	}

	go func(folderPath string) {
		defer wg.Done()
		localConfFile = folderPath + string(os.PathSeparator) + localConfFile
		// check if localconffile exists
		if _, err := os.Stat(localConfFile); errors.Is(err, fs.ErrNotExist) {
			fmt.Println(localConfFile, "doesnt exist")
			os.Exit(1)
		}

		if !IsExist("create-spdx-2.2", localConfFile) {
			InsertStringToFile(localConfFile, inheritText+"\n", posToInsertLocalConfFile)
		} else {
			fmt.Println("Skipping modifying", localConfFile)
		}
	}(folderPath)

	wg.Wait()
	fmt.Println(strings.Repeat("=", 10), "End of file", path.Base(os.Args[0]), strings.Repeat("=", 10))
}
