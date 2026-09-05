package azureapplication

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Password for the application configured in Azure Active Directory. The password is specified in the Azure CLI or generated in the Azure portal (GUI).",
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

type AzureapplicationDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Clientid      types.String `tfsdk:"clientid"`
	Clientsecret  types.String `tfsdk:"clientsecret"`
	Name          types.String `tfsdk:"name"`
	Tenantid      types.String `tfsdk:"tenantid"`
	Tokenendpoint types.String `tfsdk:"tokenendpoint"`
	Vaultresource types.String `tfsdk:"vaultresource"`
}

func azureapplicationDataSourceSetAttrFromGet(ctx context.Context, data *AzureapplicationDataSourceModel, getResponseData map[string]interface{}) *AzureapplicationDataSourceModel {
	tflog.Debug(ctx, "In azureapplicationDataSourceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientid"]; ok && val != nil {
		data.Clientid = types.StringValue(val.(string))
	} else {
		data.Clientid = types.StringNull()
	}
	// clientsecret is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["tenantid"]; ok && val != nil {
		data.Tenantid = types.StringValue(val.(string))
	} else {
		data.Tenantid = types.StringNull()
	}
	if val, ok := getResponseData["tokenendpoint"]; ok && val != nil {
		data.Tokenendpoint = types.StringValue(val.(string))
	} else {
		data.Tokenendpoint = types.StringNull()
	}
	if val, ok := getResponseData["vaultresource"]; ok && val != nil {
		data.Vaultresource = types.StringValue(val.(string))
	} else {
		data.Vaultresource = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
