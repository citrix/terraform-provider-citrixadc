package spilloverpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/spillover"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SpilloverpolicyResourceModel describes the resource data model.
type SpilloverpolicyResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Action  types.String `tfsdk:"action"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Rule    types.String `tfsdk:"rule"`
}

func (r *SpilloverpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the spilloverpolicy resource.",
			},
			// SDK v2: Required (updateable, not ForceNew).
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action for the spillover policy. Action is created using add spillover action command",
			},
			// SDK v2: Optional + Computed (updateable, no default). NITRO echoes
			// comment back on GET, so Optional+Computed is safe here.
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// NITRO unsets comment back to empty. An Optional+Computed attr with no
				// Default is sticky on config-removal (no plan diff -> Update never runs
				// -> unset never fires), so pin the server default ("" = no comment) to
				// make removal produce a diff that drives the unset.
				Default:     stringdefault.StaticString(""),
				Description: "Any comments that you might want to associate with the spillover policy.",
			},
			// SDK v2: Required + ForceNew -> RequiresReplace. Changing the name
			// itself recreates the resource (a rename is done via newname instead).
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the spillover policy.",
			},
			// newname is the rename trigger (NITRO ?action=rename). Changing it must
			// NOT force replacement - it drives an in-place rename via Update. Not
			// Computed: it is a pure user input, never echoed back by GET. Not present
			// in SDK v2 (which had no rename); kept as an additive, backward-compatible
			// capability that the NITRO API supports.
			"newname": schema.StringAttribute{
				Optional:    true,
				Description: "New name for the spillover policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nChoose a name that reflects the function that the policy performs. \n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			// SDK v2: Required (updateable, not ForceNew).
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression to be used by the spillover policy.",
			},
		},
	}
}

func spilloverpolicyGetThePayloadFromthePlan(ctx context.Context, data *SpilloverpolicyResourceModel) spillover.Spilloverpolicy {
	tflog.Debug(ctx, "In spilloverpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	spilloverpolicy := spillover.Spilloverpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		spilloverpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		spilloverpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		spilloverpolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add/update payload, so it is deliberately excluded here.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		spilloverpolicy.Rule = data.Rule.ValueString()
	}

	return spilloverpolicy
}

// spilloverpolicySetAttrFromGet is the RESOURCE-side state setter. It preserves
// prior plan/state values where NITRO omits or diverges (Pattern 7) and never
// touches the ID (managed by Create/Read/Update via data.Id).
func spilloverpolicySetAttrFromGet(ctx context.Context, data *SpilloverpolicyResourceModel, getResponseData map[string]interface{}) *SpilloverpolicyResourceModel {
	tflog.Debug(ctx, "In spilloverpolicySetAttrFromGet Function")

	// Convert API response to model.
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else if data.Action.IsUnknown() {
		data.Action = types.StringNull()
	}
	// comment: only overwrite when NITRO actually returns it; otherwise preserve a
	// known configured value and only null it when still unknown (Computed omit-case).
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	// name is the user-facing key. Once a rename has happened (via newname), the live
	// object name (tracked by data.Id) diverges from the configured name, and GET
	// returns the live (new) name. Overwriting name from GET would clobber the user's
	// configured value and trigger a spurious RequiresReplace diff. Only adopt the GET
	// value when we don't already have one (e.g. on import, where state carries only
	// the ID); otherwise preserve.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else if data.Rule.IsUnknown() {
		data.Rule = types.StringNull()
	}

	return data
}

// spilloverpolicySetAttrFromGetForDatasource faithfully copies every field from
// the GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func spilloverpolicySetAttrFromGetForDatasource(ctx context.Context, data *SpilloverpolicyResourceModel, getResponseData map[string]interface{}) *SpilloverpolicyResourceModel {
	tflog.Debug(ctx, "In spilloverpolicySetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
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
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
