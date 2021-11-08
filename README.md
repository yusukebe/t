# t

`t` is a command line tool for testing on your terminal.

![Terminal](https://user-images.githubusercontent.com/10682/140833442-541da082-8bbd-4637-9bbd-c57440a455a2.png)

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

Eval:

```
$ t -e 10 "5*2" # => PASS
```

With operator

```
$ t -o ">" 10 "3*2" # => PASS
```

Regular expression:

```
$ t -m 'hello+' helloooooo # => PASS
```

With [jq](https://stedolan.github.io/jq/):

```
$ echo '{ "message": "hello" }' | jq -r .message | t hello # => PASS
```

```
$ echo '{ "number": 5 }' | jq -r .number | t 10 -o '>' # => PASS
```

With [rj](https://github.com/yusukebe/rj), test the status is not error:

```
$ rj http://example.com/ | jq -r .code | t 500 -o '>'
```

## Author

Yusuke Wada <https://github.com/yusukebe>

## License

MIT