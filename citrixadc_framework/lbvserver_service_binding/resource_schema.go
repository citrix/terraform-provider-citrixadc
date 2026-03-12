package lbvserver_service_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbvserverServiceBindingResourceModel describes the resource data model.
type LbvserverServiceBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Order       types.Int64  `tfsdk:"order"`
	Servicename types.String `tfsdk:"servicename"`
	Weight      types.Int64  `tfsdk:"weight"`
}

func (r *LbvserverServiceBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbvserver_service_binding resource.",
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
			"servicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service to bind to the virtual server.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the specified service.",
			},
		},
	}
}

func lbvserver_service_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LbvserverServiceBindingResourceModel) lb.Lbvserverservicebinding {
	tflog.Debug(ctx, "In lbvserver_service_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbvserver_service_binding := lb.Lbvserverservicebinding{}
	if !data.Name.IsNull() {
		lbvserver_service_binding.Name = data.Name.ValueString()
	}
	if !data.Order.IsNull() {
		lbvserver_service_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Servicename.IsNull() {
		lbvserver_service_binding.Servicename = data.Servicename.ValueString()
	}
	if !data.Weight.IsNull() {
		lbvserver_service_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return lbvserver_service_binding
}

func lbvserver_service_bindingSetAttrFromGet(ctx context.Context, data *LbvserverServiceBindingResourceModel, getResponseData map[string]interface{}) *LbvserverServiceBindingResourceModel {
	tflog.Debug(ctx, "In lbvserver_service_bindingSetAttrFromGet Function")

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
