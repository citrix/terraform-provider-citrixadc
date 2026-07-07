package locationdata

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LocationdataResourceModel describes the resource data model.
//
// locationdata is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// "locationdata" object exposes no read/write properties and only the clear
// action (POST /nitro/v1/config/locationdata?action=clear). The model therefore
// carries only the synthetic id.
type LocationdataResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *LocationdataResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the locationdata resource.",
			},
		},
	}
}

// locationdataGetThePayloadFromthePlan builds the (empty) NITRO payload for the
// clear action. locationdata has no read/write attributes, so the payload is an
// empty basic.Locationdata struct.
func locationdataGetThePayloadFromthePlan(ctx context.Context, data *LocationdataResourceModel) basic.Locationdata {
	tflog.Debug(ctx, "In locationdataGetThePayloadFromthePlan Function")
	locationdata := basic.Locationdata{}
	return locationdata
}
