package azureapplication

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/azure"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AzureapplicationResourceModel describes the resource data model.
type AzureapplicationResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Clientid              types.String `tfsdk:"clientid"`
	Clientsecret          types.String `tfsdk:"clientsecret"`
	ClientsecretWo        types.String `tfsdk:"clientsecret_wo"`
	ClientsecretWoVersion types.Int64  `tfsdk:"clientsecret_wo_version"`
	Name                  types.String `tfsdk:"name"`
	Tenantid              types.String `tfsdk:"tenantid"`
	Tokenendpoint         types.String `tfsdk:"tokenendpoint"`
	Vaultresource         types.String `tfsdk:"vaultresource"`
}

func (r *AzureapplicationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the azureapplication resource.",
			},
			"clientid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Application ID that is generated when an application is created in Azure Active Directory using either the Azure CLI or the Azure portal (GUI)",
			},
			"clientsecret": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password for the application configured in Azure Active Directory. The password is specified in the Azure CLI or generated in the Azure portal (GUI).",
			},
			"clientsecret_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password for the application configured in Azure Active Directory. The password is specified in the Azure CLI or generated in the Azure portal (GUI).",
			},
			"clientsecret_wo_version": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Increment this version to signal a clientsecret_wo update.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the application. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the application is created.',\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my application\" or 'my application').",
			},
			"tenantid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ID of the directory inside Azure Active Directory in which the application was created",
			},
			"tokenendpoint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL from where access token can be obtained. If the token end point is not specified, the default value is https://login.microsoftonline.com/<tenant id>.",
			},
			"vaultresource": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Vault resource for which access token is granted. Example : vault.azure.net",
			},
		},
	}
}

func azureapplicationGetThePayloadFromthePlan(ctx context.Context, data *AzureapplicationResourceModel) azure.Azureapplication {
	tflog.Debug(ctx, "In azureapplicationGetThePayloadFromthePlan Function")

	azureapplication := azure.Azureapplication{}
	if !data.Clientid.IsNull() && !data.Clientid.IsUnknown() {
		azureapplication.Clientid = data.Clientid.ValueString()
	}
	if !data.Clientsecret.IsNull() && !data.Clientsecret.IsUnknown() {
		azureapplication.Clientsecret = data.Clientsecret.ValueString()
	}
	// Skip write-only attribute: clientsecret_wo
	// Skip version tracker attribute: clientsecret_wo_version
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		azureapplication.Name = data.Name.ValueString()
	}
	if !data.Tenantid.IsNull() && !data.Tenantid.IsUnknown() {
		azureapplication.Tenantid = data.Tenantid.ValueString()
	}
	if !data.Tokenendpoint.IsNull() && !data.Tokenendpoint.IsUnknown() {
		azureapplication.Tokenendpoint = data.Tokenendpoint.ValueString()
	}
	if !data.Vaultresource.IsNull() && !data.Vaultresource.IsUnknown() {
		azureapplication.Vaultresource = data.Vaultresource.ValueString()
	}

	return azureapplication
}

func azureapplicationGetThePayloadFromtheConfig(ctx context.Context, data *AzureapplicationResourceModel, payload *azure.Azureapplication) {
	tflog.Debug(ctx, "In azureapplicationGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: clientsecret_wo -> clientsecret
	if !data.ClientsecretWo.IsNull() {
		clientsecretWo := data.ClientsecretWo.ValueString()
		if clientsecretWo != "" {
			payload.Clientsecret = clientsecretWo
		}
	}
}

func azureapplicationSetAttrFromGet(ctx context.Context, data *AzureapplicationResourceModel, getResponseData map[string]interface{}) *AzureapplicationResourceModel {
	tflog.Debug(ctx, "In azureapplicationSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientid"]; ok && val != nil {
		data.Clientid = types.StringValue(val.(string))
	} else {
		data.Clientid = types.StringNull()
	}
	// clientsecret is not returned by NITRO API (secret/ephemeral) - retain from config
	// clientsecret_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// clientsecret_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
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
