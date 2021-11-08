# t

`t` is a command line tool for testing on your terminal.

## Installation

```
$ go get github.com/yusukebe/t/cmd/t
```

## Usage

Basic usage:

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

With operator

```
$ t -o ">" 10 "3*2" # => PASS
```

```
$ echo '{ "number": 5 }' | jq -r .number | t 10 -o '>'
```

## Author

Yusuke Wada <https://github.com/yusukebe>

## License

MIT