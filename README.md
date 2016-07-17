# NNDB
Go package for processing data from the National Nutritional Database.  

## Status
[![CircleCI](https://circleci.com/gh/rtyer/nndb.svg?style=svg)](https://circleci.com/gh/rtyer/nndb)

## Package Overview
- `nndb` contains basic model information
- `parser` contains the code necessary to parse the file formats from the National Nutritional Database into nndb model objects
- `cmd/nndb` will provide command line utility for converting from National Nutritional Database file format into a single outfile (csv, etc)

## Getting Started
- Check this out to the correct location (`$GOPATH/src/github.com/rtyer/nndb`)
- Ensure you have wget (`brew install wget`)
- Ensure you have glide (`brew install glide`)
- Execute `make prepare`

Prepare will download go linting tools and the latest NNDB data file and unzip it.  At this point, you can run any and all commands in the Makefile.

Basic commands:
- `make test`
- `make build`  