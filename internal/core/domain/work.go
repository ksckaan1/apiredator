package domain

import "time"

type Stat struct {
	Completed        uint64         `json:"completed"`
	RequestPerSecond uint64         `json:"request_per_second"`
	StatusCodes      map[int]uint64 `json:"status_codes"`
	StartedAt        time.Time      `json:"started_at"`
	EndedAt          time.Time      `json:"ended_at"`
	Duration         time.Duration  `json:"duration"`
}
