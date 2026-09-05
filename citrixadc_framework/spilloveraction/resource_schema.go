package spilloveraction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/spillover"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SpilloveractionResourceModel describes the resource data model.
type SpilloveractionResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Action  types.String `tfsdk:"action"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
}

func (r *SpilloveractionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the spilloveraction resource.",
			},
			// SDK v2 parity: `action` was Optional + ForceNew (no Default). NITRO
			// exposes no set/update verb for spilloveraction (only add/delete/get/
			// rename), so a change to action cannot be applied in place. Optional +
			// Computed reads the value back from the ADC; RequiresReplaceIfConfigured
			// reproduces the SDK v2 ForceNew-when-set behaviour (recreate on change).
			"action": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Spillover action. Currently only type SPILLOVER is supported",
			},
			// SDK v2 parity: `name` was Required + ForceNew -> Required + RequiresReplace.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the spillover action.",
			},
			// newname is the rename trigger (NITRO ?action=rename). It is Optional only:
			// changing it must NOT force replacement - it drives an in-place rename via
			// Update. Not Computed - it is a pure user input, never echoed back by GET.
			"newname": schema.StringAttribute{
				Optional:    true,
				Description: "New name for the spillover action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters. \nChoose a name that can be correlated with the function that the action performs. \n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
		},
	}
}

func spilloveractionGetThePayloadFromthePlan(ctx context.Context, data *SpilloveractionResourceModel) spillover.Spilloveraction {
	tflog.Debug(ctx, "In spilloveractionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	spilloveraction := spillover.Spilloveraction{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		spilloveraction.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		spilloveraction.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.

	return spilloveraction
}

// spilloveractionSetAttrFromGet is the RESOURCE state setter. It preserves the
// prior plan/state values for identity-bearing attributes so a rename (which
// makes the live object name diverge from the configured `name`) does not clobber
// the user-facing configuration.
func spilloveractionSetAttrFromGet(ctx context.Context, data *SpilloveractionResourceModel, getResponseData map[string]interface{}) *SpilloveractionResourceModel {
	tflog.Debug(ctx, "In spilloveractionSetAttrFromGet Function")

	// Convert API response to model.
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	}
	// name is the user-facing key. After a rename (via newname) the live object
	// name (tracked by data.Id) diverges from the configured name, and GET returns
	// the live (new) name. Overwriting name from GET would clobber the user's
	// configured value and trigger a spurious RequiresReplace diff. So only adopt
	// the GET value when we don't already have one (e.g. on import, where state
	// carries only the ID); otherwise preserve.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.

	return data
}

// spilloveractionSetAttrFromGetForDatasource faithfully copies every field from
// the GET response. The datasource has no prior plan/state to preserve, so it
// must populate the model directly from the API response and set the ID itself.
func spilloveractionSetAttrFromGetForDatasource(ctx context.Context, data *SpilloveractionResourceModel, getResponseData map[string]interface{}) *SpilloveractionResourceModel {
	tflog.Debug(ctx, "In spilloveractionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
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

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
