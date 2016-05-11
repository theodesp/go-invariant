# go-invariant

A way to provide descriptive errors in development but generic errors in production.

Install
=============

```bash

go get github.com/theodesp/go-invariant

```


Usage
====

It's intended for projects which require a simple invariant check and display
of error messages depending on the current env variables

Examples:

```go
// main.go
package main

import (
	"github.com/theodesp/go-invariant"
	"fmt"
)

func main() {
	res:= invariant.Invariant(1!=1,"")
	fmt.Println(res)
}

```
and run on the command line
```bash
$ go run main.go
    invariant requires an error message argument

$ go run main.go -env production
    invariant exception in production environment. Please use development flag to see the full error message
```

```go
// main.go
package main

import (
	"github.com/theodesp/go-invariant"
	"fmt"
)

func main() {
	res:= invariant.Invariant(1!=1,"You must say %s when talking to Dragons", "your majesty")
	fmt.Println(res)
}

```
and run on the command line
```bash
$ go run main.go
    You must say your majesty when talking to Dragons

$ go run main.go -env production
    invariant exception in production environment. Please use development flag to see the full error message
```

Note: When **env** environment variable or command line switch
is not *production*, the message is required.
If omitted, invariant will throw regardless of the truthness of the condition.
When **env** is *production*, the message will be a generic one so it can be stripped away
By default **env** is assumed to be *development*.
