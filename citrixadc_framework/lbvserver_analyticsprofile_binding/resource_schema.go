package lbvserver_analyticsprofile_binding

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

// LbvserverAnalyticsprofileBindingResourceModel describes the resource data model.
type LbvserverAnalyticsprofileBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Analyticsprofile types.String `tfsdk:"analyticsprofile"`
	Name             types.String `tfsdk:"name"`
	Order            types.Int64  `tfsdk:"order"`
}

func (r *LbvserverAnalyticsprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbvserver_analyticsprofile_binding resource.",
			},
			"analyticsprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the analytics profile bound to the LB vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.",
			},
		},
	}
}

func lbvserver_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LbvserverAnalyticsprofileBindingResourceModel) lb.Lbvserveranalyticsprofilebinding {
	tflog.Debug(ctx, "In lbvserver_analyticsprofile_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbvserver_analyticsprofile_binding := lb.Lbvserveranalyticsprofilebinding{}
	if !data.Analyticsprofile.IsNull() {
		lbvserver_analyticsprofile_binding.Analyticsprofile = data.Analyticsprofile.ValueString()
	}
	if !data.Name.IsNull() {
		lbvserver_analyticsprofile_binding.Name = data.Name.ValueString()
	}
	if !data.Order.IsNull() {
		lbvserver_analyticsprofile_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}

	return lbvserver_analyticsprofile_binding
}

func lbvserver_analyticsprofile_bindingSetAttrFromGet(ctx context.Context, data *LbvserverAnalyticsprofileBindingResourceModel, getResponseData map[string]interface{}) *LbvserverAnalyticsprofileBindingResourceModel {
	tflog.Debug(ctx, "In lbvserver_analyticsprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["analyticsprofile"]; ok && val != nil {
		data.Analyticsprofile = types.StringValue(val.(string))
	} else {
		data.Analyticsprofile = types.StringNull()
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("analyticsprofile:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Analyticsprofile.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
