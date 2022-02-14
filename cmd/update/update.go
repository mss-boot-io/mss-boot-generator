package update

import (
	"github.com/spf13/cobra"

	"github.com/lwnmengjing/micro-service-gen-tool/pkg"
)

var (
	StartCmd = &cobra.Command{
		Use:     "update",
		Short:   "Update generate-tool",
		Example: "generate-tool update",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	pkg.Update()
	return nil
}