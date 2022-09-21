package lib

import "time"

type LightHouseResult struct {
	Id      string
	Url     string
	TimeRun time.Duration
	RunOn   time.Time
	Resp    LightHouseResp
}
