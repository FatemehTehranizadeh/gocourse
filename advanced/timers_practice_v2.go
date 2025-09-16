package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cpuMonitor struct {
	ticker        time.Ticker
	alertTimerCh  <-chan time.Time
	alertTimer    time.Timer
	isAlertActive bool
}

func newCpuMonitor() *cpuMonitor {
	return &cpuMonitor{
		ticker: *time.NewTicker(time.Second * 2),
	}
}

func (m *cpuMonitor) Run() {
	for {
		select {
		case <-m.ticker.C:
			m.checkCpuUsage()
		case <-m.alertTimerCh:
			fmt.Println("ALERT: High CPU usage sustained for 10 seconds!")
			m.isAlertActive = false
			m.alertTimerCh = nil
		}
	}
}

func (m *cpuMonitor) checkCpuUsage() {
	cpuUsage := rand.Intn(100)
	fmt.Printf("CPU Usage: %d%%\n", cpuUsage)

	if cpuUsage > 80 {
		if !m.isAlertActive {
			timer := time.NewTimer(10 * time.Second)
			m.isAlertActive = true
			m.alertTimerCh = timer.C
			fmt.Println("Timer is started.")

		}
	} else {
		if m.isAlertActive {
			fmt.Println("CPU usage normalized. Cancelling alert.")
			m.alertTimerCh = nil
			m.isAlertActive = false
		}
	}

}

func main() {
	monitor := newCpuMonitor()
	monitor.Run()
	time.Sleep(20 * time.Second)
	fmt.Println("Monitoring stopped.")
}
