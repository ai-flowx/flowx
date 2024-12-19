package tool

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type FlowX struct {
	Type string
}

func (f *FlowX) Init(_ context.Context) error {
	return nil
}

func (f *FlowX) Deinit(_ context.Context) error {
	return nil
}

func (f *FlowX) Run(ctx context.Context, invokes []*Invoke) error {
	var err error

	fmt.Println("FlowX Tools")

	for _, invoke := range invokes {
		_ = f.dump(invoke)
		if invoke.Path != "" {
			err = errors.New("TBD: FIXME\n")
		} else if invoke.Func != nil {
			if invoke.Result, err = invoke.Func(ctx, invoke.Args); err != nil {
				break
			}
		} else {
			continue
		}
	}

	return err
}

func (f *FlowX) dump(invoke *Invoke) error {
	var args string

	if buf, err := json.Marshal(invoke.Args); err == nil {
		args = string(buf)
	}

	fmt.Printf("       Name: %s\n", invoke.Name)
	fmt.Printf("Description: %s\n", invoke.Description)
	fmt.Printf("       Path: %s\n", invoke.Path)
	fmt.Printf("       Func: %T\n", invoke.Func)
	fmt.Printf("       Args: %T\n", args)
	fmt.Printf("     Result: %s\n", invoke.Result)

	return nil
}
