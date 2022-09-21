package lib

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type LightHouseResults struct {
	results    []string
	mut        sync.Mutex
	OutputPath string
}

func InitLightHouseResults(outputPath string) (res *LightHouseResults) {
	return &LightHouseResults{
		results:    []string{},
		mut:        sync.Mutex{},
		OutputPath: outputPath,
	}
}

func (res *LightHouseResults) LoadAvailableResults() (results []string, err error) {

	files, err := GetFileList(res.OutputPath)
	if err != nil {
		return
	}

	res.results = []string{}

	for _, file := range files {
		fileName := strings.ReplaceAll(file, res.OutputPath+"/", "")
		uuid := strings.ReplaceAll(fileName, ".json", "")
		res.results = append(res.results, uuid)
	}

	results = res.results

	return
}

func (res *LightHouseResults) LoadResult(uuid string) (result LightHouseResult, err error) {

	result, err = FileParse[LightHouseResult](filepath.Join(res.OutputPath, uuid+".json"))

	return
}

func (res *LightHouseResults) dumResponse(result LightHouseResult) (err error) {

	storePath := filepath.Join(res.OutputPath, result.Id+".json")

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return
	}

	err = os.WriteFile(storePath, jsonBytes, 0777)

	return
}

func (res *LightHouseResults) AddResp(url string, timeRun time.Duration, resp LightHouseResp) (result LightHouseResult, err error) {
	res.mut.Lock()
	defer res.mut.Unlock()

	result = LightHouseResult{
		Id:      CreateUUID("result_"),
		Url:     url,
		TimeRun: timeRun,
		RunOn:   time.Now(),
		Resp:    resp,
	}

	err = res.dumResponse(result)
	if err != nil {
		return
	}

	res.results = append(res.results, result.Id)

	return
}

func (res *LightHouseResults) GetResults() (resp []string) {
	res.mut.Lock()
	defer res.mut.Unlock()

	resp = res.results

	return
}
