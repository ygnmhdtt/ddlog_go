# ddlog_go - Golang logger for Datadog original format

ddlog_go enables you to formatting log for [Datadog](https://docs.datadoghq.com/guides/logs/#datadog-canonical-log-format).
If you use Golang original logger, you need to prepare log [custom parser for Datadog](https://docs.datadoghq.com/guides/logs/#writing-parsing-functions).


## Table of Contents

* [Installation](#installation)
* [Examples](#examples)
* [License](#license)

## Installation

This library doesn't depends on any other packages.

```
$ go get github.com/ygnmhdtt/ddlog_go
```

(optional) To run unit tests:

```
$ cd $GOPATH/src/github.com/ygnmhdtt/ddlog_go
$ go test -v
```

## Examples

```
package main

import (
  "os"
  "github.com/ygnmhdtt/ddlog_go
)

func main() {
	// You can specify metric_name and where to output
  ddl := ddlog_go.NewddLogger("test.metric", os.Stderr)

	// Set attributes
	ddl.Attr("env", "production")

	// This line prints "test.metric 967809600 1 loglevel=INFO env=production"
	ddl.INFO("1")
	// This line prints "test.metric 967809601 2 loglevel=WARN env=production"
	ddl.WARN("2")

	// ClearAttr clears all attributes
	ddl.ClearAttr()

	// This line prints "test.metric 967809600 1 loglevel=INFO env=production hoge=fuga"
	ddl.Attr("env", "production").Attr("hoge", "fuga").INFO("1")
}
```

List of log level:
* DEBUG
* INFO
* WARN
* ERROR
* FATAL

## License

[MIT](https://github.com/ygnmhdtt/ddlog_go/blob/master/LICENSE)
