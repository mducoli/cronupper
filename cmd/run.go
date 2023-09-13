package cmd

import (
	"fmt"

	"github.com/mducoli/cronupper/pkg/config"
	"github.com/mducoli/cronupper/pkg/executer"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [job name]",
	Short: "Executes a job immediately",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cobra.CheckErr(fmt.Errorf("please specify the job name"))
		}

		config, err := config.Parse(cfgFile)
		cobra.CheckErr(err)

		job, has := config.Jobs[args[0]]
		if !has {
			cobra.CheckErr(fmt.Errorf("job not found"))
		}

		err = executer.Execute(&job)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
