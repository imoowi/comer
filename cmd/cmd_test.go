package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func TestRoot(t *testing.T) {
	output, _ := executeCommand(rootCmd)
	fmt.Println(output)
}

func TestCreate(t *testing.T) {
	output, _ := executeCommand(rootCmd, "create", "-m", "tmp/comer-example/user")
	fmt.Println(output)
}

func TestVersion(t *testing.T) {
	output, _ := executeCommand(rootCmd, "version")
	fmt.Println(output)
}

func TestAddapp(t *testing.T) {
	rootCmd.AddCommand(addappCmd)

	output, _ := executeCommand(rootCmd, "addapp", "--app", "student")
	// output, _ := executeCommand(rootCmd, "addapp", "--app", "user")
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
