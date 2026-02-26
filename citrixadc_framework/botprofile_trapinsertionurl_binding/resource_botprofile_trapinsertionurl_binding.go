package botprofile_trapinsertionurl_binding

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
var _ resource.Resource = &BotprofileTrapinsertionurlBindingResource{}
var _ resource.ResourceWithConfigure = (*BotprofileTrapinsertionurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileTrapinsertionurlBindingResource)(nil)

func NewBotprofileTrapinsertionurlBindingResource() resource.Resource {
	return &BotprofileTrapinsertionurlBindingResource{}
}

// BotprofileTrapinsertionurlBindingResource defines the resource implementation.
type BotprofileTrapinsertionurlBindingResource struct {
	client *service.NitroClient
}

func (r *BotprofileTrapinsertionurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileTrapinsertionurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_trapinsertionurl_binding"
}

func (r *BotprofileTrapinsertionurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileTrapinsertionurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileTrapinsertionurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile_trapinsertionurl_binding resource")

	// botprofile_trapinsertionurl_binding := botprofile_trapinsertionurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botprofile_trapinsertionurl_binding.Type(), &botprofile_trapinsertionurl_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile_trapinsertionurl_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("botprofile_trapinsertionurl_binding-config")

	tflog.Trace(ctx, "Created botprofile_trapinsertionurl_binding resource")

	// Read the updated state back
	r.readBotprofileTrapinsertionurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileTrapinsertionurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileTrapinsertionurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile_trapinsertionurl_binding resource")

	r.readBotprofileTrapinsertionurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileTrapinsertionurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BotprofileTrapinsertionurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating botprofile_trapinsertionurl_binding resource")

	// Create API request body from the model
	// botprofile_trapinsertionurl_binding := botprofile_trapinsertionurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botprofile_trapinsertionurl_binding.Type(), &botprofile_trapinsertionurl_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botprofile_trapinsertionurl_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated botprofile_trapinsertionurl_binding resource")

	// Read the updated state back
	r.readBotprofileTrapinsertionurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileTrapinsertionurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileTrapinsertionurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile_trapinsertionurl_binding resource")

	// For botprofile_trapinsertionurl_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted botprofile_trapinsertionurl_binding resource from state")
}

// Helper function to read botprofile_trapinsertionurl_binding data from API
func (r *BotprofileTrapinsertionurlBindingResource) readBotprofileTrapinsertionurlBindingFromApi(ctx context.Context, data *BotprofileTrapinsertionurlBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Botprofile_trapinsertionurl_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_trapinsertionurl_binding, got error: %s", err))
		return
	}

	botprofile_trapinsertionurl_bindingSetAttrFromGet(ctx, data, getResponseData)

}
