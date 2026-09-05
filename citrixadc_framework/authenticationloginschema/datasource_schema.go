package authenticationloginschema

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationloginschemaDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationloginschemaResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the read/write
// attributes (as Computed outputs) AND the read-only attributes the resource
// deliberately omits (builtin, feature). The Framework's per-attribute model
// <-> schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type AuthenticationloginschemaDataSourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Authenticationschema    types.String `tfsdk:"authenticationschema"`
	Authenticationstrength  types.Int64  `tfsdk:"authenticationstrength"`
	Name                    types.String `tfsdk:"name"`
	Passwdexpression        types.String `tfsdk:"passwdexpression"`
	Passwordcredentialindex types.Int64  `tfsdk:"passwordcredentialindex"`
	Ssocredentials          types.String `tfsdk:"ssocredentials"`
	Usercredentialindex     types.Int64  `tfsdk:"usercredentialindex"`
	Userexpression          types.String `tfsdk:"userexpression"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/authenticationloginschema.json). Never settable.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AuthenticationloginschemaDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"authenticationschema": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the file for reading authentication schema to be sent for Login Page UI. This file should contain xml definition of elements as per Citrix Forms Authentication Protocol to be able to render login form. If administrator does not want to prompt users for additional credentials but continue with previously obtained credentials, then \"noschema\" can be given as argument. Please note that this applies only to loginSchemas that are used with user-defined factors, and not the vserver factor.",
			},
			"authenticationstrength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight of the current authentication",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new login schema. Login schema defines the way login form is rendered. It provides a way to customize the fields that are shown to the user. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"passwdexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression for password extraction during login. This can be any relevant advanced policy expression.",
			},
			"passwordcredentialindex": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The index at which user entered password should be stored in session.",
			},
			"ssocredentials": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option indicates whether current factor credentials are the default SSO (SingleSignOn) credentials.",
			},
			"usercredentialindex": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The index at which user entered username should be stored in session.",
			},
			"userexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression for username extraction during login. This can be any relevant advanced policy expression.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// authenticationloginschemaDataSourceSetAttrFromGet projects a NITRO
// authenticationloginschema GET response onto the data-source model. Attributes
// are simply filled from the GET (or left Null when the GET omits them) via the
// shared utils.MapGet* helpers.
func authenticationloginschemaDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationloginschemaDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationloginschemaDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Authenticationschema = utils.MapGetString(g, "authenticationschema")
	data.Authenticationstrength = utils.MapGetInt64(g, "authenticationstrength")
	data.Passwdexpression = utils.MapGetString(g, "passwdexpression")
	data.Passwordcredentialindex = utils.MapGetInt64(g, "passwordcredentialindex")
	data.Ssocredentials = utils.MapGetString(g, "ssocredentials")
	data.Usercredentialindex = utils.MapGetInt64(g, "usercredentialindex")
	data.Userexpression = utils.MapGetString(g, "userexpression")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
