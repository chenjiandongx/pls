package cmd

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/spf13/cobra"
)

const (
	maxCurrency = 8
	maxRetry    = 5
)

func NewUpgradeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade all commands from remote.",
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ := cmd.Flags().GetString("directory")
			upgradeCmd(dir)
		},
	}
	cmd.Flags().StringP("directory", "d", "", "specify the command files directory (absolute path).")
	return cmd
}

func upgradeCmd(dir string) {
	var num int64
	l := len(commands)
	ch := make(chan string, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < maxCurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for cmd := range ch {
				retryDownloadCmd(cmd, dir)
				atomic.AddInt64(&num, 1)
				fmt.Printf("[busy working]: upgrade command:<%d/%d> => %s\n", atomic.LoadInt64(&num), l, cmd)
			}
		}()
	}

	for _, c := range commands {
		ch <- c
	}
	close(ch)
	wg.Wait()
	fmt.Println("[clap]: all commands are upgraded.")
}
