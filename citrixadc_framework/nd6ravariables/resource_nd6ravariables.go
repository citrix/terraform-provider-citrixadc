package nd6ravariables

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
var _ resource.Resource = &Nd6ravariablesResource{}
var _ resource.ResourceWithConfigure = (*Nd6ravariablesResource)(nil)
var _ resource.ResourceWithImportState = (*Nd6ravariablesResource)(nil)

func NewNd6ravariablesResource() resource.Resource {
	return &Nd6ravariablesResource{}
}

// Nd6ravariablesResource defines the resource implementation.
type Nd6ravariablesResource struct {
	client *service.NitroClient
}

func (r *Nd6ravariablesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nd6ravariablesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nd6ravariables"
}

func (r *Nd6ravariablesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nd6ravariablesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nd6ravariablesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nd6ravariables resource")

	// nd6ravariables := nd6ravariablesGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nd6ravariables.Type(), &nd6ravariables)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nd6ravariables, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nd6ravariables-config")

	tflog.Trace(ctx, "Created nd6ravariables resource")

	// Read the updated state back
	r.readNd6ravariablesFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6ravariablesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nd6ravariablesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nd6ravariables resource")

	r.readNd6ravariablesFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6ravariablesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Nd6ravariablesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nd6ravariables resource")

	// Create API request body from the model
	// nd6ravariables := nd6ravariablesGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nd6ravariables.Type(), &nd6ravariables)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nd6ravariables, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nd6ravariables resource")

	// Read the updated state back
	r.readNd6ravariablesFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6ravariablesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nd6ravariablesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nd6ravariables resource")

	// For nd6ravariables, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nd6ravariables resource from state")
}

// Helper function to read nd6ravariables data from API
func (r *Nd6ravariablesResource) readNd6ravariablesFromApi(ctx context.Context, data *Nd6ravariablesResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nd6ravariables.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nd6ravariables, got error: %s", err))
		return
	}

	nd6ravariablesSetAttrFromGet(ctx, data, getResponseData)

}
