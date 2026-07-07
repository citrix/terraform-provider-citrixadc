package dbsmonitors

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DbsmonitorsResourceModel describes the resource data model.
//
// dbsmonitors is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO "dbsmonitors"
// object exposes no read/write properties and only the restart action
// (POST /nitro/v1/config/dbsmonitors?action=restart). The model therefore
// carries only the synthetic id.
type DbsmonitorsResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *DbsmonitorsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dbsmonitors resource.",
			},
		},
	}
}

// dbsmonitorsGetThePayloadFromthePlan builds the (empty) NITRO payload for the
// restart action. dbsmonitors has no read/write attributes, so the payload is an
// empty basic.Dbsmonitors struct.
func dbsmonitorsGetThePayloadFromthePlan(ctx context.Context, data *DbsmonitorsResourceModel) basic.Dbsmonitors {
	tflog.Debug(ctx, "In dbsmonitorsGetThePayloadFromthePlan Function")
	dbsmonitors := basic.Dbsmonitors{}
	return dbsmonitors
}
