package maintenance

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ImpactType enumerates the values for impact type.
type ImpactType string

const (
	// Freeze ...
	Freeze ImpactType = "Freeze"
	// None ...
	None ImpactType = "None"
	// Redeploy ...
	Redeploy ImpactType = "Redeploy"
	// Restart ...
	Restart ImpactType = "Restart"
)

// PossibleImpactTypeValues returns an array of possible values for the ImpactType const type.
func PossibleImpactTypeValues() []ImpactType {
	return []ImpactType{Freeze, None, Redeploy, Restart}
}

// Scope enumerates the values for scope.
type Scope string

const (
	// ScopeAll ...
	ScopeAll Scope = "All"
	// ScopeHost ...
	ScopeHost Scope = "Host"
	// ScopeInResource ...
	ScopeInResource Scope = "InResource"
	// ScopeResource ...
	ScopeResource Scope = "Resource"
)

// PossibleScopeValues returns an array of possible values for the Scope const type.
func PossibleScopeValues() []Scope {
	return []Scope{ScopeAll, ScopeHost, ScopeInResource, ScopeResource}
}

// UpdateStatus enumerates the values for update status.
type UpdateStatus string

const (
	// Completed ...
	Completed UpdateStatus = "Completed"
	// InProgress ...
	InProgress UpdateStatus = "InProgress"
	// Pending ...
	Pending UpdateStatus = "Pending"
	// RetryLater ...
	RetryLater UpdateStatus = "RetryLater"
	// RetryNow ...
	RetryNow UpdateStatus = "RetryNow"
)

// PossibleUpdateStatusValues returns an array of possible values for the UpdateStatus const type.
func PossibleUpdateStatusValues() []UpdateStatus {
	return []UpdateStatus{Completed, InProgress, Pending, RetryLater, RetryNow}
}
