package linkset

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
var _ resource.Resource = &LinksetResource{}
var _ resource.ResourceWithConfigure = (*LinksetResource)(nil)
var _ resource.ResourceWithImportState = (*LinksetResource)(nil)

func NewLinksetResource() resource.Resource {
	return &LinksetResource{}
}

// LinksetResource defines the resource implementation.
type LinksetResource struct {
	client *service.NitroClient
}

func (r *LinksetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LinksetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_linkset"
}

func (r *LinksetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LinksetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LinksetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating linkset resource")

	// linkset := linksetGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Linkset.Type(), &linkset)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create linkset, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("linkset-config")

	tflog.Trace(ctx, "Created linkset resource")

	// Read the updated state back
	r.readLinksetFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LinksetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LinksetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading linkset resource")

	r.readLinksetFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LinksetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LinksetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating linkset resource")

	// Create API request body from the model
	// linkset := linksetGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Linkset.Type(), &linkset)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update linkset, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated linkset resource")

	// Read the updated state back
	r.readLinksetFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LinksetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LinksetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting linkset resource")

	// For linkset, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted linkset resource from state")
}

// Helper function to read linkset data from API
func (r *LinksetResource) readLinksetFromApi(ctx context.Context, data *LinksetResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Linkset.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read linkset, got error: %s", err))
		return
	}

	linksetSetAttrFromGet(ctx, data, getResponseData)

}
