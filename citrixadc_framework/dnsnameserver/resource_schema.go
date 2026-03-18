package dnsnameserver

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsnameserverResourceModel describes the resource data model.
type DnsnameserverResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Dnsprofilename types.String `tfsdk:"dnsprofilename"`
	Dnsvservername types.String `tfsdk:"dnsvservername"`
	Ip             types.String `tfsdk:"ip"`
	Local          types.Bool   `tfsdk:"local"`
	State          types.String `tfsdk:"state"`
	Type           types.String `tfsdk:"type"`
}

func (r *DnsnameserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsnameserver resource.",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the name server",
			},
			"dnsvservername": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of a DNS virtual server. Overrides any IP address-based name servers configured on the Citrix ADC.",
			},
			"ip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of an external name server or, if the Local parameter is set, IP address of a local DNS server (LDNS).",
			},
			"local": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Mark the IP address as one that belongs to a local recursive DNS server on the Citrix ADC. The appliance recursively resolves queries received on an IP address that is marked as being local. For recursive resolution to work, the global DNS parameter, Recursion, must also be set.\n\nIf no name server is marked as being local, the appliance functions as a stub resolver and load balances the name servers.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Administrative state of the name server.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("UDP"),
				Description: "Protocol used by the name server. UDP_TCP is not valid if the name server is a DNS virtual server configured on the appliance.",
			},
		},
	}
}

func dnsnameserverGetThePayloadFromtheConfig(ctx context.Context, data *DnsnameserverResourceModel) dns.Dnsnameserver {
	tflog.Debug(ctx, "In dnsnameserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnsnameserver := dns.Dnsnameserver{}
	if !data.Dnsprofilename.IsNull() {
		dnsnameserver.Dnsprofilename = data.Dnsprofilename.ValueString()
	}
	if !data.Dnsvservername.IsNull() {
		dnsnameserver.Dnsvservername = data.Dnsvservername.ValueString()
	}
	if !data.Ip.IsNull() {
		dnsnameserver.Ip = data.Ip.ValueString()
	}
	if !data.Local.IsNull() {
		dnsnameserver.Local = data.Local.ValueBool()
	}
	if !data.State.IsNull() {
		dnsnameserver.State = data.State.ValueString()
	}
	if !data.Type.IsNull() {
		dnsnameserver.Type = data.Type.ValueString()
	}

	return dnsnameserver
}

func dnsnameserverSetAttrFromGet(ctx context.Context, data *DnsnameserverResourceModel, getResponseData map[string]interface{}) *DnsnameserverResourceModel {
	tflog.Debug(ctx, "In dnsnameserverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["dnsprofilename"]; ok && val != nil {
		data.Dnsprofilename = types.StringValue(val.(string))
	} else {
		data.Dnsprofilename = types.StringNull()
	}
	if val, ok := getResponseData["dnsvservername"]; ok && val != nil {
		data.Dnsvservername = types.StringValue(val.(string))
	} else {
		data.Dnsvservername = types.StringNull()
	}
	if val, ok := getResponseData["ip"]; ok && val != nil {
		data.Ip = types.StringValue(val.(string))
	} else {
		data.Ip = types.StringNull()
	}
	if val, ok := getResponseData["local"]; ok && val != nil {
		data.Local = types.BoolValue(val.(bool))
	} else {
		data.Local = types.BoolNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource based on which identifiers are present
	var idParts []string
	if !data.Dnsvservername.IsNull() {
		idParts = append(idParts, fmt.Sprintf("dnsvservername:%s", data.Dnsvservername.ValueString()))
	}
	if !data.Ip.IsNull() {
		idParts = append(idParts, fmt.Sprintf("ip:%s", data.Ip.ValueString()))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
