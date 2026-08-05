package netprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "Response will be sent using learnt info if enabled. When creating a netprofile, if you do not set this parameter, the netprofile inherits the global MBF setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the netprofile",
			},
			"name": schema.StringAttribute{
				// SDK v2: Optional+Computed+ForceNew. Keep Optional+Computed (a name is
				// auto-generated in Create when omitted, mirroring the SDK v2 resource).
				// UseStateForUnknown keeps the (possibly generated) name stable across
				// refreshes; RequiresReplaceIfConfigured reproduces ForceNew only when the
				// user actually configured the name and changes it.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Name for the net profile. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created. Choose a name that helps identify the net profile.",
			},
			"overridelsn": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "USNIP/USIP settings override LSN settings for configured\n              service/virtual server traffic..",
			},
			"proxyprotocol": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "Proxy Protocol Action (Enabled/Disabled)",
			},
			"proxyprotocolaftertlshandshake": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "ADC doesnt look for proxy header before TLS handshake, if enabled. Proxy protocol parsed after TLS handshake",
			},
			"proxyprotocoltxversion": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "Proxy Protocol Version (V1/V2)",
			},
			"srcip": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "IP address or the name of an IP set.",
			},
			"srcippersistency": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "When the net profile is associated with a virtual server or its bound services, this option enables the Citrix ADC to use the same  address, specified in the net profile, to communicate to servers for all sessions initiated from a particular client to the virtual server.",
			},
			"td": schema.Int64Attribute{
				// SDK v2: Optional+Computed, NOT ForceNew. The auto-gen added a spurious
				// RequiresReplace (from is_updateable:false) that is removed here to match
				// the SDK v2 contract. td is create-only at the NITRO layer (excluded from
				// the update payload), so it is not sent on Update.
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func netprofileGetThePayloadFromthePlan(ctx context.Context, data *NetprofileResourceModel) network.Netprofile {
	tflog.Debug(ctx, "In netprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model (used for the NITRO add). Skip
	// unknown/null values so that Optional+Computed attributes the user omitted
	// fall through to the appliance defaults.
	netprofile := network.Netprofile{}
	if !data.Mbf.IsNull() && !data.Mbf.IsUnknown() {
		netprofile.Mbf = data.Mbf.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		netprofile.Name = data.Name.ValueString()
	}
	if !data.Overridelsn.IsNull() && !data.Overridelsn.IsUnknown() {
		netprofile.Overridelsn = data.Overridelsn.ValueString()
	}
	if !data.Proxyprotocol.IsNull() && !data.Proxyprotocol.IsUnknown() {
		netprofile.Proxyprotocol = data.Proxyprotocol.ValueString()
	}
	if !data.Proxyprotocolaftertlshandshake.IsNull() && !data.Proxyprotocolaftertlshandshake.IsUnknown() {
		netprofile.Proxyprotocolaftertlshandshake = data.Proxyprotocolaftertlshandshake.ValueString()
	}
	if !data.Proxyprotocoltxversion.IsNull() && !data.Proxyprotocoltxversion.IsUnknown() {
		netprofile.Proxyprotocoltxversion = data.Proxyprotocoltxversion.ValueString()
	}
	if !data.Srcip.IsNull() && !data.Srcip.IsUnknown() {
		netprofile.Srcip = data.Srcip.ValueString()
	}
	if !data.Srcippersistency.IsNull() && !data.Srcippersistency.IsUnknown() {
		netprofile.Srcippersistency = data.Srcippersistency.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		netprofile.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return netprofile
}

func netprofileGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *NetprofileResourceModel) network.Netprofile {
	tflog.Debug(ctx, "In netprofileGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body restricted to NITRO-updatable fields. Per the NITRO
	// doc, the netprofile update payload does NOT include td, so it is omitted here
	// (td is create-only). name is always sent as the resource key.
	netprofile := network.Netprofile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		netprofile.Name = data.Name.ValueString()
	}
	if !data.Mbf.IsNull() && !data.Mbf.IsUnknown() {
		netprofile.Mbf = data.Mbf.ValueString()
	}
	if !data.Overridelsn.IsNull() && !data.Overridelsn.IsUnknown() {
		netprofile.Overridelsn = data.Overridelsn.ValueString()
	}
	if !data.Proxyprotocol.IsNull() && !data.Proxyprotocol.IsUnknown() {
		netprofile.Proxyprotocol = data.Proxyprotocol.ValueString()
	}
	if !data.Proxyprotocolaftertlshandshake.IsNull() && !data.Proxyprotocolaftertlshandshake.IsUnknown() {
		netprofile.Proxyprotocolaftertlshandshake = data.Proxyprotocolaftertlshandshake.ValueString()
	}
	if !data.Proxyprotocoltxversion.IsNull() && !data.Proxyprotocoltxversion.IsUnknown() {
		netprofile.Proxyprotocoltxversion = data.Proxyprotocoltxversion.ValueString()
	}
	if !data.Srcip.IsNull() && !data.Srcip.IsUnknown() {
		netprofile.Srcip = data.Srcip.ValueString()
	}
	if !data.Srcippersistency.IsNull() && !data.Srcippersistency.IsUnknown() {
		netprofile.Srcippersistency = data.Srcippersistency.ValueString()
	}

	return netprofile
}

func netprofileSetAttrFromGet(ctx context.Context, data *NetprofileResourceModel, getResponseData map[string]interface{}) *NetprofileResourceModel {
	tflog.Debug(ctx, "In netprofileSetAttrFromGet Function")

	// Convert API response to model. The else-branches only null a value when the
	// current model value is unknown; a known configured value that NITRO happens
	// to omit from GET (omit-on-default) is preserved to avoid a spurious diff /
	// "inconsistent result after apply".
	if val, ok := getResponseData["mbf"]; ok && val != nil {
		data.Mbf = types.StringValue(val.(string))
	} else if data.Mbf.IsUnknown() {
		data.Mbf = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["overridelsn"]; ok && val != nil {
		data.Overridelsn = types.StringValue(val.(string))
	} else if data.Overridelsn.IsUnknown() {
		data.Overridelsn = types.StringNull()
	}
	if val, ok := getResponseData["proxyprotocol"]; ok && val != nil {
		data.Proxyprotocol = types.StringValue(val.(string))
	} else if data.Proxyprotocol.IsUnknown() {
		data.Proxyprotocol = types.StringNull()
	}
	if val, ok := getResponseData["proxyprotocolaftertlshandshake"]; ok && val != nil {
		data.Proxyprotocolaftertlshandshake = types.StringValue(val.(string))
	} else if data.Proxyprotocolaftertlshandshake.IsUnknown() {
		data.Proxyprotocolaftertlshandshake = types.StringNull()
	}
	if val, ok := getResponseData["proxyprotocoltxversion"]; ok && val != nil {
		data.Proxyprotocoltxversion = types.StringValue(val.(string))
	} else if data.Proxyprotocoltxversion.IsUnknown() {
		data.Proxyprotocoltxversion = types.StringNull()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else if data.Srcip.IsUnknown() {
		data.Srcip = types.StringNull()
	}
	if val, ok := getResponseData["srcippersistency"]; ok && val != nil {
		data.Srcippersistency = types.StringValue(val.(string))
	} else if data.Srcippersistency.IsUnknown() {
		data.Srcippersistency = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value (name) as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
