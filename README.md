# ZIPp

`zipp` is a Go package for creating and extracting ZIP archives,
where the second "p" helps avoid conflicts and aids in memorization—think
of it as representing "package" or "pack"
It provides simple functions for working with ZIP files in a cross-platform way.

## Installation

Install the package using `go get`:

```bash
go get github.com/woozymasta/zipp
```

## Usage

```go
import "github.com/woozymasta/zipp"

// To create a ZIP archive from a directory, use the Pack function:
err := zipp.Pack("path/to/sourceDir", "path/to/archive.zip")
if err != nil {
    log.Fatal(err)
}

// To extract the contents of a ZIP archive to a specified directory, use the Unpack function:
err := ziputil.Unpack("path/to/archive.zip", "path/to/destinationDir")
if err != nil {
    log.Fatal(err)
}
```

## Other archive packages

* [https://github.com/WoozyMasta/tgz](TGZ Package) -
  simple way to create and extract tar.gz archives
