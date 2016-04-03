# Getting started with Webqueue

## Requirements:

* [MongoDB](https://www.mongodb.org/) (3.x)
* [RabbitMQ](https://www.rabbitmq.com/) (3.x, with the management plugin enabled)
* [Configured Golang environment](https://golang.org/doc/install)

## Installation

Clone this repository and checkout the master branch.
``` bash
$ git clone git@github.com:Enrise/webqueue.git
```

Install the dependencies:
``` bash
$ go get ./...
```

Run the golang app:

``` bash
$ go run cmd/webqueue.go
```

## Configuration
Webqueue can be configured by editing `webqueue.yml`.

Running it with a separate configuration file can be done like so:
``` bash
$ go run cmd/webqueue.go -c foo.yml
```

## System-wide installation

When installed system-wide using the instructions above you can start Webqueue by running:

``` bash
$ webqueue
```

Or when you want to provide a custom configuration file:

``` bash
$ webqueue -c webqueue.yml
```