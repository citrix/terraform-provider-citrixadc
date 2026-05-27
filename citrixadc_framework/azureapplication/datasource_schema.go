package azureapplication

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AzureapplicationDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Application ID that is generated when an application is created in Azure Active Directory using either the Azure CLI or the Azure portal (GUI)",
			},
			"clientsecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password for the application configured in Azure Active Directory. The password is specified in the Azure CLI or generated in the Azure portal (GUI).",
			},
			"clientsecret_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Password for the application configured in Azure Active Directory. The password is specified in the Azure CLI or generated in the Azure portal (GUI).",
			},
			"clientsecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a clientsecret_wo update.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the application. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the application is created.',\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my application\" or 'my application').",
			},
			"tenantid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the directory inside Azure Active Directory in which the application was created",
			},
			"tokenendpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL from where access token can be obtained. If the token end point is not specified, the default value is https://login.microsoftonline.com/<tenant id>.",
			},
			"vaultresource": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Vault resource for which access token is granted. Example : vault.azure.net",
			},
		},
	}
}
