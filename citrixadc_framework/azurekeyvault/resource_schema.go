package azurekeyvault

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/azure"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AzurekeyvaultResourceModel describes the resource data model.
type AzurekeyvaultResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Azureapplication types.String `tfsdk:"azureapplication"`
	Azurevaultname   types.String `tfsdk:"azurevaultname"`
	Name             types.String `tfsdk:"name"`
}

func (r *AzurekeyvaultResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the azurekeyvault resource.",
			},
			"azureapplication": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Azure Application object created on the ADC appliance. This object will be used for authentication with Azure Active Directory",
			},
			"azurevaultname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Key Vault configured in Azure cloud using either the Azure CLI or the Azure portal (GUI) with complete domain name. Example: Test.vault.azure.net.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Key Vault. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the Key Vault is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my keyvault\" or 'my keyvault').",
			},
		},
	}
}

func azurekeyvaultGetThePayloadFromthePlan(ctx context.Context, data *AzurekeyvaultResourceModel) azure.Azurekeyvault {
	tflog.Debug(ctx, "In azurekeyvaultGetThePayloadFromthePlan Function")

	azurekeyvault := azure.Azurekeyvault{}
	if !data.Azureapplication.IsNull() && !data.Azureapplication.IsUnknown() {
		azurekeyvault.Azureapplication = data.Azureapplication.ValueString()
	}
	if !data.Azurevaultname.IsNull() && !data.Azurevaultname.IsUnknown() {
		azurekeyvault.Azurevaultname = data.Azurevaultname.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		azurekeyvault.Name = data.Name.ValueString()
	}

	return azurekeyvault
}

func azurekeyvaultSetAttrFromGet(ctx context.Context, data *AzurekeyvaultResourceModel, getResponseData map[string]interface{}) *AzurekeyvaultResourceModel {
	tflog.Debug(ctx, "In azurekeyvaultSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["azureapplication"]; ok && val != nil {
		data.Azureapplication = types.StringValue(val.(string))
	} else {
		data.Azureapplication = types.StringNull()
	}
	if val, ok := getResponseData["azurevaultname"]; ok && val != nil {
		data.Azurevaultname = types.StringValue(val.(string))
	} else {
		data.Azurevaultname = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
