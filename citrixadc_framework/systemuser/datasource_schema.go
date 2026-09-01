package systemuser

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemuserDataSourceModel is the data-source-specific model, decoupled from
// SystemuserResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the configurable attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed.
type SystemuserDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Allowedmanagementinterface types.List   `tfsdk:"allowedmanagementinterface"`
	Externalauth               types.String `tfsdk:"externalauth"`
	Hashedpassword             types.String `tfsdk:"hashedpassword"`
	Logging                    types.String `tfsdk:"logging"`
	Maxsession                 types.Int64  `tfsdk:"maxsession"`
	Password                   types.String `tfsdk:"password"`
	PasswordWo                 types.String `tfsdk:"password_wo"`
	PasswordWoVersion          types.Int64  `tfsdk:"password_wo_version"`
	Promptstring               types.String `tfsdk:"promptstring"`
	Timeout                    types.Int64  `tfsdk:"timeout"`
	Username                   types.String `tfsdk:"username"` // Required lookup key
	Cmdpolicybinding           types.Set    `tfsdk:"cmdpolicybinding"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/systemuser.json). Never settable; populated from GET.
	Encrypted                      types.Bool   `tfsdk:"encrypted"`
	Hashmethod                     types.String `tfsdk:"hashmethod"`
	Promptinheritedfrom            types.String `tfsdk:"promptinheritedfrom"`
	Timeoutkind                    types.String `tfsdk:"timeoutkind"`
	Allowedmanagementinterfacekind types.String `tfsdk:"allowedmanagementinterfacekind"`
	Lastpwdchangetimestamp         types.Int64  `tfsdk:"lastpwdchangetimestamp"`
	Daystoexpirekind               types.String `tfsdk:"daystoexpirekind"`
}

func SystemuserDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allowedmanagementinterface": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Allowed Management interfaces to the system user. By default user is allowed from both API and CLI interfaces. If management interface for a user is set to API, then user is not allowed to access NS through CLI. GUI interface will come under API interface",
			},
			"externalauth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether to use external authentication servers for the system user authentication or not",
			},
			"hashedpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hashed password for the system user, as returned by the NITRO API.",
			},
			"logging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Users logging privilege",
			},
			"maxsession": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of client connection allowed per user",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password for the system user. Can include any ASCII character.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Password for the system user. Can include any ASCII character.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a password_wo update.",
			},
			"promptstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:\n* %u - Will be replaced by the user name.\n* %h - Will be replaced by the hostname of the Citrix ADC.\n* %t - Will be replaced by the current time in 12-hour format.\n* %T - Will be replaced by the current time in 24-hour format.\n* %d - Will be replaced by the current date.\n* %s - Will be replaced by the state of the Citrix ADC.\n\nNote: The 63-character limit for the length of the string does not apply to the characters that replace the variables.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds. If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name for a user. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the user is added.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my user\" or 'my user').",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"encrypted": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the password stored on the appliance is encrypted.",
			},
			"hashmethod": schema.StringAttribute{
				Computed:    true,
				Description: "Hash method used for the system user password (SHA1, SHA512, PBKDF2).",
			},
			"promptinheritedfrom": schema.StringAttribute{
				Computed:    true,
				Description: "From where the prompt has been inherited (User, Group, Global, Climode).",
			},
			"timeoutkind": schema.StringAttribute{
				Computed:    true,
				Description: "From where the timeout has been inherited (User, Group, Global, Climode).",
			},
			"allowedmanagementinterfacekind": schema.StringAttribute{
				Computed:    true,
				Description: "Value of allowed interface which can be inherited from Global, Group or User (User, Group, Global, Climode).",
			},
			"lastpwdchangetimestamp": schema.Int64Attribute{
				Computed:    true,
				Description: "Timestamp for the last password change for the system user.",
			},
			"daystoexpirekind": schema.StringAttribute{
				Computed:    true,
				Description: "From where the daystoexpire value has been inherited (User, Group, Global, Climode).",
			},
		},
		Blocks: map[string]schema.Block{
			"cmdpolicybinding": schema.SetNestedBlock{
				Description: "Inline command policy bindings for the system user.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"policyname": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The name of command policy.",
						},
						"priority": schema.Int64Attribute{
							Optional:    true,
							Computed:    true,
							Description: "The priority of the policy.",
						},
					},
				},
			},
		},
	}
}

// systemuserDataSourceSetAttrFromGet projects a NITRO systemuser GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits them).
// The shared utils.MapGet* helpers implement that projection.
func systemuserDataSourceSetAttrFromGet(ctx context.Context, data *SystemuserDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In systemuserDataSourceSetAttrFromGet Function")

	if v, ok := g["username"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Username = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Allowedmanagementinterface = utils.MapGetStringList(g, "allowedmanagementinterface")
	data.Externalauth = utils.MapGetString(g, "externalauth")
	data.Logging = utils.MapGetString(g, "logging")
	data.Maxsession = utils.MapGetInt64(g, "maxsession")
	data.Promptstring = utils.MapGetString(g, "promptstring")
	data.Timeout = utils.MapGetInt64(g, "timeout")

	// The NITRO API returns the hashed password in the "password" response field;
	// surface it as hashedpassword. The plaintext password / write-only inputs are
	// never returned by GET -> Null.
	data.Hashedpassword = utils.MapGetString(g, "password")
	data.Password = types.StringNull()
	data.PasswordWo = types.StringNull()
	data.PasswordWoVersion = types.Int64Null()

	// Bindings are fetched via a separate NITRO endpoint; the base GET does not
	// return them -> typed Null.
	data.Cmdpolicybinding = types.SetNull(types.ObjectType{AttrTypes: map[string]attr.Type{
		"policyname": types.StringType,
		"priority":   types.Int64Type,
	}})

	// Read-only (GET-only) attributes.
	data.Encrypted = utils.MapGetBool(g, "encrypted")
	data.Hashmethod = utils.MapGetString(g, "hashmethod")
	data.Promptinheritedfrom = utils.MapGetString(g, "promptinheritedfrom")
	data.Timeoutkind = utils.MapGetString(g, "timeoutkind")
	data.Allowedmanagementinterfacekind = utils.MapGetString(g, "allowedmanagementinterfacekind")
	data.Lastpwdchangetimestamp = utils.MapGetInt64(g, "lastpwdchangetimestamp")
	data.Daystoexpirekind = utils.MapGetString(g, "daystoexpirekind")
}
