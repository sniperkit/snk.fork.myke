/*
Sniperkit-Bot
- Status: analyzed
*/

package cmd

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/sniperkit/snk.fork.myke/core"
)

// License prints all open source licenses
func License(opts *mykeOpts) error {
	data, err := core.Asset("tmp/LICENSES")
	if err != nil {
		return errors.Wrap(err, "error showing licenses")
	}

	fmt.Fprintln(opts.Writer, string(data))
	return nil
}
