// Copyright Red Hat

package main

import (
	goflag "flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	registeredcluster "github.com/stolostron/compute-operator/controllers/cluster-registration"
	"github.com/stolostron/compute-operator/controllers/installer"

	"github.com/stolostron/compute-operator/webhook"

	utilflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	klog.InitFlags(goflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	logs.InitLogs()
	defer logs.FlushLogs()

	command := newWorkCommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func newWorkCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "compute-operator",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
			os.Exit(1)
		},
	}

	cmd.AddCommand(installer.NewInstaller())
	cmd.AddCommand(registeredcluster.NewManager())
	cmd.AddCommand(webhook.NewAdmissionHook())

	return cmd
}
