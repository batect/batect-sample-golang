# batect-sample-golang

[![CircleCI branch](https://img.shields.io/circleci/project/github/batect/batect-sample-golang/main.svg)](https://circleci.com/gh/batect/batect-sample-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/batect/batect-sample-golang)](https://goreportcard.com/report/github.com/batect/batect-sample-golang)
[![License](https://img.shields.io/github/license/batect/batect-sample-golang.svg)](https://opensource.org/licenses/Apache-2.0)

A very simple Golang application with [Batect](https://github.com/batect/batect)-based build and testing environments.

## Building, testing, running etc.

Run `./batect --list-tasks` to see the available commands and their descriptions, then `./batect <task>` to run `<task>`.

Most tasks require you to have downloaded the application's dependencies first. `./batect setup` will do this for you.

All of this is controlled by [batect.yml](batect.yml), and it shows a number of common patterns you might adopt in your own
application.

## More information

You can find further sample projects on [this page in the batect documentation](https://batect.dev/docs/getting-started/sample-projects).
