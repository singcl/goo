## Vscode GO 环境安装

### 1. 一下是安装完 vscode 时需要安装的插件，注意后面的名称

```sh
github.com/mdempsky/gocode  		gocode
github.com/uudashr/gopkgs/cmd/gopkgs	gopkgs
github.com/ramya-rao-a/go-outline  	go-outline
github.com/acroca/go-symbols  		go-symbols
golang.org/x/tools/cmd/guru  		guru
golang.org/x/tools/cmd/gorename  	gorename
github.com/derekparker/delve/cmd/dlv   	dlv
github.com/stamblerre/gocode  		gocode-gomod
github.com/rogpeppe/godef		godef
github.com/ianthehat/godef  		godef-gomod
github.com/sqs/goreturns  		goreturns
golang.org/x/lint/golint  		golint

```

### 2. 自行下载安装不要通过 vscode，应为 golang.org 被和谐了。go get 会自动编译安装文件到 GOPATH/bin 目录下。

新建一个批处理文件，运行以下内容：

```bat
mkdir  %GOPATH%\\src\\golang.org\\x
git clone https://github.com/golang/tools.git %GOPATH%\\src\\golang.org\\x\\tools

go get -v github.com/mdempsky/gocode
go get -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -v github.com/ramya-rao-a/go-outline
go get -v github.com/acroca/go-symbols
go get -v golang.org/x/tools/cmd/guru
go get -v golang.org/x/tools/cmd/gorename
go get -v github.com/derekparker/delve/cmd/dlv
go get -v github.com/stamblerre/gocode
go get -v github.com/rogpeppe/godef
go get -v github.com/ianthehat/godef
go get -v github.com/sqs/goreturns
%go get -v github.com/golang/lint%
git clone https://github.com/golang/lint.git %GOPATH%\\src\\golang.org\\x\\lint

go build -o %GOPATH%\\bin\\gocode.exe github.com/mdempsky/gocode
%go build -o %GOPATH%\\bin\\gopkgs.exe github.com/uudashr/gopkgs/cmd/gopkgs%
%go build -o %GOPATH%\\bin\\go-outline.exe github.com/ramya-rao-a/go-outline%
%go build -o %GOPATH%\\bin\\go-symbols.exe github.com/acroca/go-symbols%
%go build -o %GOPATH%\\bin\\guru.exe golang.org/x/tools/cmd/guru%
%go build -o %GOPATH%\\bin\\gorename.exe golang.org/x/tools/cmd/gorename%
%go build -o %GOPATH%\\bin\\dlv.exe github.com/derekparker/delve/cmd/dlv%
go build -o %GOPATH%\\bin\\gocode-gomod.exe github.com/stamblerre/gocode
go build -o %GOPATH%\\bin\\godef.exe github.com/rogpeppe/godef
go build -o %GOPATH%\\bin\\godef-gomod.exe github.com/ianthehat/godef
%go build -o %GOPATH%\\bin\\goreturns.exe github.com/sqs/goreturns%
go build -o %GOPATH%\\bin\\golint.exe golang.org/x/lint/golint

pause


```

### 3. vscode 的配置， 文件->首选项->设置

```json
"go.buildOnSave": "workspace",
"go.lintOnSave": "package",
"go.vetOnSave": "package",
"go.buildTags": "",
"go.buildFlags": [],
"go.lintFlags": [],
"go.vetFlags": [],
"go.coverOnSave": false,
"go.useCodeSnippetsOnFunctionSuggest": true,
"go.formatTool": "goreturns",
"go.goroot": "D:\\MySoft\\Go",
"go.gopath": "D:\\MySoft\\Go\\gopath",
"go.gocodeAutoBuild": false,
"terminal.integrated.shell.windows": "C:\\windows\\System32\\cmd.exe",
"go.autocompleteUnimportedPackages": true,
"files.autoSave": "off",
"go.docsTool": "guru",
"go.gocodePackageLookupMode": "go",
"files.associations": {
    "*.tpl": "html"
},
"go.inferGopath": true,

```

setting.json vscode 的配置文件

launch.json 调试

task.json 任务

## 代理 GOPROXY

-   https://goproxy.io
-   https://goproxy.cn
