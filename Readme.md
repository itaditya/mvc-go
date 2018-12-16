# MVC CRUD App in Golang

Resources on how to build complete Golang applications are scarce and not easily found. Coming from a Node.js background, I have built this project to demo a CRUD application in which we can add and delete users with complete frontend integrations following best practises of the Web Development world.

:warning:  **Disclaimer-** I am in the process of learning Golang so I can't guarantee that the Go code is idiomatic. Any suggestions and nitpickings are welcome.

## Setup Guide

1. Install Golang(v 1.11.3 or higher), setup `GOPATH` and make sure `GOPATH/bin` is in the `PATH` env variable.
1. Install `beego` cli tool.
1. Clone this project in a folder outside `GOPATH` (because we are using Go Modules).
1. `cd` into the project.
1. Install all dependencies with `go get -u`.
1. Rename `.env.example` to `.env` and provide the necessary information like Database URI etc.
1. Start server with `bee run`.
