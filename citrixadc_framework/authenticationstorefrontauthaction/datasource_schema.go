package authenticationstorefrontauthaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationstorefrontauthactionDataSourceModel is the data-source-specific
// model, decoupled from AuthenticationstorefrontauthactionResourceModel. Every
// non-key attribute is Computed, and it additionally exposes read-only
// (GET-only) attributes the resource deliberately omits (success, failure).
type AuthenticationstorefrontauthactionDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Domain                     types.String `tfsdk:"domain"`
	Name                       types.String `tfsdk:"name"`
	Serverurl                  types.String `tfsdk:"serverurl"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/authenticationstorefrontauthaction.json). Never
	// settable; populated from GET.
	Success types.Int64 `tfsdk:"success"`
	Failure types.Int64 `tfsdk:"failure"`
}

func AuthenticationstorefrontauthactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain of the server that is used for authentication. If users enter name without domain, this parameter is added to username in the authentication request to server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Storefront Authentication action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the Storefront server. This is the FQDN of the Storefront server. example: https://storefront.com/.  Authentication endpoints are learned dynamically by Gateway.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"success": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of successful authentications.",
			},
			"failure": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of failed authentications.",
			},
		},
	}
}

// authenticationstorefrontauthactionDataSourceSetAttrFromGet projects a NITRO
// authenticationstorefrontauthaction GET response onto the data-source model.
// Because a data source has no plan/apply reconciliation, attributes are simply
// filled from the GET (or left Null when the GET omits them). The shared
// utils.MapGet* helpers implement that projection.
func authenticationstorefrontauthactionDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationstorefrontauthactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationstorefrontauthactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Domain = utils.MapGetString(g, "domain")
	data.Serverurl = utils.MapGetString(g, "serverurl")

	// Read-only attributes.
	data.Success = utils.MapGetInt64(g, "success")
	data.Failure = utils.MapGetInt64(g, "failure")
}
