package clusternodegroup_clusternode_binding

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

// ClusternodegroupClusternodeBindingResourceModel describes the resource data model.
type ClusternodegroupClusternodeBindingResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Node types.Int64  `tfsdk:"node"`
}

func (r *ClusternodegroupClusternodeBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusternodegroup_clusternode_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
			"node": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Nodes in the nodegroup",
			},
		},
	}
}

func clusternodegroup_clusternode_bindingGetThePayloadFromtheConfig(ctx context.Context, data *ClusternodegroupClusternodeBindingResourceModel) cluster.Clusternodegroupclusternodebinding {
	tflog.Debug(ctx, "In clusternodegroup_clusternode_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	clusternodegroup_clusternode_binding := cluster.Clusternodegroupclusternodebinding{}
	if !data.Name.IsNull() {
		clusternodegroup_clusternode_binding.Name = data.Name.ValueString()
	}
	if !data.Node.IsNull() {
		clusternodegroup_clusternode_binding.Node = utils.IntPtr(int(data.Node.ValueInt64()))
	}

	return clusternodegroup_clusternode_binding
}

func clusternodegroup_clusternode_bindingSetAttrFromGet(ctx context.Context, data *ClusternodegroupClusternodeBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupClusternodeBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_clusternode_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["node"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Node = types.Int64Value(intVal)
		}
	} else {
		data.Node = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("node:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Node.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
