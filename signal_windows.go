// +build windows

package prompt

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (p *Prompt) handleSignals(exitCh chan int, winSizeCh chan *WinSize) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(
		sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	for {
		s, ok := <-sigCh
		if !ok {
			log.Println("[INFO] stop handleSignals")
			signal.Stop(sigCh)
			return
		}

		switch s {
		case syscall.SIGINT: // kill -SIGINT XXXX or Ctrl+c
			log.Println("[SIGNAL] Catch SIGINT")
			exitCh <- ExitSuccess

		case syscall.SIGTERM: // kill -SIGTERM XXXX
			log.Println("[SIGNAL] Catch SIGTERM")
			exitCh <- ExitKilled

		case syscall.SIGQUIT: // kill -SIGQUIT XXXX
			log.Println("[SIGNAL] Catch SIGQUIT")
			exitCh <- ExitSuccess
		}
	}
}
