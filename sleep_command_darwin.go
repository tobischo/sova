//go:build darwin

package main

import "os/exec"

func (a *App) executeSleep() {
	cmd := exec.CommandContext(a.ctx, "pmset", "sleepnow")
	_ = cmd.Run()
}
