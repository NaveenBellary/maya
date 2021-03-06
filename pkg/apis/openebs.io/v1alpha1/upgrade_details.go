/*
Copyright 2019 The OpenEBS Authors.

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

package v1alpha1

// VersionDetails provides the details for upgrade
type VersionDetails struct {
	// If AutoUpgrade is set to true then the resource is
	// upgraded automatically without any manual steps
	AutoUpgrade bool `json:"autoUpgrade"`
	// Current is the version of resource
	Current string `json:"current"`
	// Desired is the version that we want to
	// upgrade or the control plane version
	Desired string `json:"desired"`
	// DependentsUpgraded gives the details whether all children
	// of a resource are upgraded to desired version or not
	DependentsUpgraded bool `json:"dependentsUpgraded"`
}
