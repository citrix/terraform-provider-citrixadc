package nsextension

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsextensionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the extension object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the extension object on the Citrix ADC.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported extension.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
			"trace": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables tracing to the NS log file of extension execution:\n   off   - turns off tracing (equivalent to unset ns extension <extension-name> -trace)\n   calls - traces extension function calls with arguments and function returns with the first return value\n   lines - traces the above plus line numbers for executed extension lines\n   all   - traces the above plus local variables changed by executed extension lines\nNote that the DEBUG log level must be enabled to see extension tracing.\nThis can be done by set audit syslogParams -loglevel ALL or -loglevel DEBUG.",
			},
			"tracefunctions": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of extension functions to trace. By default, all extension functions are traced.",
			},
			"tracevariables": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of variables (in traced extension functions) to trace. By default, all variables are traced.",
			},
		},
	}
}
