package client

import (
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubecfg"
	"github.com/openshift/origin/pkg/deploy/api"
)

var deploymentColumns = []string{"ID", "State"}
var deploymentConfigColumns = []string{"ID", "Trigger Policy"}

// RegisterPrintHandlers registers human-readable printers for deploy types.
func RegisterPrintHandlers(printer *kubecfg.HumanReadablePrinter) {
	printer.Handler(deploymentColumns, printDeployment)
	printer.Handler(deploymentColumns, printDeploymentList)
	printer.Handler(deploymentConfigColumns, printDeploymentConfig)
	printer.Handler(deploymentConfigColumns, printDeploymentConfigList)
}

func printDeployment(d *api.Deployment, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s\t%s\n", d.ID, d.State)
	return err
}

func printDeploymentList(list *api.DeploymentList, w io.Writer) error {
	for _, d := range list.Items {
		if err := printDeployment(&d, w); err != nil {
			return err
		}
	}

	return nil
}

func printDeploymentConfig(dc *api.DeploymentConfig, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s\t%s\n", dc.ID, dc.TriggerPolicy)
	return err
}

func printDeploymentConfigList(list *api.DeploymentConfigList, w io.Writer) error {
	for _, dc := range list.Items {
		if err := printDeploymentConfig(&dc, w); err != nil {
			return err
		}
	}

	return nil
}
