package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/blang/semver"
	"github.com/jakewarren/semverfmt"
	"github.com/spf13/pflag"
)

var (
	// build information set by ldflags
	appName    = "semverfmt"
	appVersion = "(ﾉ☉ヮ⚆)ﾉ ⌒*:･ﾟ✧"
	commit     = "(ﾉ☉ヮ⚆)ﾉ ⌒*:･ﾟ✧"
	buildDate  = "(ﾉ☉ヮ⚆)ﾉ ⌒*:･ﾟ✧"
)

type config struct {
	gitDir string
}

const usageMessage = `Usage: semverfmt [flags] <format string>`

const usageExample = `Examples:
	- Read from stdin:
	echo "v1.2.3" | semverfmt v%M.%m
	
	- Read version info from git tags:
	semverfmt --git ~/path/to/foo "v%M"`

func main() {
	c := config{}

	pflag.Usage = func() {
		_, _ = fmt.Fprintln(os.Stderr, usageMessage)
		_, _ = fmt.Fprintln(os.Stderr, "")
		_, _ = fmt.Fprintln(os.Stderr, "Flags:")
		pflag.PrintDefaults()
		_, _ = fmt.Fprintln(os.Stderr, "")
		_, _ = fmt.Fprintln(os.Stderr, usageExample)
		_, _ = fmt.Fprintln(os.Stderr, "")
		_, _ = fmt.Fprintln(os.Stderr, "URL: https://github.com/jakewarren/semverfmt")
	}

	displayHelp := pflag.BoolP("help", "h", false, "display help")
	displayVersion := pflag.BoolP("version", "V", false, "display version information")
	pflag.StringVarP(&c.gitDir, "git", "g", "", "directory path to query tag via git. use '.' to specify to current directory")
	pflag.Parse()

	// override the default usage display
	if *displayHelp {
		pflag.Usage()
		os.Exit(0)
	}

	if *displayVersion {
		fmt.Printf(`%s:
    version     : %s
    git hash    : %s
    build date  : %s 
    go version  : %s
    go compiler : %s
    platform    : %s/%s
`, appName, appVersion, commit, buildDate, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	if pflag.NArg() == 0 {
		log.Fatalln("format string not provided")
	}

	var (
		version string
		format  = pflag.Arg(0)
	)

	if c.gitDir == "" {
		// read from stdin
		data, readErr := ioutil.ReadAll(os.Stdin)
		if readErr != nil {
			log.Fatalf("error getting the tag from stdin: %s", readErr)
		}
		version = string(data)
	} else {
		// read from the directory
		tag, gitErr := tagFromRepo(c.gitDir)
		if gitErr != nil {
			log.Fatalf("error getting the tag from git: %s", gitErr)
		}
		version = tag
	}

	v, parseErr := semver.ParseTolerant(version)
	if parseErr != nil {
		log.Fatalf("error parsing version: %s", parseErr)
	}

	fmt.Println(semverfmt.Sprintf(v, format))
}

// function sourced from: https://github.com/caarlos0/svu/
func tagFromRepo(path string) (string, error) {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	cmd.Dir = path
	bts, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.Split(string(bts), "\n")[0], nil
}
