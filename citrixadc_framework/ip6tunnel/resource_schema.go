package ip6tunnel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ip6tunnelResourceModel describes the resource data model.
type Ip6tunnelResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Local      types.String `tfsdk:"local"`
	Name       types.String `tfsdk:"name"`
	Ownergroup types.String `tfsdk:"ownergroup"`
	Remote     types.String `tfsdk:"remote"`
}

func (r *Ip6tunnelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ip6tunnel resource.",
			},
			"local": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "An IPv6 address of the local Citrix ADC used to set up the tunnel.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the IPv6 Tunnel. Cannot be changed after the service group is created. Must begin with a number or letter, and can consist of letters, numbers, and the @ _ - . (period) : (colon) # and space ( ) characters.",
			},
			"ownergroup": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("DEFAULT_NG"),
				Description: "The owner node group in a Cluster for the tunnel.",
			},
			"remote": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "An IPv6 address of the remote Citrix ADC used to set up the tunnel.",
			},
		},
	}
}

func ip6tunnelGetThePayloadFromtheConfig(ctx context.Context, data *Ip6tunnelResourceModel) network.Ip6tunnel {
	tflog.Debug(ctx, "In ip6tunnelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ip6tunnel := network.Ip6tunnel{}
	if !data.Local.IsNull() {
		ip6tunnel.Local = data.Local.ValueString()
	}
	if !data.Name.IsNull() {
		ip6tunnel.Name = data.Name.ValueString()
	}
	if !data.Ownergroup.IsNull() {
		ip6tunnel.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Remote.IsNull() {
		ip6tunnel.Remote = data.Remote.ValueString()
	}

	return ip6tunnel
}

func ip6tunnelSetAttrFromGet(ctx context.Context, data *Ip6tunnelResourceModel, getResponseData map[string]interface{}) *Ip6tunnelResourceModel {
	tflog.Debug(ctx, "In ip6tunnelSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["local"]; ok && val != nil {
		data.Local = types.StringValue(val.(string))
	} else {
		data.Local = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["remoteip"]; ok && val != nil {
		data.Remote = types.StringValue(val.(string))
	} else if val, ok := getResponseData["remote"]; ok && val != nil {
		data.Remote = types.StringValue(val.(string))
	} else {
		data.Remote = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Name.ValueString(), data.Remote.ValueString()))

	return data
}
