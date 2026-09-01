package authenticationtacacsaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationtacacsactionDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationtacacsactionResourceModel. Every non-key
// attribute is Computed, and it additionally exposes read-only (GET-only)
// attributes the resource deliberately omits (success, failure).
type AuthenticationtacacsactionDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Accounting                 types.String `tfsdk:"accounting"`
	Attribute1                 types.String `tfsdk:"attribute1"`
	Attribute10                types.String `tfsdk:"attribute10"`
	Attribute11                types.String `tfsdk:"attribute11"`
	Attribute12                types.String `tfsdk:"attribute12"`
	Attribute13                types.String `tfsdk:"attribute13"`
	Attribute14                types.String `tfsdk:"attribute14"`
	Attribute15                types.String `tfsdk:"attribute15"`
	Attribute16                types.String `tfsdk:"attribute16"`
	Attribute2                 types.String `tfsdk:"attribute2"`
	Attribute3                 types.String `tfsdk:"attribute3"`
	Attribute4                 types.String `tfsdk:"attribute4"`
	Attribute5                 types.String `tfsdk:"attribute5"`
	Attribute6                 types.String `tfsdk:"attribute6"`
	Attribute7                 types.String `tfsdk:"attribute7"`
	Attribute8                 types.String `tfsdk:"attribute8"`
	Attribute9                 types.String `tfsdk:"attribute9"`
	Attributes                 types.String `tfsdk:"attributes"`
	Auditfailedcmds            types.String `tfsdk:"auditfailedcmds"`
	Authorization              types.String `tfsdk:"authorization"`
	Authtimeout                types.Int64  `tfsdk:"authtimeout"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Groupattrname              types.String `tfsdk:"groupattrname"`
	Name                       types.String `tfsdk:"name"`
	Serverip                   types.String `tfsdk:"serverip"`
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Tacacssecret               types.String `tfsdk:"tacacssecret"`
	TacacssecretWo             types.String `tfsdk:"tacacssecret_wo"`
	TacacssecretWoVersion      types.Int64  `tfsdk:"tacacssecret_wo_version"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/authenticationtacacsaction.json). Never settable;
	// populated from GET.
	Success types.Int64 `tfsdk:"success"`
	Failure types.Int64 `tfsdk:"failure"`
}

func AuthenticationtacacsactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the TACACS+ server is currently accepting accounting messages.",
			},
			"attribute1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '1' (where '1' changes for each attribute)",
			},
			"attribute10": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '10' (where '10' changes for each attribute)",
			},
			"attribute11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '11' (where '11' changes for each attribute)",
			},
			"attribute12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '12' (where '12' changes for each attribute)",
			},
			"attribute13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '13' (where '13' changes for each attribute)",
			},
			"attribute14": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '14' (where '14' changes for each attribute)",
			},
			"attribute15": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '15' (where '15' changes for each attribute)",
			},
			"attribute16": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '16' (where '16' changes for each attribute)",
			},
			"attribute2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '2' (where '2' changes for each attribute)",
			},
			"attribute3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '3' (where '3' changes for each attribute)",
			},
			"attribute4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '4' (where '4' changes for each attribute)",
			},
			"attribute5": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '5' (where '5' changes for each attribute)",
			},
			"attribute6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '6' (where '6' changes for each attribute)",
			},
			"attribute7": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '7' (where '7' changes for each attribute)",
			},
			"attribute8": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '8' (where '8' changes for each attribute)",
			},
			"attribute9": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the custom attribute to be extracted from server and stored at index '9' (where '9' changes for each attribute)",
			},
			"attributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of attribute names separated by ',' which needs to be fetched from tacacs server.\nNote that preceeding and trailing spaces will be removed.\nAttribute name can be 127 bytes and total length of this string should not cross 2047 bytes.\nThese attributes have multi-value support separated by ',' and stored as key-value pair in AAA session",
			},
			"auditfailedcmds": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the TACACS+ server that will receive accounting messages.",
			},
			"authorization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use streaming authorization on the TACACS+ server.",
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of seconds the Citrix ADC waits for a response from the TACACS+ server.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupattrname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TACACS+ group attribute name.\nUsed for group extraction on the TACACS+ server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the TACACS+ profile (action).\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after TACACS profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'y authentication action').",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address assigned to the TACACS+ server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number on which the TACACS+ server listens for connections.",
			},
			"tacacssecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Key shared between the TACACS+ server and the Citrix ADC.\nRequired for allowing the Citrix ADC to communicate with the TACACS+ server.",
			},
			"tacacssecret_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Key shared between the TACACS+ server and the Citrix ADC.\nRequired for allowing the Citrix ADC to communicate with the TACACS+ server.",
			},
			"tacacssecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a tacacssecret_wo update.",
			},
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

// authenticationtacacsactionDataSourceSetAttrFromGet projects a NITRO
// authenticationtacacsaction GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func authenticationtacacsactionDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationtacacsactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationtacacsactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Accounting = utils.MapGetString(g, "accounting")
	data.Attribute1 = utils.MapGetString(g, "attribute1")
	data.Attribute10 = utils.MapGetString(g, "attribute10")
	data.Attribute11 = utils.MapGetString(g, "attribute11")
	data.Attribute12 = utils.MapGetString(g, "attribute12")
	data.Attribute13 = utils.MapGetString(g, "attribute13")
	data.Attribute14 = utils.MapGetString(g, "attribute14")
	data.Attribute15 = utils.MapGetString(g, "attribute15")
	data.Attribute16 = utils.MapGetString(g, "attribute16")
	data.Attribute2 = utils.MapGetString(g, "attribute2")
	data.Attribute3 = utils.MapGetString(g, "attribute3")
	data.Attribute4 = utils.MapGetString(g, "attribute4")
	data.Attribute5 = utils.MapGetString(g, "attribute5")
	data.Attribute6 = utils.MapGetString(g, "attribute6")
	data.Attribute7 = utils.MapGetString(g, "attribute7")
	data.Attribute8 = utils.MapGetString(g, "attribute8")
	data.Attribute9 = utils.MapGetString(g, "attribute9")
	data.Attributes = utils.MapGetString(g, "attributes")
	data.Auditfailedcmds = utils.MapGetString(g, "auditfailedcmds")
	data.Authorization = utils.MapGetString(g, "authorization")
	data.Authtimeout = utils.MapGetInt64(g, "authtimeout")
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Groupattrname = utils.MapGetString(g, "groupattrname")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Tacacssecret = utils.MapGetString(g, "tacacssecret")

	// tacacssecret_wo (write-only) and its TF-only version tracker are never
	// returned by GET -> Null.
	data.TacacssecretWo = types.StringNull()
	data.TacacssecretWoVersion = types.Int64Null()

	// Read-only attributes.
	data.Success = utils.MapGetInt64(g, "success")
	data.Failure = utils.MapGetInt64(g, "failure")
}
