package dnsaction64

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Dnsaction64ResourceModel describes the resource data model.
type Dnsaction64ResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Actionname  types.String `tfsdk:"actionname"`
	Excluderule types.String `tfsdk:"excluderule"`
	Mappedrule  types.String `tfsdk:"mappedrule"`
	Prefix      types.String `tfsdk:"prefix"`
}

func (r *Dnsaction64Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsaction64 resource.",
			},
			"actionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dns64 action.",
			},
			"excluderule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The expression to select the criteria for eliminating the corresponding ipv6 addresses from the response.",
			},
			"mappedrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The expression to select the criteria for ipv4 addresses to be used for synthesis.\n                      Only if the mappedrule is evaluated to true the corresponding ipv4 address is used for synthesis using respective prefix,\n                      otherwise the A RR is discarded",
			},
			"prefix": schema.StringAttribute{
				Required:    true,
				Description: "The dns64 prefix to be used if the after evaluating the rules",
			},
		},
	}
}

func dnsaction64GetThePayloadFromtheConfig(ctx context.Context, data *Dnsaction64ResourceModel) dns.Dnsaction64 {
	tflog.Debug(ctx, "In dnsaction64GetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnsaction64 := dns.Dnsaction64{}
	if !data.Actionname.IsNull() {
		dnsaction64.Actionname = data.Actionname.ValueString()
	}
	if !data.Excluderule.IsNull() {
		dnsaction64.Excluderule = data.Excluderule.ValueString()
	}
	if !data.Mappedrule.IsNull() {
		dnsaction64.Mappedrule = data.Mappedrule.ValueString()
	}
	if !data.Prefix.IsNull() {
		dnsaction64.Prefix = data.Prefix.ValueString()
	}

	return dnsaction64
}

func dnsaction64SetAttrFromGet(ctx context.Context, data *Dnsaction64ResourceModel, getResponseData map[string]interface{}) *Dnsaction64ResourceModel {
	tflog.Debug(ctx, "In dnsaction64SetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["actionname"]; ok && val != nil {
		data.Actionname = types.StringValue(val.(string))
	} else {
		data.Actionname = types.StringNull()
	}
	if val, ok := getResponseData["excluderule"]; ok && val != nil {
		data.Excluderule = types.StringValue(val.(string))
	} else {
		data.Excluderule = types.StringNull()
	}
	if val, ok := getResponseData["mappedrule"]; ok && val != nil {
		data.Mappedrule = types.StringValue(val.(string))
	} else {
		data.Mappedrule = types.StringNull()
	}
	if val, ok := getResponseData["prefix"]; ok && val != nil {
		data.Prefix = types.StringValue(val.(string))
	} else {
		data.Prefix = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Actionname.ValueString())

	return data
}
