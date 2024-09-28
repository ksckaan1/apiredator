package models

import "time"

type Stat struct {
	SentCount      uint64         `json:"sent_count"`
	RPS            RPS            `json:"rps"`
	StatusCodes    map[int]uint64 `json:"status_codes"`
	StartedAt      time.Time      `json:"started_at"`
	EndedAt        time.Time      `json:"ended_at"`
	PassedDuration string         `json:"passed_duration"`
	ResponseInfo   *ResponseInfo
}

type RPS struct {
	List   []uint64 `json:"list"`
	Latest uint64   `json:"latest"`
	Min    uint64   `json:"min"`
	Avg    float64  `json:"avg"`
	Max    uint64   `json:"max"`
}

type ResponseInfo struct {
	Durations []time.Duration
}
