package lsnclient

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LsnclientResourceModel describes the resource data model.
type LsnclientResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Clientname types.String `tfsdk:"clientname"`
}

func (r *LsnclientResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnclient resource.",
			},
			"clientname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn client1\" or 'lsn client1').",
			},
		},
	}
}

func lsnclientGetThePayloadFromtheConfig(ctx context.Context, data *LsnclientResourceModel) lsn.Lsnclient {
	tflog.Debug(ctx, "In lsnclientGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnclient := lsn.Lsnclient{}
	if !data.Clientname.IsNull() && !data.Clientname.IsUnknown() {
		lsnclient.Clientname = data.Clientname.ValueString()
	}

	return lsnclient
}

func lsnclientSetAttrFromGet(ctx context.Context, data *LsnclientResourceModel, getResponseData map[string]interface{}) *LsnclientResourceModel {
	tflog.Debug(ctx, "In lsnclientSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientname"]; ok && val != nil {
		data.Clientname = types.StringValue(val.(string))
	} else if data.Clientname.IsUnknown() {
		// Only null when unknown; never clobber a known configured value.
		data.Clientname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Clientname.ValueString()))

	return data
}
