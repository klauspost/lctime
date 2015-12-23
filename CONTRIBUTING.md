# Contributing

Hey! Thanks for being interested in helping, amigo! Let's bring native dates to
the world! Here's how you can help!

## Create an issue

I'm only familar with the `en_US` and `es_MX` locales, so it's harder for me to
verify that the other locales work. If you notice something wrong with your
locale, create an issue to let me know.

## Submit a pull request

You don't have to be a coder to submit a pull request, but code is _definitely_
welcome! If you notice documentation lacking, a bad translation, or missing
locale information send a PR!

But before you send anything, create an issue first to get other people's
feedback. After everything checks out, you can start working. Setting up is
easy.

### Set up a dev environment

`lctime` is a Go project, so I'm assuming you already have a working Go
installation and you set up your `GOPATH`. If not, then read Go's [Getting
Started] page.

#### Fork lctime

Create a fork of `lctime` on your own GitHub account. That way you can make
whatever changes you like to your own copy.

#### Download the source

```
cd $GOPATH/src/github.com/yourusername
git clone git@github.com:yourusername/lctime.git
```

#### Hack!

For example: add a test case, fix a parser bug, correct a translation, add
missing information to docs, simplify the code—make `lctime` better so people
can have friggin' _dates_ in their language and format!

#### Run tests

```
cd $GOPATH/src/github.com/yourusername/lctime
go test
```

There should be no failing tests. If there is, then something is broken. Fix it.

#### If necessary, regenerate locale data

If you changed any of the locale data, then you'll need to regenerate the Go
code for that data. Use these commands.

```
cd internal/locales
rm 1data.go .DS_Store
go-bindata -pkg locale -o 1data.go .
```

Next, verify that only the locale files were picked up by `go-bindata` by
opening `1data.go` and reading the `sources` comment at the top of the file.

Thank you!

[Getting Started]: https://golang.org/doc/install
