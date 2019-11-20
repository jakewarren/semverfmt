# semverfmt
[![CircleCI](https://circleci.com/gh/jakewarren/semverfmt.svg?style=shield)](https://circleci.com/gh/jakewarren/semverfmt)
[![GitHub release](http://img.shields.io/github/release/jakewarren/semverfmt.svg?style=flat-square)](https://github.com/jakewarren/semverfmt/releases])
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/jakewarren/semverfmt/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakewarren/semverfmt)](https://goreportcard.com/report/github.com/jakewarren/semverfmt)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields)](http://makeapullrequest.com)

`semverfmt` is a utility to parse and reformat semantic versions. Semantic versions can be reformatted using `printf` style specifications.

## Install
### Option 1: Binary

Download the latest release from [https://github.com/jakewarren/semverfmt/releases/latest](https://github.com/jakewarren/semverfmt/releases/latest)

### Option 2: From source

```
go get github.com/jakewarren/semverfmt/...
```

## Usage

```
❯ semverfmt -h
Usage: semverfmt [flags] <format string>

Flags:
  -g, --git string   directory path to query tag via git. use '.' to specify to current directory
  -h, --help         display help
  -V, --version      display version information

Examples:
	- Read from stdin:
	echo "v1.2.3" | semverfmt v%M.%m
	
	- Read version info from git tags:
	semverfmt --git ~/path/to/foo "v%M"
```

Format string specification:

| Format | Description            |
|--------|------------------------|
| %M     | Major version          |
| %m     | Minor version          |
| %p     | Patch version          |

### Read from stdin

```
❯ echo "v1.2.3" | semverfmt v%M.%m
v1.2
```

### Read version info from git tags

`semverfmt` can also query git directly to get a tag. 

```
❯ semverfmt --git ~/golang/src/github.com/jakewarren/fixme "v%M"
v1
```
## Changes

All notable changes to this project will be documented in the [changelog].

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this project adheres to [Semantic Versioning](http://semver.org/).

## License

MIT © 2019 Jake Warren

[changelog]: https://github.com/jakewarren/semverfmt/blob/master/CHANGELOG.md
