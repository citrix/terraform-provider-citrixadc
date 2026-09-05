package transformpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/transform"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TransformpolicylabelResourceModel describes the resource data model.
type TransformpolicylabelResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Labelname       types.String `tfsdk:"labelname"`
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`
}

func (r *TransformpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the transformpolicylabel resource.",
			},
			"labelname": schema.StringAttribute{
				// SDK v2: Required + ForceNew. Primary key / ID source.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the URL Transformation policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policylabel or my transform policylabel).",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it
				// must NOT force replacement - it drives an in-place rename via Update.
				// Not Computed: it is a pure user input, never echoed back by GET.
				Optional:    true,
				Description: "New name for the policy label.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policylabel or my transform policylabel).",
			},
			"policylabeltype": schema.StringAttribute{
				// SDK v2: Required + ForceNew (tfdata is_required=true, is_updateable=false).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Types of transformations allowed by the policies bound to the label. For URL transformation, always http_req (HTTP Request).",
			},
		},
	}
}

func transformpolicylabelGetThePayloadFromthePlan(ctx context.Context, data *TransformpolicylabelResourceModel) transform.Transformpolicylabel {
	tflog.Debug(ctx, "In transformpolicylabelGetThePayloadFromthePlan Function")

	// Create API request body from the model
	transformpolicylabel := transform.Transformpolicylabel{}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		transformpolicylabel.Labelname = data.Labelname.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.
	if !data.Policylabeltype.IsNull() && !data.Policylabeltype.IsUnknown() {
		transformpolicylabel.Policylabeltype = data.Policylabeltype.ValueString()
	}

	return transformpolicylabel
}

func transformpolicylabelSetAttrFromGet(ctx context.Context, data *TransformpolicylabelResourceModel, getResponseData map[string]interface{}) *TransformpolicylabelResourceModel {
	tflog.Debug(ctx, "In transformpolicylabelSetAttrFromGet Function")

	// labelname is the user-facing key. Once a rename has happened (via newname),
	// the live object name (tracked by data.Id) diverges from the configured
	// labelname, and GET returns the live (new) name. Overwriting labelname from
	// GET would clobber the user's configured value and trigger a spurious
	// RequiresReplace diff. So only adopt the GET value when we don't already have
	// one (e.g. on import, where state carries only the ID); otherwise preserve.
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

// transformpolicylabelSetAttrFromGetForDatasource faithfully copies every field
// from the GET response. The datasource has no prior plan/state to preserve, so it
// must populate the model directly from the API response and set the ID itself.
func transformpolicylabelSetAttrFromGetForDatasource(ctx context.Context, data *TransformpolicylabelResourceModel, getResponseData map[string]interface{}) *TransformpolicylabelResourceModel {
	tflog.Debug(ctx, "In transformpolicylabelSetAttrFromGetForDatasource Function")

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
