//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditLogConfiguration) DeepCopyInto(out *AuditLogConfiguration) {
	*out = *in
	out.AuditPolicy = in.AuditPolicy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditLogConfiguration.
func (in *AuditLogConfiguration) DeepCopy() *AuditLogConfiguration {
	if in == nil {
		return nil
	}
	out := new(AuditLogConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityMonitoringConfiguration) DeepCopyInto(out *AvailabilityMonitoringConfiguration) {
	*out = *in
	if in.AvailabilityServiceConfiguration != nil {
		in, out := &in.AvailabilityServiceConfiguration, &out.AvailabilityServiceConfiguration
		*out = new(AvailabilityServiceConfiguration)
		**out = **in
	}
	out.PeriodicCheckInterval = in.PeriodicCheckInterval
	out.LSHealthCheckTimeout = in.LSHealthCheckTimeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityMonitoringConfiguration.
func (in *AvailabilityMonitoringConfiguration) DeepCopy() *AvailabilityMonitoringConfiguration {
	if in == nil {
		return nil
	}
	out := new(AvailabilityMonitoringConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityServiceConfiguration) DeepCopyInto(out *AvailabilityServiceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityServiceConfiguration.
func (in *AvailabilityServiceConfiguration) DeepCopy() *AvailabilityServiceConfiguration {
	if in == nil {
		return nil
	}
	out := new(AvailabilityServiceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlane) DeepCopyInto(out *ControlPlane) {
	*out = *in
	out.HighAvailability = in.HighAvailability
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlane.
func (in *ControlPlane) DeepCopy() *ControlPlane {
	if in == nil {
		return nil
	}
	out := new(ControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrdManagementConfiguration) DeepCopyInto(out *CrdManagementConfiguration) {
	*out = *in
	if in.DeployCustomResourceDefinitions != nil {
		in, out := &in.DeployCustomResourceDefinitions, &out.DeployCustomResourceDefinitions
		*out = new(bool)
		**out = **in
	}
	if in.ForceUpdate != nil {
		in, out := &in.ForceUpdate, &out.ForceUpdate
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrdManagementConfiguration.
func (in *CrdManagementConfiguration) DeepCopy() *CrdManagementConfiguration {
	if in == nil {
		return nil
	}
	out := new(CrdManagementConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenerConfiguration) DeepCopyInto(out *GardenerConfiguration) {
	*out = *in
	out.ServiceAccountKubeconfig = in.ServiceAccountKubeconfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenerConfiguration.
func (in *GardenerConfiguration) DeepCopy() *GardenerConfiguration {
	if in == nil {
		return nil
	}
	out := new(GardenerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailability) DeepCopyInto(out *HighAvailability) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HighAvailability.
func (in *HighAvailability) DeepCopy() *HighAvailability {
	if in == nil {
		return nil
	}
	out := new(HighAvailability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeAPIServerConfig) DeepCopyInto(out *KubeAPIServerConfig) {
	*out = *in
	out.OIDCConfig = in.OIDCConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeAPIServerConfig.
func (in *KubeAPIServerConfig) DeepCopy() *KubeAPIServerConfig {
	if in == nil {
		return nil
	}
	out := new(KubeAPIServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscaperServiceComponentConfiguration) DeepCopyInto(out *LandscaperServiceComponentConfiguration) {
	*out = *in
	in.RepositoryContext.DeepCopyInto(&out.RepositoryContext)
	if in.RegistryPullSecrets != nil {
		in, out := &in.RegistryPullSecrets, &out.RegistryPullSecrets
		*out = make([]v1.SecretReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscaperServiceComponentConfiguration.
func (in *LandscaperServiceComponentConfiguration) DeepCopy() *LandscaperServiceComponentConfiguration {
	if in == nil {
		return nil
	}
	out := new(LandscaperServiceComponentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscaperServiceConfiguration) DeepCopyInto(out *LandscaperServiceConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(MetricsConfiguration)
		**out = **in
	}
	in.AvailabilityMonitoring.DeepCopyInto(&out.AvailabilityMonitoring)
	in.CrdManagement.DeepCopyInto(&out.CrdManagement)
	in.LandscaperServiceComponent.DeepCopyInto(&out.LandscaperServiceComponent)
	out.GardenerConfiguration = in.GardenerConfiguration
	in.ShootConfiguration.DeepCopyInto(&out.ShootConfiguration)
	if in.AuditLogConfig != nil {
		in, out := &in.AuditLogConfig, &out.AuditLogConfig
		*out = new(AuditLogConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscaperServiceConfiguration.
func (in *LandscaperServiceConfiguration) DeepCopy() *LandscaperServiceConfiguration {
	if in == nil {
		return nil
	}
	out := new(LandscaperServiceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LandscaperServiceConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsConfiguration) DeepCopyInto(out *MetricsConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsConfiguration.
func (in *MetricsConfiguration) DeepCopy() *MetricsConfiguration {
	if in == nil {
		return nil
	}
	out := new(MetricsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCConfig) DeepCopyInto(out *OIDCConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCConfig.
func (in *OIDCConfig) DeepCopy() *OIDCConfig {
	if in == nil {
		return nil
	}
	out := new(OIDCConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootAutoUpdateConfig) DeepCopyInto(out *ShootAutoUpdateConfig) {
	*out = *in
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(bool)
		**out = **in
	}
	if in.MachineImageVersion != nil {
		in, out := &in.MachineImageVersion, &out.MachineImageVersion
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootAutoUpdateConfig.
func (in *ShootAutoUpdateConfig) DeepCopy() *ShootAutoUpdateConfig {
	if in == nil {
		return nil
	}
	out := new(ShootAutoUpdateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootConfiguration) DeepCopyInto(out *ShootConfiguration) {
	*out = *in
	out.Provider = in.Provider
	in.Workers.DeepCopyInto(&out.Workers)
	out.Kubernetes = in.Kubernetes
	in.Maintenance.DeepCopyInto(&out.Maintenance)
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(ControlPlane)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootConfiguration.
func (in *ShootConfiguration) DeepCopy() *ShootConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootKubernetesConfig) DeepCopyInto(out *ShootKubernetesConfig) {
	*out = *in
	out.KubeAPIServer = in.KubeAPIServer
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootKubernetesConfig.
func (in *ShootKubernetesConfig) DeepCopy() *ShootKubernetesConfig {
	if in == nil {
		return nil
	}
	out := new(ShootKubernetesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootMachineConfiguration) DeepCopyInto(out *ShootMachineConfiguration) {
	*out = *in
	out.Image = in.Image
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootMachineConfiguration.
func (in *ShootMachineConfiguration) DeepCopy() *ShootMachineConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootMachineConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootMachineImage) DeepCopyInto(out *ShootMachineImage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootMachineImage.
func (in *ShootMachineImage) DeepCopy() *ShootMachineImage {
	if in == nil {
		return nil
	}
	out := new(ShootMachineImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootMaintenanceConfig) DeepCopyInto(out *ShootMaintenanceConfig) {
	*out = *in
	in.AutoUpdate.DeepCopyInto(&out.AutoUpdate)
	out.TimeWindow = in.TimeWindow
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootMaintenanceConfig.
func (in *ShootMaintenanceConfig) DeepCopy() *ShootMaintenanceConfig {
	if in == nil {
		return nil
	}
	out := new(ShootMaintenanceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootMaintenanceTimeWindow) DeepCopyInto(out *ShootMaintenanceTimeWindow) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootMaintenanceTimeWindow.
func (in *ShootMaintenanceTimeWindow) DeepCopy() *ShootMaintenanceTimeWindow {
	if in == nil {
		return nil
	}
	out := new(ShootMaintenanceTimeWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootProviderConfiguration) DeepCopyInto(out *ShootProviderConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootProviderConfiguration.
func (in *ShootProviderConfiguration) DeepCopy() *ShootProviderConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootProviderConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootWorkerVolumeConfiguration) DeepCopyInto(out *ShootWorkerVolumeConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootWorkerVolumeConfiguration.
func (in *ShootWorkerVolumeConfiguration) DeepCopy() *ShootWorkerVolumeConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootWorkerVolumeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootWorkersConfiguration) DeepCopyInto(out *ShootWorkersConfiguration) {
	*out = *in
	out.Machine = in.Machine
	out.Volume = in.Volume
	if in.Minimum != nil {
		in, out := &in.Minimum, &out.Minimum
		*out = new(int32)
		**out = **in
	}
	if in.Maximum != nil {
		in, out := &in.Maximum, &out.Maximum
		*out = new(int32)
		**out = **in
	}
	if in.MaxSurge != nil {
		in, out := &in.MaxSurge, &out.MaxSurge
		*out = new(int32)
		**out = **in
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootWorkersConfiguration.
func (in *ShootWorkersConfiguration) DeepCopy() *ShootWorkersConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootWorkersConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetShootSidecarConfiguration) DeepCopyInto(out *TargetShootSidecarConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(MetricsConfiguration)
		**out = **in
	}
	in.CrdManagement.DeepCopyInto(&out.CrdManagement)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetShootSidecarConfiguration.
func (in *TargetShootSidecarConfiguration) DeepCopy() *TargetShootSidecarConfiguration {
	if in == nil {
		return nil
	}
	out := new(TargetShootSidecarConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TargetShootSidecarConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
