package shutdown

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ShutdownResourceModel describes the resource data model.
//
// shutdown is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO "shutdown"
// object exposes no read/write properties and only the Shutdown action
// (POST /nitro/v1/config/shutdown). The model therefore carries only the
// synthetic id.
type ShutdownResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *ShutdownResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the shutdown resource.",
			},
		},
	}
}

// shutdownGetThePayloadFromthePlan builds the (empty) NITRO payload for the
// Shutdown action. shutdown has no read/write attributes, so the payload is an
// empty ns.Shutdown struct.
func shutdownGetThePayloadFromthePlan(ctx context.Context, data *ShutdownResourceModel) ns.Shutdown {
	tflog.Debug(ctx, "In shutdownGetThePayloadFromthePlan Function")
	shutdown := ns.Shutdown{}
	return shutdown
}
