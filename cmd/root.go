package cmd

import (
	"github.com/mducoli/cronupper/pkg/config"
	"github.com/mducoli/cronupper/pkg/scheduler"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cronupper",
	Short: "Schedule cron jobs to take and upload backups",
  Long: "Schedule cron jobs to take and upload backups\nMore informations: https://github.com/mducoli/cronupper",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.Parse(cfgFile)
		cobra.CheckErr(err)

		scheduler.CronJobBlockingLog(config.Jobs)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

var (
	cfgFile string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "/etc/cronupper/config.yml", "config file location")
}
