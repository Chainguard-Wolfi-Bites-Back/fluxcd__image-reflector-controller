//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/fluxcd/pkg/apis/acl"
	"github.com/fluxcd/pkg/apis/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlphabeticalPolicy) DeepCopyInto(out *AlphabeticalPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlphabeticalPolicy.
func (in *AlphabeticalPolicy) DeepCopy() *AlphabeticalPolicy {
	if in == nil {
		return nil
	}
	out := new(AlphabeticalPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicy) DeepCopyInto(out *ImagePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicy.
func (in *ImagePolicy) DeepCopy() *ImagePolicy {
	if in == nil {
		return nil
	}
	out := new(ImagePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicyChoice) DeepCopyInto(out *ImagePolicyChoice) {
	*out = *in
	if in.SemVer != nil {
		in, out := &in.SemVer, &out.SemVer
		*out = new(SemVerPolicy)
		**out = **in
	}
	if in.Alphabetical != nil {
		in, out := &in.Alphabetical, &out.Alphabetical
		*out = new(AlphabeticalPolicy)
		**out = **in
	}
	if in.Numerical != nil {
		in, out := &in.Numerical, &out.Numerical
		*out = new(NumericalPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicyChoice.
func (in *ImagePolicyChoice) DeepCopy() *ImagePolicyChoice {
	if in == nil {
		return nil
	}
	out := new(ImagePolicyChoice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicyList) DeepCopyInto(out *ImagePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImagePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicyList.
func (in *ImagePolicyList) DeepCopy() *ImagePolicyList {
	if in == nil {
		return nil
	}
	out := new(ImagePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicySpec) DeepCopyInto(out *ImagePolicySpec) {
	*out = *in
	out.ImageRepositoryRef = in.ImageRepositoryRef
	in.Policy.DeepCopyInto(&out.Policy)
	if in.FilterTags != nil {
		in, out := &in.FilterTags, &out.FilterTags
		*out = new(TagFilter)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicySpec.
func (in *ImagePolicySpec) DeepCopy() *ImagePolicySpec {
	if in == nil {
		return nil
	}
	out := new(ImagePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicyStatus) DeepCopyInto(out *ImagePolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicyStatus.
func (in *ImagePolicyStatus) DeepCopy() *ImagePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ImagePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRepository) DeepCopyInto(out *ImageRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRepository.
func (in *ImageRepository) DeepCopy() *ImageRepository {
	if in == nil {
		return nil
	}
	out := new(ImageRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRepositoryList) DeepCopyInto(out *ImageRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRepositoryList.
func (in *ImageRepositoryList) DeepCopy() *ImageRepositoryList {
	if in == nil {
		return nil
	}
	out := new(ImageRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRepositorySpec) DeepCopyInto(out *ImageRepositorySpec) {
	*out = *in
	out.Interval = in.Interval
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(meta.LocalObjectReference)
		**out = **in
	}
	if in.CertSecretRef != nil {
		in, out := &in.CertSecretRef, &out.CertSecretRef
		*out = new(meta.LocalObjectReference)
		**out = **in
	}
	if in.AccessFrom != nil {
		in, out := &in.AccessFrom, &out.AccessFrom
		*out = new(acl.AccessFrom)
		(*in).DeepCopyInto(*out)
	}
	if in.ExclusionList != nil {
		in, out := &in.ExclusionList, &out.ExclusionList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRepositorySpec.
func (in *ImageRepositorySpec) DeepCopy() *ImageRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(ImageRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRepositoryStatus) DeepCopyInto(out *ImageRepositoryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastScanResult != nil {
		in, out := &in.LastScanResult, &out.LastScanResult
		*out = new(ScanResult)
		(*in).DeepCopyInto(*out)
	}
	if in.ObservedExclusionList != nil {
		in, out := &in.ObservedExclusionList, &out.ObservedExclusionList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRepositoryStatus.
func (in *ImageRepositoryStatus) DeepCopy() *ImageRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(ImageRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumericalPolicy) DeepCopyInto(out *NumericalPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumericalPolicy.
func (in *NumericalPolicy) DeepCopy() *NumericalPolicy {
	if in == nil {
		return nil
	}
	out := new(NumericalPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanResult) DeepCopyInto(out *ScanResult) {
	*out = *in
	in.ScanTime.DeepCopyInto(&out.ScanTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanResult.
func (in *ScanResult) DeepCopy() *ScanResult {
	if in == nil {
		return nil
	}
	out := new(ScanResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SemVerPolicy) DeepCopyInto(out *SemVerPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SemVerPolicy.
func (in *SemVerPolicy) DeepCopy() *SemVerPolicy {
	if in == nil {
		return nil
	}
	out := new(SemVerPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagFilter) DeepCopyInto(out *TagFilter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagFilter.
func (in *TagFilter) DeepCopy() *TagFilter {
	if in == nil {
		return nil
	}
	out := new(TagFilter)
	in.DeepCopyInto(out)
	return out
}
