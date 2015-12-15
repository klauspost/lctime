# lctime

Finally, simple, familiar, locale-based datetime formatting.

[![GoDoc](https://godoc.org/github.com/variadico/lctime?status.svg)](https://godoc.org/github.com/variadico/lctime)
[![GoCover](http://gocover.io/_badge/github.com/variadico/lctime)](http://gocover.io/github.com/variadico/lctime)
[![ReportCard](http://goreportcard.com/badge/variadico/lctime)](http://goreportcard.com/report/variadico/lctime)

**Work in progress! Not prod ready! Beware!**

## Install

```
go get -u github.com/variadico/lctime
```

## Usage

```go
// Default locale is based on env vars or en_US if none are set.
fmt.Println(lctime.Strftime("%c", time.Now()))
// prints: Mon 14 Dec 2015 10:31:56 PM PST

lctime.SetLocale("es_MX")
fmt.Println(lctime.Strftime("%c", time.Now()))
// prints: lun 14 dic 2015 22:31:56 PST
```

## The problem with the Go standard library

Go's standard library `time` is fine most of the time. However, it's currently
impossible for the Go standard library to format dates based on a user's locale
or langauge setting. Unfortunately, that means code like this, doesn't work.

```go
d := time.Date(2006, 5, 4, 3, 2, 1, 0, time.UTC)
fmt.Println(d.Format("02 de enero de 2006"))
// got:  04 de enero de 2006
// want: 04 de mayo de 2006
```

## The solution offered by lctime

`lctime` brings the familiar `setlocale` and `strftime` functions found in other
programming languages like C, Python, or PHP. The formats and translations used
for this package come from the [glibc locale files], with close to 300 locales.

This may not be the most elegant answer, like Go's `2006 January 02` thing, but
it sure is simple and effective. Parsing codes is way easier than natural
language. Most important of all, non-English speakers will get the luxury of
dates they're familiar with.

## Known issues


* Locale information is in a folder of JSON files in `lctime`. The files are
found by checking `$GOPATH/src/github.com/variadico/lctime/locales`. If the
files aren't there, then you won't have any locale info.
* Having a folder of JSON files also makes it harder to have a single
statically-linked binary of your app. Though there might be a solution with
`go-bindata`. Maybe...
* Not all directives have been fully implemented or well tested. Though, mostly
for locales that don't start their week on Sunday or other non-Western
calendars.

[glibc locale files]: http://lh.2xlibre.net/
