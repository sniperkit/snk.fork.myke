/*
Sniperkit-Bot
- Status: analyzed
*/

package test

import (
	"testing"

	. "github.com/sniperkit/snk.fork.myke/examples/util"
)

var tests = []TestTable{
	{Arg: `-v=0 error`, Out: `error`, Err: true},
	{Arg: `-v=0 error`, Out: `(Failed){0}`, Err: true},
	{Arg: `-v=0 error`, Out: `(foobar.*not found)`, Err: true},
	{Arg: `-v=0 echo`, Out: `(Running){0}`},
	{Arg: `-v=0 echo`, Out: `(echo){0}`},
	{Arg: `-v=1 error`, Out: `error`, Err: true},
	{Arg: `-v=1 error`, Out: `(Failed){0}`, Err: true},
	{Arg: `-v=1 echo`, Out: `(Running){0}`},
	{Arg: `-v=1 echo`, Out: `(echo)`},
	{Arg: `subshell`, Out: `subshell works`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples/shell", tests)
}
