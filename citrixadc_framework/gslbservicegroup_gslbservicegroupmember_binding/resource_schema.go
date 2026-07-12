package gslbservicegroup_gslbservicegroupmember_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbservicegroupGslbservicegroupmemberBindingResourceModel describes the resource data model.
type GslbservicegroupGslbservicegroupmemberBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	Ip               types.String `tfsdk:"ip"`
	Order            types.Int64  `tfsdk:"order"`
	Port             types.Int64  `tfsdk:"port"`
	Publicip         types.String `tfsdk:"publicip"`
	Publicport       types.Int64  `tfsdk:"publicport"`
	Servername       types.String `tfsdk:"servername"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Siteprefix       types.String `tfsdk:"siteprefix"`
	State            types.String `tfsdk:"state"`
	Weight           types.Int64  `tfsdk:"weight"`
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbservicegroup_gslbservicegroupmember_binding resource.",
			},
			"hashid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.",
			},
			"ip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP Address.",
			},
			"order": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Order number to be assigned to the gslb servicegroup member",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Server port number.",
			},
			"publicip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.",
			},
			"publicport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.",
			},
			"servername": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the server to which to bind the service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the GSLB service group.",
			},
			"siteprefix": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("ENABLED"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Initial state of the GSLB service group.",
			},
			"weight": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},
		},
	}
}

func gslbservicegroup_gslbservicegroupmember_bindingGetThePayloadFromtheConfig(ctx context.Context, data *GslbservicegroupGslbservicegroupmemberBindingResourceModel) gslb.Gslbservicegroupgslbservicegroupmemberbinding {
	tflog.Debug(ctx, "In gslbservicegroup_gslbservicegroupmember_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbservicegroup_gslbservicegroupmember_binding := gslb.Gslbservicegroupgslbservicegroupmemberbinding{}
	if !data.Hashid.IsNull() && !data.Hashid.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.Ip = data.Ip.ValueString()
	}
	if !data.Order.IsNull() && !data.Order.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Publicip.IsNull() && !data.Publicip.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.Publicip = data.Publicip.ValueString()
	}
	if !data.Publicport.IsNull() && !data.Publicport.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.Publicport = utils.IntPtr(int(data.Publicport.ValueInt64()))
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.Servername = data.Servername.ValueString()
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Siteprefix.IsNull() && !data.Siteprefix.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.Siteprefix = data.Siteprefix.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.State = data.State.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		gslbservicegroup_gslbservicegroupmember_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return gslbservicegroup_gslbservicegroupmember_binding
}

// gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGet is the RESOURCE-side
// state setter. It preserves the prior plan/state value for attributes the NITRO GET
// response does not echo back so that an Optional+Computed input the user supplied is
// not nulled and does not trigger an "inconsistent result after apply" error.
// It re-derives data.Id in the new key:value format so a legacy SDK v2 id is upgraded
// on Read.
func gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGet(ctx context.Context, data *GslbservicegroupGslbservicegroupmemberBindingResourceModel, getResponseData map[string]interface{}) *GslbservicegroupGslbservicegroupmemberBindingResourceModel {
	tflog.Debug(ctx, "In gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGet Function")

	// Convert API response to model. Every Optional+Computed attribute MUST end up
	// with a KNOWN value after apply, otherwise Terraform errors with "provider still
	// indicated an unknown value". So assign the echoed value when present, and a typed
	// Null when the GET response omits it -- never leave a Computed attribute unknown.
	if val, ok := getResponseData["hashid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hashid = types.Int64Value(intVal)
		}
	} else {
		data.Hashid = types.Int64Null()
	}
	if val, ok := getResponseData["ip"]; ok && val != nil {
		data.Ip = types.StringValue(val.(string))
	} else {
		data.Ip = types.StringNull()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["publicip"]; ok && val != nil {
		data.Publicip = types.StringValue(val.(string))
	} else {
		data.Publicip = types.StringNull()
	}
	if val, ok := getResponseData["publicport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Publicport = types.Int64Value(intVal)
		}
	} else {
		data.Publicport = types.Int64Null()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["siteprefix"]; ok && val != nil {
		data.Siteprefix = types.StringValue(val.(string))
	} else {
		data.Siteprefix = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	// Identity mirrors SDK v2: servicegroupname + (servername OR ip) + optional port.
	effectiveServername := data.Servername.ValueString()
	if effectiveServername == "" {
		effectiveServername = data.Ip.ValueString()
	}
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(data.Servicegroupname.ValueString())))
	idParts = append(idParts, fmt.Sprintf("servername:%s", utils.UrlEncode(effectiveServername)))
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		idParts = append(idParts, fmt.Sprintf("port:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Port.ValueInt64()))))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter. The datasource has no prior plan/state to preserve, so it
// faithfully copies every field from the GET response (nulling absent fields) and
// computes data.Id itself.
func gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGetForDatasource(ctx context.Context, data *GslbservicegroupGslbservicegroupmemberBindingResourceModel, getResponseData map[string]interface{}) *GslbservicegroupGslbservicegroupmemberBindingResourceModel {
	tflog.Debug(ctx, "In gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["hashid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hashid = types.Int64Value(intVal)
		}
	} else {
		data.Hashid = types.Int64Null()
	}
	if val, ok := getResponseData["ip"]; ok && val != nil {
		data.Ip = types.StringValue(val.(string))
	} else {
		data.Ip = types.StringNull()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["publicip"]; ok && val != nil {
		data.Publicip = types.StringValue(val.(string))
	} else {
		data.Publicip = types.StringNull()
	}
	if val, ok := getResponseData["publicport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Publicport = types.Int64Value(intVal)
		}
	} else {
		data.Publicport = types.Int64Null()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["siteprefix"]; ok && val != nil {
		data.Siteprefix = types.StringValue(val.(string))
	} else {
		data.Siteprefix = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the datasource (no Create to set it).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ip.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("port:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Port.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("servername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
