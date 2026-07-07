package nstestlicense

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// nstestlicenseResourceType is the NITRO resource-type string. nstestlicense is
// not present in the adc-nitro-go service enum, so the type is declared here.
const nstestlicenseResourceType = "nstestlicense"

// NstestlicenseResourceModel describes the resource data model.
//
// nstestlicense is a ZERO-ATTRIBUTE, ACTION-ONLY resource (apply). The model
// carries only the synthetic id.
type NstestlicenseResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *NstestlicenseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstestlicense resource. It is a synthetic value (nstestlicense-config).",
			},
		},
	}
}

// nstestlicenseGetThePayloadFromthePlan builds the (empty) NITRO payload for the
// apply action. nstestlicense has no read/write attributes.
func nstestlicenseGetThePayloadFromthePlan(ctx context.Context, data *NstestlicenseResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nstestlicenseGetThePayloadFromthePlan Function")
	nstestlicense := make(map[string]interface{})
	return nstestlicense
}
