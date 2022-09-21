package lib

import (
	"sync"
	"time"
)

type LightHouseStats struct {
	TotalResults int64
	AvgRunTime   time.Duration
	JsLibs       map[string]int64

	mutex sync.Mutex
}

func InitLightHouseStats() (stats *LightHouseStats) {
	stats = &LightHouseStats{
		TotalResults: 0,
		AvgRunTime:   0,
		JsLibs:       map[string]int64{},
		mutex:        sync.Mutex{},
	}
	return
}

func (stats *LightHouseStats) AddResult(result LightHouseResult) {
	stats.mutex.Lock()
	defer stats.mutex.Unlock()

	stats.TotalResults++
	stats.AvgRunTime = (stats.AvgRunTime*time.Duration(stats.TotalResults-int64(1)) + result.TimeRun) / time.Duration(stats.TotalResults)

	for _, item := range result.Resp.Audits.JsLibraries.Details.Items {
		mapIdx := item.Npm + "@" + item.Version

		_, exists := stats.JsLibs[mapIdx]
		if !exists {
			stats.JsLibs[mapIdx] = 0
		}
		stats.JsLibs[mapIdx] = stats.JsLibs[mapIdx] + 1
	}
}
