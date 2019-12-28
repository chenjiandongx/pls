package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	commandDir  = ".commands"
	commandPath = filepath.Join(getHomedir(), commandDir)
	commandUrl  = "https://raw.githubusercontent.com/jaywcjlove/linux-command/master/command/%s.md"
)

func NewShowCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show <command>",
		Short: "Show the specified command usage.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("[sorry]: the show command does not accept any arguments")
				return
			}
			force, _ := cmd.Flags().GetBool("force")
			showCmd(args[0], force)
		},
	}
	cmd.Flags().BoolP("force", "f", false, "force to refresh command usage from remote.")
	return cmd
}

func showCmd(cmd string, force bool) {
	cmd = strings.ToLower(cmd)
	if force {
		retryDownloadCmd(cmd)
	}

	p := path.Join(getHomedir(), commandDir, fmt.Sprintf("%s.md", cmd))
	if !isFileExist(p) {
		status, err := retryDownloadCmd(cmd)
		if err != nil {
			fmt.Printf("[sorry]: fetch command <%s> error\n", cmd)
			return
		}
		if status == http.StatusNotFound {
			fmt.Printf("[sorry]: could not found command <%s>\n", cmd)
			return
		}
	}
	source, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Printf("[sorry]: open file <%s> error\n", p)
		return
	}
	markdown.BlueBgItalic = color.New(color.FgBlue).SprintFunc()
	result := markdown.Render(string(source), 80, 6)
	fmt.Println(string(result))
}

func getHomedir() string {
	home, _ := homedir.Expand("~")
	return home
}

func makeCmdDir() error {
	if _, err := os.Stat(commandPath); err != nil && !os.IsExist(err) {
		return os.Mkdir(commandPath, 0755)
	}
	return nil
}

func isFileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func retryDownloadCmd(cmd string) (int, error) {
	var err error
	var status int
	for j := 0; j < maxRetry; j++ {
		if err, status = downloadCmd(cmd); err != nil {
			continue
		}
		break
	}
	return status, err
}

func downloadCmd(cmd string) (error, int) {
	if err := makeCmdDir(); err != nil {
		return err, 0
	}

	resp, err := http.Get(fmt.Sprintf(commandUrl, cmd))
	if err != nil {
		return err, 0
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, http.StatusNotFound
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, 0
	}

	fp := path.Join(commandPath, fmt.Sprintf("%s.md", cmd))
	return ioutil.WriteFile(fp, content, 0666), 0
}
