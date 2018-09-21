/*
Sniperkit-Bot
- Status: analyzed
*/

package cmd

import (
	"github.com/pkg/errors"
	"github.com/sniperkit/snk.fork.myke/core"
)

// Run runs the given tasks
func Run(opts *mykeOpts, tasks []string) error {
	queries := core.ParseQueries(tasks)
	w, err := loadWorkspace(opts.File)
	if err != nil {
		return err
	}

	for _, q := range queries {
		err := core.ExecuteQuery(&w, q, opts.DryRun, opts.Verbose)
		if err != nil {
			return errors.Wrap(err, "error running command")
		}
	}

	return nil
}
