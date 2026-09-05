package azurekeyvault

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AzurekeyvaultDataSourceModel is the data-source-specific model, decoupled from
// AzurekeyvaultResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime/state attributes that the resource
// deliberately omits. Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type AzurekeyvaultDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Azureapplication types.String `tfsdk:"azureapplication"`
	Azurevaultname   types.String `tfsdk:"azurevaultname"`
	Name             types.String `tfsdk:"name"`

	// Read-only (GET-only) runtime/state attributes from the NITRO doc read-only
	// set (zion73x_readonly/azurekeyvault.json). Never settable; populated from
	// GET.
	State types.String `tfsdk:"state"`
}

func AzurekeyvaultDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"azureapplication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Azure Application object created on the ADC appliance. This object will be used for authentication with Azure Active Directory",
			},
			"azurevaultname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Key Vault configured in Azure cloud using either the Azure CLI or the Azure portal (GUI) with complete domain name. Example: Test.vault.azure.net.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Key Vault. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the Key Vault is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my keyvault\" or 'my keyvault').",
			},

			// Read-only (GET-only) runtime/state attributes surfaced by the data
			// source (these are intentionally NOT modeled on the resource). All Computed.
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "Current state of keyvault.",
			},
		},
	}
}

// azurekeyvaultDataSourceSetAttrFromGet projects a NITRO azurekeyvault GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) — no unknown->null resolution or plan preservation is
// required. The shared utils.MapGet* helpers implement that projection.
func azurekeyvaultDataSourceSetAttrFromGet(ctx context.Context, data *AzurekeyvaultDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In azurekeyvaultDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Azureapplication = utils.MapGetString(g, "azureapplication")
	data.Azurevaultname = utils.MapGetString(g, "azurevaultname")

	// Read-only runtime/state attributes.
	data.State = utils.MapGetString(g, "state")
}
