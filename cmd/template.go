/*
Sniperkit-Bot
- Status: analyzed
*/

package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/sniperkit/snk.fork.myke/core"
)

// Template renders the given template file
func Template(opts *mykeOpts) error {
	bytes, err := ioutil.ReadFile(opts.Template)
	if err != nil {
		return errors.Wrap(err, "error rendering template")
	}

	rendered, err := core.RenderTemplate(string(bytes), core.OsEnv(), map[string]string{})
	if err != nil {
		return errors.Wrap(err, "error rendering template")
	}

	fmt.Fprint(opts.Writer, rendered)
	return nil
}
