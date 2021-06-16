package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// 结构使用 https://raw.githubusercontent.com/jaywcjlove/linux-command/master/dist/data.json
type commandData struct {
	Name        string `json:"n"`
	Path        string `json:"p"`
	Description string `json:"d"`
}

func NewSearchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search <command>",
		Short: "Search command",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("[sorry]: the search command does not accept any arguments")
				return
			}
			dir, _ := cmd.Flags().GetString("directory")
			rebuild, _ := cmd.Flags().GetBool("rebuild")
			index, _ := cmd.Flags().GetString("index")
			showSearch(args[0], dir, rebuild, index)
		},
	}
	cmd.Flags().StringP("directory", "d", "", "specify the command files directory (absolute path).")
	cmd.Flags().BoolP("rebuild", "r", false, "rebuild index")
	cmd.Flags().StringP("index", "i", "", "index file path")
	return cmd
}

func showSearch(text, dir string, rebuildIndex bool, indexPath string) {
	text = strings.ToLower(text)
	d := commandPath
	if dir != "" {
		d = dir
	}
	index, err := readAndSaveIndex(d, rebuildIndex, indexPath)
	if err != nil {
		fmt.Println("read or save index error:", err)
	}
	for k, v := range index {
		if strings.Contains(strings.ToLower(v.Description), text) || strings.Contains(strings.ToLower(k), text) {
			fmt.Println(k, ":", v.Description)
		}
	}
}

func readAndSaveIndex(d string, rebuildIndex bool, indexPath string) (index map[string]commandData, err error) {
	if indexPath == "" {
		indexPath = path.Join(d, "command_index.json")
	}
	index = make(map[string]commandData)

	if isFileExist(indexPath) {
		if c, err := ioutil.ReadFile(indexPath); err == nil {
			if err := json.Unmarshal(c, &index); err != nil {
				rebuildIndex = true
			}
		} else {
			rebuildIndex = true
		}
	} else {
		rebuildIndex = true
	}

	if rebuildIndex || len(index) == 0 {
		index, err = buildIndex(d)
		if len(index) > 0 && err == nil {
			text, _ := json.MarshalIndent(index, "", "  ")
			err = ioutil.WriteFile(indexPath, text, 0666)
		}
	}
	return
}

// 解析方式,
// 第一行命令名称
// `===`后第一行是简介
func parseCommand(root, path string) (name string, command commandData) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	p := strings.TrimPrefix(path, root)
	command.Path = p
	nameFlag := true
	introductionFlag := false
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			return
		}
		l := strings.TrimSpace(line)
		if len(l) == 0 {
			continue
		}
		if nameFlag {
			name = l
			command.Name = l
			nameFlag = false
			continue
		}
		if l == "===" {
			introductionFlag = true
			continue
		}
		if introductionFlag {
			command.Description = l
			return
		}
	}
}

func buildIndex(root string) (map[string]commandData, error) {
	commandIndex := make(map[string]commandData)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			name, introduction := parseCommand(root, path)
			if name != "" {
				commandIndex[name] = introduction
			}
		}
		return err
	})
	if err == nil && len(commandIndex) == 0 {
		return commandIndex, fmt.Errorf("no command file, please 'upgrade' or '-d path'")
	}
	return commandIndex, err
}
