package gslbvserver_gslbservice_binding

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

// GslbvserverGslbserviceBindingResourceModel describes the resource data model.
type GslbvserverGslbserviceBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Domainname  types.String `tfsdk:"domainname"`
	Name        types.String `tfsdk:"name"`
	Order       types.Int64  `tfsdk:"order"`
	Servicename types.String `tfsdk:"servicename"`
	Weight      types.Int64  `tfsdk:"weight"`
}

func (r *GslbvserverGslbserviceBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbvserver_gslbservice_binding resource.",
			},
			"domainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
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
			"servicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the GSLB service for which to change the weight.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight for the service.",
			},
		},
	}
}

func gslbvserver_gslbservice_bindingGetThePayloadFromtheConfig(ctx context.Context, data *GslbvserverGslbserviceBindingResourceModel) gslb.Gslbvservergslbservicebinding {
	tflog.Debug(ctx, "In gslbvserver_gslbservice_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbvserver_gslbservice_binding := gslb.Gslbvservergslbservicebinding{}
	if !data.Domainname.IsNull() {
		gslbvserver_gslbservice_binding.Domainname = data.Domainname.ValueString()
	}
	if !data.Name.IsNull() {
		gslbvserver_gslbservice_binding.Name = data.Name.ValueString()
	}
	if !data.Order.IsNull() {
		gslbvserver_gslbservice_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Servicename.IsNull() {
		gslbvserver_gslbservice_binding.Servicename = data.Servicename.ValueString()
	}
	if !data.Weight.IsNull() {
		gslbvserver_gslbservice_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return gslbvserver_gslbservice_binding
}

func gslbvserver_gslbservice_bindingSetAttrFromGet(ctx context.Context, data *GslbvserverGslbserviceBindingResourceModel, getResponseData map[string]interface{}) *GslbvserverGslbserviceBindingResourceModel {
	tflog.Debug(ctx, "In gslbvserver_gslbservice_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["domainname"]; ok && val != nil {
		data.Domainname = types.StringValue(val.(string))
	} else {
		data.Domainname = types.StringNull()
	}
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
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
