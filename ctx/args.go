package ctx

import (
	"fmt"
	"os"

	"github.com/nixys/nxs-go-appctx-example-counter/misc"

	"github.com/pborman/getopt/v2"
)

const (
	confPathDefault = "./example.conf"
)

// Args contains arguments value read from command line
type Args struct {
	ConfigPath      string
	CounterInterval *string
}

var version string

// ArgsRead reads arguments from command line
func ArgsRead() (Args, error) {

	args := getopt.New()

	helpFlag := args.BoolLong(
		"help",
		'h',
		"Show help")

	versionFlag := args.BoolLong(
		"version",
		'v',
		"Show program version")

	confPath := args.StringLong(
		"conf",
		'c',
		confPathDefault,
		"Config file path")

	counterInterval := args.StringLong(
		"counter-interval",
		'i',
		"",
		"Counter interval")

	args.Parse(os.Args)

	/* Show help */
	if *helpFlag == true {
		argsHelp(args)
		return Args{}, misc.ErrArgSuccessExit
	}

	/* Show version */
	if *versionFlag == true {
		argsVersion()
		return Args{}, misc.ErrArgSuccessExit
	}

	var counter *string
	if args.IsSet("counter-interval") == true {
		counter = counterInterval
	}

	return Args{
		ConfigPath:      *confPath,
		CounterInterval: counter,
	}, nil
}

func argsHelp(args *getopt.Set) {

	additionalDescription := `
	
Additional description

  Just a sample.
`

	args.PrintUsage(os.Stdout)
	fmt.Println(additionalDescription)
}

func argsVersion() {
	fmt.Println(version)
}
