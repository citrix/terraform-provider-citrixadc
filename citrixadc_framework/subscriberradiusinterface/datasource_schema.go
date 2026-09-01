package subscriberradiusinterface

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SubscriberradiusinterfaceDataSourceModel is the data-source-specific model,
// decoupled from SubscriberradiusinterfaceResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// The Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares.
type SubscriberradiusinterfaceDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Listeningservice     types.String `tfsdk:"listeningservice"`
	Radiusinterimasstart types.String `tfsdk:"radiusinterimasstart"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/subscriberradiusinterface.json). Never settable; from GET.
	Svrstate types.String `tfsdk:"svrstate"`
}

func SubscriberradiusinterfaceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"listeningservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of RADIUS LISTENING service that will process RADIUS accounting requests.",
			},
			"radiusinterimasstart": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Treat radius interim message as start radius messages.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"svrstate": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the radius service.",
			},
		},
	}
}

// subscriberradiusinterfaceDataSourceSetAttrFromGet projects a NITRO
// subscriberradiusinterface GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled from
// the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func subscriberradiusinterfaceDataSourceSetAttrFromGet(ctx context.Context, data *SubscriberradiusinterfaceDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In subscriberradiusinterfaceDataSourceSetAttrFromGet Function")

	data.Listeningservice = utils.MapGetString(g, "listeningservice")
	data.Radiusinterimasstart = utils.MapGetString(g, "radiusinterimasstart")

	// Read-only attributes.
	data.Svrstate = utils.MapGetString(g, "svrstate")

	// Singleton resource: static ID.
	data.Id = types.StringValue("subscriberradiusinterface-config")
}
