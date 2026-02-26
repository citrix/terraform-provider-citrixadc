package botprofile_whitelist_binding

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
var _ resource.Resource = &BotprofileWhitelistBindingResource{}
var _ resource.ResourceWithConfigure = (*BotprofileWhitelistBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileWhitelistBindingResource)(nil)

func NewBotprofileWhitelistBindingResource() resource.Resource {
	return &BotprofileWhitelistBindingResource{}
}

// BotprofileWhitelistBindingResource defines the resource implementation.
type BotprofileWhitelistBindingResource struct {
	client *service.NitroClient
}

func (r *BotprofileWhitelistBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileWhitelistBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_whitelist_binding"
}

func (r *BotprofileWhitelistBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileWhitelistBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileWhitelistBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile_whitelist_binding resource")

	// botprofile_whitelist_binding := botprofile_whitelist_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botprofile_whitelist_binding.Type(), &botprofile_whitelist_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile_whitelist_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("botprofile_whitelist_binding-config")

	tflog.Trace(ctx, "Created botprofile_whitelist_binding resource")

	// Read the updated state back
	r.readBotprofileWhitelistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileWhitelistBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileWhitelistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile_whitelist_binding resource")

	r.readBotprofileWhitelistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileWhitelistBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BotprofileWhitelistBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating botprofile_whitelist_binding resource")

	// Create API request body from the model
	// botprofile_whitelist_binding := botprofile_whitelist_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botprofile_whitelist_binding.Type(), &botprofile_whitelist_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botprofile_whitelist_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated botprofile_whitelist_binding resource")

	// Read the updated state back
	r.readBotprofileWhitelistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileWhitelistBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileWhitelistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile_whitelist_binding resource")

	// For botprofile_whitelist_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted botprofile_whitelist_binding resource from state")
}

// Helper function to read botprofile_whitelist_binding data from API
func (r *BotprofileWhitelistBindingResource) readBotprofileWhitelistBindingFromApi(ctx context.Context, data *BotprofileWhitelistBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Botprofile_whitelist_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_whitelist_binding, got error: %s", err))
		return
	}

	botprofile_whitelist_bindingSetAttrFromGet(ctx, data, getResponseData)

}
