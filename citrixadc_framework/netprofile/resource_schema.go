package netprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NetprofileResourceModel describes the resource data model.
type NetprofileResourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Mbf                            types.String `tfsdk:"mbf"`
	Name                           types.String `tfsdk:"name"`
	Overridelsn                    types.String `tfsdk:"overridelsn"`
	Proxyprotocol                  types.String `tfsdk:"proxyprotocol"`
	Proxyprotocolaftertlshandshake types.String `tfsdk:"proxyprotocolaftertlshandshake"`
	Proxyprotocoltxversion         types.String `tfsdk:"proxyprotocoltxversion"`
	Srcip                          types.String `tfsdk:"srcip"`
	Srcippersistency               types.String `tfsdk:"srcippersistency"`
	Td                             types.Int64  `tfsdk:"td"`
}

func (r *NetprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the netprofile resource.",
			},
			"mbf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Response will be sent using learnt info if enabled. When creating a netprofile, if you do not set this parameter, the netprofile inherits the global MBF setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the netprofile",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the net profile. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created. Choose a name that helps identify the net profile.",
			},
			"overridelsn": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "USNIP/USIP settings override LSN settings for configured\n              service/virtual server traffic..",
			},
			"proxyprotocol": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Proxy Protocol Action (Enabled/Disabled)",
			},
			"proxyprotocolaftertlshandshake": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "ADC doesnt look for proxy header before TLS handshake, if enabled. Proxy protocol parsed after TLS handshake",
			},
			"proxyprotocoltxversion": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("V1"),
				Description: "Proxy Protocol Version (V1/V2)",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or the name of an IP set.",
			},
			"srcippersistency": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "When the net profile is associated with a virtual server or its bound services, this option enables the Citrix ADC to use the same  address, specified in the net profile, to communicate to servers for all sessions initiated from a particular client to the virtual server.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func netprofileGetThePayloadFromtheConfig(ctx context.Context, data *NetprofileResourceModel) network.Netprofile {
	tflog.Debug(ctx, "In netprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	netprofile := network.Netprofile{}
	if !data.Mbf.IsNull() {
		netprofile.Mbf = data.Mbf.ValueString()
	}
	if !data.Name.IsNull() {
		netprofile.Name = data.Name.ValueString()
	}
	if !data.Overridelsn.IsNull() {
		netprofile.Overridelsn = data.Overridelsn.ValueString()
	}
	if !data.Proxyprotocol.IsNull() {
		netprofile.Proxyprotocol = data.Proxyprotocol.ValueString()
	}
	if !data.Proxyprotocolaftertlshandshake.IsNull() {
		netprofile.Proxyprotocolaftertlshandshake = data.Proxyprotocolaftertlshandshake.ValueString()
	}
	if !data.Proxyprotocoltxversion.IsNull() {
		netprofile.Proxyprotocoltxversion = data.Proxyprotocoltxversion.ValueString()
	}
	if !data.Srcip.IsNull() {
		netprofile.Srcip = data.Srcip.ValueString()
	}
	if !data.Srcippersistency.IsNull() {
		netprofile.Srcippersistency = data.Srcippersistency.ValueString()
	}
	if !data.Td.IsNull() {
		netprofile.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return netprofile
}

func netprofileSetAttrFromGet(ctx context.Context, data *NetprofileResourceModel, getResponseData map[string]interface{}) *NetprofileResourceModel {
	tflog.Debug(ctx, "In netprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["mbf"]; ok && val != nil {
		data.Mbf = types.StringValue(val.(string))
	} else {
		data.Mbf = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["overridelsn"]; ok && val != nil {
		data.Overridelsn = types.StringValue(val.(string))
	} else {
		data.Overridelsn = types.StringNull()
	}
	if val, ok := getResponseData["proxyprotocol"]; ok && val != nil {
		data.Proxyprotocol = types.StringValue(val.(string))
	} else {
		data.Proxyprotocol = types.StringNull()
	}
	if val, ok := getResponseData["proxyprotocolaftertlshandshake"]; ok && val != nil {
		data.Proxyprotocolaftertlshandshake = types.StringValue(val.(string))
	} else {
		data.Proxyprotocolaftertlshandshake = types.StringNull()
	}
	if val, ok := getResponseData["proxyprotocoltxversion"]; ok && val != nil {
		data.Proxyprotocoltxversion = types.StringValue(val.(string))
	} else {
		data.Proxyprotocoltxversion = types.StringNull()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else {
		data.Srcip = types.StringNull()
	}
	if val, ok := getResponseData["srcippersistency"]; ok && val != nil {
		data.Srcippersistency = types.StringValue(val.(string))
	} else {
		data.Srcippersistency = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
