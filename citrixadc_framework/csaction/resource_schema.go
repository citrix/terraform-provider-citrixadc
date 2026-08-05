package csaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsactionResourceModel describes the resource data model.
type CsactionResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Comment           types.String `tfsdk:"comment"`
	Name              types.String `tfsdk:"name"`
	Newname           types.String `tfsdk:"newname"`
	Targetlbvserver   types.String `tfsdk:"targetlbvserver"`
	Targetvserver     types.String `tfsdk:"targetvserver"`
	Targetvserverexpr types.String `tfsdk:"targetvserverexpr"`
}

func (r *CsactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csaction resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this cs action.",
			},
			"name": schema.StringAttribute{
				Required: true,
				// SDK v2 marked name as ForceNew. Preserve that contract with
				// RequiresReplace: changing the primary key recreates the resource.
				// (An in-place rename is offered separately via the newname attribute.)
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the content switching action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the content switching action is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				// newname is the rename trigger (NITRO ?action=rename). Changing it
				// must NOT force replacement - it drives an in-place rename via Update.
				// Not Computed: it is a pure user input, never echoed back by GET, so a
				// Computed flag would leave it unknown-after-apply (framework rejects that).
				Description: "New name for the content switching action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"targetlbvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing virtual server to which the content is switched.",
			},
			"targetvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the VPN, GSLB or Authentication virtual server to which the content is switched.",
			},
			"targetvserverexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Information about this content switching action.",
			},
		},
	}
}

func csactionGetThePayloadFromthePlan(ctx context.Context, data *CsactionResourceModel) cs.Csaction {
	tflog.Debug(ctx, "In csactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	csaction := cs.Csaction{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		csaction.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		csaction.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add/update payload, so it is deliberately excluded from the create body.
	if !data.Targetlbvserver.IsNull() && !data.Targetlbvserver.IsUnknown() {
		csaction.Targetlbvserver = data.Targetlbvserver.ValueString()
	}
	if !data.Targetvserver.IsNull() && !data.Targetvserver.IsUnknown() {
		csaction.Targetvserver = data.Targetvserver.ValueString()
	}
	if !data.Targetvserverexpr.IsNull() && !data.Targetvserverexpr.IsUnknown() {
		csaction.Targetvserverexpr = data.Targetvserverexpr.ValueString()
	}

	return csaction
}

func csactionSetAttrFromGet(ctx context.Context, data *CsactionResourceModel, getResponseData map[string]interface{}) *CsactionResourceModel {
	tflog.Debug(ctx, "In csactionSetAttrFromGet Function")

	// Convert API response to model.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	// name is the user-facing key. Once a rename has happened (via newname), the
	// live object name (tracked by data.Id) diverges from the configured name, and
	// GET returns the live (new) name. Overwriting name from GET would clobber the
	// user's configured value and trigger a spurious RequiresReplace diff. So only
	// adopt the GET value when we don't already have one (e.g. on import, where
	// state carries only the ID); otherwise preserve the existing value.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["targetlbvserver"]; ok && val != nil {
		data.Targetlbvserver = types.StringValue(val.(string))
	} else {
		data.Targetlbvserver = types.StringNull()
	}
	if val, ok := getResponseData["targetvserver"]; ok && val != nil {
		data.Targetvserver = types.StringValue(val.(string))
	} else {
		data.Targetvserver = types.StringNull()
	}
	if val, ok := getResponseData["targetvserverexpr"]; ok && val != nil {
		data.Targetvserverexpr = types.StringValue(val.(string))
	} else {
		data.Targetvserverexpr = types.StringNull()
	}

	return data
}

// csactionSetAttrFromGetForDatasource faithfully copies every field from the GET
// response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func csactionSetAttrFromGetForDatasource(ctx context.Context, data *CsactionResourceModel, getResponseData map[string]interface{}) *CsactionResourceModel {
	tflog.Debug(ctx, "In csactionSetAttrFromGetForDatasource Function")

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
	// newname is rename-only and never returned by GET.
	data.Newname = types.StringNull()
	if val, ok := getResponseData["targetlbvserver"]; ok && val != nil {
		data.Targetlbvserver = types.StringValue(val.(string))
	} else {
		data.Targetlbvserver = types.StringNull()
	}
	if val, ok := getResponseData["targetvserver"]; ok && val != nil {
		data.Targetvserver = types.StringValue(val.(string))
	} else {
		data.Targetvserver = types.StringNull()
	}
	if val, ok := getResponseData["targetvserverexpr"]; ok && val != nil {
		data.Targetvserverexpr = types.StringValue(val.(string))
	} else {
		data.Targetvserverexpr = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
