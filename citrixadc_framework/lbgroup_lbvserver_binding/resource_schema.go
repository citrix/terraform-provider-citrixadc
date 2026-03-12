package lbgroup_lbvserver_binding

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

// LbgroupLbvserverBindingResourceModel describes the resource data model.
type LbgroupLbvserverBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Vservername types.String `tfsdk:"vservername"`
}

func (r *LbgroupLbvserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbgroup_lbvserver_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the load balancing virtual server group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lbgroup\" or 'my lbgroup').",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Virtual server name.",
			},
		},
	}
}

func lbgroup_lbvserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LbgroupLbvserverBindingResourceModel) lb.Lbgrouplbvserverbinding {
	tflog.Debug(ctx, "In lbgroup_lbvserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbgroup_lbvserver_binding := lb.Lbgrouplbvserverbinding{}
	if !data.Name.IsNull() {
		lbgroup_lbvserver_binding.Name = data.Name.ValueString()
	}
	if !data.Vservername.IsNull() {
		lbgroup_lbvserver_binding.Vservername = data.Vservername.ValueString()
	}

	return lbgroup_lbvserver_binding
}

func lbgroup_lbvserver_bindingSetAttrFromGet(ctx context.Context, data *LbgroupLbvserverBindingResourceModel, getResponseData map[string]interface{}) *LbgroupLbvserverBindingResourceModel {
	tflog.Debug(ctx, "In lbgroup_lbvserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
