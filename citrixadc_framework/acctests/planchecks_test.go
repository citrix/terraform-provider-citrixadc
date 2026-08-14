/*
Copyright 2026 Citrix Systems, Inc

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
package citrixadc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// expectNoReplaceCheck is a plan check that fails if ANY resource in the plan is
// scheduled for destroy+recreate (a "replace"). In-place updates, creates, deletes
// and no-ops are all allowed.
//
// It is used as a PreApply ConfigPlanCheck on the step-2 (current-provider) apply of
// the _sdkv2StateUpgrade tests. Those tests otherwise only assert the resource still
// exists and the trailing plan is empty, so a *clean* spurious destroy+recreate on
// the SDK v2 -> Framework state upgrade (the GH#1436 "bare RequiresReplace on an
// Optional+Computed attribute" class) would pass silently. This check turns such a
// spurious replace into a test failure.
type expectNoReplaceCheck struct{}

var _ plancheck.PlanCheck = expectNoReplaceCheck{}

func (expectNoReplaceCheck) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, rc := range req.Plan.ResourceChanges {
		if rc == nil || rc.Change == nil {
			continue
		}
		if rc.Change.Actions.Replace() {
			resp.Error = fmt.Errorf(
				"spurious destroy+recreate on SDK v2 -> Framework upgrade: %s is planned for replacement (actions=%v)",
				rc.Address, rc.Change.Actions,
			)
			return
		}
	}
}

// expectNoReplace returns a PlanCheck asserting no resource is planned for replacement.
func expectNoReplace() plancheck.PlanCheck { return expectNoReplaceCheck{} }
