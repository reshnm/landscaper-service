// SPDX-FileCopyrightText: 2022 "SAP SE or an SAP affiliate company and Gardener contributors"
//
// SPDX-License-Identifier: Apache-2.0

package tests

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"

	cliutil "github.com/gardener/landscapercli/pkg/util"
	"k8s.io/apimachinery/pkg/types"

	"github.com/gardener/landscaper-service/test/integration/pkg/test"
)

type UninstallLAASTestRunner struct {
	ctx            context.Context
	log            logr.Logger
	config         *test.TestConfig
	clusterClients *test.ClusterClients
	clusterTargets *test.ClusterTargets
	testObjects    *test.SharedTestObjects
}

func (r *UninstallLAASTestRunner) Init(
	ctx context.Context, log logr.Logger, config *test.TestConfig,
	clusterClients *test.ClusterClients, clusterTargets *test.ClusterTargets, testObjects *test.SharedTestObjects) {
	r.ctx = ctx
	r.log = log.WithName(r.Name())
	r.config = config
	r.clusterClients = clusterClients
	r.clusterTargets = clusterTargets
	r.testObjects = testObjects
}

func (r *UninstallLAASTestRunner) Name() string {
	return "UninstallLAAS"
}

func (r *UninstallLAASTestRunner) Description() string {
	description := `This test uninstalls the Landscaper-As-A-Service controller.
The test succeeds when the installation has been deleted before the timeout expires.
`
	return description
}

func (r *UninstallLAASTestRunner) String() string {
	return r.Name()
}

func (r *UninstallLAASTestRunner) Run() error {
	r.log.Info("deleting laas service target config")
	if err := r.deleteServiceTargetConfig(); err != nil {
		return err
	}

	r.log.Info("deleting laas installation")
	if err := r.deleteInstallation(); err != nil {
		return err
	}

	return nil
}

func (r *UninstallLAASTestRunner) deleteServiceTargetConfig() error {
	serviceTargetConfig := r.testObjects.ServiceTargetConfigs[types.NamespacedName{Name: r.clusterTargets.LaasCluster.Name, Namespace: r.config.LaasNamespace}.String()]

	if err := r.clusterClients.TestCluster.Delete(r.ctx, serviceTargetConfig); err != nil {
		return fmt.Errorf("failed to delete service hostingTarget config: %w", err)
	}

	return nil
}

func (r *UninstallLAASTestRunner) deleteInstallation() error {
	installation := r.testObjects.Installations[types.NamespacedName{Name: "laas", Namespace: r.config.LaasNamespace}.String()]

	if err := r.clusterClients.TestCluster.Delete(r.ctx, installation); err != nil {
		return fmt.Errorf("failed to delete laas installation: %w", err)
	}

	timeout, err := cliutil.CheckAndWaitUntilObjectNotExistAnymore(
		r.clusterClients.TestCluster,
		types.NamespacedName{Name: installation.Name, Namespace: installation.Namespace},
		installation,
		r.config.SleepTime,
		r.config.MaxRetries)

	if err != nil {
		return fmt.Errorf("failed to wait for laas installation to be deleted: %w", err)
	}
	if timeout {
		return fmt.Errorf("waiting for laas installation to be deleted timed out")
	}

	return nil
}
