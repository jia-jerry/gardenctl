// +build !ignore_autogenerated

/*
Copyright 2017 The Gardener Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConnectionConfiguration) DeepCopyInto(out *ClientConnectionConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConnectionConfiguration.
func (in *ClientConnectionConfiguration) DeepCopy() *ClientConnectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClientConnectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerReconciliationConfiguration) DeepCopyInto(out *ControllerReconciliationConfiguration) {
	*out = *in
	out.ResyncPeriod = in.ResyncPeriod
	if in.RetryDuration != nil {
		in, out := &in.RetryDuration, &out.RetryDuration
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerReconciliationConfiguration.
func (in *ControllerReconciliationConfiguration) DeepCopy() *ControllerReconciliationConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerReconciliationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenConstraints) DeepCopyInto(out *GardenConstraints) {
	*out = *in
	if in.CloudProviders != nil {
		in, out := &in.CloudProviders, &out.CloudProviders
		*out = make([]GardenOperatorCloudProviderConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KubernetesVersions != nil {
		in, out := &in.KubernetesVersions, &out.KubernetesVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DNSProviders != nil {
		in, out := &in.DNSProviders, &out.DNSProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenConstraints.
func (in *GardenConstraints) DeepCopy() *GardenConstraints {
	if in == nil {
		return nil
	}
	out := new(GardenConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenOperatorCloudProviderConfiguration) DeepCopyInto(out *GardenOperatorCloudProviderConfiguration) {
	*out = *in
	if in.VolumeTypes != nil {
		in, out := &in.VolumeTypes, &out.VolumeTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MachineTypes != nil {
		in, out := &in.MachineTypes, &out.MachineTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenOperatorCloudProviderConfiguration.
func (in *GardenOperatorCloudProviderConfiguration) DeepCopy() *GardenOperatorCloudProviderConfiguration {
	if in == nil {
		return nil
	}
	out := new(GardenOperatorCloudProviderConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenOperatorConfiguration) DeepCopyInto(out *GardenOperatorConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ClientConnection = in.ClientConnection
	in.Controller.DeepCopyInto(&out.Controller)
	in.Constraints.DeepCopyInto(&out.Constraints)
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]GardenOperatorImagesConfiguration, len(*in))
		copy(*out, *in)
	}
	out.LeaderElection = in.LeaderElection
	out.Server = in.Server
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenOperatorConfiguration.
func (in *GardenOperatorConfiguration) DeepCopy() *GardenOperatorConfiguration {
	if in == nil {
		return nil
	}
	out := new(GardenOperatorConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GardenOperatorConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenOperatorControllerConfiguration) DeepCopyInto(out *GardenOperatorControllerConfiguration) {
	*out = *in
	out.HealthCheckPeriod = in.HealthCheckPeriod
	in.Reconciliation.DeepCopyInto(&out.Reconciliation)
	if in.WatchNamespace != nil {
		in, out := &in.WatchNamespace, &out.WatchNamespace
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenOperatorControllerConfiguration.
func (in *GardenOperatorControllerConfiguration) DeepCopy() *GardenOperatorControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(GardenOperatorControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GardenOperatorImagesConfiguration) DeepCopyInto(out *GardenOperatorImagesConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GardenOperatorImagesConfiguration.
func (in *GardenOperatorImagesConfiguration) DeepCopy() *GardenOperatorImagesConfiguration {
	if in == nil {
		return nil
	}
	out := new(GardenOperatorImagesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionConfiguration) DeepCopyInto(out *LeaderElectionConfiguration) {
	*out = *in
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionConfiguration.
func (in *LeaderElectionConfiguration) DeepCopy() *LeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}