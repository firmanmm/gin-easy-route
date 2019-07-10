# Gin Easy Route
To make routing with Golang Gin Web Framework easier. This packages allow you to define complex url in its original form (ex : `/api/stuff/another/newer/stuff` instead of having to manually grouping them. It will automatically make group that belongs to parent group. For example all `/api` might have to pass some kind of **authorization** before being processed, the original solution *(as far as i know)* you have to manually make group from `/api` and from there you have to manually add `/stuff` into that group. While that might work but it makes code less readable by default. But with *Gin Easy Route* you can ignore grouping all together and let this package do it's job.

## Installation
Get the package via 
```
go get github.com/firmanmm/gin-easy-route
```

## Usage
Please see [Example](example/main.go) for how to use it.