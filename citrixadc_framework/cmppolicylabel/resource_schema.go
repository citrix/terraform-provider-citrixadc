package cmppolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CmppolicylabelResourceModel describes the resource data model.
type CmppolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
	Type      types.String `tfsdk:"type"`
}

func (r *CmppolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cmppolicylabel resource.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the HTTP compression policy label. Must begin with a letter, number, or the underscore character (_). Additional characters allowed, after the first character, are the hyphen (-), period (.) pound sign (#), space ( ), at sign (@), equals (=), and colon (:). The name must be unique within the list of policy labels for compression policies. Can be renamed after the policy label is created.\n\n            The following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp policylabel\" or 'my cmp policylabel').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the compression policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\n                        The following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp policylabel\" or 'my cmp policylabel').",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of packets (request packets or response) against which to match the policies bound to this policy label.",
			},
		},
	}
}

func cmppolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *CmppolicylabelResourceModel) cmp.Cmppolicylabel {
	tflog.Debug(ctx, "In cmppolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cmppolicylabel := cmp.Cmppolicylabel{}
	if !data.Labelname.IsNull() {
		cmppolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		cmppolicylabel.Newname = data.Newname.ValueString()
	}
	if !data.Type.IsNull() {
		cmppolicylabel.Type = data.Type.ValueString()
	}

	return cmppolicylabel
}

func cmppolicylabelSetAttrFromGet(ctx context.Context, data *CmppolicylabelResourceModel, getResponseData map[string]interface{}) *CmppolicylabelResourceModel {
	tflog.Debug(ctx, "In cmppolicylabelSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
