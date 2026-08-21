package utils

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// WoVersionUpgradeState builds the StateUpgrader map for a Framework resource
// whose schema Version was bumped solely to migrate pre-write-only state.
//
// Background (GH #1441): the "*_wo_version" tracker attributes that pair with
// write-only ("*_wo") secrets are Optional+Computed with a static Default of 1.
// When a user upgrades from a provider release that predated write-only secrets
// the stored state has no value for these attributes, so they decode to null.
// At plan time the framework's TransformDefaults stamps the Default onto the
// null, producing a spurious "null -> 1" diff that cascades other
// Optional+Computed attributes to "(known after apply)".
//
// The fix pairs a schema Version bump with this upgrader, and neither half works
// alone:
//   - The Version bump makes the stored version less than the current version,
//     which is what causes Terraform to run the upgrade path at all. A bump with
//     no registered upgrader hard-errors ("no upgrader for version N").
//   - The seed closure sets each "*_wo_version" attribute to 1 when it is null,
//     so TransformDefaults finds the value already present and plans no change.
//     An upgrader with no version bump never fires (upgraders only run when the
//     stored version is less than the current version).
//
// The same seed upgrader is registered under BOTH version 0 and version 1: a
// pre-write-only state may have been written by an older SDKv2 provider
// (version 0) or an early Framework release (version 1). Registering both
// guarantees the upgrade path finds a matching upgrader for whatever old version
// the state holds and never hard-errors. Extra keys are harmless — the framework
// only looks up the stored version, and a key equal to the current version is
// never used.
//
// currentSchema is passed as PriorSchema: the Version bump does not change the
// attribute set (it only seeds a value), so the current schema decodes the prior
// state directly. The framework unmarshals prior state with
// IgnoreUndefinedAttributes, so a missing "*_wo_version" key simply decodes to
// null, which the seed closure then fills. No separate old-schema definition is
// needed.
func WoVersionUpgradeState(currentSchema schema.Schema, seed func(context.Context, resource.UpgradeStateRequest, *resource.UpgradeStateResponse)) map[int64]resource.StateUpgrader {
	upgrader := resource.StateUpgrader{
		PriorSchema:   &currentSchema,
		StateUpgrader: seed,
	}
	return map[int64]resource.StateUpgrader{
		0: upgrader,
		1: upgrader,
	}
}
