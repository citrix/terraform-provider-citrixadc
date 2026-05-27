package azureapplication

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AzureapplicationResource{}
var _ resource.ResourceWithConfigure = (*AzureapplicationResource)(nil)
var _ resource.ResourceWithImportState = (*AzureapplicationResource)(nil)

func NewAzureapplicationResource() resource.Resource {
	return &AzureapplicationResource{}
}

// AzureapplicationResource defines the resource implementation.
type AzureapplicationResource struct {
	client *service.NitroClient
}

func (r *AzureapplicationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AzureapplicationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_azureapplication"
}

func (r *AzureapplicationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AzureapplicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AzureapplicationResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating azureapplication resource")
	// Get payload from plan (regular attributes)
	azureapplication := azureapplicationGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	azureapplicationGetThePayloadFromtheConfig(ctx, &config, &azureapplication)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Azureapplication.Type(), name_value, &azureapplication)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create azureapplication, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created azureapplication resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAzureapplicationFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AzureapplicationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AzureapplicationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading azureapplication resource")

	r.readAzureapplicationFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AzureapplicationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AzureapplicationResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating azureapplication resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute clientsecret or its version tracker
	if !data.Clientsecret.Equal(state.Clientsecret) {
		tflog.Debug(ctx, fmt.Sprintf("clientsecret has changed for azureapplication"))
		hasChange = true
	} else if !data.ClientsecretWoVersion.Equal(state.ClientsecretWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("clientsecret_wo_version has changed for azureapplication"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		azureapplication := azureapplicationGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		azureapplicationGetThePayloadFromtheConfig(ctx, &config, &azureapplication)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Azureapplication.Type(), name_value, &azureapplication)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update azureapplication, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated azureapplication resource")
	} else {
		tflog.Debug(ctx, "No changes detected for azureapplication resource, skipping update")
	}

	// Read the updated state back
	r.readAzureapplicationFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AzureapplicationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AzureapplicationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting azureapplication resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Azureapplication.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete azureapplication, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted azureapplication resource")
}

// Helper function to read azureapplication data from API
func (r *AzureapplicationResource) readAzureapplicationFromApi(ctx context.Context, data *AzureapplicationResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Azureapplication.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read azureapplication, got error: %s", err))
		return
	}

	azureapplicationSetAttrFromGet(ctx, data, getResponseData)

}
