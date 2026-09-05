package ssldynamicclientcertcache

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SsldynamicclientcertcacheFlushResourceModel describes the resource data model.
//
// ssldynamicclientcertcache is a ZERO-ATTRIBUTE, ACTION-ONLY NITRO object: it
// exposes ONLY the POST /config/ssldynamicclientcertcache?action=flush operation
// (see nitro-rest-73x/ssl/ssldynamicclientcertcache.html and the empty
// service/config/ssl.Ssldynamicclientcertcache{} struct). There is no
// add/set/update/delete/unset endpoint and no GET endpoint. Because the object
// is action-only, the TF resource is named with the action suffix
// `citrixadc_ssldynamicclientcertcache_flush` (mirroring
// citrixadc_appfwarchive_export / citrixadc_nsmemrecovery_start). The underlying
// NITRO type is still "ssldynamicclientcertcache"
// (service.Ssldynamicclientcertcache.Type()). The model carries only the
// synthetic id.
type SsldynamicclientcertcacheFlushResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SsldynamicclientcertcacheFlushResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ssldynamicclientcertcache_flush resource. It is a synthetic value (ssldynamicclientcertcache).",
			},
		},
	}
}

// ssldynamicclientcertcacheFlushGetThePayloadFromthePlan builds the (empty)
// NITRO payload for the flush action. flush has no read/write attributes.
func ssldynamicclientcertcacheFlushGetThePayloadFromthePlan(ctx context.Context, data *SsldynamicclientcertcacheFlushResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In ssldynamicclientcertcacheFlushGetThePayloadFromthePlan Function")
	ssldynamicclientcertcache := make(map[string]interface{})
	return ssldynamicclientcertcache
}
