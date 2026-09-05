package authenticationdfaaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationdfaactionDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationdfaactionResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the read/write
// attributes (as Computed outputs) AND the read-only attributes the resource
// deliberately omits (success, failure). The Framework's per-attribute model
// <-> schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type AuthenticationdfaactionDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Clientid                   types.String `tfsdk:"clientid"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Name                       types.String `tfsdk:"name"`
	Passphrase                 types.String `tfsdk:"passphrase"`
	Serverurl                  types.String `tfsdk:"serverurl"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/authenticationdfaaction.json). Never settable.
	Success types.Int64 `tfsdk:"success"`
	Failure types.Int64 `tfsdk:"failure"`
}

func AuthenticationdfaactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If configured, this string is sent to the DFA server as the X-Citrix-Exchange header value.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DFA action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the DFA action is added.",
			},
			"passphrase": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Key shared between the DFA server and the Citrix ADC.\nRequired to allow the Citrix ADC to communicate with the DFA server.",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DFA Server URL",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"success": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of successful authentications through this DFA action.",
			},
			"failure": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of failed authentications through this DFA action.",
			},
		},
	}
}

// authenticationdfaactionDataSourceSetAttrFromGet projects a NITRO
// authenticationdfaaction GET response onto the data-source model. Attributes
// are simply filled from the GET (or left Null when the GET omits them) via the
// shared utils.MapGet* helpers.
func authenticationdfaactionDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationdfaactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationdfaactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Clientid = utils.MapGetString(g, "clientid")
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Serverurl = utils.MapGetString(g, "serverurl")

	// passphrase is a secret input the GET never returns -> Null.
	data.Passphrase = types.StringNull()

	// Read-only metadata.
	data.Success = utils.MapGetInt64(g, "success")
	data.Failure = utils.MapGetInt64(g, "failure")
}
