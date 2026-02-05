package cluster

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ClusterResourceModel describes the resource data model.
type ClusterResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Clip     types.String `tfsdk:"clip"`
	Password types.String `tfsdk:"password"`
}

func (r *ClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cluster resource.",
			},
			"clip": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Cluster IP address to which to add the node.",
			},
			"password": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password for the nsroot account of the configuration coordinator (CCO).",
			},
		},
	}
}

func clusterGetThePayloadFromtheConfig(ctx context.Context, data *ClusterResourceModel) cluster.Cluster {
	tflog.Debug(ctx, "In clusterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cluster := cluster.Cluster{}
	if !data.Clip.IsNull() {
		cluster.Clip = data.Clip.ValueString()
	}
	if !data.Password.IsNull() {
		cluster.Password = data.Password.ValueString()
	}

	return cluster
}

func clusterSetAttrFromGet(ctx context.Context, data *ClusterResourceModel, getResponseData map[string]interface{}) *ClusterResourceModel {
	tflog.Debug(ctx, "In clusterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clip"]; ok && val != nil {
		data.Clip = types.StringValue(val.(string))
	} else {
		data.Clip = types.StringNull()
	}
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Clip.ValueString())

	return data
}
