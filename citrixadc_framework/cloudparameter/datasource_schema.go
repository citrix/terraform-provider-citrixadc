package cloudparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CloudparameterDataSourceModel is the data-source-specific model, decoupled
// from CloudparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type CloudparameterDataSourceModel struct {
	Id types.String `tfsdk:"id"`

	// Existing read/write attributes, surfaced here as Computed outputs.
	Activationcode     types.String `tfsdk:"activationcode"`
	Connectorresidence types.String `tfsdk:"connectorresidence"`
	Controllerfqdn     types.String `tfsdk:"controllerfqdn"`
	Controllerport     types.Int64  `tfsdk:"controllerport"`
	Customerid         types.String `tfsdk:"customerid"`
	Deployment         types.String `tfsdk:"deployment"`
	Instanceid         types.String `tfsdk:"instanceid"`
	Resourcelocation   types.String `tfsdk:"resourcelocation"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/cloudparameter.json). Never settable; populated from GET.
	Controlconnectionstatus types.String `tfsdk:"controlconnectionstatus"`
}

func CloudparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to read global cloud parameter configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"activationcode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Activation code for the NGS Connector instance",
			},
			"connectorresidence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Identifies whether the connector is located Onprem, Aws or Azure",
			},
			"controllerfqdn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FQDN of the controller to which the Citrix ADC SDProxy Connects",
			},
			"controllerport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number of the controller to which the Citrix ADC SDProxy connects",
			},
			"customerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Customer ID of the citrix cloud customer",
			},
			"deployment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Describes if the customer is a Staging/Production or Dev Citrix Cloud customer",
			},
			"instanceid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Instance ID of the customer provided by Trust",
			},
			"resourcelocation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Resource Location of the customer provided by Trust",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (this is intentionally NOT modeled on the resource). All Computed.
			"controlconnectionstatus": schema.StringAttribute{
				Computed:    true,
				Description: "Status of the control connection: In case of successful registration to the controller, connection status will be shown as \"Registered\" else \"Unregistered\".",
			},
		},
	}
}

// cloudparameterDataSourceSetAttrFromGet projects a NITRO cloudparameter GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection. cloudparameter is a singleton, so the ID is a static string.
func cloudparameterDataSourceSetAttrFromGet(ctx context.Context, data *CloudparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cloudparameterDataSourceSetAttrFromGet Function")

	// cloudparameter is a singleton (no unique lookup key) -> static ID.
	data.Id = types.StringValue("cloudparameter-config")

	// activationcode is write-only (never returned by GET) -> leave it null.
	data.Activationcode = types.StringNull()

	// Existing read/write attributes as read-back outputs.
	data.Connectorresidence = utils.MapGetString(g, "connectorresidence")
	data.Controllerfqdn = utils.MapGetString(g, "controllerfqdn")
	data.Controllerport = utils.MapGetInt64(g, "controllerport")
	data.Customerid = utils.MapGetString(g, "customerid")
	data.Deployment = utils.MapGetString(g, "deployment")
	data.Instanceid = utils.MapGetString(g, "instanceid")
	data.Resourcelocation = utils.MapGetString(g, "resourcelocation")

	// Read-only metadata.
	data.Controlconnectionstatus = utils.MapGetString(g, "controlconnectionstatus")
}
