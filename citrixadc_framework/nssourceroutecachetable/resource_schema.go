package nssourceroutecachetable

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// nssourceroutecachetableResourceType is the NITRO resource-type string.
// nssourceroutecachetable is not present in the adc-nitro-go service enum, so the
// type is declared here.
const nssourceroutecachetableResourceType = "nssourceroutecachetable"

// NssourceroutecachetableResourceModel describes the resource data model.
//
// nssourceroutecachetable is a ZERO-ATTRIBUTE, ACTION-ONLY resource (flush).
// The model carries only the synthetic id.
type NssourceroutecachetableResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *NssourceroutecachetableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nssourceroutecachetable resource. It is a synthetic value (nssourceroutecachetable-config).",
			},
		},
	}
}

// nssourceroutecachetableGetThePayloadFromthePlan builds the (empty) NITRO payload
// for the flush action. nssourceroutecachetable has no read/write attributes.
func nssourceroutecachetableGetThePayloadFromthePlan(ctx context.Context, data *NssourceroutecachetableResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nssourceroutecachetableGetThePayloadFromthePlan Function")
	nssourceroutecachetable := make(map[string]interface{})
	return nssourceroutecachetable
}
