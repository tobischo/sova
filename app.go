package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// App struct
type App struct {
	ctx       context.Context
	sleepTime *time.Time
	abortChan chan any
	mutex     sync.Mutex
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		abortChan: make(chan any),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CurrentTime() string {
	now := time.Now().Local()
	return now.Format("15:04:05")
}

func (a *App) RemainingTime() string {
	if a.sleepTime == nil {
		return ""
	}

	now := time.Now().Local()

	remainingDuration := a.sleepTime.Sub(now)

	nullTime := time.Time{}

	return nullTime.Add(remainingDuration).Format("15:04:05")
}

func (a *App) ScheduleDuration(durationStr string) string {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Sprintf("Failed to parse duration '%s': %v", durationStr, err)
	}

	a.CancelSchedule()

	now := time.Now().Local()
	t := now.Add(duration)
	a.sleepTime = &t

	go a.triggerSleepDelayed(duration)

	return fmt.Sprintf("Scheduled for %s", t.Format("15:04:05"))
}

func (a *App) IsScheduled() bool {
	if a.sleepTime == nil {
		return false
	}

	return true
}

func (a *App) CancelSchedule() string {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.IsScheduled() {
		a.sleepTime = nil
		a.abortChan <- true
		<-time.After(20 * time.Millisecond)
	}

	return fmt.Sprintf("Scheduled sleep has been canceled")
}

func (a *App) triggerSleepDelayed(duration time.Duration) {
	select {
	case <-time.After(duration):
		defer a.CancelSchedule()

		if a.sleepTime != nil && a.sleepTime.Before(time.Now()) {
			a.sleepTime = nil
			a.executeSleep()
		}
	case <-a.abortChan:
	}
}
