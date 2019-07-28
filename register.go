package flaggs

import (
	"flag"
	"time"
)

func registerUsage() {
	if usageRegistered == true {
		return
	}

	usageRegistered = true
	flag.Usage = func() {
		PrintUsage()
	}

	helpFlag := betterFlag{
		shortName:   "h",
		longName:    "help",
		description: "Prints this help information."}
	flags = append(flags, helpFlag)
}

// RegisterStringFlag registers a string flag
func RegisterStringFlag(reference *string, shortName string, longName string, description string) {
	registerUsage()

	defaultValue := *reference
	flag.StringVar(reference, longName, *reference, description)
	if shortName != "" {
		flag.StringVar(reference, shortName, *reference, description)
	}

	var usage string
	if defaultValue != "" {
		usage = "[string]"
	} else {
		usage = "string"
	}

	f := betterFlag{usage, shortName, longName, description}
	flags = append(flags, f)
}

// RegisterBoolFlag registers a boolean flag
func RegisterBoolFlag(reference *bool, shortName string, longName string, description string) {
	registerUsage()

	flag.BoolVar(reference, longName, *reference, description)
	if shortName != "" {
		flag.BoolVar(reference, shortName, *reference, description)
	}

	usage := ""

	f := betterFlag{usage, shortName, longName, description}
	flags = append(flags, f)
}

// RegisterUintFlag registers a uint flag
func RegisterUintFlag(reference *uint, shortName string, longName string, description string) {
	registerUsage()

	defaultValue := *reference
	flag.UintVar(reference, longName, *reference, description)
	if shortName != "" {
		flag.UintVar(reference, shortName, *reference, description)
	}

	var usage string
	if defaultValue != 0 {
		usage = "[uint]"
	} else {
		usage = "uint"
	}

	f := betterFlag{usage, shortName, longName, description}
	flags = append(flags, f)
}

// RegisterIntFlag registers an int flag
func RegisterIntFlag(reference *int, shortName string, longName string, description string) {
	registerUsage()

	defaultValue := *reference
	flag.IntVar(reference, longName, *reference, description)
	if shortName != "" {
		flag.IntVar(reference, shortName, *reference, description)
	}

	var usage string
	if defaultValue != 0 {
		usage = "[int]"
	} else {
		usage = "int"
	}

	f := betterFlag{usage, shortName, longName, description}
	flags = append(flags, f)
}

// RegisterDurationFlag registers a time.Duration flag
func RegisterDurationFlag(reference *time.Duration, shortName string, longName string, description string) {
	registerUsage()

	defaultValue := *reference
	flag.DurationVar(reference, longName, *reference, description)
	if shortName != "" {
		flag.DurationVar(reference, shortName, *reference, description)
	}

	var usage string
	if defaultValue != 0 {
		usage = "[duration]"
	} else {
		usage = "duration"
	}

	f := betterFlag{usage, shortName, longName, description}
	flags = append(flags, f)
}
