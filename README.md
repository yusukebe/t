# t

`t` is a command line tool for testing on your terminal.

## Installation

```
$ go get github.com/yusukebe/t/cmd/t
```

## Usage

Basic ussage:

```
$ t hello hello # => PASS
```

```
$ t hello hell # => FAIL
```

With pipe:

```
$ echo "hello" | t hello # => PASS
```

With [jq](https://stedolan.github.io/jq/):

```
$ echo '{ "message": "hello" }' | jq -r .message | t hello # => PASS
```

## Author

Yusuke Wada <https://github.com/yusukebe>

## License

MIT