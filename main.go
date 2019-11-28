package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"./compiler"

	"github.com/spf13/cobra"
)

var outFile string
var inFile string
var source string
var compileOnly bool
var verbose bool
var deprecated bool

var rootCmd = &cobra.Command{
	Use:   "tealang",
	Short: "Tealang compiler to TEAL",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 {
			return errors.New("requires a source file name")
		}
		inFile = args[0]
		srcBytes, err := ioutil.ReadFile(inFile)
		if err != nil {
			return err
		}
		source = string(srcBytes)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var result string
		if deprecated {
			result = compiler.Compile(source)
		}
		prog, parseErrors := compiler.Parse(source)
		if len(parseErrors) > 0 {
			for _, e := range parseErrors {
				fmt.Printf("%s\n", e.String())
			}
			os.Exit(1)
		}

		result = compiler.Codegen(prog)
		if compileOnly {
			if outFile == "" {
				ext := path.Ext(inFile)
				outFile = inFile[0:len(inFile)-len(ext)] + ".teal"
			}
			if verbose {
				fmt.Printf("Writing result to %s\n", outFile)
			}
			ioutil.WriteFile(outFile, []byte(result), 0644)
		} else {
			fmt.Println("assembling to tealc not implemented yet\n Use -c to see TEAL output")
		}
	},
}

func main() {
	rootCmd.Flags().StringVarP(&outFile, "output", "o", "", "Output file")
	rootCmd.Flags().BoolVarP(&compileOnly, "compile", "c", false, "Compile to TEAL and stop")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.Flags().BoolVarP(&deprecated, "experimental", "e", false, "Experimental feature")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
