# Helm Trash

[![Build Status](https://api.travis-ci.com/jroucheton/helm-trash.svg?branch=master)](https://travis-ci.com/jroucheton/helm-trash)

![helm-trash](https://jroucheton.github.io/helm-trash/logo/helm-trash_150x150.png)

## What is Helm Trash?

This is a Helm plugin to delete releases previously installed through helm spray.

It works like `helm delete --purge`, except that it is processed for all releases using the same chart.

## Continuous Integration & Delivery

Helm Trash is building and delivering under Travis.

[![Build Status](https://api.travis-ci.com/jrouheton/helm-trash.svg?branch=master)](https://travis-ci.com/jroucheton/helm-trash)


## Usage

```
  $ helm trash CHART
```

Helm Trash shall always be called on the umbrella chart.

### Flags:

```
      --debug              enable verbose output
      --dry-run            simulate a trash
  -h, --help               help for helm
  -n, --namespace string   namespace to trash the chart into. (default "default")
      --version string     specify the exact chart version to install. If this is not specified, the latest version is installed
```

## Install

```
  $ helm plugin install https://github.com/jroucheton/helm-trash
```

The above will fetch the latest binary release of `helm trash` and install it.

## Developer (From Source) Install

If you would like to handle the build yourself, instead of fetching a binary,
this is how recommend doing it.

First, set up your environment:

- You need to have [Go](http://golang.org) installed. Make sure to set `$GOPATH`
- If you don't have [Glide](http://glide.sh) installed, this will install it into
  `$GOPATH/bin` for you.

Clone this repo into your `$GOPATH`. You can use `go get -d github.com/jroucheton/helm-trash`
for that.

```
$ cd $GOPATH/src/github.com/jroucheton/helm-trash
$ make bootstrap build
$ SKIP_BIN_INSTALL=1 helm plugin install $GOPATH/src/github.com/jroucheton/helm-trash
```

That last command will skip fetching the binary install and use the one you
built.







