//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&NetworkTrafficArgs{}, func(obj interface{}) {
		SetObjectDefaultNetworkTrafficArgs(obj.(*NetworkTrafficArgs)) 
	})
	scheme.AddTypeDefaultingFunc(&CoschedulingArgs{}, func(obj interface{}) { SetObjectDefaults_CoschedulingArgs(obj.(*CoschedulingArgs)) })
	scheme.AddTypeDefaultingFunc(&LoadVariationRiskBalancingArgs{}, func(obj interface{}) {
		SetObjectDefaults_LoadVariationRiskBalancingArgs(obj.(*LoadVariationRiskBalancingArgs))
	})
	scheme.AddTypeDefaultingFunc(&LowRiskOverCommitmentArgs{}, func(obj interface{}) { SetObjectDefaults_LowRiskOverCommitmentArgs(obj.(*LowRiskOverCommitmentArgs)) })
	scheme.AddTypeDefaultingFunc(&NetworkOverheadArgs{}, func(obj interface{}) { SetObjectDefaults_NetworkOverheadArgs(obj.(*NetworkOverheadArgs)) })
	scheme.AddTypeDefaultingFunc(&NodeResourceTopologyMatchArgs{}, func(obj interface{}) {
		SetObjectDefaults_NodeResourceTopologyMatchArgs(obj.(*NodeResourceTopologyMatchArgs))
	})
	scheme.AddTypeDefaultingFunc(&NodeResourcesAllocatableArgs{}, func(obj interface{}) {
		SetObjectDefaults_NodeResourcesAllocatableArgs(obj.(*NodeResourcesAllocatableArgs))
	})
	scheme.AddTypeDefaultingFunc(&PreemptionTolerationArgs{}, func(obj interface{}) { SetObjectDefaults_PreemptionTolerationArgs(obj.(*PreemptionTolerationArgs)) })
	scheme.AddTypeDefaultingFunc(&SySchedArgs{}, func(obj interface{}) { SetObjectDefaults_SySchedArgs(obj.(*SySchedArgs)) })
	scheme.AddTypeDefaultingFunc(&TargetLoadPackingArgs{}, func(obj interface{}) { SetObjectDefaults_TargetLoadPackingArgs(obj.(*TargetLoadPackingArgs)) })
	scheme.AddTypeDefaultingFunc(&TopologicalSortArgs{}, func(obj interface{}) { SetObjectDefaults_TopologicalSortArgs(obj.(*TopologicalSortArgs)) })
	return nil
}

func SetObjectDefaultNetworkTrafficArgs(in *NetworkTrafficArgs) {
	SetDefaultNetworkTrafficArgs(in)
}

func SetObjectDefaults_CoschedulingArgs(in *CoschedulingArgs) {
	SetDefaults_CoschedulingArgs(in)
}

func SetObjectDefaults_LoadVariationRiskBalancingArgs(in *LoadVariationRiskBalancingArgs) {
	SetDefaults_LoadVariationRiskBalancingArgs(in)
}

func SetObjectDefaults_LowRiskOverCommitmentArgs(in *LowRiskOverCommitmentArgs) {
	SetDefaults_LowRiskOverCommitmentArgs(in)
}

func SetObjectDefaults_NetworkOverheadArgs(in *NetworkOverheadArgs) {
	SetDefaults_NetworkOverheadArgs(in)
}

func SetObjectDefaults_NodeResourceTopologyMatchArgs(in *NodeResourceTopologyMatchArgs) {
	SetDefaults_NodeResourceTopologyMatchArgs(in)
}

func SetObjectDefaults_NodeResourcesAllocatableArgs(in *NodeResourcesAllocatableArgs) {
	SetDefaults_NodeResourcesAllocatableArgs(in)
}

func SetObjectDefaults_PreemptionTolerationArgs(in *PreemptionTolerationArgs) {
	SetDefaults_PreemptionTolerationArgs(in)
}

func SetObjectDefaults_SySchedArgs(in *SySchedArgs) {
	SetDefaults_SySchedArgs(in)
}

func SetObjectDefaults_TargetLoadPackingArgs(in *TargetLoadPackingArgs) {
	SetDefaults_TargetLoadPackingArgs(in)
}

func SetObjectDefaults_TopologicalSortArgs(in *TopologicalSortArgs) {
	SetDefaults_TopologicalSortArgs(in)
}
