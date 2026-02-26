package clusternode_routemonitor_binding

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

// ClusternodeRoutemonitorBindingResourceModel describes the resource data model.
type ClusternodeRoutemonitorBindingResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Netmask      types.String `tfsdk:"netmask"`
	Nodeid       types.Int64  `tfsdk:"nodeid"`
	Routemonitor types.String `tfsdk:"routemonitor"`
}

func (r *ClusternodeRoutemonitorBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusternode_routemonitor_binding resource.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask.",
			},
			"nodeid": schema.Int64Attribute{
				Required:    true,
				Description: "A number that uniquely identifies the cluster node.",
			},
			"routemonitor": schema.StringAttribute{
				Required:    true,
				Description: "The IP address (IPv4 or IPv6).",
			},
		},
	}
}

func clusternode_routemonitor_bindingGetThePayloadFromtheConfig(ctx context.Context, data *ClusternodeRoutemonitorBindingResourceModel) cluster.Clusternoderoutemonitorbinding {
	tflog.Debug(ctx, "In clusternode_routemonitor_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	clusternode_routemonitor_binding := cluster.Clusternoderoutemonitorbinding{}
	if !data.Netmask.IsNull() {
		clusternode_routemonitor_binding.Netmask = data.Netmask.ValueString()
	}
	if !data.Nodeid.IsNull() {
		clusternode_routemonitor_binding.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Routemonitor.IsNull() {
		clusternode_routemonitor_binding.Routemonitor = data.Routemonitor.ValueString()
	}

	return clusternode_routemonitor_binding
}

func clusternode_routemonitor_bindingSetAttrFromGet(ctx context.Context, data *ClusternodeRoutemonitorBindingResourceModel, getResponseData map[string]interface{}) *ClusternodeRoutemonitorBindingResourceModel {
	tflog.Debug(ctx, "In clusternode_routemonitor_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["routemonitor"]; ok && val != nil {
		data.Routemonitor = types.StringValue(val.(string))
	} else {
		data.Routemonitor = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("nodeid:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Nodeid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("routemonitor:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Routemonitor.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
