package lib

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func RunCmd(execPath string, args []string, folder string) (stdOut []byte, err error) {

	execCmd := exec.Command(execPath, args...)
	execCmd.Dir = folder

	stdOut, err = execCmd.Output()

	return
}

// FileExists
// return true if a file or dir exists
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func FileParse[K any](file string) (parseTo K, err error) {

	if !FileExists(file) {
		err = fmt.Errorf("file does not exist - %s", file)
		return
	}

	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(fileBytes, &parseTo)

	return
}

// CreateUUID create a UUID with the given prefix
func CreateUUID(prefix string) string {
	return prefix + uuid.New().String()
}

func GetFileList(dir string) (files []string, err error) {

	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dirFiles {
		if !file.IsDir() {
			files = append(files, filepath.Join(dir, file.Name()))
		}
	}

	return
}
