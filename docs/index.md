# Webqueue

The job queue with insight.

Webqueue makes you more productive when it comes to job queueing.

> **Note** Webqueue is currently in development, be careful with use in production!

## How does webqueue work?

Webqueue uses [RabbitMQ](https://www.rabbitmq.com/) and HTTP to make asynchronous processing easier for you.
At its core Webqueue is a message queue consumer that executes an HTTP-POST to a URL for every message it gets. You don't have to worry about managing a queue consumer.

## Why do I need this?

When you have a background in PHP for the web, long-running processes like a queue consumer are a different story compared to short-living web requests. Developing long running processes require an eye for memory- and processmanagement and a different type of error recovery.

Webqueue takes this trouble away from you. Your jobs (or messages) are executed as the same short-living calls your web application is written in!

## How can I get started?

There is a [getting started guide](getting-started.md) included in the documentation.