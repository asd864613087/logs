package utils

import (
	"fmt"
	"github.com/asd864613087/logs/consts"
	"os"
	"strings"
)

func GetPsm() string {
	hostName := os.Getenv("HOSTNAME")
	if hostName == "" || len(hostName) == 0 {
		return ""
	}
	psm := strings.Split(hostName,"-")[0]
	return psm
}

func GetUnixPath() string {
	hostName := os.Getenv("HOSTNAME")
	if hostName == "" || len(hostName) == 0 {
		file := consts.DEFAULT_LOGAGENT_UNIX_PATH_TEST
		if !isFileExist(file) {
			createDir(file)
		}
		return file
	}

	file := fmt.Sprintf(consts.DEFAULT_LOGAGENT_UNIX_PATH_K8s, hostName)
	if !isFileExist(file) {
		createDir(file)
	}

	return file

}

func isFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func createDir(path string)  {
	idx := strings.LastIndex(path, "/")
	err := os.MkdirAll(path[:idx], os.ModePerm)
	if err != nil {
		panic(err)
	}
}
