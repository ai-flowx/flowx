package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/ai-flowx/flowx/config"
	"github.com/ai-flowx/flowx/flow"
)

const (
	routineNum = -1
)

var (
	configFile string
	configData config.Config
	listenAddr string
)

var rootCmd = &cobra.Command{
	Use:     "flowx",
	Version: config.Version + "-build-" + config.Build,
	Short:   "ai framework",
	Long:    fmt.Sprintf("ai framework %s (%s %s)", config.Version, config.Commit, config.Build),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		if err := viper.Unmarshal(&configData); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		f, err := initFlow(ctx)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if err := runFlow(ctx, f); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// nolint:gochecknoinits
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config-file", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&listenAddr, "listen-addr", "u", "127.0.0.1:8080", "listen address")

	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

func initConfig() {
	if configFile == "" {
		return
	}

	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func initFlow(ctx context.Context) (flow.Flow, error) {
	c := flow.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config\n")
	}

	c.Addr = listenAddr
	c.Cache = configData.Cache
	c.Gpt = configData.Gpt
	c.Store = configData.Store

	return flow.New(ctx, c), nil
}

func runFlow(ctx context.Context, _flow flow.Flow) error {
	if err := _flow.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init\n")
	}

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(routineNum)

	g.Go(func() error {
		if err := _flow.Run(ctx); err != nil {
			return errors.Wrap(err, "failed to run\n")
		}
		return nil
	})

	s := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be caught, so don't need add it
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	g.Go(func() error {
		<-s
		if err := _flow.Deinit(ctx); err != nil {
			return errors.Wrap(err, "failed to deinit\n")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
