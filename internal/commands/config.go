package commands

const (
	EXIT_SUCCESS = iota
	EXIT_FAILURE
)

// contains flag config values...
type Config struct {
	InputDir    string
	OutputDir   string
	ResizeWidth int
	Format      string
	Quality     int
	Workers     int
	MaxInflight int
	Watermark   bool
	StripEXIF   bool
}
