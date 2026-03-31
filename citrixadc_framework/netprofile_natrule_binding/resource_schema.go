package netprofile_natrule_binding

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

// NetprofileNatruleBindingResourceModel describes the resource data model.
type NetprofileNatruleBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Natrule   types.String `tfsdk:"natrule"`
	Netmask   types.String `tfsdk:"netmask"`
	Rewriteip types.String `tfsdk:"rewriteip"`
}

func (r *NetprofileNatruleBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the netprofile_natrule_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the netprofile to which to bind port ranges.",
			},
			"natrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 network address on whose traffic you want the Citrix ADC to do rewrite ip prefix.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"rewriteip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
		},
	}
}

func netprofile_natrule_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NetprofileNatruleBindingResourceModel) network.Netprofilenatrulebinding {
	tflog.Debug(ctx, "In netprofile_natrule_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	netprofile_natrule_binding := network.Netprofilenatrulebinding{}
	if !data.Name.IsNull() {
		netprofile_natrule_binding.Name = data.Name.ValueString()
	}
	if !data.Natrule.IsNull() {
		netprofile_natrule_binding.Natrule = data.Natrule.ValueString()
	}
	if !data.Netmask.IsNull() {
		netprofile_natrule_binding.Netmask = data.Netmask.ValueString()
	}
	if !data.Rewriteip.IsNull() {
		netprofile_natrule_binding.Rewriteip = data.Rewriteip.ValueString()
	}

	return netprofile_natrule_binding
}

func netprofile_natrule_bindingSetAttrFromGet(ctx context.Context, data *NetprofileNatruleBindingResourceModel, getResponseData map[string]interface{}) *NetprofileNatruleBindingResourceModel {
	tflog.Debug(ctx, "In netprofile_natrule_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["natrule"]; ok && val != nil {
		data.Natrule = types.StringValue(val.(string))
	} else {
		data.Natrule = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["rewriteip"]; ok && val != nil {
		data.Rewriteip = types.StringValue(val.(string))
	} else {
		data.Rewriteip = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("natrule:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Natrule.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
