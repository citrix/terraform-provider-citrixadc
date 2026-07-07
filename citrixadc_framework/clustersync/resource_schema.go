package clustersync

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ClustersyncResourceModel describes the resource data model.
//
// clustersync is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO "clustersync"
// object exposes no read/write properties and only the Force action
// (POST /nitro/v1/config/clustersync?action=Force). The model therefore carries
// only the synthetic id.
type ClustersyncResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *ClustersyncResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clustersync resource.",
			},
		},
	}
}

// clustersyncGetThePayloadFromthePlan builds the (empty) NITRO payload for the
// Force action. clustersync has no read/write attributes, so the payload is an
// empty cluster.Clustersync struct.
func clustersyncGetThePayloadFromthePlan(ctx context.Context, data *ClustersyncResourceModel) cluster.Clustersync {
	tflog.Debug(ctx, "In clustersyncGetThePayloadFromthePlan Function")
	clustersync := cluster.Clustersync{}
	return clustersync
}
