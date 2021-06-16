
# pls

> Impressive Linux commands cheat sheet cli. [Python 版本](https://github.com/chenjiandongx/how)

### Installation

* 使用 `go get` 安装 
    ```shell
    $ go get -u github.com/chenjiandongx/pls
    ```

* 使用编译好的二进制版本: [releases](https://github.com/chenjiandongx/pls/releases)

### Usages

```shell
~ 🐶 pls --help
Impressive Linux commands cheat sheet cli.

Usage:
  pls [command]

Available Commands:
  help        Help about any command
  search      Search command by keywords
  show        Show the specified command usage.
  upgrade     Upgrade all commands from remote.
  version     Prints the version of pls

Flags:
  -h, --help   help for pls

Use "pls [command] --help" for more information about a command.
```

### Examples

> Note: 建议第一次使用的时候先初始化所有命令，可以使用 -d 指定命令文档文件夹下载位置
```shell
$ pls upgrade
```

> Tip: 可以将输出结果传入到 less 管道
```shell
$ pls show curl | less
```

效果图

![](https://user-images.githubusercontent.com/19553554/122259619-f1e3f780-cf04-11eb-949e-763d82a4e3b9.png)
![](https://user-images.githubusercontent.com/19553554/122258451-a0873880-cf03-11eb-865f-067416787cb7.png)


### LICENSE

MIT [©chenjiandongx](https://github.com/chenjiandongx)
