# t

`t` is a command line tool for testing on your terminal.

![Terminal](https://user-images.githubusercontent.com/10682/140739013-ed862433-72a6-4e36-b253-aba36052b8f6.png)

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
$ echo '{ "number": 5 }' | jq -r .number | t 10 -o '>' # => PASS
```

With [rj](https://github.com/yusukebe/rj), test status is OK:

```
$ rj http://example.com/ | jq -r .code | t 200 -o '<='
```

## Author

Yusuke Wada <https://github.com/yusukebe>

## License

MIT