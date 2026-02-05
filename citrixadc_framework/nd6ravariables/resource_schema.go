package nd6ravariables

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
				Default:     int64default.StaticInt64(64),
				Description: "Current Hop limit.",
			},
			"defaultlifetime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1800),
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
				Default:     int64default.StaticInt64(600),
				Description: "Maximum time allowed between unsolicited multicast RAs, in seconds.",
			},
			"minrtadvinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(198),
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
				Default:     stringdefault.StaticString("True"),
				Description: "Include source link layer address option in RA messages.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "The VLAN number.",
			},
		},
	}
}

func nd6ravariablesGetThePayloadFromtheConfig(ctx context.Context, data *Nd6ravariablesResourceModel) network.Nd6ravariables {
	tflog.Debug(ctx, "In nd6ravariablesGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nd6ravariables := network.Nd6ravariables{}
	if !data.Ceaserouteradv.IsNull() {
		nd6ravariables.Ceaserouteradv = data.Ceaserouteradv.ValueString()
	}
	if !data.Currhoplimit.IsNull() {
		nd6ravariables.Currhoplimit = utils.IntPtr(int(data.Currhoplimit.ValueInt64()))
	}
	if !data.Defaultlifetime.IsNull() {
		nd6ravariables.Defaultlifetime = utils.IntPtr(int(data.Defaultlifetime.ValueInt64()))
	}
	if !data.Linkmtu.IsNull() {
		nd6ravariables.Linkmtu = utils.IntPtr(int(data.Linkmtu.ValueInt64()))
	}
	if !data.Managedaddrconfig.IsNull() {
		nd6ravariables.Managedaddrconfig = data.Managedaddrconfig.ValueString()
	}
	if !data.Maxrtadvinterval.IsNull() {
		nd6ravariables.Maxrtadvinterval = utils.IntPtr(int(data.Maxrtadvinterval.ValueInt64()))
	}
	if !data.Minrtadvinterval.IsNull() {
		nd6ravariables.Minrtadvinterval = utils.IntPtr(int(data.Minrtadvinterval.ValueInt64()))
	}
	if !data.Onlyunicastrtadvresponse.IsNull() {
		nd6ravariables.Onlyunicastrtadvresponse = data.Onlyunicastrtadvresponse.ValueString()
	}
	if !data.Otheraddrconfig.IsNull() {
		nd6ravariables.Otheraddrconfig = data.Otheraddrconfig.ValueString()
	}
	if !data.Reachabletime.IsNull() {
		nd6ravariables.Reachabletime = utils.IntPtr(int(data.Reachabletime.ValueInt64()))
	}
	if !data.Retranstime.IsNull() {
		nd6ravariables.Retranstime = utils.IntPtr(int(data.Retranstime.ValueInt64()))
	}
	if !data.Sendrouteradv.IsNull() {
		nd6ravariables.Sendrouteradv = data.Sendrouteradv.ValueString()
	}
	if !data.Srclinklayeraddroption.IsNull() {
		nd6ravariables.Srclinklayeraddroption = data.Srclinklayeraddroption.ValueString()
	}
	if !data.Vlan.IsNull() {
		nd6ravariables.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return nd6ravariables
}

func nd6ravariablesSetAttrFromGet(ctx context.Context, data *Nd6ravariablesResourceModel, getResponseData map[string]interface{}) *Nd6ravariablesResourceModel {
	tflog.Debug(ctx, "In nd6ravariablesSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ceaserouteradv"]; ok && val != nil {
		data.Ceaserouteradv = types.StringValue(val.(string))
	} else {
		data.Ceaserouteradv = types.StringNull()
	}
	if val, ok := getResponseData["currhoplimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Currhoplimit = types.Int64Value(intVal)
		}
	} else {
		data.Currhoplimit = types.Int64Null()
	}
	if val, ok := getResponseData["defaultlifetime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Defaultlifetime = types.Int64Value(intVal)
		}
	} else {
		data.Defaultlifetime = types.Int64Null()
	}
	if val, ok := getResponseData["linkmtu"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Linkmtu = types.Int64Value(intVal)
		}
	} else {
		data.Linkmtu = types.Int64Null()
	}
	if val, ok := getResponseData["managedaddrconfig"]; ok && val != nil {
		data.Managedaddrconfig = types.StringValue(val.(string))
	} else {
		data.Managedaddrconfig = types.StringNull()
	}
	if val, ok := getResponseData["maxrtadvinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxrtadvinterval = types.Int64Value(intVal)
		}
	} else {
		data.Maxrtadvinterval = types.Int64Null()
	}
	if val, ok := getResponseData["minrtadvinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minrtadvinterval = types.Int64Value(intVal)
		}
	} else {
		data.Minrtadvinterval = types.Int64Null()
	}
	if val, ok := getResponseData["onlyunicastrtadvresponse"]; ok && val != nil {
		data.Onlyunicastrtadvresponse = types.StringValue(val.(string))
	} else {
		data.Onlyunicastrtadvresponse = types.StringNull()
	}
	if val, ok := getResponseData["otheraddrconfig"]; ok && val != nil {
		data.Otheraddrconfig = types.StringValue(val.(string))
	} else {
		data.Otheraddrconfig = types.StringNull()
	}
	if val, ok := getResponseData["reachabletime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Reachabletime = types.Int64Value(intVal)
		}
	} else {
		data.Reachabletime = types.Int64Null()
	}
	if val, ok := getResponseData["retranstime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retranstime = types.Int64Value(intVal)
		}
	} else {
		data.Retranstime = types.Int64Null()
	}
	if val, ok := getResponseData["sendrouteradv"]; ok && val != nil {
		data.Sendrouteradv = types.StringValue(val.(string))
	} else {
		data.Sendrouteradv = types.StringNull()
	}
	if val, ok := getResponseData["srclinklayeraddroption"]; ok && val != nil {
		data.Srclinklayeraddroption = types.StringValue(val.(string))
	} else {
		data.Srclinklayeraddroption = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Vlan.ValueInt64()))

	return data
}
