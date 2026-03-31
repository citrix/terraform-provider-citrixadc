package lbvserver_servicegroup_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbvserverServicegroupBindingResourceModel describes the resource data model.
type LbvserverServicegroupBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Order            types.Int64  `tfsdk:"order"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Weight           types.Int64  `tfsdk:"weight"`
}

func (r *LbvserverServicegroupBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbvserver_servicegroup_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"servicegroupname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The service group name bound to the selected load balancing virtual server.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Integer specifying the weight of the service. A larger number specifies a greater weight. Defines the capacity of the service relative to the other services in the load balancing configuration. Determines the priority given to the service in load balancing decisions.",
			},
		},
	}
}

func lbvserver_servicegroup_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LbvserverServicegroupBindingResourceModel) lb.Lbvserverservicegroupbinding {
	tflog.Debug(ctx, "In lbvserver_servicegroup_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbvserver_servicegroup_binding := lb.Lbvserverservicegroupbinding{}
	if !data.Name.IsNull() {
		lbvserver_servicegroup_binding.Name = data.Name.ValueString()
	}
	if !data.Order.IsNull() {
		lbvserver_servicegroup_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Servicegroupname.IsNull() {
		lbvserver_servicegroup_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Weight.IsNull() {
		lbvserver_servicegroup_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return lbvserver_servicegroup_binding
}

func lbvserver_servicegroup_bindingSetAttrFromGet(ctx context.Context, data *LbvserverServicegroupBindingResourceModel, getResponseData map[string]interface{}) *LbvserverServicegroupBindingResourceModel {
	tflog.Debug(ctx, "In lbvserver_servicegroup_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
