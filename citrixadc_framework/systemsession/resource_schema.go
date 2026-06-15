package systemsession

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SystemsessionResourceModel describes the resource data model.
//
// systemsession is an ACTION-ONLY resource. NITRO exposes get/get(all)/count and a
// single write verb `kill` (?action=kill, POST). There is NO add/update/delete API.
// The only inputs accepted by the kill action are `sid` and `all`; every other field
// listed in the NITRO struct (username, logintime, clientipaddress, ...) is a
// GET-response-only attribute and is intentionally excluded from this resource model.
type SystemsessionResourceModel struct {
	Id  types.String `tfsdk:"id"`
	All types.Bool   `tfsdk:"all"`
	Sid types.Int64  `tfsdk:"sid"`
}

func (r *SystemsessionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		// DESTRUCTIVE ONE-SHOT ACTION RESOURCE.
		// `kill systemsession -all` terminates ALL admin sessions, INCLUDING the
		// provider's own NITRO session. Killing a specific `sid` terminates that one
		// session. This resource performs the kill action on Create and has no GET-based
		// drift detection (sessions are transient); Read/Update/Delete are no-ops.
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "The ID of the systemsession resource (the sid value, or \"all\" when all=true).",
			},
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					// Action input — re-running with a different value is a new kill action.
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Terminate all the system sessions except the current session.",
			},
			"sid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the system session to kill.",
			},
		},
	}
}

// systemsessionGetThePayloadFromthePlan builds the kill action payload. Only `sid`
// and `all` are valid write parameters for ?action=kill (Pattern 15: all read-only
// GET fields are excluded).
func systemsessionGetThePayloadFromthePlan(ctx context.Context, data *SystemsessionResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In systemsessionGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		payload["all"] = data.All.ValueBool()
	}
	if !data.Sid.IsNull() && !data.Sid.IsUnknown() {
		payload["sid"] = utils.IntPtr(int(data.Sid.ValueInt64()))
	}

	return payload
}
