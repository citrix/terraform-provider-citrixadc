package authenticationadfsproxyprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationadfsproxyprofileDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationadfsproxyprofileResourceModel. A data source is a
// pure read surface, so it can expose the FULL GET projection: the read/write
// attributes (as Computed outputs) AND the read-only metadata attributes the
// resource deliberately omits.
type AuthenticationadfsproxyprofileDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Certkeyname       types.String `tfsdk:"certkeyname"`
	Name              types.String `tfsdk:"name"` // Required lookup key
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Serverurl         types.String `tfsdk:"serverurl"`
	Username          types.String `tfsdk:"username"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/authenticationadfsproxyprofile.json). Populated from GET.
	Adfstruststatus types.String `tfsdk:"adfstruststatus"`
}

func AuthenticationadfsproxyprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"certkeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL certificate of the proxy that is registered at adfs server for trust.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the adfs proxy profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my push service\" or 'my push service').",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "This is the password of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "This is the password of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a password_wo update.",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified url of the adfs server.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the name of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (this is intentionally NOT modeled on the resource). Computed.
			"adfstruststatus": schema.StringAttribute{
				Computed:    true,
				Description: "Describes status of ADFS trust. Possible values = INIT, FAILED, ESTABLISHED, ESTABLISHED/CONFIGURED, ESTABLISHED_RENEW_SUCCESS, ESTABLISHED_RENEW_FAILED, RENEWED/CONFIGURED.",
			},
		},
	}
}

// authenticationadfsproxyprofileDataSourceSetAttrFromGet projects a NITRO
// authenticationadfsproxyprofile GET response onto the data-source model.
// Attributes are filled from the GET (or left Null when the GET omits them) using
// the shared utils.MapGet* helpers.
func authenticationadfsproxyprofileDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationadfsproxyprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationadfsproxyprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Certkeyname = utils.MapGetString(g, "certkeyname")
	data.Serverurl = utils.MapGetString(g, "serverurl")
	data.Username = utils.MapGetString(g, "username")

	// password / password_wo(+version) are write-only/secret inputs the GET never
	// returns -> Null.
	data.Password = types.StringNull()
	data.PasswordWo = types.StringNull()
	data.PasswordWoVersion = types.Int64Null()

	// Read-only metadata.
	data.Adfstruststatus = utils.MapGetString(g, "adfstruststatus")
}
