# clean-cache

用 `go` 写的清理 `node` 项目缓存工具，速度很快

<br />

## Usage

### install

```shell
go install github.com/markthree/clean-cache
```

### clean

默认清理当前目录下的
`.nuxt`，`cache`，`.cache`，`@cache`，`temp`，`.temp`，`@temp` 目录

```shell
clean-cache
```

#### 同时移除 dist 和 .output

```shell
clean-cache -d
```

#### 同时移除 node_modules

```shell
clean-cache -n
```

#### 查看帮助

```shell
clean-cache -h
```

<br />

## License

Made with [markthree](https://github.com/markthee)

Published under [MIT License](./LICENSE).
