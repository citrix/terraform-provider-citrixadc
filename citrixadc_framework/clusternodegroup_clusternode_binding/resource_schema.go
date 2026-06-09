package clusternodegroup_clusternode_binding

import (
	"context"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
			"node": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Nodes in the nodegroup",
			},
		},
	}
}

func clusternodegroup_clusternode_bindingGetThePayloadFromthePlan(ctx context.Context, data *ClusternodegroupClusternodeBindingResourceModel) cluster.Clusternodegroupclusternodebinding {
	tflog.Debug(ctx, "In clusternodegroup_clusternode_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	clusternodegroup_clusternode_binding := cluster.Clusternodegroupclusternodebinding{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		clusternodegroup_clusternode_binding.Name = data.Name.ValueString()
	}
	if !data.Node.IsNull() && !data.Node.IsUnknown() {
		clusternodegroup_clusternode_binding.Node = utils.IntPtr(int(data.Node.ValueInt64()))
	}

	return clusternodegroup_clusternode_binding
}

// clusternodegroup_clusternode_bindingComposeId builds the composite ID for the
// binding. The string key `name` is URL-encoded; the integer key `node` is
// formatted with strconv (integers need no URL-encoding).
func clusternodegroup_clusternode_bindingComposeId(name string, node int64) string {
	idParts := []string{}
	idParts = append(idParts, "name:"+utils.UrlEncode(name))
	idParts = append(idParts, "node:"+strconv.FormatInt(node, 10))
	return strings.Join(idParts, ",")
}

// clusternodegroup_clusternode_bindingSetAttrFromGet is used by the resource Read
// path. The binding GET response only echoes back the two key attributes, both of
// which are RequiresReplace, so we preserve the existing plan/state values and the
// already-set ID rather than recomputing them.
func clusternodegroup_clusternode_bindingSetAttrFromGet(ctx context.Context, data *ClusternodegroupClusternodeBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupClusternodeBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_clusternode_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["node"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Node = types.Int64Value(intVal)
		}
	}

	return data
}

// clusternodegroup_clusternode_bindingSetAttrFromGetForDatasource is used by the
// datasource Read path. The datasource has no prior plan/state, so it faithfully
// copies every field from the GET response and composes the ID itself.
func clusternodegroup_clusternode_bindingSetAttrFromGetForDatasource(ctx context.Context, data *ClusternodegroupClusternodeBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupClusternodeBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_clusternode_bindingSetAttrFromGetForDatasource Function")

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

	// Set ID for the datasource (no Create runs for a datasource).
	data.Id = types.StringValue(clusternodegroup_clusternode_bindingComposeId(data.Name.ValueString(), data.Node.ValueInt64()))

	return data
}
