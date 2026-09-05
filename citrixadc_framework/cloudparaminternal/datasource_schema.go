package cloudparaminternal

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CloudparaminternalDataSourceModel is the data-source-specific model, decoupled
// from CloudparaminternalResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
type CloudparaminternalDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Nonftumode types.String `tfsdk:"nonftumode"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/cloudparaminternal.json). Never settable; populated from GET.
	Iamperm types.String `tfsdk:"iamperm"`
}

func CloudparaminternalDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nonftumode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates if GUI in in FTU mode or not",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"iamperm": schema.StringAttribute{
				Computed:    true,
				Description: "Indicates if user has sufficient IAM previliges.",
			},
		},
	}
}

// cloudparaminternalDataSourceSetAttrFromGet projects a NITRO cloudparaminternal
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func cloudparaminternalDataSourceSetAttrFromGet(ctx context.Context, data *CloudparaminternalDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cloudparaminternalDataSourceSetAttrFromGet Function")

	// Read/write attribute as read-back output.
	data.Nonftumode = utils.MapGetString(g, "nonftumode")

	// Read-only attribute.
	data.Iamperm = utils.MapGetString(g, "iamperm")

	// cloudparaminternal is a singleton with no lookup key; use a static ID.
	data.Id = types.StringValue("cloudparaminternal-config")
}
