package netprofile_srcportset_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NetprofileSrcportsetBindingResourceModel describes the resource data model.
type NetprofileSrcportsetBindingResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Srcportrange types.String `tfsdk:"srcportrange"`
}

func (r *NetprofileSrcportsetBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the netprofile_srcportset_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the netprofile to which to bind port ranges.",
			},
			"srcportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When the source port range is configured and associated with the netprofile bound to a service group, Citrix ADC will choose a port from the range configured for connection establishment at the backend servers.",
			},
		},
	}
}

func netprofile_srcportset_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NetprofileSrcportsetBindingResourceModel) network.Netprofilesrcportsetbinding {
	tflog.Debug(ctx, "In netprofile_srcportset_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	netprofile_srcportset_binding := network.Netprofilesrcportsetbinding{}
	if !data.Name.IsNull() {
		netprofile_srcportset_binding.Name = data.Name.ValueString()
	}
	if !data.Srcportrange.IsNull() {
		netprofile_srcportset_binding.Srcportrange = data.Srcportrange.ValueString()
	}

	return netprofile_srcportset_binding
}

func netprofile_srcportset_bindingSetAttrFromGet(ctx context.Context, data *NetprofileSrcportsetBindingResourceModel, getResponseData map[string]interface{}) *NetprofileSrcportsetBindingResourceModel {
	tflog.Debug(ctx, "In netprofile_srcportset_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["srcportrange"]; ok && val != nil {
		data.Srcportrange = types.StringValue(val.(string))
	} else {
		data.Srcportrange = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("srcportrange:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Srcportrange.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
