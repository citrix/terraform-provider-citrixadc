package dnszone

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnszoneResourceModel describes the resource data model.
type DnszoneResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Dnssecoffload types.String `tfsdk:"dnssecoffload"`
	Keyname       types.List   `tfsdk:"keyname"`
	Nsec          types.String `tfsdk:"nsec"`
	Proxymode     types.String `tfsdk:"proxymode"`
	Type          types.String `tfsdk:"type"`
	Zonename      types.String `tfsdk:"zonename"`
}

func (r *DnszoneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnszone resource.",
			},
			"dnssecoffload": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable dnssec offload for this zone.",
			},
			"keyname": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name of the public/private DNS key pair with which to sign the zone. You can sign a zone with up to four keys.",
			},
			"nsec": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable nsec generation for dnssec offload.",
			},
			"proxymode": schema.StringAttribute{
				Required:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Deploy the zone in proxy mode. Enable in the following scenarios:\n* The load balanced DNS servers are authoritative for the zone and all resource records that are part of the zone.\n* The load balanced DNS servers are authoritative for the zone, but the Citrix ADC owns a subset of the resource records that belong to the zone (partial zone ownership configuration). Typically seen in global server load balancing (GSLB) configurations, in which the appliance responds authoritatively to queries for GSLB domain names but forwards queries for other domain names in the zone to the load balanced servers.\nIn either scenario, do not create the zone's Start of Authority (SOA) and name server (NS) resource records on the appliance.\nDisable if the appliance is authoritative for the zone, but make sure that you have created the SOA and NS records on the appliance before you create the zone.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of zone to display. Mutually exclusive with the DNS Zone (zoneName) parameter. Available settings function as follows:\n* ADNS - Display all the zones for which the Citrix ADC is authoritative.\n* PROXY - Display all the zones for which the Citrix ADC is functioning as a proxy server.\n* ALL - Display all the zones configured on the appliance.",
			},
			"zonename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the zone to create.",
			},
		},
	}
}

func dnszoneGetThePayloadFromtheConfig(ctx context.Context, data *DnszoneResourceModel) dns.Dnszone {
	tflog.Debug(ctx, "In dnszoneGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnszone := dns.Dnszone{}
	if !data.Dnssecoffload.IsNull() {
		dnszone.Dnssecoffload = data.Dnssecoffload.ValueString()
	}
	if !data.Nsec.IsNull() {
		dnszone.Nsec = data.Nsec.ValueString()
	}
	if !data.Proxymode.IsNull() {
		dnszone.Proxymode = data.Proxymode.ValueString()
	}
	if !data.Type.IsNull() {
		dnszone.Type = data.Type.ValueString()
	}
	if !data.Zonename.IsNull() {
		dnszone.Zonename = data.Zonename.ValueString()
	}

	return dnszone
}

func dnszoneSetAttrFromGet(ctx context.Context, data *DnszoneResourceModel, getResponseData map[string]interface{}) *DnszoneResourceModel {
	tflog.Debug(ctx, "In dnszoneSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["dnssecoffload"]; ok && val != nil {
		data.Dnssecoffload = types.StringValue(val.(string))
	} else {
		data.Dnssecoffload = types.StringNull()
	}
	if val, ok := getResponseData["nsec"]; ok && val != nil {
		data.Nsec = types.StringValue(val.(string))
	} else {
		data.Nsec = types.StringNull()
	}
	if val, ok := getResponseData["proxymode"]; ok && val != nil {
		data.Proxymode = types.StringValue(val.(string))
	} else {
		data.Proxymode = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["zonename"]; ok && val != nil {
		data.Zonename = types.StringValue(val.(string))
	} else {
		data.Zonename = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Type.ValueString(), data.Zonename.ValueString()))

	return data
}
