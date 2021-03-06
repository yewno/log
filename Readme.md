
# log

Modified version of TJ Holowaychuk's log component [apex/log](http://github.com/apex/log)


Package log implements a simple structured logging API inspired by Logrus, designed with centralization in mind. Read more on [Medium](https://medium.com/@tjholowaychuk/apex-log-e8d9627f4a9a#.rav8yhkud).

## Handlers

- __cli__ – human-friendly CLI output
- __text__ – human-friendly colored output
- __logfmt__ – logfmt plain-text formatter
- __memory__ – in-memory handler for tests
- __discard__ – discards all logs
- __es__ – Elasticsearch handler
- __graylog__ – Graylog handler
- __json__ – JSON output handler
- __kinesis__ – AWS Kinesis handler
- __multi__ – fan-out to multiple handlers
- __level__ – level filter handler

## Badges

[![Build Status](https://semaphoreci.com/api/v1/projects/d8a8b1c0-45b0-4b89-b066-99d788d0b94c/642077/badge.svg)](https://semaphoreci.com/tj/log)
[![GoDoc](https://godoc.org/github.com/apex/log?status.svg)](https://godoc.org/github.com/apex/log)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)
[![](http://apex.sh/images/badge.svg)](https://apex.sh/ping/)

---

> [tjholowaychuk.com](http://tjholowaychuk.com) &nbsp;&middot;&nbsp;
> GitHub [@tj](https://github.com/tj) &nbsp;&middot;&nbsp;
> Twitter [@tjholowaychuk](https://twitter.com/tjholowaychuk)
