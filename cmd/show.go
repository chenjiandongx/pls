package cmd

import (
	"bufio"
	"fmt"
	"io"
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
	commandUrl  = "https://unpkg.com/linux-command/command/%s.md"
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
			dir, _ := cmd.Flags().GetString("directory")
			showCmd(args[0], dir, force)
		},
	}
	cmd.Flags().BoolP("force", "f", false, "force to refresh command usage from remote.")
	cmd.Flags().StringP("directory", "d", "", "specify the command files directory (absolute path).")
	return cmd
}

func showCmd(cmd, dir string, force bool) {
	cmd = strings.ToLower(cmd)
	d := commandPath
	if dir != "" {
		d = dir
	}

	if force {
		retryDownloadCmd(cmd, d)
	}

	p := path.Join(d, fmt.Sprintf("%s.md", cmd))
	if !isFileExist(p) {
		status, err := retryDownloadCmd(cmd, d)
		if err != nil {
			fmt.Printf("[sorry]: failed to retrieve command <%s>\n", cmd)
			return
		}
		if status == http.StatusNotFound {
			fmt.Printf("[sorry]: could not found command <%s>\n", cmd)
			return
		}
	}
	source, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Printf("[sorry]: failed to open file <%s>\n", p)
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

func makeCmdDir(dir string) error {
	if _, err := os.Stat(dir); err != nil && !os.IsExist(err) {
		return os.Mkdir(dir, 0755)
	}
	return nil
}

func isFileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func retryDownloadCmd(cmd, dir string) (int, error) {
	var err error
	var status int
	for j := 0; j < maxRetry; j++ {
		if err, status = downloadCmd(cmd, dir); err != nil {
			continue
		}
		break
	}
	return status, err
}

func downloadCmd(cmd, dir string) (error, int) {
	d := commandPath
	if dir != "" {
		d = dir
	}

	if err := makeCmdDir(d); err != nil {
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

	content := make([]byte, 0)
	reader := bufio.NewReader(resp.Body)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			return err, 0
		}
		if err == io.EOF {
			break
		}
		content = append(content, line...)
		content = append(content, []byte("\n")...)
	}

	fp := path.Join(d, fmt.Sprintf("%s.md", cmd))
	return ioutil.WriteFile(fp, content, 0666), 0
}
