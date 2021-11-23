# t

`t` is a command line tool for testing on your terminal.

![Terminal](https://user-images.githubusercontent.com/10682/140833442-541da082-8bbd-4637-9bbd-c57440a455a2.png)

## Installation

```
$ go install github.com/yusukebe/t/cmd/t@latest
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

With [**jq**](https://stedolan.github.io/jq/):

```
$ echo '{ "message": "hello" }' | jq -r .message | t hello # => PASS
```

```
$ echo '{ "number": 5 }' | jq -r .number | t 10 -o '>' # => PASS
```

With [**rj**](https://github.com/yusukebe/rj), test the status is not error:

```
$ rj http://example.com/ | jq -r .code | t 500 -o '>'
```

CircleCI, you can test if the http status is ok on CircleCI with **only** config.yml:

```yaml
version: 2.1
jobs:
  test:
    parameters:
      url:
        type: string
    docker:
      - image: cimg/go:1.17.3
    steps:
      - run:
          command: go install github.com/yusukebe/rj/cmd/rj@latest
      - run:
          command: go install github.com/yusukebe/t/cmd/t@latest
      - run:
          name: Test << parameters.url >>
          command: rj <<parameters.url>> | jq '.code' | t 200
workflows:
  test_workflow:
    jobs:
      - test:
          url: https://example.com/
```

## Author

Yusuke Wada <https://github.com/yusukebe>

## License

MIT