package nd6ravariables

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Nd6ravariablesResourceModel describes the resource data model.
type Nd6ravariablesResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Ceaserouteradv           types.String `tfsdk:"ceaserouteradv"`
	Currhoplimit             types.Int64  `tfsdk:"currhoplimit"`
	Defaultlifetime          types.Int64  `tfsdk:"defaultlifetime"`
	Linkmtu                  types.Int64  `tfsdk:"linkmtu"`
	Managedaddrconfig        types.String `tfsdk:"managedaddrconfig"`
	Maxrtadvinterval         types.Int64  `tfsdk:"maxrtadvinterval"`
	Minrtadvinterval         types.Int64  `tfsdk:"minrtadvinterval"`
	Onlyunicastrtadvresponse types.String `tfsdk:"onlyunicastrtadvresponse"`
	Otheraddrconfig          types.String `tfsdk:"otheraddrconfig"`
	Reachabletime            types.Int64  `tfsdk:"reachabletime"`
	Retranstime              types.Int64  `tfsdk:"retranstime"`
	Sendrouteradv            types.String `tfsdk:"sendrouteradv"`
	Srclinklayeraddroption   types.String `tfsdk:"srclinklayeraddroption"`
	Vlan                     types.Int64  `tfsdk:"vlan"`
}

func (r *Nd6ravariablesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nd6ravariables resource.",
			},
			"ceaserouteradv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cease router advertisements on this vlan.",
			},
			"currhoplimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Current Hop limit.",
			},
			"defaultlifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Default life time, in seconds.",
			},
			"linkmtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The Link MTU.",
			},
			"managedaddrconfig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value to be placed in the Managed address configuration flag field.",
			},
			"maxrtadvinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time allowed between unsolicited multicast RAs, in seconds.",
			},
			"minrtadvinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum time interval between RA messages, in seconds.",
			},
			"onlyunicastrtadvresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send only Unicast Router Advertisements in respond to Router Solicitations.",
			},
			"otheraddrconfig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value to be placed in the Other configuration flag field.",
			},
			"reachabletime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Reachable time, in milliseconds.",
			},
			"retranstime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Retransmission time, in milliseconds.",
			},
			"sendrouteradv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "whether the router sends periodic RAs and responds to Router Solicitations.",
			},
			"srclinklayeraddroption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include source link layer address option in RA messages.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "The VLAN number.",
			},
		},
	}
}

