package nsdhcpip

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsdhcpipResourceModel describes the resource data model.
//
// nsdhcpip is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO "nsdhcpip"
// object exposes no read/write properties and only the release action
// (POST /nitro/v1/config/nsdhcpip?action=release). The model therefore carries
// only the synthetic id.
type NsdhcpipResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *NsdhcpipResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsdhcpip resource.",
			},
		},
	}
}

// nsdhcpipGetThePayloadFromthePlan builds the (empty) NITRO payload for the
// release action. nsdhcpip has no read/write attributes, so the payload is an
// empty ns.Nsdhcpip struct.
func nsdhcpipGetThePayloadFromthePlan(ctx context.Context, data *NsdhcpipResourceModel) ns.Nsdhcpip {
	tflog.Debug(ctx, "In nsdhcpipGetThePayloadFromthePlan Function")
	nsdhcpip := ns.Nsdhcpip{}
	return nsdhcpip
}
