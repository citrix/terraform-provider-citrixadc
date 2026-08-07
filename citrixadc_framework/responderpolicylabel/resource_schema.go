package responderpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ResponderpolicylabelResourceModel describes the resource data model.
type ResponderpolicylabelResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Comment         types.String `tfsdk:"comment"`
	Labelname       types.String `tfsdk:"labelname"`
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`
}

func (r *ResponderpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the responderpolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// SDK v2 had comment as Optional+Computed+ForceNew and stored "" when
				// unset. Keep Optional+Computed to preserve that contract: an upgraded
				// SDK v2 state carries comment="", so making comment Optional-only would
				// force a spurious RequiresReplace (config-null vs state-"") that
				// destroys the parent and drops any child binding. NITRO omits comment
				// from GET when empty, so SetAttrFromGet resolves the absent value to ""
				// (a known value) to avoid "still indicated an unknown value" after
				// apply. UseStateForUnknown (before RequiresReplace) preserves the
				// Computed value across plans instead of churning it to unknown, exactly
				// as done for policylabeltype below; RequiresReplace preserves the SDK v2
				// ForceNew contract.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this responder policy label.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the responder policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the responder policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder policy label\" or my responder policy label').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				// newname is the rename trigger (NITRO ?action=rename). Changing it must
				// NOT force replacement - it drives an in-place rename via Update. Not
				// Computed: it is a pure user input, never echoed back by GET.
				Description: "New name for the responder policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"policylabeltype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// SDK v2 had policylabeltype as Optional+Computed+ForceNew with NO
				// Default (the auto-gen wrongly added Default("HTTP")). NITRO returns the
				// value on GET, so keep Optional+Computed. UseStateForUnknown before
				// RequiresReplace preserves the Computed value across plans instead of
				// churning it to unknown.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of responses sent by the policies bound to this policy label. Types are:\n* HTTP - HTTP responses.\n* OTHERTCP - NON-HTTP TCP responses.\n* SIP_UDP - SIP responses.\n* RADIUS - RADIUS responses.\n* MYSQL - SQL responses in MySQL format.\n* MSSQL - SQL responses in Microsoft SQL format.\n* NAT - NAT response.\n* MQTT - Trigger policies bind with MQTT type.\n* MQTT_JUMBO - Trigger policies bind with MQTT Jumbo type.",
			},
		},
	}
}

func responderpolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *ResponderpolicylabelResourceModel) responder.Responderpolicylabel {
	tflog.Debug(ctx, "In responderpolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	responderpolicylabel := responder.Responderpolicylabel{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		responderpolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		responderpolicylabel.Labelname = data.Labelname.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of the
	// add payload, so it is deliberately excluded from the create POST body.
	if !data.Policylabeltype.IsNull() && !data.Policylabeltype.IsUnknown() {
		responderpolicylabel.Policylabeltype = data.Policylabeltype.ValueString()
	}

	return responderpolicylabel
}

func responderpolicylabelSetAttrFromGet(ctx context.Context, data *ResponderpolicylabelResourceModel, getResponseData map[string]interface{}) *ResponderpolicylabelResourceModel {
	tflog.Debug(ctx, "In responderpolicylabelSetAttrFromGet Function")

	// Convert API response to model.
	// comment is Optional+Computed (see schema). NITRO omits "comment" from the GET
	// response when it is empty, so when the field is absent resolve it to "" (the
	// value SDK v2 stored and a known value) rather than leaving the Computed
	// planned value unknown, which would otherwise fail with "still indicated an
	// unknown value" after a fresh apply where comment is unset.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringValue("")
	}
	// labelname is the user-facing key. Once a rename has happened (via newname), the
	// live object name (tracked by data.Id) diverges from the configured labelname,
	// and GET returns the live (new) name. Overwriting labelname from GET would
	// clobber the user's configured value and trigger a spurious RequiresReplace
	// diff. So only adopt the GET value when we don't already have one (e.g. on
	// import, where state carries only the ID); otherwise preserve.
	if data.Labelname.IsNull() || data.Labelname.IsUnknown() || data.Labelname.ValueString() == "" {
		if val, ok := getResponseData["labelname"]; ok && val != nil {
			data.Labelname = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["policylabeltype"]; ok && val != nil {
		data.Policylabeltype = types.StringValue(val.(string))
	}

	return data
}

// responderpolicylabelSetAttrFromGetForDatasource faithfully copies every field from
// the GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func responderpolicylabelSetAttrFromGetForDatasource(ctx context.Context, data *ResponderpolicylabelResourceModel, getResponseData map[string]interface{}) *ResponderpolicylabelResourceModel {
	tflog.Debug(ctx, "In responderpolicylabelSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["policylabeltype"]; ok && val != nil {
		data.Policylabeltype = types.StringValue(val.(string))
	} else {
		data.Policylabeltype = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	return data
}
