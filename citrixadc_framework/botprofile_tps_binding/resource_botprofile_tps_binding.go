package botprofile_tps_binding

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
var _ resource.Resource = &BotprofileTpsBindingResource{}
var _ resource.ResourceWithConfigure = (*BotprofileTpsBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileTpsBindingResource)(nil)

func NewBotprofileTpsBindingResource() resource.Resource {
	return &BotprofileTpsBindingResource{}
}

// BotprofileTpsBindingResource defines the resource implementation.
type BotprofileTpsBindingResource struct {
	client *service.NitroClient
}

func (r *BotprofileTpsBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileTpsBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_tps_binding"
}

func (r *BotprofileTpsBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileTpsBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileTpsBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile_tps_binding resource")

	// botprofile_tps_binding := botprofile_tps_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botprofile_tps_binding.Type(), &botprofile_tps_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile_tps_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("botprofile_tps_binding-config")

	tflog.Trace(ctx, "Created botprofile_tps_binding resource")

	// Read the updated state back
	r.readBotprofileTpsBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileTpsBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileTpsBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile_tps_binding resource")

	r.readBotprofileTpsBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileTpsBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BotprofileTpsBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating botprofile_tps_binding resource")

	// Create API request body from the model
	// botprofile_tps_binding := botprofile_tps_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botprofile_tps_binding.Type(), &botprofile_tps_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botprofile_tps_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated botprofile_tps_binding resource")

	// Read the updated state back
	r.readBotprofileTpsBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileTpsBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileTpsBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile_tps_binding resource")

	// For botprofile_tps_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted botprofile_tps_binding resource from state")
}

// Helper function to read botprofile_tps_binding data from API
func (r *BotprofileTpsBindingResource) readBotprofileTpsBindingFromApi(ctx context.Context, data *BotprofileTpsBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Botprofile_tps_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_tps_binding, got error: %s", err))
		return
	}

	botprofile_tps_bindingSetAttrFromGet(ctx, data, getResponseData)

}
