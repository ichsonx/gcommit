# gcommit
一个轻量级，简单易用的提交git commit message的工具
- 使用标准的`header`、`body`、`footer`格式提交
- 即使`body`、`footer`为空，也依然会提交body、footer字样的message。目的是完整性和易读性。
- 中文提示
- 只 commit 在 __stage区__的文件
- 只生成并提交 commit message 的内容
- 可在commit后打tag
- 可在commit、tag后push到远程仓库



## install

```shell
go get -u github.com/ichsonx/gcommit
go install github.com/ichsonx/gcommit
```

**将编译后的程序加入到环境变量中，方便命令行使用。**



## useage

当需要commit的时候，直接运行即可。会有提示一步一步边界快速完成标准的git message。

- 值得注意的是，在body、footer输入时，可直接【回车】跳过。但若是已经有输入内容了，要结束时，需要`回车`新起一行输入`eof`来结束当前输入。
- 即使`body`、`footer`都没有内容，但最终生成的git message会仍然保留它们的标签与格式内容。目的是为了易读性。
- 关闭issue和不兼容声明，需要自己主动输入（footer部分会有提示）。因为issue输入在哪个部分相对于某些平台也是生效的，这里不做严格规定。

```
【0】feat：新功能（feature）
【1】fix：修补bug
【2】docs：文档（documentation）
【3】style：格式（不影响代码运行的变动）
【4】refactor：重构（即不是新增功能，也不是修改bug的代码变动）
【5】test：增加测试
【6】build：（以前称chore）构建系统（涉及脚本、配置或工具）和包依赖项相关的开发更改
【7】perf：性能提升相关的更改
【8】vendor：更新依赖项、包的版本
```

1. 基本用法

```shell
gcommit
```

2. 命令行说明

```shell
gcommit -h
```

3. 提交后自动push

```shell
gcommit -p
```

4. commit同时打tag

```shell
# 为刚刚提交的commit打tag版本
gcommit -t v0.1.0

# 在打tag的同时附加tag注解。
# 单独使用-tm是会被忽略的
gcommit -t v0.1.1 -tm 这是tag的附注信息
```



