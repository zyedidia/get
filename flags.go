package main

type Flags struct {
	Tag         string   `short:"t" long:"tag" description:"tagged release to use instead of latest"`
	Prerelease  bool     `long:"pre-release" description:"include pre-releases when fetching the latest version"`
	Output      string   `long:"to" description:"move to given location after extracting"`
	System      string   `short:"s" long:"system" description:"target system to download for (use \"all\" for all choices)"`
	ExtractFile string   `short:"f" long:"file" description:"glob to select files for extraction"`
	All         bool     `long:"all" description:"extract all candidate files"`
	Quiet       bool     `short:"q" long:"quiet" description:"only print essential output"`
	DLOnly      bool     `short:"d" long:"download-only" description:"stop after downloading the asset (no extraction)"`
	UpgradeOnly bool     `long:"upgrade-only" description:"only download if release is more recent than current version"`
	BinaryName  string   `short:"b" long:"binary" description:"binary name of current version"`
	Asset       []string `short:"a" long:"asset" description:"download a specific asset containing the given string; can be specified multiple times for additional filtering; use ^ for anti-match"`
	Hash        bool     `long:"sha256" description:"show the SHA-256 hash of the downloaded asset"`
	Verify      string   `long:"verify-sha256" description:"verify the downloaded asset checksum against the one provided"`
	Rate        bool     `long:"rate" description:"show GitHub API rate limiting information"`
	Remove      bool     `short:"r" long:"remove" description:"remove the given file from $EGET_BIN or the current directory"`
	Version     bool     `short:"v" long:"version" description:"show version information"`
	Help        bool     `short:"h" long:"help" description:"show this help message"`
}
