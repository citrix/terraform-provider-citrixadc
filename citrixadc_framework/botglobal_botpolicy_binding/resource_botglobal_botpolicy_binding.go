package botglobal_botpolicy_binding

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
var _ resource.Resource = &BotglobalBotpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*BotglobalBotpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotglobalBotpolicyBindingResource)(nil)

func NewBotglobalBotpolicyBindingResource() resource.Resource {
	return &BotglobalBotpolicyBindingResource{}
}

// BotglobalBotpolicyBindingResource defines the resource implementation.
type BotglobalBotpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *BotglobalBotpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotglobalBotpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botglobal_botpolicy_binding"
}

func (r *BotglobalBotpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotglobalBotpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotglobalBotpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botglobal_botpolicy_binding resource")

	// botglobal_botpolicy_binding := botglobal_botpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botglobal_botpolicy_binding.Type(), &botglobal_botpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botglobal_botpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("botglobal_botpolicy_binding-config")

	tflog.Trace(ctx, "Created botglobal_botpolicy_binding resource")

	// Read the updated state back
	r.readBotglobalBotpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotglobalBotpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotglobalBotpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botglobal_botpolicy_binding resource")

	r.readBotglobalBotpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotglobalBotpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BotglobalBotpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating botglobal_botpolicy_binding resource")

	// Create API request body from the model
	// botglobal_botpolicy_binding := botglobal_botpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botglobal_botpolicy_binding.Type(), &botglobal_botpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botglobal_botpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated botglobal_botpolicy_binding resource")

	// Read the updated state back
	r.readBotglobalBotpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotglobalBotpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotglobalBotpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botglobal_botpolicy_binding resource")

	// For botglobal_botpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted botglobal_botpolicy_binding resource from state")
}

// Helper function to read botglobal_botpolicy_binding data from API
func (r *BotglobalBotpolicyBindingResource) readBotglobalBotpolicyBindingFromApi(ctx context.Context, data *BotglobalBotpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Botglobal_botpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botglobal_botpolicy_binding, got error: %s", err))
		return
	}

	botglobal_botpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
