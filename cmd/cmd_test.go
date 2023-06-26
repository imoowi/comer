package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func TestRoot(t *testing.T) {
	output, _ := executeCommand(rootCmd, "--module", "user", "--path", "tmp/comer-example")
	fmt.Println(output)
}

func TestVersion(t *testing.T) {
	output, _ := executeCommand(rootCmd, "version")
	fmt.Println(output)
}

func TestGenapp(t *testing.T) {
	rootCmd.AddCommand(genappCmd)
	output, _ := executeCommand(rootCmd, "genapp", "--app", "user")
	fmt.Println(`output is : `, output)
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	err = root.Execute()
	if err != nil {
		fmt.Println(err)
	}

	return buf.String(), err
}
