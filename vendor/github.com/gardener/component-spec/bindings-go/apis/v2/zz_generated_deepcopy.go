//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v2

import (
	json "encoding/json"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDescriptor) DeepCopyInto(out *ComponentDescriptor) {
	*out = *in
	out.Metadata = in.Metadata
	in.ComponentSpec.DeepCopyInto(&out.ComponentSpec)
	if in.Signatures != nil {
		in, out := &in.Signatures, &out.Signatures
		*out = make([]Signature, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDescriptor.
func (in *ComponentDescriptor) DeepCopy() *ComponentDescriptor {
	if in == nil {
		return nil
	}
	out := new(ComponentDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentReference) DeepCopyInto(out *ComponentReference) {
	*out = *in
	if in.ExtraIdentity != nil {
		in, out := &in.ExtraIdentity, &out.ExtraIdentity
		*out = make(Identity, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Digest != nil {
		in, out := &in.Digest, &out.Digest
		*out = new(DigestSpec)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(Labels, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentReference.
func (in *ComponentReference) DeepCopy() *ComponentReference {
	if in == nil {
		return nil
	}
	out := new(ComponentReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.RepositoryContexts != nil {
		in, out := &in.RepositoryContexts, &out.RepositoryContexts
		*out = make([]*UnstructuredTypedObject, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]Source, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ComponentReferences != nil {
		in, out := &in.ComponentReferences, &out.ComponentReferences
		*out = make([]ComponentReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]Resource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpec.
func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DigestSpec) DeepCopyInto(out *DigestSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DigestSpec.
func (in *DigestSpec) DeepCopy() *DigestSpec {
	if in == nil {
		return nil
	}
	out := new(DigestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Identity) DeepCopyInto(out *Identity) {
	{
		in := &in
		*out = make(Identity, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity.
func (in Identity) DeepCopy() Identity {
	if in == nil {
		return nil
	}
	out := new(Identity)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObjectMeta) DeepCopyInto(out *IdentityObjectMeta) {
	*out = *in
	if in.ExtraIdentity != nil {
		in, out := &in.ExtraIdentity, &out.ExtraIdentity
		*out = make(Identity, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(Labels, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObjectMeta.
func (in *IdentityObjectMeta) DeepCopy() *IdentityObjectMeta {
	if in == nil {
		return nil
	}
	out := new(IdentityObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Label) DeepCopyInto(out *Label) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Label.
func (in *Label) DeepCopy() *Label {
	if in == nil {
		return nil
	}
	out := new(Label)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Labels) DeepCopyInto(out *Labels) {
	{
		in := &in
		*out = make(Labels, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Labels.
func (in Labels) DeepCopy() Labels {
	if in == nil {
		return nil
	}
	out := new(Labels)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(Labels, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectType) DeepCopyInto(out *ObjectType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectType.
func (in *ObjectType) DeepCopy() *ObjectType {
	if in == nil {
		return nil
	}
	out := new(ObjectType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	in.IdentityObjectMeta.DeepCopyInto(&out.IdentityObjectMeta)
	if in.Digest != nil {
		in, out := &in.Digest, &out.Digest
		*out = new(DigestSpec)
		**out = **in
	}
	if in.SourceRef != nil {
		in, out := &in.SourceRef, &out.SourceRef
		*out = make([]SourceRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Signature) DeepCopyInto(out *Signature) {
	*out = *in
	out.Digest = in.Digest
	out.Signature = in.Signature
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Signature.
func (in *Signature) DeepCopy() *Signature {
	if in == nil {
		return nil
	}
	out := new(Signature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureSpec) DeepCopyInto(out *SignatureSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureSpec.
func (in *SignatureSpec) DeepCopy() *SignatureSpec {
	if in == nil {
		return nil
	}
	out := new(SignatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	in.IdentityObjectMeta.DeepCopyInto(&out.IdentityObjectMeta)
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceRef) DeepCopyInto(out *SourceRef) {
	*out = *in
	if in.IdentitySelector != nil {
		in, out := &in.IdentitySelector, &out.IdentitySelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(Labels, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceRef.
func (in *SourceRef) DeepCopy() *SourceRef {
	if in == nil {
		return nil
	}
	out := new(SourceRef)
	in.DeepCopyInto(out)
	return out
}
