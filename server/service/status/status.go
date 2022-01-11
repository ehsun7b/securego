package status

import "time"

type status struct {
	FirstRun bool      `json: firstRun`
	LastRun  time.Time `json: lastRun`
}
