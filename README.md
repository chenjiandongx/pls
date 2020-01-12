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

使用 `go get` 安装 
```shell
$ go get -u github.com/chenjiandongx/pls
```

使用编译好的二进制版本
```shell
# https://github.com/chenjiandongx/pls/releases

# linux
$ wget https://github.com/chenjiandongx/pls/releases/download/v0.1.1/pls_linux_amd64

# macos
$ wget https://github.com/chenjiandongx/pls/releases/download/v0.1.1/pls_darwin_amd64

# windows
$ wget https://github.com/chenjiandongx/pls/releases/download/v0.1.1/pls_windows_amd64.exe
```

### 📏 使用

```shell
~ 🐶 pls --help
Impressive Linux commands cheat sheet cli.

Usage:
  pls [command]

Available Commands:
  help        Help about any command
  show        Show the specified command usage.
  upgrade     Upgrade all commands from remote.
  version     Prints the version of pls

Flags:
  -h, --help   help for pls

Use "pls [command] --help" for more information about a command.
```

### 🔖 示例

> Note: 建议第一次使用的时候先初始化所有命令
```shell
$ pls upgrade
```

> Tip: 可以将输出结果传入到 less 管道
```shell
$ pls show curl | less
```

效果图

![image](https://user-images.githubusercontent.com/19553554/71540604-caebdb80-2987-11ea-909c-f1f1488ef226.png)


### 📃 LICENSE

MIT [©chenjiandongx](https://github.com/chenjiandongx)
