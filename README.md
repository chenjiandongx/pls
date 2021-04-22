<p align="center">
    <img src="https://user-images.githubusercontent.com/19553554/61995478-bd21e980-b0bb-11e9-8206-5a5958e27b25.png" alt="Linux logo" width=180 />
</p>

<h1 align="center">📝 pls</h1>
<p align="center">
    <em>Impressive Linux commands cheat sheet cli.</em>
</p>

### 💡 IDEA

Linux 是每位开发者必备的技能，如何高效地掌握 Linux 命令就成为一件很重要的事了。[jaywcjlove/linux-command](https://github.com/jaywcjlove/linux-command) 项目收集和整理了 500+ 的 Linux 命令使用文档，不过缺少了一个命令行版本，`pls` 决定来填补这个空缺。

* Python 版本: [chenjiandongx/how](https://github.com/chenjiandongx/how)

### 🔰 安装

* 使用 `go get` 安装 
    ```shell
    $ go get -u github.com/chenjiandongx/pls
    ```

* 使用编译好的二进制版本: [releases](https://github.com/chenjiandongx/pls/releases)

### 📏 使用

```shell
~ 🐶 pls --help
Impressive Linux commands cheat sheet cli.

Usage:
  pls [command]

Available Commands:
  help        Help about any command
  search      Search command
  show        Show the specified command usage.
  upgrade     Upgrade all commands from remote.
  version     Prints the version of pls

Flags:
  -h, --help   help for pls

Use "pls [command] --help" for more information about a command.
```

### 🔖 示例

> Note: 建议第一次使用的时候先初始化所有命令，可以使用 -d 指定命令文档文件夹下载位置
```shell
$ pls upgrade
```

> Tip: 可以将输出结果传入到 less 管道
```shell
$ pls show curl | less
```

效果图

![image](https://user-images.githubusercontent.com/19553554/72659887-52fe5780-3a01-11ea-89b2-dfaf9faf8dac.png)


#### 辅助脚本

写一个脚本, source到环境中. 方便指定文件目录.

```shell
_cman() {
    pls show -d ~/command $1 | less
}
_cmans() {
    pls search -d ~/command $1
}

alias ,cman=_cman
alias ,cmans=_cmans
```

> Tip: 搜索和展示使用

```shell
# 展示
,cman sort

# 搜索
,cmans 排序
```


### 📃 LICENSE

MIT [©chenjiandongx](https://github.com/chenjiandongx)
