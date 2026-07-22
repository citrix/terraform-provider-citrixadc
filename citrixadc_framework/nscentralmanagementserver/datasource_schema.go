package nscentralmanagementserver

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NscentralmanagementserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"activationcode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Activation code is used to register to ADM service",
			},
			"adcpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ADC password used to create device profile on ADM",
			},
			"adcpassword_wo": schema.StringAttribute{
				Optional:    true,
				Description: "ADC password used to create device profile on ADM",
			},
			"adcpassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a adcpassword_wo update.",
			},
			"adcusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ADC username used to create device profile on ADM",
			},
			"deviceprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Device profile is created on ADM and contains the user name and password of the instance(s).",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ip Address of central management server.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password for access to central management server. Required for any user account.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Password for access to central management server. Required for any user account.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a password_wo update.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name of the central management server or service-url to locate ADM service.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Type of the central management server. Must be either CLOUD or ONPREM depending on whether the server is on the cloud or on premise.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username for access to central management server. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or\nsingle quotation marks (for example, \"my ns centralmgmtserver\" or \"my ns centralmgmtserver\").",
			},
			"validatecert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "validate the server certificate for secure SSL connections.",
			},
		},
	}
}
