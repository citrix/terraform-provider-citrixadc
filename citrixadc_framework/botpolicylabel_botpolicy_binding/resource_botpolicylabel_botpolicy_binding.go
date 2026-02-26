package botpolicylabel_botpolicy_binding

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
var _ resource.Resource = &BotpolicylabelBotpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*BotpolicylabelBotpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotpolicylabelBotpolicyBindingResource)(nil)

func NewBotpolicylabelBotpolicyBindingResource() resource.Resource {
	return &BotpolicylabelBotpolicyBindingResource{}
}

// BotpolicylabelBotpolicyBindingResource defines the resource implementation.
type BotpolicylabelBotpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *BotpolicylabelBotpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotpolicylabelBotpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botpolicylabel_botpolicy_binding"
}

func (r *BotpolicylabelBotpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotpolicylabelBotpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotpolicylabelBotpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botpolicylabel_botpolicy_binding resource")

	// botpolicylabel_botpolicy_binding := botpolicylabel_botpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botpolicylabel_botpolicy_binding.Type(), &botpolicylabel_botpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botpolicylabel_botpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("botpolicylabel_botpolicy_binding-config")

	tflog.Trace(ctx, "Created botpolicylabel_botpolicy_binding resource")

	// Read the updated state back
	r.readBotpolicylabelBotpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotpolicylabelBotpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotpolicylabelBotpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botpolicylabel_botpolicy_binding resource")

	r.readBotpolicylabelBotpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotpolicylabelBotpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BotpolicylabelBotpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating botpolicylabel_botpolicy_binding resource")

	// Create API request body from the model
	// botpolicylabel_botpolicy_binding := botpolicylabel_botpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botpolicylabel_botpolicy_binding.Type(), &botpolicylabel_botpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botpolicylabel_botpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated botpolicylabel_botpolicy_binding resource")

	// Read the updated state back
	r.readBotpolicylabelBotpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotpolicylabelBotpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotpolicylabelBotpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botpolicylabel_botpolicy_binding resource")

	// For botpolicylabel_botpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted botpolicylabel_botpolicy_binding resource from state")
}

// Helper function to read botpolicylabel_botpolicy_binding data from API
func (r *BotpolicylabelBotpolicyBindingResource) readBotpolicylabelBotpolicyBindingFromApi(ctx context.Context, data *BotpolicylabelBotpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Botpolicylabel_botpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botpolicylabel_botpolicy_binding, got error: %s", err))
		return
	}

	botpolicylabel_botpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
