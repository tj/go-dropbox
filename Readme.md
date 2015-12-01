
[![GoDoc](https://godoc.org/github.com/tj/go-dropbox?status.svg)](https://godoc.org/github.com/tj/go-dropbox) [![Build Status](https://semaphoreci.com/api/v1/projects/bc0bfd8b-73c9-45ba-b988-00f9e285e6ef/617305/badge.svg)](https://semaphoreci.com/tj/go-dropbox)

# Dropbox

 Simple Dropbox v2 client for Go.

 For a higher level client take a look at [go-dropy](https://github.com/tj/go-dropy).

## About

 Modelled more or less 1:1 with the API for consistency and parity with the [official documentation](https://www.dropbox.com/developers/documentation/http). More sugar should be implemented on top.

## Testing

 To manually run tests use the test account access token:

```
$ export DROPBOX_ACCESS_TOKEN=oENFkq_oIVAAAAAAAAAAC8gE3wIUFMEraPBL-D71Aq2C4zuh1l4oDn5FiWSdVVlL
$ go test -v
```

# License

MIT