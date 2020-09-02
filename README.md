This library helps you to programmatically pilot vagrant using Go.

It provides an easy API around the vagrant command line tool.

## Getting Started

There is a test written inside `vagrant_test.go` that helps you to figure out
what this library provides. A few hints:

* `machine, err := vagrant.Up(ctx)` helps you to provision a new virtual
  machine.
* `machine.Destroy(ctx)` helps you to cleanup after yourself.
* `machine.Exec(ctx, "touch", "/tmp/file.text")` execute a command inside the
  virtual machine via ssh.

## Run tests

This library does not replace the `testing` package. It just provide a few
function that you can use to leverage vagrant as part of your testing
environment.

You can run the test I wrote but you need a functional vagrant installation in
your laptop:

```
$ git clone git@github.com:gianarb/vagrant-go
$ cd vagrant-go
$ go test --timeout 30m -v ./...
```

Timeout is optional but it depends on how fast you can download the Ubuntu image
from the internet.

NOTE: Parallelization is not supported.
