package sslocspresponder

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
var _ resource.Resource = &SslocspresponderResource{}
var _ resource.ResourceWithConfigure = (*SslocspresponderResource)(nil)
var _ resource.ResourceWithImportState = (*SslocspresponderResource)(nil)

func NewSslocspresponderResource() resource.Resource {
	return &SslocspresponderResource{}
}

// SslocspresponderResource defines the resource implementation.
type SslocspresponderResource struct {
	client *service.NitroClient
}

func (r *SslocspresponderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslocspresponderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslocspresponder"
}

func (r *SslocspresponderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslocspresponderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslocspresponderResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslocspresponder resource")

	// sslocspresponder := sslocspresponderGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslocspresponder.Type(), &sslocspresponder)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslocspresponder, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslocspresponder-config")

	tflog.Trace(ctx, "Created sslocspresponder resource")

	// Read the updated state back
	r.readSslocspresponderFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslocspresponderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslocspresponderResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslocspresponder resource")

	r.readSslocspresponderFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslocspresponderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslocspresponderResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslocspresponder resource")

	// Create API request body from the model
	// sslocspresponder := sslocspresponderGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslocspresponder.Type(), &sslocspresponder)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslocspresponder, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslocspresponder resource")

	// Read the updated state back
	r.readSslocspresponderFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslocspresponderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslocspresponderResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslocspresponder resource")

	// For sslocspresponder, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslocspresponder resource from state")
}

// Helper function to read sslocspresponder data from API
func (r *SslocspresponderResource) readSslocspresponderFromApi(ctx context.Context, data *SslocspresponderResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslocspresponder.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslocspresponder, got error: %s", err))
		return
	}

	sslocspresponderSetAttrFromGet(ctx, data, getResponseData)

}
