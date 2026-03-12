package botprofile_blacklist_binding

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
var _ resource.Resource = &BotprofileBlacklistBindingResource{}
var _ resource.ResourceWithConfigure = (*BotprofileBlacklistBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileBlacklistBindingResource)(nil)

func NewBotprofileBlacklistBindingResource() resource.Resource {
	return &BotprofileBlacklistBindingResource{}
}

// BotprofileBlacklistBindingResource defines the resource implementation.
type BotprofileBlacklistBindingResource struct {
	client *service.NitroClient
}

func (r *BotprofileBlacklistBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileBlacklistBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_blacklist_binding"
}

func (r *BotprofileBlacklistBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileBlacklistBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileBlacklistBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile_blacklist_binding resource")

	// botprofile_blacklist_binding := botprofile_blacklist_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botprofile_blacklist_binding.Type(), &botprofile_blacklist_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile_blacklist_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("botprofile_blacklist_binding-config")

	tflog.Trace(ctx, "Created botprofile_blacklist_binding resource")

	// Read the updated state back
	r.readBotprofileBlacklistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileBlacklistBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileBlacklistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile_blacklist_binding resource")

	r.readBotprofileBlacklistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileBlacklistBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BotprofileBlacklistBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating botprofile_blacklist_binding resource")

	// Create API request body from the model
	// botprofile_blacklist_binding := botprofile_blacklist_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botprofile_blacklist_binding.Type(), &botprofile_blacklist_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botprofile_blacklist_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated botprofile_blacklist_binding resource")

	// Read the updated state back
	r.readBotprofileBlacklistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileBlacklistBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileBlacklistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile_blacklist_binding resource")

	// For botprofile_blacklist_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted botprofile_blacklist_binding resource from state")
}

// Helper function to read botprofile_blacklist_binding data from API
func (r *BotprofileBlacklistBindingResource) readBotprofileBlacklistBindingFromApi(ctx context.Context, data *BotprofileBlacklistBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Botprofile_blacklist_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_blacklist_binding, got error: %s", err))
		return
	}

	botprofile_blacklist_bindingSetAttrFromGet(ctx, data, getResponseData)

}
