package utils

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_logus"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

func GetCurrentFile() utils_types.FilePath {
	_, filename, _, _ := runtime.Caller(1)
	return utils_types.FilePath(filename)
}

func GetCurrrentFolder() utils_types.FilePath {
	_, filename, _, _ := runtime.Caller(1)
	directory := filepath.Dir(filename)
	return utils_types.FilePath(directory)
}

func GetProjectDir() utils_types.FilePath {
	path, err := os.Getwd()
	if folder_override, ok := os.LookupEnv("AUTOGIT_PROJECT_FOLDER"); ok {
		path = folder_override
	}
	utils_logus.Log.CheckFatal(err, "unable to get workdir")
	return utils_types.FilePath(path)
}
