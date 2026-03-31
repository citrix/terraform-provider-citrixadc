package gslbvserver_gslbservicegroup_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbvserverGslbservicegroupBindingResourceModel describes the resource data model.
type GslbvserverGslbservicegroupBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Order            types.Int64  `tfsdk:"order"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
}

func (r *GslbvserverGslbservicegroupBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbvserver_gslbservicegroup_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server on which to perform the binding operation.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"servicegroupname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The GSLB service group name bound to the selected GSLB virtual server.",
			},
		},
	}
}

func gslbvserver_gslbservicegroup_bindingGetThePayloadFromtheConfig(ctx context.Context, data *GslbvserverGslbservicegroupBindingResourceModel) gslb.Gslbvservergslbservicegroupbinding {
	tflog.Debug(ctx, "In gslbvserver_gslbservicegroup_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbvserver_gslbservicegroup_binding := gslb.Gslbvservergslbservicegroupbinding{}
	if !data.Name.IsNull() {
		gslbvserver_gslbservicegroup_binding.Name = data.Name.ValueString()
	}
	if !data.Order.IsNull() {
		gslbvserver_gslbservicegroup_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Servicegroupname.IsNull() {
		gslbvserver_gslbservicegroup_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}

	return gslbvserver_gslbservicegroup_binding
}

func gslbvserver_gslbservicegroup_bindingSetAttrFromGet(ctx context.Context, data *GslbvserverGslbservicegroupBindingResourceModel, getResponseData map[string]interface{}) *GslbvserverGslbservicegroupBindingResourceModel {
	tflog.Debug(ctx, "In gslbvserver_gslbservicegroup_bindingSetAttrFromGet Function")

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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