func nd6ravariablesGetThePayloadFromthePlan(ctx context.Context, data *Nd6ravariablesResourceModel) network.Nd6ravariables {
	tflog.Debug(ctx, "In nd6ravariablesGetThePayloadFromthePlan Function")

	// Create API request body from the model. Only send attributes the user
	// actually configured (non-null AND non-unknown), mirroring the SDK v2
	// GetRawConfig().IsNull() guards so that unconfigured Optional+Computed
	// attributes are not overwritten with a zero value on the ADC.
	nd6ravariables := network.Nd6ravariables{}
	if !data.Ceaserouteradv.IsNull() && !data.Ceaserouteradv.IsUnknown() {
		nd6ravariables.Ceaserouteradv = data.Ceaserouteradv.ValueString()
	}
	if !data.Currhoplimit.IsNull() && !data.Currhoplimit.IsUnknown() {
		nd6ravariables.Currhoplimit = utils.IntPtr(int(data.Currhoplimit.ValueInt64()))
	}
	if !data.Defaultlifetime.IsNull() && !data.Defaultlifetime.IsUnknown() {
		nd6ravariables.Defaultlifetime = utils.IntPtr(int(data.Defaultlifetime.ValueInt64()))
	}
	if !data.Linkmtu.IsNull() && !data.Linkmtu.IsUnknown() {
		nd6ravariables.Linkmtu = utils.IntPtr(int(data.Linkmtu.ValueInt64()))
	}
	if !data.Managedaddrconfig.IsNull() && !data.Managedaddrconfig.IsUnknown() {
		nd6ravariables.Managedaddrconfig = data.Managedaddrconfig.ValueString()
	}
	if !data.Maxrtadvinterval.IsNull() && !data.Maxrtadvinterval.IsUnknown() {
		nd6ravariables.Maxrtadvinterval = utils.IntPtr(int(data.Maxrtadvinterval.ValueInt64()))
	}
	if !data.Minrtadvinterval.IsNull() && !data.Minrtadvinterval.IsUnknown() {
		nd6ravariables.Minrtadvinterval = utils.IntPtr(int(data.Minrtadvinterval.ValueInt64()))
	}
	if !data.Onlyunicastrtadvresponse.IsNull() && !data.Onlyunicastrtadvresponse.IsUnknown() {
		nd6ravariables.Onlyunicastrtadvresponse = data.Onlyunicastrtadvresponse.ValueString()
	}
	if !data.Otheraddrconfig.IsNull() && !data.Otheraddrconfig.IsUnknown() {
		nd6ravariables.Otheraddrconfig = data.Otheraddrconfig.ValueString()
	}
	if !data.Reachabletime.IsNull() && !data.Reachabletime.IsUnknown() {
		nd6ravariables.Reachabletime = utils.IntPtr(int(data.Reachabletime.ValueInt64()))
	}
	if !data.Retranstime.IsNull() && !data.Retranstime.IsUnknown() {
		nd6ravariables.Retranstime = utils.IntPtr(int(data.Retranstime.ValueInt64()))
	}
	if !data.Sendrouteradv.IsNull() && !data.Sendrouteradv.IsUnknown() {
		nd6ravariables.Sendrouteradv = data.Sendrouteradv.ValueString()
	}
	if !data.Srclinklayeraddroption.IsNull() && !data.Srclinklayeraddroption.IsUnknown() {
		nd6ravariables.Srclinklayeraddroption = data.Srclinklayeraddroption.ValueString()
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		nd6ravariables.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return nd6ravariables
}

func nd6ravariablesSetAttrFromGet(ctx context.Context, data *Nd6ravariablesResourceModel, getResponseData map[string]interface{}) *Nd6ravariablesResourceModel {
	tflog.Debug(ctx, "In nd6ravariablesSetAttrFromGet Function")

	// Convert API response to model. The else branches only null the value when
	// it is Unknown (omit-on-default guard): never clobber a known configured
	// value that NITRO happens to omit from the GET response.
	if val, ok := getResponseData["ceaserouteradv"]; ok && val != nil {
		data.Ceaserouteradv = types.StringValue(val.(string))
	} else if data.Ceaserouteradv.IsUnknown() {
		data.Ceaserouteradv = types.StringNull()
	}
	if val, ok := getResponseData["currhoplimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Currhoplimit = types.Int64Value(intVal)
		}
	} else if data.Currhoplimit.IsUnknown() {
		data.Currhoplimit = types.Int64Null()
	}
	if val, ok := getResponseData["defaultlifetime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Defaultlifetime = types.Int64Value(intVal)
		}
	} else if data.Defaultlifetime.IsUnknown() {
		data.Defaultlifetime = types.Int64Null()
	}
	if val, ok := getResponseData["linkmtu"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Linkmtu = types.Int64Value(intVal)
		}
	} else if data.Linkmtu.IsUnknown() {
		data.Linkmtu = types.Int64Null()
	}
	if val, ok := getResponseData["managedaddrconfig"]; ok && val != nil {
		data.Managedaddrconfig = types.StringValue(val.(string))
	} else if data.Managedaddrconfig.IsUnknown() {
		data.Managedaddrconfig = types.StringNull()
	}
	if val, ok := getResponseData["maxrtadvinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxrtadvinterval = types.Int64Value(intVal)
		}
	} else if data.Maxrtadvinterval.IsUnknown() {
		data.Maxrtadvinterval = types.Int64Null()
	}
	if val, ok := getResponseData["minrtadvinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minrtadvinterval = types.Int64Value(intVal)
		}
	} else if data.Minrtadvinterval.IsUnknown() {
		data.Minrtadvinterval = types.Int64Null()
	}
	if val, ok := getResponseData["onlyunicastrtadvresponse"]; ok && val != nil {
		data.Onlyunicastrtadvresponse = types.StringValue(val.(string))
	} else if data.Onlyunicastrtadvresponse.IsUnknown() {
		data.Onlyunicastrtadvresponse = types.StringNull()
	}
	if val, ok := getResponseData["otheraddrconfig"]; ok && val != nil {
		data.Otheraddrconfig = types.StringValue(val.(string))
	} else if data.Otheraddrconfig.IsUnknown() {
		data.Otheraddrconfig = types.StringNull()
	}
	if val, ok := getResponseData["reachabletime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Reachabletime = types.Int64Value(intVal)
		}
	} else if data.Reachabletime.IsUnknown() {
		data.Reachabletime = types.Int64Null()
	}
	if val, ok := getResponseData["retranstime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retranstime = types.Int64Value(intVal)
		}
	} else if data.Retranstime.IsUnknown() {
		data.Retranstime = types.Int64Null()
	}
	if val, ok := getResponseData["sendrouteradv"]; ok && val != nil {
		data.Sendrouteradv = types.StringValue(val.(string))
	} else if data.Sendrouteradv.IsUnknown() {
		data.Sendrouteradv = types.StringNull()
	}
	if val, ok := getResponseData["srclinklayeraddroption"]; ok && val != nil {
		data.Srclinklayeraddroption = types.StringValue(val.(string))
	} else if data.Srclinklayeraddroption.IsUnknown() {
		data.Srclinklayeraddroption = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else if data.Vlan.IsUnknown() {
		data.Vlan = types.Int64Null()
	}

	// Set ID for the resource.
	// Single unique attribute (vlan) - matches SDK v2 d.SetId(strconv.Itoa(vlan)).
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Vlan.ValueInt64()))

	return data
}
