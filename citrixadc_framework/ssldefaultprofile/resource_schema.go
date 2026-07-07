package ssldefaultprofile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ssldefaultprofileResourceType is the NITRO resource-type string. There is no
// service.Ssldefaultprofile enum in the vendored adc-nitro-go, so the raw type
// string is used with a map payload. NO vendor/ edits.
const ssldefaultprofileResourceType = "ssldefaultprofile"

// SsldefaultprofileResourceModel describes the resource data model.
//
// ssldefaultprofile is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// "ssldefaultprofile" object exposes no read/write properties and only the
// convert action (POST /nitro/v1/config/ssldefaultprofile?action=convert). The
// model therefore carries only the synthetic id.
type SsldefaultprofileResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SsldefaultprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ssldefaultprofile resource.",
			},
		},
	}
}

// ssldefaultprofileGetThePayloadFromthePlan builds the (empty) NITRO payload for
// the convert action. ssldefaultprofile has no read/write attributes, so the
// payload is an empty map.
func ssldefaultprofileGetThePayloadFromthePlan(ctx context.Context, data *SsldefaultprofileResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In ssldefaultprofileGetThePayloadFromthePlan Function")
	ssldefaultprofile := make(map[string]interface{})
	return ssldefaultprofile
}
