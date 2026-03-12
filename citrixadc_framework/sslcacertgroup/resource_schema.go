package sslcacertgroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcacertgroupResourceModel describes the resource data model.
type SslcacertgroupResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Cacertgroupname types.String `tfsdk:"cacertgroupname"`
}

func (r *SslcacertgroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcacertgroup resource.",
			},
			"cacertgroupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name given to the CA certificate group. The name will be used to add the CA certificates to the group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
		},
	}
}

func sslcacertgroupGetThePayloadFromtheConfig(ctx context.Context, data *SslcacertgroupResourceModel) ssl.Sslcacertgroup {
	tflog.Debug(ctx, "In sslcacertgroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslcacertgroup := ssl.Sslcacertgroup{}
	if !data.Cacertgroupname.IsNull() {
		sslcacertgroup.Cacertgroupname = data.Cacertgroupname.ValueString()
	}

	return sslcacertgroup
}

func sslcacertgroupSetAttrFromGet(ctx context.Context, data *SslcacertgroupResourceModel, getResponseData map[string]interface{}) *SslcacertgroupResourceModel {
	tflog.Debug(ctx, "In sslcacertgroupSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacertgroupname"]; ok && val != nil {
		data.Cacertgroupname = types.StringValue(val.(string))
	} else {
		data.Cacertgroupname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Cacertgroupname.ValueString())

	return data
}
