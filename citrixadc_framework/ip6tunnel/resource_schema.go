package ip6tunnel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
				// SDK v2: Optional + Computed + ForceNew. NITRO defaults this to
				// "DEFAULT_NG" server-side, so it is Computed rather than carrying a
				// framework Default (a Default without Computed panics at schema build).
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436: avoid spurious destroy+recreate on upgrade for
					// Optional+Computed; ip6tunnel has no update op (create-only).
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
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
	if !data.Local.IsNull() && !data.Local.IsUnknown() {
		ip6tunnel.Local = data.Local.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		ip6tunnel.Name = data.Name.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		ip6tunnel.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Remote.IsNull() && !data.Remote.IsUnknown() {
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
	} else if data.Ownergroup.IsUnknown() {
		// Only resolve an unknown (unconfigured) value to null. Preserve any
		// configured/prior value when the ADC omits ownergroup from the GET
		// response, to avoid "inconsistent result after apply".
		data.Ownergroup = types.StringNull()
	}
	// NITRO echoes the configured "remote" write property back as the read-only
	// "remoteip" property, so map state's "remote" from "remoteip" first
	// (matching the SDK v2 resource), falling back to "remote".
	if val, ok := getResponseData["remoteip"]; ok && val != nil {
		data.Remote = types.StringValue(val.(string))
	} else if val, ok := getResponseData["remote"]; ok && val != nil {
		data.Remote = types.StringValue(val.(string))
	} else {
		data.Remote = types.StringNull()
	}

	// Set ID for the resource.
	// Named resource keyed on "name" - the ID is the plain name value
	// (matches SDK v2 d.SetId(name)).
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
