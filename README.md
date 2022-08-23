# Vanity Cruncher

A tool to solve Vanity challenge from [Paradigm CTF 2022](https://ctf.paradigm.xyz/). It finds a SHA-256 hash meeting
certain requirement: first 4 bytes must be `1626ba7e`.

I haven't benchmarked it yet, but it should be fast enough.

Inspired by [crunchvanity](https://github.com/hrkrshnn/crunchvanity) which is written in Rust.

## Usage

``` shell
$ go install github.com/Jeiwan/vanitycruncher-go@latest
$ vanitycruncher-go
```