package cmppolicylabel

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
				// newname is the rename trigger (NITRO ?action=rename). Changing it
				// must NOT force replacement - it drives an in-place rename via Update.
				// Not Computed: it is a pure user input, never echoed back by GET.
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

func cmppolicylabelGetThePayloadFromthePlan(ctx context.Context, data *CmppolicylabelResourceModel) cmp.Cmppolicylabel {
	tflog.Debug(ctx, "In cmppolicylabelGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cmppolicylabel := cmp.Cmppolicylabel{}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		cmppolicylabel.Labelname = data.Labelname.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		cmppolicylabel.Type = data.Type.ValueString()
	}

	return cmppolicylabel
}

func cmppolicylabelSetAttrFromGet(ctx context.Context, data *CmppolicylabelResourceModel, getResponseData map[string]interface{}) *CmppolicylabelResourceModel {
	tflog.Debug(ctx, "In cmppolicylabelSetAttrFromGet Function")

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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	}

	return data
}

// cmppolicylabelSetAttrFromGetForDatasource faithfully copies every field from the
// GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func cmppolicylabelSetAttrFromGetForDatasource(ctx context.Context, data *CmppolicylabelResourceModel, getResponseData map[string]interface{}) *CmppolicylabelResourceModel {
	tflog.Debug(ctx, "In cmppolicylabelSetAttrFromGetForDatasource Function")

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

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	return data
}
