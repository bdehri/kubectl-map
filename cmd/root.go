package cmd

import (
	//egress "github.com/bdehri/kubectl-map/cmd/egress"
	ingress "github.com/bdehri/kubectl-map/cmd/ingress"
	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
)

var rootCmd = cobra.Command{
	Use:   "map",
	Short: "Get network policy map",
	Long:  "Get visual representation of network policies as a map",
}

func init() {
	rootCmd.AddCommand(ingress.NewIngressCommand())
	rootCmd.AddCommand()
}

func Execute() error {
	return rootCmd.Execute()
}
