// +build !ignore_autogenerated

/*

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

package v1

import (
	"github.com/newrelic/newrelic-client-go/pkg/alerts"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APMSpecificSpec) DeepCopyInto(out *APMSpecificSpec) {
	*out = *in
	out.UserDefined = in.UserDefined
	if in.Entities != nil {
		in, out := &in.Entities, &out.Entities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APMSpecificSpec.
func (in *APMSpecificSpec) DeepCopy() *APMSpecificSpec {
	if in == nil {
		return nil
	}
	out := new(APMSpecificSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionTerm) DeepCopyInto(out *AlertConditionTerm) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionTerm.
func (in *AlertConditionTerm) DeepCopy() *AlertConditionTerm {
	if in == nil {
		return nil
	}
	out := new(AlertConditionTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsAPMCondition) DeepCopyInto(out *AlertsAPMCondition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsAPMCondition.
func (in *AlertsAPMCondition) DeepCopy() *AlertsAPMCondition {
	if in == nil {
		return nil
	}
	out := new(AlertsAPMCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertsAPMCondition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsAPMConditionList) DeepCopyInto(out *AlertsAPMConditionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertsAPMCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsAPMConditionList.
func (in *AlertsAPMConditionList) DeepCopy() *AlertsAPMConditionList {
	if in == nil {
		return nil
	}
	out := new(AlertsAPMConditionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertsAPMConditionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsAPMConditionSpec) DeepCopyInto(out *AlertsAPMConditionSpec) {
	*out = *in
	in.AlertsGenericConditionSpec.DeepCopyInto(&out.AlertsGenericConditionSpec)
	in.AlertsAPMSpecificSpec.DeepCopyInto(&out.AlertsAPMSpecificSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsAPMConditionSpec.
func (in *AlertsAPMConditionSpec) DeepCopy() *AlertsAPMConditionSpec {
	if in == nil {
		return nil
	}
	out := new(AlertsAPMConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsAPMConditionStatus) DeepCopyInto(out *AlertsAPMConditionStatus) {
	*out = *in
	if in.AppliedSpec != nil {
		in, out := &in.AppliedSpec, &out.AppliedSpec
		*out = new(AlertsAPMConditionSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsAPMConditionStatus.
func (in *AlertsAPMConditionStatus) DeepCopy() *AlertsAPMConditionStatus {
	if in == nil {
		return nil
	}
	out := new(AlertsAPMConditionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsAPMSpecificSpec) DeepCopyInto(out *AlertsAPMSpecificSpec) {
	*out = *in
	out.UserDefined = in.UserDefined
	if in.Entities != nil {
		in, out := &in.Entities, &out.Entities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsAPMSpecificSpec.
func (in *AlertsAPMSpecificSpec) DeepCopy() *AlertsAPMSpecificSpec {
	if in == nil {
		return nil
	}
	out := new(AlertsAPMSpecificSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsGenericConditionSpec) DeepCopyInto(out *AlertsGenericConditionSpec) {
	*out = *in
	out.APIKeySecret = in.APIKeySecret
	if in.Terms != nil {
		in, out := &in.Terms, &out.Terms
		*out = make([]AlertsNrqlConditionTerm, len(*in))
		copy(*out, *in)
	}
	if in.APMTerms != nil {
		in, out := &in.APMTerms, &out.APMTerms
		*out = make([]AlertConditionTerm, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsGenericConditionSpec.
func (in *AlertsGenericConditionSpec) DeepCopy() *AlertsGenericConditionSpec {
	if in == nil {
		return nil
	}
	out := new(AlertsGenericConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsNrqlCondition) DeepCopyInto(out *AlertsNrqlCondition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsNrqlCondition.
func (in *AlertsNrqlCondition) DeepCopy() *AlertsNrqlCondition {
	if in == nil {
		return nil
	}
	out := new(AlertsNrqlCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertsNrqlCondition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsNrqlConditionList) DeepCopyInto(out *AlertsNrqlConditionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertsNrqlCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsNrqlConditionList.
func (in *AlertsNrqlConditionList) DeepCopy() *AlertsNrqlConditionList {
	if in == nil {
		return nil
	}
	out := new(AlertsNrqlConditionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertsNrqlConditionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsNrqlConditionQuery) DeepCopyInto(out *AlertsNrqlConditionQuery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsNrqlConditionQuery.
func (in *AlertsNrqlConditionQuery) DeepCopy() *AlertsNrqlConditionQuery {
	if in == nil {
		return nil
	}
	out := new(AlertsNrqlConditionQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsNrqlConditionSpec) DeepCopyInto(out *AlertsNrqlConditionSpec) {
	*out = *in
	in.AlertsGenericConditionSpec.DeepCopyInto(&out.AlertsGenericConditionSpec)
	in.AlertsNrqlSpecificSpec.DeepCopyInto(&out.AlertsNrqlSpecificSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsNrqlConditionSpec.
func (in *AlertsNrqlConditionSpec) DeepCopy() *AlertsNrqlConditionSpec {
	if in == nil {
		return nil
	}
	out := new(AlertsNrqlConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsNrqlConditionStatus) DeepCopyInto(out *AlertsNrqlConditionStatus) {
	*out = *in
	if in.AppliedSpec != nil {
		in, out := &in.AppliedSpec, &out.AppliedSpec
		*out = new(AlertsNrqlConditionSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsNrqlConditionStatus.
func (in *AlertsNrqlConditionStatus) DeepCopy() *AlertsNrqlConditionStatus {
	if in == nil {
		return nil
	}
	out := new(AlertsNrqlConditionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsNrqlConditionTerm) DeepCopyInto(out *AlertsNrqlConditionTerm) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsNrqlConditionTerm.
func (in *AlertsNrqlConditionTerm) DeepCopy() *AlertsNrqlConditionTerm {
	if in == nil {
		return nil
	}
	out := new(AlertsNrqlConditionTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsNrqlSpecificSpec) DeepCopyInto(out *AlertsNrqlSpecificSpec) {
	*out = *in
	out.Nrql = in.Nrql
	if in.ValueFunction != nil {
		in, out := &in.ValueFunction, &out.ValueFunction
		*out = new(alerts.NrqlConditionValueFunction)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsNrqlSpecificSpec.
func (in *AlertsNrqlSpecificSpec) DeepCopy() *AlertsNrqlSpecificSpec {
	if in == nil {
		return nil
	}
	out := new(AlertsNrqlSpecificSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsPolicy) DeepCopyInto(out *AlertsPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsPolicy.
func (in *AlertsPolicy) DeepCopy() *AlertsPolicy {
	if in == nil {
		return nil
	}
	out := new(AlertsPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertsPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsPolicyCondition) DeepCopyInto(out *AlertsPolicyCondition) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsPolicyCondition.
func (in *AlertsPolicyCondition) DeepCopy() *AlertsPolicyCondition {
	if in == nil {
		return nil
	}
	out := new(AlertsPolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsPolicyConditionSpec) DeepCopyInto(out *AlertsPolicyConditionSpec) {
	*out = *in
	in.AlertsGenericConditionSpec.DeepCopyInto(&out.AlertsGenericConditionSpec)
	in.AlertsNrqlSpecificSpec.DeepCopyInto(&out.AlertsNrqlSpecificSpec)
	in.AlertsAPMSpecificSpec.DeepCopyInto(&out.AlertsAPMSpecificSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsPolicyConditionSpec.
func (in *AlertsPolicyConditionSpec) DeepCopy() *AlertsPolicyConditionSpec {
	if in == nil {
		return nil
	}
	out := new(AlertsPolicyConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsPolicyList) DeepCopyInto(out *AlertsPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertsPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsPolicyList.
func (in *AlertsPolicyList) DeepCopy() *AlertsPolicyList {
	if in == nil {
		return nil
	}
	out := new(AlertsPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertsPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsPolicySpec) DeepCopyInto(out *AlertsPolicySpec) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AlertsPolicyCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.APIKeySecret = in.APIKeySecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsPolicySpec.
func (in *AlertsPolicySpec) DeepCopy() *AlertsPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AlertsPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsPolicyStatus) DeepCopyInto(out *AlertsPolicyStatus) {
	*out = *in
	if in.AppliedSpec != nil {
		in, out := &in.AppliedSpec, &out.AppliedSpec
		*out = new(AlertsPolicySpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsPolicyStatus.
func (in *AlertsPolicyStatus) DeepCopy() *AlertsPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AlertsPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApmAlertCondition) DeepCopyInto(out *ApmAlertCondition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApmAlertCondition.
func (in *ApmAlertCondition) DeepCopy() *ApmAlertCondition {
	if in == nil {
		return nil
	}
	out := new(ApmAlertCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApmAlertCondition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApmAlertConditionList) DeepCopyInto(out *ApmAlertConditionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApmAlertCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApmAlertConditionList.
func (in *ApmAlertConditionList) DeepCopy() *ApmAlertConditionList {
	if in == nil {
		return nil
	}
	out := new(ApmAlertConditionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApmAlertConditionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApmAlertConditionSpec) DeepCopyInto(out *ApmAlertConditionSpec) {
	*out = *in
	in.GenericConditionSpec.DeepCopyInto(&out.GenericConditionSpec)
	in.APMSpecificSpec.DeepCopyInto(&out.APMSpecificSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApmAlertConditionSpec.
func (in *ApmAlertConditionSpec) DeepCopy() *ApmAlertConditionSpec {
	if in == nil {
		return nil
	}
	out := new(ApmAlertConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApmAlertConditionStatus) DeepCopyInto(out *ApmAlertConditionStatus) {
	*out = *in
	if in.AppliedSpec != nil {
		in, out := &in.AppliedSpec, &out.AppliedSpec
		*out = new(ApmAlertConditionSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApmAlertConditionStatus.
func (in *ApmAlertConditionStatus) DeepCopy() *ApmAlertConditionStatus {
	if in == nil {
		return nil
	}
	out := new(ApmAlertConditionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionSpec) DeepCopyInto(out *ConditionSpec) {
	*out = *in
	in.GenericConditionSpec.DeepCopyInto(&out.GenericConditionSpec)
	out.NrqlSpecificSpec = in.NrqlSpecificSpec
	in.APMSpecificSpec.DeepCopyInto(&out.APMSpecificSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionSpec.
func (in *ConditionSpec) DeepCopy() *ConditionSpec {
	if in == nil {
		return nil
	}
	out := new(ConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericConditionSpec) DeepCopyInto(out *GenericConditionSpec) {
	*out = *in
	if in.Terms != nil {
		in, out := &in.Terms, &out.Terms
		*out = make([]AlertConditionTerm, len(*in))
		copy(*out, *in)
	}
	out.APIKeySecret = in.APIKeySecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericConditionSpec.
func (in *GenericConditionSpec) DeepCopy() *GenericConditionSpec {
	if in == nil {
		return nil
	}
	out := new(GenericConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in InvalidAttributeSlice) DeepCopyInto(out *InvalidAttributeSlice) {
	{
		in := &in
		*out = make(InvalidAttributeSlice, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvalidAttributeSlice.
func (in InvalidAttributeSlice) DeepCopy() InvalidAttributeSlice {
	if in == nil {
		return nil
	}
	out := new(InvalidAttributeSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NewRelicAPIKeySecret) DeepCopyInto(out *NewRelicAPIKeySecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NewRelicAPIKeySecret.
func (in *NewRelicAPIKeySecret) DeepCopy() *NewRelicAPIKeySecret {
	if in == nil {
		return nil
	}
	out := new(NewRelicAPIKeySecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NrqlAlertCondition) DeepCopyInto(out *NrqlAlertCondition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NrqlAlertCondition.
func (in *NrqlAlertCondition) DeepCopy() *NrqlAlertCondition {
	if in == nil {
		return nil
	}
	out := new(NrqlAlertCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NrqlAlertCondition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NrqlAlertConditionList) DeepCopyInto(out *NrqlAlertConditionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NrqlAlertCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NrqlAlertConditionList.
func (in *NrqlAlertConditionList) DeepCopy() *NrqlAlertConditionList {
	if in == nil {
		return nil
	}
	out := new(NrqlAlertConditionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NrqlAlertConditionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NrqlAlertConditionSpec) DeepCopyInto(out *NrqlAlertConditionSpec) {
	*out = *in
	in.GenericConditionSpec.DeepCopyInto(&out.GenericConditionSpec)
	out.NrqlSpecificSpec = in.NrqlSpecificSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NrqlAlertConditionSpec.
func (in *NrqlAlertConditionSpec) DeepCopy() *NrqlAlertConditionSpec {
	if in == nil {
		return nil
	}
	out := new(NrqlAlertConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NrqlAlertConditionStatus) DeepCopyInto(out *NrqlAlertConditionStatus) {
	*out = *in
	if in.AppliedSpec != nil {
		in, out := &in.AppliedSpec, &out.AppliedSpec
		*out = new(NrqlAlertConditionSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NrqlAlertConditionStatus.
func (in *NrqlAlertConditionStatus) DeepCopy() *NrqlAlertConditionStatus {
	if in == nil {
		return nil
	}
	out := new(NrqlAlertConditionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NrqlQuery) DeepCopyInto(out *NrqlQuery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NrqlQuery.
func (in *NrqlQuery) DeepCopy() *NrqlQuery {
	if in == nil {
		return nil
	}
	out := new(NrqlQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NrqlSpecificSpec) DeepCopyInto(out *NrqlSpecificSpec) {
	*out = *in
	out.Nrql = in.Nrql
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NrqlSpecificSpec.
func (in *NrqlSpecificSpec) DeepCopy() *NrqlSpecificSpec {
	if in == nil {
		return nil
	}
	out := new(NrqlSpecificSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyCondition) DeepCopyInto(out *PolicyCondition) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyCondition.
func (in *PolicyCondition) DeepCopy() *PolicyCondition {
	if in == nil {
		return nil
	}
	out := new(PolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	out.APIKeySecret = in.APIKeySecret
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PolicyCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
	if in.AppliedSpec != nil {
		in, out := &in.AppliedSpec, &out.AppliedSpec
		*out = new(PolicySpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}
