//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

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
func (in *LandscaperServiceConfiguration) DeepCopyInto(out *LandscaperServiceConfiguration) {
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
