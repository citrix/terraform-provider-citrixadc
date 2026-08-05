package cmppolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CmppolicyResourceModel describes the resource data model.
type CmppolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Newname   types.String `tfsdk:"newname"`
	Resaction types.String `tfsdk:"resaction"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *CmppolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cmppolicy resource.",
			},
			"name": schema.StringAttribute{
				// Primary key. SDK v2 marked this ForceNew, so a change must recreate
				// the resource (same contract as the legacy provider).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the HTTP compression policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nCan be changed after the policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp policy\" or 'my cmp policy').",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it must
				// NOT force replacement - it drives an in-place rename via Update. It is a
				// pure user input, never echoed back by GET, so it is Optional only (no
				// Computed, no RequiresReplace).
				Optional:    true,
				Description: "New name for the compression policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nChoose a name that reflects the function that the policy performs.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp policy\" or 'my cmp policy').",
			},
			"resaction": schema.StringAttribute{
				// SDK v2 contract: Optional+Computed and updateable (in-place via PUT).
				Optional:    true,
				Computed:    true,
				Description: "The built-in or user-defined compression action to apply to the response when the policy matches a request or response.",
			},
			"rule": schema.StringAttribute{
				// SDK v2 contract: Optional+Computed and updateable (in-place via PUT).
				Optional:    true,
				Computed:    true,
				Description: "Expression that determines which HTTP requests or responses match the compression policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func cmppolicyGetThePayloadFromthePlan(ctx context.Context, data *CmppolicyResourceModel) cmp.Cmppolicy {
	tflog.Debug(ctx, "In cmppolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cmppolicy := cmp.Cmppolicy{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		cmppolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add/update payload, so it is deliberately excluded here.
	if !data.Resaction.IsNull() && !data.Resaction.IsUnknown() {
		cmppolicy.Resaction = data.Resaction.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		cmppolicy.Rule = data.Rule.ValueString()
	}

	return cmppolicy
}

func cmppolicySetAttrFromGet(ctx context.Context, data *CmppolicyResourceModel, getResponseData map[string]interface{}) *CmppolicyResourceModel {
	tflog.Debug(ctx, "In cmppolicySetAttrFromGet Function")

	// name is the user-facing key. Once a rename has happened (via newname), the live
	// object name (tracked by data.Id) diverges from the configured name, and GET
	// returns the live (new) name. Overwriting name from GET would clobber the user's
	// configured value and trigger a spurious RequiresReplace diff. So only adopt the
	// GET value when we don't already have one (e.g. on import, where state carries
	// only the ID); otherwise preserve.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["resaction"]; ok && val != nil {
		data.Resaction = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	}

	// NOTE: do NOT set data.Id here. The ID tracks the CURRENT LIVE name (== name at
	// create, == newname after a rename); the resource CRUD funcs manage it. Resetting
	// it to data.Name would break reads/deletes after a rename.

	return data
}

// cmppolicySetAttrFromGetForDatasource faithfully copies every field from the GET
// response. The datasource has no prior plan/state to preserve, so it must populate
// the model directly from the API response and set the ID itself.
func cmppolicySetAttrFromGetForDatasource(ctx context.Context, data *CmppolicyResourceModel, getResponseData map[string]interface{}) *CmppolicyResourceModel {
	tflog.Debug(ctx, "In cmppolicySetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["resaction"]; ok && val != nil {
		data.Resaction = types.StringValue(val.(string))
	} else {
		data.Resaction = types.StringNull()
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
