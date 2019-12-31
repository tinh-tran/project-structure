package router_defined

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	global "school_project/internal/app/init_global"
	"school_project/internal/app/models"
	"school_project/internal/pkg/config"
)

func ReadRouterFile() []models.GroupRoute {
	var files []string
	var baseDir, _ = os.Getwd()
	ctx := config.PathSaveFile().RouterFile
	root := fmt.Sprintf("%s/%s", baseDir, ctx+"internal/app/router/router_defined/")

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".json" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Print(err)
	}
	var routers []models.GroupRoute
	for _, file := range files {
		routers = append(routers, readFile(file)...)
	}
	global.RoutersArr = routers
	return routers
}

func readFile(file string) []models.GroupRoute {
	var routers []models.GroupRoute
	jsonFile, _ := os.Open(file)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal(byteValue, &routers)
	return routers
}

func IsPublicAPI(handleName string) bool {
	for _, val := range global.RoutersArr {
		for _, valR := range val.Data {
			if handleName == valR.Handler {
				if valR.PublicApi == true {
					return true
				}
			}
		}
	}
	return false
}
