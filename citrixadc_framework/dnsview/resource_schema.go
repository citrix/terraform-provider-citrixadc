package dnsview

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsviewResourceModel describes the resource data model.
type DnsviewResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Viewname types.String `tfsdk:"viewname"`
}

func (r *DnsviewResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsview resource.",
			},
			"viewname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the DNS view.",
			},
		},
	}
}

func dnsviewGetThePayloadFromtheConfig(ctx context.Context, data *DnsviewResourceModel) dns.Dnsview {
	tflog.Debug(ctx, "In dnsviewGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnsview := dns.Dnsview{}
	if !data.Viewname.IsNull() {
		dnsview.Viewname = data.Viewname.ValueString()
	}

	return dnsview
}

func dnsviewSetAttrFromGet(ctx context.Context, data *DnsviewResourceModel, getResponseData map[string]interface{}) *DnsviewResourceModel {
	tflog.Debug(ctx, "In dnsviewSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["viewname"]; ok && val != nil {
		data.Viewname = types.StringValue(val.(string))
	} else {
		data.Viewname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Viewname.ValueString())

	return data
}
