/*
Sniperkit-Bot
- Status: analyzed
*/

package cmd

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/sniperkit/snk.fork.myke/core"
)

// Version prints myke version
func Version(opts *mykeOpts) error {
	data, err := core.Asset("tmp/version")
	if err != nil {
		return errors.Wrap(err, "error showing version")
	}

	version := strings.TrimSpace(string(data))
	fmt.Fprintf(opts.Writer, "myke version %s\n", version)
	return nil
}
