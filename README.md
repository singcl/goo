### 编译输出到指定目录

\$ go build -o ./exe/values.exe ./values.go

### 实例来源

https://gobyexample.com/

### About Learning Go

https://www.jianshu.com/p/5ee783bbdcac

### VScode Go 环境搭建

https://www.cnblogs.com/Leo_wl/p/8242628.html

### <<GO 入门指南>>

https://github.com/Unknwon/the-way-to-go_ZH_CN

### JS 深入系列

https://github.com/mqyqingfeng/Blog

### go

~~貌似 import 不能 import 相对路径--网上很多文章都说可以？？the-way-to-go 实例中也说可以？？~~

~~只能 import 绝对路径--即现在$GOPATH 的 src 下查找再在$GOROOT 的 src 下查找。~~

**/advanced/13.4_panic_package.go 中就使用了相对路径包。**

一个文件夹下只能包含一个自定义的 package，但是一个 package 可以在这个文件夹下分成多个文件。

go install 有 main 包的 go 项目和 没有 main 包的 go 项目

go install 不带参数可以直接在包所在的文件下执行 或者 go install + 包所在的文件夹名字 来安装

### vscode GO 環境安裝 https://blog.csdn.net/wyy626562203/article/details/83833592

### GO MODULES

**wiki: https://github.com/golang/go/wiki/Modules**

### GOPROXY 设置
`golang.org`在国内无法访问，导致很多包下载不了。

解决方法一：
在github上找到镜像仓库clone到$GOPATH/src/golang.org/x/  下面然后 go install 自行安装。

解决方式二：
设置goproxy：https://goproxy.cn/

```sh
go env -w GOPROXY=https://goproxy.cn,direct
echo "export GOPROXY=https://goproxy.cn" >> ~/.bash_profile    (MAC/LINUX)
source ./.bash_profile
```
