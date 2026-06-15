package servicegroup_servicegroupmember_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ServicegroupServicegroupmemberBindingResourceModel describes the resource data model.
type ServicegroupServicegroupmemberBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Customserverid   types.String `tfsdk:"customserverid"`
	Dbsttl           types.Int64  `tfsdk:"dbsttl"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	Ip               types.String `tfsdk:"ip"`
	Nameserver       types.String `tfsdk:"nameserver"`
	Order            types.Int64  `tfsdk:"order"`
	Port             types.Int64  `tfsdk:"port"`
	Serverid         types.Int64  `tfsdk:"serverid"`
	Servername       types.String `tfsdk:"servername"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	State            types.String `tfsdk:"state"`
	Weight           types.Int64  `tfsdk:"weight"`
}

func (r *ServicegroupServicegroupmemberBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the servicegroup_servicegroupmember_binding resource.",
			},
			"customserverid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.",
			},
			"dbsttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors",
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
			"nameserver": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver",
			},
			"order": schema.Int64Attribute{
				// "order" is never echoed back by the NITRO GET response (only
				// "orderstr" is returned), so it must NOT be Computed — otherwise a
				// config that omits it leaves the value unknown after apply
				// ("still indicated an unknown value", Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Order number to be assigned to the servicegroup member",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Server port number.",
			},
			"serverid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The  identifier for the service. This is used when the persistency type is set to Custom Server ID.",
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
				Description: "Name of the service group.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Initial state of the service group.",
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

func servicegroup_servicegroupmember_bindingGetThePayloadFromthePlan(ctx context.Context, data *ServicegroupServicegroupmemberBindingResourceModel) basic.Servicegroupservicegroupmemberbinding {
	tflog.Debug(ctx, "In servicegroup_servicegroupmember_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	servicegroup_servicegroupmember_binding := basic.Servicegroupservicegroupmemberbinding{}
	if !data.Customserverid.IsNull() && !data.Customserverid.IsUnknown() {
		servicegroup_servicegroupmember_binding.Customserverid = data.Customserverid.ValueString()
	}
	if !data.Dbsttl.IsNull() && !data.Dbsttl.IsUnknown() {
		servicegroup_servicegroupmember_binding.Dbsttl = utils.IntPtr(int(data.Dbsttl.ValueInt64()))
	}
	if !data.Hashid.IsNull() && !data.Hashid.IsUnknown() {
		servicegroup_servicegroupmember_binding.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		servicegroup_servicegroupmember_binding.Ip = data.Ip.ValueString()
	}
	if !data.Nameserver.IsNull() && !data.Nameserver.IsUnknown() {
		servicegroup_servicegroupmember_binding.Nameserver = data.Nameserver.ValueString()
	}
	if !data.Order.IsNull() && !data.Order.IsUnknown() {
		servicegroup_servicegroupmember_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		servicegroup_servicegroupmember_binding.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Serverid.IsNull() && !data.Serverid.IsUnknown() {
		servicegroup_servicegroupmember_binding.Serverid = utils.IntPtr(int(data.Serverid.ValueInt64()))
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		servicegroup_servicegroupmember_binding.Servername = data.Servername.ValueString()
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		servicegroup_servicegroupmember_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		servicegroup_servicegroupmember_binding.State = data.State.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		servicegroup_servicegroupmember_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return servicegroup_servicegroupmember_binding
}

// servicegroup_servicegroupmember_bindingSetAttrFromGet is the RESOURCE-side state
// setter. It preserves the prior plan/state value for attributes the NITRO GET
// response does not echo back (notably "order", which only appears as "orderstr"
// in GET) so that an Optional+Computed input the user supplied is not nulled and
// does not trigger an "inconsistent result after apply" error (Pattern 7 / 13).
// It does NOT recompute data.Id; the ID is set once in Create.
func servicegroup_servicegroupmember_bindingSetAttrFromGet(ctx context.Context, data *ServicegroupServicegroupmemberBindingResourceModel, getResponseData map[string]interface{}) *ServicegroupServicegroupmemberBindingResourceModel {
	tflog.Debug(ctx, "In servicegroup_servicegroupmember_bindingSetAttrFromGet Function")

	// Convert API response to model. Only overwrite a field when the GET response
	// actually echoes it; otherwise preserve the existing plan/state value.
	if val, ok := getResponseData["customserverid"]; ok && val != nil {
		data.Customserverid = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["dbsttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dbsttl = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["hashid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hashid = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["ip"]; ok && val != nil {
		data.Ip = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["nameserver"]; ok && val != nil {
		data.Nameserver = types.StringValue(val.(string))
	}
	// "order" is never echoed by the NITRO GET response (only "orderstr" is
	// returned). Preserve the existing plan/state value rather than nulling it.
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["serverid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverid = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	}

	return data
}

// servicegroup_servicegroupmember_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter (Pattern 7). The datasource has no prior plan/state to
// preserve, so it faithfully copies every field from the GET response (nulling
// absent fields) and computes data.Id itself.
func servicegroup_servicegroupmember_bindingSetAttrFromGetForDatasource(ctx context.Context, data *ServicegroupServicegroupmemberBindingResourceModel, getResponseData map[string]interface{}) *ServicegroupServicegroupmemberBindingResourceModel {
	tflog.Debug(ctx, "In servicegroup_servicegroupmember_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["customserverid"]; ok && val != nil {
		data.Customserverid = types.StringValue(val.(string))
	} else {
		data.Customserverid = types.StringNull()
	}
	if val, ok := getResponseData["dbsttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dbsttl = types.Int64Value(intVal)
		}
	} else {
		data.Dbsttl = types.Int64Null()
	}
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
	if val, ok := getResponseData["nameserver"]; ok && val != nil {
		data.Nameserver = types.StringValue(val.(string))
	} else {
		data.Nameserver = types.StringNull()
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
	if val, ok := getResponseData["serverid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverid = types.Int64Value(intVal)
		}
	} else {
		data.Serverid = types.Int64Null()
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
