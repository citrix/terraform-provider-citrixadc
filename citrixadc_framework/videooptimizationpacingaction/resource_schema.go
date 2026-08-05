package videooptimizationpacingaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VideooptimizationpacingactionResourceModel describes the resource data model.
type VideooptimizationpacingactionResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Rate    types.Int64  `tfsdk:"rate"`
}

func (r *VideooptimizationpacingactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the videooptimizationpacingaction resource.",
			},
			"comment": schema.StringAttribute{
				// SDK v2 parity: Optional + Computed (no server-side default).
				Optional:    true,
				Computed:    true,
				Description: "Comment. Any type of information about this video optimization detection action.",
			},
			"name": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew -> Required + RequiresReplace.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the video optimization pacing action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it
				// must NOT force replacement - it drives an in-place rename via Update.
				// Not Computed: it is a pure user input, never echoed back by GET.
				Optional:    true,
				Description: "New name for the videooptimization pacing action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"rate": schema.Int64Attribute{
				// SDK v2 parity: Required, not Computed, NO Default (a Default is invalid
				// without Computed, and SDK v2 declared none - user must supply rate).
				Required:    true,
				Description: "ABR Video Optimization Pacing Rate (in Kbps)",
			},
		},
	}
}

// videooptimizationpacingactionGetThePayloadFromthePlan builds the NITRO body for
// create/update. newname is deliberately excluded (rename-only argument).
func videooptimizationpacingactionGetThePayloadFromthePlan(ctx context.Context, data *VideooptimizationpacingactionResourceModel) videooptimization.Videooptimizationpacingaction {
	tflog.Debug(ctx, "In videooptimizationpacingactionGetThePayloadFromthePlan Function")

	videooptimizationpacingaction := videooptimization.Videooptimizationpacingaction{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		videooptimizationpacingaction.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		videooptimizationpacingaction.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add/update payload, so it is deliberately excluded from the body.
	if !data.Rate.IsNull() && !data.Rate.IsUnknown() {
		videooptimizationpacingaction.Rate = utils.IntPtr(int(data.Rate.ValueInt64()))
	}

	return videooptimizationpacingaction
}

// videooptimizationpacingactionSetAttrFromGet updates a resource model from a GET
// response while preserving user-facing/plan values (Pattern 7). It does NOT set
// data.Id - the ID (== the current live name) is owned by Create/Update/Read.
func videooptimizationpacingactionSetAttrFromGet(ctx context.Context, data *VideooptimizationpacingactionResourceModel, getResponseData map[string]interface{}) *VideooptimizationpacingactionResourceModel {
	tflog.Debug(ctx, "In videooptimizationpacingactionSetAttrFromGet Function")

	// comment: Optional+Computed. NITRO echoes a configured comment; when empty it is
	// omitted, in which case null resolves the Computed value (no lingering unknown).
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	// name is the user-facing key. After a rename (via newname) the live object name
	// (tracked by data.Id) diverges from the configured name, and GET returns the live
	// (new) name. Overwriting name from GET would clobber the user's configured value
	// and trigger a spurious RequiresReplace diff. Only adopt the GET value when we do
	// not already have one (e.g. on import, where state carries only the ID).
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	// rate is Required and always returned; only overwrite when present so a value
	// NITRO might omit is never clobbered (omit-on-default guard).
	if val, ok := getResponseData["rate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rate = types.Int64Value(intVal)
		}
	}

	return data
}

// videooptimizationpacingactionSetAttrFromGetForDatasource faithfully copies every
// field from the GET response. The datasource has no prior plan/state to preserve,
// so it populates the model directly from the API response and sets the ID itself.
func videooptimizationpacingactionSetAttrFromGetForDatasource(ctx context.Context, data *VideooptimizationpacingactionResourceModel, getResponseData map[string]interface{}) *VideooptimizationpacingactionResourceModel {
	tflog.Debug(ctx, "In videooptimizationpacingactionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["rate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rate = types.Int64Value(intVal)
		} else {
			data.Rate = types.Int64Null()
		}
	} else {
		data.Rate = types.Int64Null()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
