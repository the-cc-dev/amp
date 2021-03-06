package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/appcelerator/amp/api/rpc/service"
	"github.com/appcelerator/amp/cli"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

type taskOptions struct {
	all   bool
	quiet bool
}

const (
	StateRunning = "RUNNING"
)

// NewServicePsCommand returns a new instance of the service ps command
func NewServicePsCommand(c cli.Interface) *cobra.Command {
	opts := taskOptions{}
	cmd := &cobra.Command{
		Use:     "ps [OPTIONS] SERVICE",
		Short:   "List running tasks of a service",
		PreRunE: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return tasks(c, args, opts)
		},
	}
	flags := cmd.Flags()
	flags.BoolVarP(&opts.all, "all", "a", false, "Display all tasks")
	flags.BoolVarP(&opts.quiet, "quiet", "q", false, "Only display task ids")
	return cmd
}

func tasks(c cli.Interface, args []string, opts taskOptions) error {
	conn := c.ClientConn()
	client := service.NewServiceClient(conn)
	request := &service.TasksRequest{
		ServiceId: args[0],
	}
	reply, err := client.Tasks(context.Background(), request)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			return errors.New(s.Message())
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, cli.Padding, ' ', 0)
	fmt.Fprintln(w, "ID\tIMAGE\tDESIRED STATE\tCURRENT STATE\tNODE ID\tERROR")
	for _, task := range reply.Tasks {
		if !opts.all && task.CurrentState != StateRunning {
			continue
		}
		if opts.quiet {
			c.Console().Println(task.Id)
		} else {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", task.Id, task.Image, task.DesiredState, task.CurrentState, task.NodeId, task.Error)
		}
	}
	if !opts.quiet {
		w.Flush()
	}
	return nil
}
