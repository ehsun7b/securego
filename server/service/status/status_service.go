package status

import "time"

func Status() status {
	return status{FirstRun: true, LastRun: time.Now()}
}

func Update() {
	return
}
