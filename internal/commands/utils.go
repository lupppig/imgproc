package commands

import (
	"path/filepath"
	"strings"
)

func isSupported(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return supportedExt[ext]
}

func outputPath(outDir, input string) string {
	base := filepath.Base(input)
	return filepath.Join(outDir, base)
}
