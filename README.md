## What is it?

A fake HTTP API built as a support for the developpement of Postmanerator.
It also demonstrates how you can use [Newman](https://github.com/postmanlabs/newman) to automate
integration tests.

The app is built with Golang and its goal is to serve data about libraries and books.

## How does it work?

You need to have Newman and Docker > 1.10 installed on your machine to run it.
When you are good to go, simply run `./run-tests.sh`, it will:

- build the docker image for the app
- launch the app using the docker-compose command line
- run the test suites with newman

![Demo](demo.gif)
