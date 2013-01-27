RandomNameGenerator
===================

A handy little tool used to generate random, short, punchy names for top-secret projects and hip startups, but hopefully not for your children.

Usage
-----

The basic usage below generates 25 of said awesome punchy names by default.

```go
go run rng.go
```

Alternatively, the number of names to be generated can be set manually with the `-c` flag (for 'count'). The example below will generate 500 names, for you eager beavers.

```go
go run rng.go -c 500
```