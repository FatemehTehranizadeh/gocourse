// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	// Create a ticker that "ticks" every 1 second.
// 	// Each tick sends the current time into ticker.C.
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()
// 	expireTimer := time.NewTimer(6 * time.Second)
// 	defer expireTimer.Stop()

// 	// The channel for our alert timer. Starts as nil (inactive).
// 	// When nil, it will never trigger in select.
// 	var alertTimerCh <-chan time.Time

// 	// Boolean flag: do we currently have an alert timer running?
// 	alertTimerActive := false

// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				// Every second, check CPU usage.
// 				checkCpuUsage(&alertTimerActive, &alertTimerCh)

// 			case <-expireTimer.C:
// 				// After 20 seconds, stop everything.
// 				// Note: this is a NEW timer created each loop!
// 				// (That's why your old version with 6s never fired:
// 				// it kept creating a fresh timer before it expired.)
// 				fmt.Println("The operations are stopped.")
// 				return

// 			case <-alertTimerCh:
// 				// If our alert timer fires, it means CPU has been
// 				// high for 10 seconds straight.
// 				fmt.Println("ALERT: High CPU usage sustained for 5 seconds!")
// 				alertTimerActive = false
// 				alertTimerCh = nil // disable the timer again
// 			}
// 		}
// 	}()

// 	// Block main goroutine forever so program doesn’t exit
// 	// select {}
// 	time.Sleep(time.Second * 10)
// }

// func checkCpuUsage(alertTimerActive *bool, alertTimerCh *<-chan time.Time) {
// 	// Random CPU usage between 0 and 99.
// 	cpuusage := rand.Intn(100)
// 	fmt.Printf("CPU Usage: %d%%\n", cpuusage)

// 	if cpuusage > 80 {
// 		// High usage detected
// 		fmt.Println("High usage of CPU detected:", cpuusage)

// 		// Only start the alert timer if it’s not already running.
// 		// !*alertTimerActive means "if not active yet".
// 		if !*alertTimerActive {
// 			alertTimer := time.NewTimer(5 * time.Second)
// 			fmt.Println("Timer has been started.")
// 			*alertTimerCh = alertTimer.C
// 			*alertTimerActive = true
// 		}

// 	} else {
// 		// CPU went back below 80. If we had a timer, cancel it.
// 		if *alertTimerActive {
// 			fmt.Println("CPU usage normalized, stopping timer alert...")
// 			*alertTimerActive = false
// 			*alertTimerCh = nil
// 		}
// 	}
// }

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Monitor struct {
	ticker           *time.Ticker
	alertTimer       *time.Timer
	alertTimerCh     <-chan time.Time
	alertTimerActive bool
}

func NewMonitor() *Monitor {
	return &Monitor{
		ticker: time.NewTicker(time.Second * 2),
	}
}

func (m *Monitor) Run() {
	for {
		select {
		case <-m.ticker.C:
			m.checkCpuUsage()

		case <-m.alertTimerCh:
			fmt.Println("ALERT: High CPU usage sustained for 10 seconds!")
			m.alertTimerActive = false
			m.alertTimerCh = nil
		}
	}
}

func (m *Monitor) checkCpuUsage() {
	usage := rand.Intn(100)
	fmt.Printf("CPU Usage: %d%%\n", usage)

	if usage > 80 {
		fmt.Println("High usage detected:", usage)
		if !m.alertTimerActive {
			m.alertTimer = time.NewTimer(10 * time.Second)
			m.alertTimerCh = m.alertTimer.C
			m.alertTimerActive = true
			fmt.Println("Alert timer started. ")
		}
	} else {
		if m.alertTimerActive {
			fmt.Println("CPU usage normalized. Cancelling alert.")
			m.alertTimer.Stop()
			m.alertTimerActive = false
			m.alertTimerCh = nil
		}
	}
}

func main() {
	monitor := NewMonitor()
	go monitor.Run()

	// stop after 20s for demo
	time.Sleep(20 * time.Second)
	fmt.Println("Monitoring stopped.")
}
