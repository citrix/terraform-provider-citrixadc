package vpnalwaysonprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnalwaysonprofileResourceModel describes the resource data model.
type VpnalwaysonprofileResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Clientcontrol             types.String `tfsdk:"clientcontrol"`
	Locationbasedvpn          types.String `tfsdk:"locationbasedvpn"`
	Name                      types.String `tfsdk:"name"`
	Networkaccessonvpnfailure types.String `tfsdk:"networkaccessonvpnfailure"`
}

func (r *VpnalwaysonprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnalwaysonprofile resource.",
			},
			"clientcontrol": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DENY"),
				Description: "Allow/Deny user to log off and connect to another Gateway",
			},
			"locationbasedvpn": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Remote"),
				Description: "Option to decide if tunnel should be established when in enterprise network. When locationBasedVPN is remote, client tries to detect if it is located in enterprise network or not and establishes the tunnel if not in enterprise network. Dns suffixes configured using -add dns suffix- are used to decide if the client is in the enterprise network or not. If the resolution of the DNS suffix results in private IP, client is said to be in enterprise network. When set to EveryWhere, the client skips the check to detect if it is on the enterprise network and tries to establish the tunnel",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "name of AlwaysON profile",
			},
			"networkaccessonvpnfailure": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("fullAccess"),
				Description: "Option to block network traffic when tunnel is not established(and the config requires that tunnel be established). When set to onlyToGateway, the network traffic to and from the client (except Gateway IP) is blocked. When set to fullAccess, the network traffic is not blocked",
			},
		},
	}
}

func vpnalwaysonprofileGetThePayloadFromtheConfig(ctx context.Context, data *VpnalwaysonprofileResourceModel) vpn.Vpnalwaysonprofile {
	tflog.Debug(ctx, "In vpnalwaysonprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnalwaysonprofile := vpn.Vpnalwaysonprofile{}
	if !data.Clientcontrol.IsNull() {
		vpnalwaysonprofile.Clientcontrol = data.Clientcontrol.ValueString()
	}
	if !data.Locationbasedvpn.IsNull() {
		vpnalwaysonprofile.Locationbasedvpn = data.Locationbasedvpn.ValueString()
	}
	if !data.Name.IsNull() {
		vpnalwaysonprofile.Name = data.Name.ValueString()
	}
	if !data.Networkaccessonvpnfailure.IsNull() {
		vpnalwaysonprofile.Networkaccessonvpnfailure = data.Networkaccessonvpnfailure.ValueString()
	}

	return vpnalwaysonprofile
}

func vpnalwaysonprofileSetAttrFromGet(ctx context.Context, data *VpnalwaysonprofileResourceModel, getResponseData map[string]interface{}) *VpnalwaysonprofileResourceModel {
	tflog.Debug(ctx, "In vpnalwaysonprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientcontrol"]; ok && val != nil {
		data.Clientcontrol = types.StringValue(val.(string))
	} else {
		data.Clientcontrol = types.StringNull()
	}
	if val, ok := getResponseData["locationbasedvpn"]; ok && val != nil {
		data.Locationbasedvpn = types.StringValue(val.(string))
	} else {
		data.Locationbasedvpn = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["networkaccessonvpnfailure"]; ok && val != nil {
		data.Networkaccessonvpnfailure = types.StringValue(val.(string))
	} else {
		data.Networkaccessonvpnfailure = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
