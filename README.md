# Webqueue

The job queue with insight.

Webqueue makes you more productive when it comes to job queueing.

## Installation

Clone this repository and checkout the master branch.
``` bash
$ git clone git@github.com:Enrise/webqueue.git
```

Install the golang app:

``` bash
$ go install cmd/webqueue.go
```


### Requirements:

* MongoDB (3.x)
* RabbitMQ (3.x)
* Golang (when compiling yourself)

## Usage

When installed system-wide using the instructions above you can start webqueue by running:

``` bash
$ webqueue
```

Or when you want to provide a custom configuration file:

``` bash
$ webqueue -c webqueue.yml
```

## Change log

Please see [CHANGELOG](CHANGELOG.md) for more information what has changed recently.

## Testing

``` bash
$ go test
```

## Credits

- [Richard Tuin][http://github.com/rtuin]
- [All Contributors][https://github.com/Enrise/webqueue/contributors]