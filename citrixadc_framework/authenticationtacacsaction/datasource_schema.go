package authenticationtacacsaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

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
		},
	}
}
