package protocolhttpband

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/protocol"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ProtocolhttpbandResourceModel describes the resource data model.
type ProtocolhttpbandResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Reqbandsize  types.Int64  `tfsdk:"reqbandsize"`
	Respbandsize types.Int64  `tfsdk:"respbandsize"`
}

func (r *ProtocolhttpbandResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the protocolhttpband resource.",
			},
			"reqbandsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Band size, in bytes, for HTTP request band statistics. For example, if you specify a band size of 100 bytes, statistics will be maintained and displayed for the following size ranges:\n0 - 99 bytes\n100 - 199 bytes\n200 - 299 bytes and so on.",
			},
			"respbandsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1024),
				Description: "Band size, in bytes, for HTTP response band statistics. For example, if you specify a band size of 100 bytes, statistics will be maintained and displayed for the following size ranges:\n0 - 99 bytes\n100 - 199 bytes\n200 - 299 bytes and so on.",
			},
		},
	}
}

func protocolhttpbandGetThePayloadFromthePlan(ctx context.Context, data *ProtocolhttpbandResourceModel) protocol.Protocolhttpband {
	tflog.Debug(ctx, "In protocolhttpbandGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// Only reqbandsize/respbandsize are settable via update/set.
	// type (show/clear filter) and nodeid (GET cluster filter) are NOT accepted by set/update (Pattern 15).
	protocolhttpband := protocol.Protocolhttpband{}
	if !data.Reqbandsize.IsNull() && !data.Reqbandsize.IsUnknown() {
		protocolhttpband.Reqbandsize = utils.IntPtr(int(data.Reqbandsize.ValueInt64()))
	}
	if !data.Respbandsize.IsNull() && !data.Respbandsize.IsUnknown() {
		protocolhttpband.Respbandsize = utils.IntPtr(int(data.Respbandsize.ValueInt64()))
	}

	return protocolhttpband
}
