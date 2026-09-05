package nsmemrecovery

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsmemrecoveryStartResourceModel describes the resource data model.
//
// nsmemrecovery is an ACTION-ONLY NITRO object: it exposes only the "start"
// action (POST /nitro/v1/config/nsmemrecovery?action=start) which recovers a
// configurable percentage of memory from the freepools. There is NO
// add/set/get/delete/unset endpoint, so the resource cannot be read back and has
// no datasource. It is therefore modelled with the action-suffixed TF name
// `citrixadc_nsmemrecovery_start` (mirroring citrixadc_appfwarchive_export) to
// make the action-only nature explicit. The underlying NITRO type is still
// "nsmemrecovery" (service.Nsmemrecovery.Type()).
type NsmemrecoveryStartResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Percentage types.Int64  `tfsdk:"percentage"`
}

func (r *NsmemrecoveryStartResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsmemrecovery_start resource.",
			},
			"percentage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Percentage of memory to be recovered from freepools. Default value: 10. Minimum value = 5. Maximum value = 90.",
			},
		},
	}
}

// nsmemrecoveryStartGetThePayloadFromthePlan builds the NITRO payload for the
// nsmemrecovery "start" action from the plan model.
func nsmemrecoveryStartGetThePayloadFromthePlan(ctx context.Context, data *NsmemrecoveryStartResourceModel) ns.Nsmemrecovery {
	tflog.Debug(ctx, "In nsmemrecoveryStartGetThePayloadFromthePlan Function")

	nsmemrecovery := ns.Nsmemrecovery{}
	if !data.Percentage.IsNull() && !data.Percentage.IsUnknown() {
		val := int(data.Percentage.ValueInt64())
		nsmemrecovery.Percentage = &val
	}
	return nsmemrecovery
}
