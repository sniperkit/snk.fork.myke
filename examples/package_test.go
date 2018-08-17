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
	{Arg: ``, Out: `(?m)^\s*PROJECT\s*\|\s*TAGS\s*\|\s*TASKS\s*$`},
	{Arg: ``, Out: `(?m)^\s*env\s*\|\s*\|\s*file_custom, file_custom_local, file_default, file_default_local, path, yml\s*$`},
	{Arg: ``, Out: `(?m)^\s*hooks\s*\|\s*\|\s*after, before, error\s*$`},
	{Arg: ``, Out: `(?m)^\s*mixin\s*\|\s*\|\s*path, task1, task2, task3\s*$`},
	{Arg: ``, Out: `(?m)^\s*retry\s*\|\s*\|\s*retry\s*$`},
	{Arg: ``, Out: `(?m)^\s*tags1\s*\|\s*tagA, tagB\s*\|\s*tag\s*$`},
	{Arg: ``, Out: `(?m)^\s*tags2\s*\|\s*tagB, tagC\s*\|\s*tag\s*$`},
	{Arg: ``, Out: `(?m)^\s*template\s*\|\s*\|\s*args, envs, file\s*$`},
	{Arg: `--help`, Out: `(?s).*Usage.*myke options.*`},
	{Arg: `--version`, Out: `.*myke version.*`},
	{Arg: `--license`, Out: `.*OPEN SOURCE LICENSES.*`},
	{Arg: `-f myke-error.yml`, Err: true, Out: `.*open.*foobar.*`},
}

func Test(t *testing.T) {
	RunCliTests(t, "examples", tests)
}
