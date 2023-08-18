# clean-cache

Clean up `node` project cache tool written in `go`, very fast

<br />

## README 🦉

[简体中文](./README.md) | English

<br />

## motivation

Want to quickly clean up the `node` project cache

<br />

## Usage

### install

```shell
go install github.com/markthree/clean-cache
```

### clean

By default, clean up the current directory `.nuxt`, `cache`, `.cache`, `@cache`,
`temp`, `.temp`, `@temp` directories

```shell
clean-cache
```

#### with dist and .output

```shell
clean-cache -d
```

#### with node_modules

```shell
clean-cache -n
```

#### help

```shell
clean-cache -h
```

<br />

## License

Made with [markthree](https://github.com/markthee)

Published under [MIT License](./LICENSE).
