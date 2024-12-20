package tool

import (
	"context"
	"os/exec"
	"path/filepath"
)

const (
	extensionBash   = ".sh"
	extensionGo     = ".go"
	extensionPython = ".py"
)

type FlowX struct{}

func (f *FlowX) Init(_ context.Context) error {
	return nil
}

func (f *FlowX) Deinit(_ context.Context) error {
	return nil
}

// nolint:gosec
func (f *FlowX) Run(ctx context.Context, invokes []*Invoke) error {
	var cmd *exec.Cmd
	var res []byte
	var err error

	for _, invoke := range invokes {
		if invoke.Path != "" {
			ext := filepath.Ext(invoke.Path)
			if ext == extensionBash {
				cmd = exec.Command("bash", "-c", invoke.Path)
				res, err = cmd.Output()
			} else if ext == extensionGo {
				cmd = exec.Command("go", "run", invoke.Path)
				res, err = cmd.Output()
			} else if ext == extensionPython {
				cmd = exec.Command("python", invoke.Path)
				res, err = cmd.Output()
			} else {
				cmd = exec.Command(invoke.Path)
				res, err = cmd.Output()
			}
		} else if invoke.Func != nil {
			res, err = invoke.Func(ctx, invoke.Args)
		}
		if err == nil {
			invoke.Result = string(res)
		} else {
			invoke.Result = err.Error()
		}
	}

	return err
}
