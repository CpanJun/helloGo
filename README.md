# helloGo

## 开发环境

设备: MacBook Air(M1, 2020)

系统: macOS Monterey 12.6

内容: 16GB
## 环境安装
 从[官方网站](https://go.dev/dl/)下载GO语言安装包:
    
    我下载的是macOS的 install 版本，安装完毕以后还需要在设置一下引用。

    [Mac 下载安装go之后 go version显示zsh: command not found: go](https://blog.csdn.net/qq_44812523/article/details/118529183)

    ```shell
    # 安装完毕后, 通过go version 确认是否安装成功.
    $ go version
    go version go1.15.2 darwin/amd64 # 正确返回版本信息即可
    ```
## 运行
```shell
$ go run helloWorld.go
Hello, World!
```