package authenticationwebauthaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationwebauthactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"attribute1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute1 from the webauth response",
			},
			"attribute10": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute10 from the webauth response",
			},
			"attribute11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute11 from the webauth response",
			},
			"attribute12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute12 from the webauth response",
			},
			"attribute13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute13 from the webauth response",
			},
			"attribute14": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute14 from the webauth response",
			},
			"attribute15": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute15 from the webauth response",
			},
			"attribute16": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute16 from the webauth response",
			},
			"attribute2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute2 from the webauth response",
			},
			"attribute3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute3 from the webauth response",
			},
			"attribute4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute4 from the webauth response",
			},
			"attribute5": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute5 from the webauth response",
			},
			"attribute6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute6 from the webauth response",
			},
			"attribute7": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute7 from the webauth response",
			},
			"attribute8": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute8 from the webauth response",
			},
			"attribute9": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute9 from the webauth response",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"fullreqexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exact HTTP request, in the form of an expression, which the Citrix ADC sends to the authentication server.\nThe Citrix ADC does not check the validity of this request. One must manually validate the request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Web Authentication action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"scheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of scheme for the web server.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the web server to be used for authentication.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the web server accepts connections.",
			},
			"successrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, that checks to see if authentication is successful.",
			},
		},
	}
}
