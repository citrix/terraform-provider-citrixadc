package clusternodegroup_vpnvserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ClusternodegroupVpnvserverBindingResourceModel describes the resource data model.
type ClusternodegroupVpnvserverBindingResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Vserver types.String `tfsdk:"vserver"`
}

func (r *ClusternodegroupVpnvserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusternodegroup_vpnvserver_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "vserver that need to be bound to this nodegroup.",
			},
		},
	}
}

func clusternodegroup_vpnvserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *ClusternodegroupVpnvserverBindingResourceModel) cluster.Clusternodegroupvpnvserverbinding {
	tflog.Debug(ctx, "In clusternodegroup_vpnvserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	clusternodegroup_vpnvserver_binding := cluster.Clusternodegroupvpnvserverbinding{}
	if !data.Name.IsNull() {
		clusternodegroup_vpnvserver_binding.Name = data.Name.ValueString()
	}
	if !data.Vserver.IsNull() {
		clusternodegroup_vpnvserver_binding.Vserver = data.Vserver.ValueString()
	}

	return clusternodegroup_vpnvserver_binding
}

func clusternodegroup_vpnvserver_bindingSetAttrFromGet(ctx context.Context, data *ClusternodegroupVpnvserverBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupVpnvserverBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_vpnvserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	} else {
		data.Vserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vserver:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vserver.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
