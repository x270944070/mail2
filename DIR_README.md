# Go项目标准布局

## Go目录

### `/cmd`

项目主要的应用程序。

对于每个应用程序来说这个目录的名字应该和项目可执行文件的名字相匹配（例如，`/cmd/myapp`）。

不要在这个目录中放太多的代码。如果目录中的代码可以被其他项目导入并使用，那么应该把他们放在`/pkg`目录。如果目录中的代码不可重用，或者不希望被他人使用，应该将代码放在`/internal`目录。显式地表明意图比较好！

通常来说，项目都应该拥有一个小的`main`函数，并在`main`函数中导入或者调用`/internal`和`/pkg`目录中的代码。

更多详情，请看[`/cmd`](https://github.com/golang-standards/project-layout/blob/master/cmd/README.md)目录中的例子。

### `/internal`

私有的应用程序代码库。这些是不希望被其他人导入的代码。请注意：这种模式是Go编译器强制执行的。详细内容情况Go 1.4的[release notes](https://golang.org/doc/go1.4#internalpackages)。再次注意，在项目的目录树中的任意位置都可以有`internal`目录，而不仅仅是在顶级目录中。

可以在内部代码包中添加一些额外的结构，来分隔共享和非共享的内部代码。这不是必选项（尤其是在小项目中），但是有一个直观的包用途是很棒的。应用程序实际的代码可以放在`/internal/app`目录（如，`internal/app/myapp`），而应用程序的共享代码放在`/internal/pkg`目录（如，`internal/pkg/myprivlib`）中。

### `/internal/repository`

数据库操作目录

### `/internal/models`

数据库操作目录

### `/internal/api/request`和`/internal/api/response`

外层的请求与响应


### `/pkg`

外部应用程序可以使用的库代码（如，`/pkg/mypubliclib`）。其他项目将会导入这些库来保证项目可以正常运行，所以在将代码放在这里前，一定要三思而行。请注意，`internal`目录是一个更好的选择来确保项目私有代码不会被其他人导入，因为这是Go强制执行的。使用`/pkg`目录来明确表示代码可以被其他人安全的导入仍然是一个好方式。Travis Jeffery撰写的关于 [I’ll take pkg over internal](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) 文章很好地概述了`pkg`和`inernal`目录以及何时使用它们。

当您的根目录包含大量非Go组件和目录时，这也是一种将Go代码分组到一个位置的方法，从而使运行各种Go工具更加容易（在如下的文章中都有提到：2018年GopherCon [Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)，[Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) ，Golab 2018 [Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)）。

点击查看`/pkg`就能看到那些使用这个布局模式的流行Go代码仓库。这是一种常见的布局模式，但未被普遍接受，并且Go社区中的某些人不推荐这样做。

如果项目确实很小并且嵌套的层次并不会带来多少价值（除非你就是想用它），那么就不要使用它。请仔细思考这种情况，当项目变得很大，并且根目录中包含的内容相当繁杂（尤其是有很多非Go的组件）。

### `/vendor`

应用程序的依赖关系（通过手动或者使用喜欢的依赖管理工具，如新增的内置[Go Modules](https://go.dev/wiki/Modules)特性）。执行`go mod vendor`命令将会在项目中创建`/vendor`目录，注意，如果使用的不是Go 1.14版本，在执行`go build`进行编译时，需要添加`-mod=vendor`命令行选项，因为它不是默认选项。

构建库文件时，不要提交应用程序依赖项。

请注意，从[1.13](https://golang.org/doc/go1.13#modules)开始，Go也启动了模块代理特性（使用`https：//proxy.golang.org`作为默认的模块代理服务器）。点击[这里](https://blog.golang.org/module-mirror-launch)阅读有关它的更多信息，来了解它是否符合所需要求和约束。如果`Go Module`满足需要，那么就不需要`vendor`目录。

## 服务端应用程序的目录

### `/api`

OpenAPI/Swagger规范，JSON模式文件，协议定义文件。

更多样例查看[`/api`](https://github.com/golang-standards/project-layout/blob/master/api/README.md)目录。

## Web应用程序的目录

### `/web`

Web应用程序特定的组件：静态Web资源，服务器端模板和单页应用（Single-Page App，SPA）。

## 通用应用程序的目录

### `/configs`

配置文件模板或默认配置。

将`confd`或者`consul-template`文件放在这里。

### `/init`

系统初始化（systemd、upstart、sysv）和进程管理（runit、supervisord）配置。

### `/scripts`

用于执行各种构建，安装，分析等操作的脚本。

这些脚本使根级别的Makefile变得更小更简单（例如<https://github.com/hashicorp/terraform/blob/main/Makefile>）。

更多样例查看[`/scripts`](https://github.com/golang-standards/project-layout/blob/master/scripts/README.md)。

### `/build`

打包和持续集成。

将云（AMI），容器（Docker），操作系统（deb，rpm，pkg）软件包配置和脚本放在`/build/package`目录中。

将CI（travis、circle、drone）配置文件和就脚本放在`build/ci`目录中。请注意，有一些CI工具（如，travis CI）对于配置文件的位置有严格的要求。尝试将配置文件放在`/build/ci`目录，然后链接到CI工具想要的位置。

### `/deployments`

IaaS，PaaS，系统和容器编排部署配置和模板（docker-compose，kubernetes/helm，mesos，terraform，bosh）。请注意，在某些存储库中（尤其是使用kubernetes部署的应用程序），该目录的名字是`/deploy`。

### `/test`

外部测试应用程序和测试数据。随时根据需要构建`/test`目录。对于较大的项目，有一个数据子目录更好一些。例如，如果需要Go忽略目录中的内容，则可以使用`/test/data`或`/test/testdata`这样的目录名字。请注意，Go还将忽略以“`.`”或“`_`”开头的目录或文件，因此可以更具灵活性的来命名测试数据目录。

更多样例查看[`/test`](https://github.com/golang-standards/project-layout/blob/master/test/README.md)。

## 其他

### `/docs`

设计和用户文档（除了godoc生成的文档）。

更多样例查看[`/docs`](https://github.com/golang-standards/project-layout/blob/master/docs/README.md)。

### `/tools`

此项目的支持工具。请注意，这些工具可以从`/pkg`和`/internal`目录导入代码。

更多样例查看[`/tools`](https://github.com/golang-standards/project-layout/blob/master/tools/README.md)。

### `/examples`

应用程序或公共库的示例。

更多样例查看[`/examples`](https://github.com/golang-standards/project-layout/blob/master/examples/README.md)。

### `/third_party`

外部辅助工具，fork的代码和其他第三方工具（例如Swagger UI）。

### `/githooks`

Git的钩子。

### `/assets`

项目中使用的其他资源（图像，Logo等）。

### `/website`

如果不使用Github pages，则在这里放置项目的网站数据。

更多样例查看[`/website`](https://github.com/golang-standards/project-layout/blob/master/website/README.md)。

## 不应该出现的目录

### `/src`

有一些Go项目确实包含`src`文件夹，但通常只有在开发者是从Java（这是Java中一个通用的模式）转过来的情况下才会有。如果可以的话请不要使用这种Java模式。你肯定不希望你的Go代码和项目看起来像Java。
