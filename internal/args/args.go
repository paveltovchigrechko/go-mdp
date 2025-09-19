package args

import "flag"

var Root string
var Ext string

// prepareRootFlag registers flags for root directory
func prepareRootFlag() {
	const (
		rootFlagDefaultValue = "."
		rootFlagUsage        = "Root directory"
	)
	flag.StringVar(&Root, "root", rootFlagDefaultValue, rootFlagUsage)
	flag.StringVar(&Root, "r", rootFlagDefaultValue, rootFlagUsage)
}

// prepareExtFlag registers flags for file extension
func prepareExtFlag() {
	const (
		extFlagDefaultValue = ".md"
		extFlagUsage        = "File extension"
	)
	flag.StringVar(&Ext, "ext", extFlagDefaultValue, extFlagUsage)
	flag.StringVar(&Ext, "e", extFlagDefaultValue, extFlagUsage)
}

// PrepareFlags registers all program flags
func PrepareFlags() {
	prepareRootFlag()
	prepareExtFlag()
}
