package botprofile_logexpression_binding

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
var _ resource.Resource = &BotprofileLogexpressionBindingResource{}
var _ resource.ResourceWithConfigure = (*BotprofileLogexpressionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileLogexpressionBindingResource)(nil)

func NewBotprofileLogexpressionBindingResource() resource.Resource {
	return &BotprofileLogexpressionBindingResource{}
}

// BotprofileLogexpressionBindingResource defines the resource implementation.
type BotprofileLogexpressionBindingResource struct {
	client *service.NitroClient
}

func (r *BotprofileLogexpressionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileLogexpressionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_logexpression_binding"
}

func (r *BotprofileLogexpressionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileLogexpressionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileLogexpressionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile_logexpression_binding resource")

	// botprofile_logexpression_binding := botprofile_logexpression_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botprofile_logexpression_binding.Type(), &botprofile_logexpression_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile_logexpression_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("botprofile_logexpression_binding-config")

	tflog.Trace(ctx, "Created botprofile_logexpression_binding resource")

	// Read the updated state back
	r.readBotprofileLogexpressionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileLogexpressionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileLogexpressionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile_logexpression_binding resource")

	r.readBotprofileLogexpressionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileLogexpressionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BotprofileLogexpressionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating botprofile_logexpression_binding resource")

	// Create API request body from the model
	// botprofile_logexpression_binding := botprofile_logexpression_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botprofile_logexpression_binding.Type(), &botprofile_logexpression_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botprofile_logexpression_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated botprofile_logexpression_binding resource")

	// Read the updated state back
	r.readBotprofileLogexpressionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileLogexpressionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileLogexpressionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile_logexpression_binding resource")

	// For botprofile_logexpression_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted botprofile_logexpression_binding resource from state")
}

// Helper function to read botprofile_logexpression_binding data from API
func (r *BotprofileLogexpressionBindingResource) readBotprofileLogexpressionBindingFromApi(ctx context.Context, data *BotprofileLogexpressionBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Botprofile_logexpression_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_logexpression_binding, got error: %s", err))
		return
	}

	botprofile_logexpression_bindingSetAttrFromGet(ctx, data, getResponseData)

}
