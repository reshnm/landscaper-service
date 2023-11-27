// SPDX-FileCopyrightText: 2023 "SAP SE or an SAP affiliate company and Gardener contributors"
//
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	goflag "flag"
	"os"

	"github.com/gardener/landscaper/controller-utils/pkg/logging"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	configinstall "github.com/gardener/landscaper-service/pkg/apis/config/install"
	"github.com/gardener/landscaper-service/pkg/apis/config/v1alpha1"

	flag "github.com/spf13/pflag"
	ctrl "sigs.k8s.io/controller-runtime"
)

// options holds the landscaper service controller options
type options struct {
	Log        logging.Logger // Log is the logger instance
	ConfigPath string         // ConfigPath is the path to the configuration file

	Config *v1alpha1.TargetShootSidecarConfiguration // Config is the parsed configuration
}

// NewOptions returns a new options instance
func NewOptions() *options {
	return &options{}
}

// AddFlags adds flags passed via command line
func (o *options) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.ConfigPath, "config", "", "Specify the path to the configuration file")
	logging.InitFlags(fs)
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

// Complete initializes the options instance and validates flags
func (o *options) Complete(ctx context.Context) error {
	log, err := logging.GetLogger()
	if err != nil {
		return err
	}
	o.Log = log
	ctrl.SetLogger(log.Logr())

	o.Config, err = o.parseConfigurationFile(ctx)
	if err != nil {
		return err
	}

	err = o.validate()
	return err
}

func (o *options) parseConfigurationFile(ctx context.Context) (*v1alpha1.TargetShootSidecarConfiguration, error) {
	configScheme := runtime.NewScheme()
	configinstall.Install(configScheme)
	decoder := serializer.NewCodecFactory(configScheme).UniversalDecoder()

	config := &v1alpha1.TargetShootSidecarConfiguration{}
	if len(o.ConfigPath) != 0 {
		data, err := os.ReadFile(o.ConfigPath)
		if err != nil {
			return nil, err
		}

		if _, _, err := decoder.Decode(data, nil, config); err != nil {
			return nil, err
		}
	}

	configScheme.Default(config)
	return config, nil
}

func (o *options) validate() error {
	return nil
}
