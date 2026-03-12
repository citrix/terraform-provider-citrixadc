package crvserver_lbvserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cr"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CrvserverLbvserverBindingResourceModel describes the resource data model.
type CrvserverLbvserverBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Lbvserver types.String `tfsdk:"lbvserver"`
	Name      types.String `tfsdk:"name"`
}

func (r *CrvserverLbvserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the crvserver_lbvserver_binding resource.",
			},
			"lbvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Default target server name.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the cache redirection virtual server to which to bind the cache redirection policy.",
			},
		},
	}
}

func crvserver_lbvserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *CrvserverLbvserverBindingResourceModel) cr.Crvserverlbvserverbinding {
	tflog.Debug(ctx, "In crvserver_lbvserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	crvserver_lbvserver_binding := cr.Crvserverlbvserverbinding{}
	if !data.Lbvserver.IsNull() {
		crvserver_lbvserver_binding.Lbvserver = data.Lbvserver.ValueString()
	}
	if !data.Name.IsNull() {
		crvserver_lbvserver_binding.Name = data.Name.ValueString()
	}

	return crvserver_lbvserver_binding
}

func crvserver_lbvserver_bindingSetAttrFromGet(ctx context.Context, data *CrvserverLbvserverBindingResourceModel, getResponseData map[string]interface{}) *CrvserverLbvserverBindingResourceModel {
	tflog.Debug(ctx, "In crvserver_lbvserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["lbvserver"]; ok && val != nil {
		data.Lbvserver = types.StringValue(val.(string))
	} else {
		data.Lbvserver = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("lbvserver:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Lbvserver.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
