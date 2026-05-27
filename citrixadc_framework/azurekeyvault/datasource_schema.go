package azurekeyvault

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AzurekeyvaultDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"azureapplication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Azure Application object created on the ADC appliance. This object will be used for authentication with Azure Active Directory",
			},
			"azurevaultname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Key Vault configured in Azure cloud using either the Azure CLI or the Azure portal (GUI) with complete domain name. Example: Test.vault.azure.net.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Key Vault. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the Key Vault is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my keyvault\" or 'my keyvault').",
			},
		},
	}
}
