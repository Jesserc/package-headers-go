package headers

import (
	"fmt"
	"strings"

	"golang.design/x/clipboard"
)

// GenerateHeader generates a Go package header and copy it to your clipboard:
func GenerateHeader(packageName, description, example *string, stdout *bool) error {
	// Generate the package header
	header := fmt.Sprintf(`/*
Package %s provides %s

%s
*/
package %s
`, *packageName, *description, formatExample(*example), *packageName)

	if err := clipboard.Init(); err != nil {
		return err
	}

	if *stdout {
		fmt.Println(header)
	}
	clipboard.Write(clipboard.FmtText, []byte(header))
	return nil
}

// formatExample formats the example usage to include new lines
func formatExample(example string) string {
	// Replace the literal `\n` sequences with actual new lines
	return strings.ReplaceAll(example, `\n`, "\n")
}
