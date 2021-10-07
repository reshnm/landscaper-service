//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package core

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Error) DeepCopyInto(out *Error) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Error.
func (in *Error) DeepCopy() *Error {
	if in == nil {
		return nil
	}
	out := new(Error)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	out.SeedConfigRef = in.SeedConfigRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	if in.LastError != nil {
		in, out := &in.LastError, &out.LastError
		*out = new(Error)
		**out = **in
	}
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(ObjectReference)
		**out = **in
	}
	if in.InstallationRef != nil {
		in, out := &in.InstallationRef, &out.InstallationRef
		*out = new(ObjectReference)
		**out = **in
	}
	if in.ClusterEndpointRef != nil {
		in, out := &in.ClusterEndpointRef, &out.ClusterEndpointRef
		*out = new(ObjectReference)
		**out = **in
	}
	if in.ClusterKubeconfigRef != nil {
		in, out := &in.ClusterKubeconfigRef, &out.ClusterKubeconfigRef
		*out = new(ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscaperDeployment) DeepCopyInto(out *LandscaperDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscaperDeployment.
func (in *LandscaperDeployment) DeepCopy() *LandscaperDeployment {
	if in == nil {
		return nil
	}
	out := new(LandscaperDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LandscaperDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscaperDeploymentSpec) DeepCopyInto(out *LandscaperDeploymentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscaperDeploymentSpec.
func (in *LandscaperDeploymentSpec) DeepCopy() *LandscaperDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(LandscaperDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscaperDeploymentStatus) DeepCopyInto(out *LandscaperDeploymentStatus) {
	*out = *in
	if in.LastError != nil {
		in, out := &in.LastError, &out.LastError
		*out = new(Error)
		**out = **in
	}
	if in.InstanceRef != nil {
		in, out := &in.InstanceRef, &out.InstanceRef
		*out = new(ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscaperDeploymentStatus.
func (in *LandscaperDeploymentStatus) DeepCopy() *LandscaperDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(LandscaperDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedConfig) DeepCopyInto(out *SeedConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedConfig.
func (in *SeedConfig) DeepCopy() *SeedConfig {
	if in == nil {
		return nil
	}
	out := new(SeedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SeedConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedConfigSpec) DeepCopyInto(out *SeedConfigSpec) {
	*out = *in
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedConfigSpec.
func (in *SeedConfigSpec) DeepCopy() *SeedConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SeedConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedConfigStatus) DeepCopyInto(out *SeedConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedConfigStatus.
func (in *SeedConfigStatus) DeepCopy() *SeedConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SeedConfigStatus)
	in.DeepCopyInto(out)
	return out
}
