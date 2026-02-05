package snmpengineid

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
var _ resource.Resource = &SnmpengineidResource{}
var _ resource.ResourceWithConfigure = (*SnmpengineidResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpengineidResource)(nil)

func NewSnmpengineidResource() resource.Resource {
	return &SnmpengineidResource{}
}

// SnmpengineidResource defines the resource implementation.
type SnmpengineidResource struct {
	client *service.NitroClient
}

func (r *SnmpengineidResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpengineidResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpengineid"
}

func (r *SnmpengineidResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpengineidResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmpengineidResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpengineid resource")

	// snmpengineid := snmpengineidGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmpengineid.Type(), &snmpengineid)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpengineid, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("snmpengineid-config")

	tflog.Trace(ctx, "Created snmpengineid resource")

	// Read the updated state back
	r.readSnmpengineidFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpengineidResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpengineidResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpengineid resource")

	r.readSnmpengineidFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpengineidResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SnmpengineidResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating snmpengineid resource")

	// Create API request body from the model
	// snmpengineid := snmpengineidGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmpengineid.Type(), &snmpengineid)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpengineid, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated snmpengineid resource")

	// Read the updated state back
	r.readSnmpengineidFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpengineidResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpengineidResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpengineid resource")

	// For snmpengineid, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted snmpengineid resource from state")
}

// Helper function to read snmpengineid data from API
func (r *SnmpengineidResource) readSnmpengineidFromApi(ctx context.Context, data *SnmpengineidResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Snmpengineid.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpengineid, got error: %s", err))
		return
	}

	snmpengineidSetAttrFromGet(ctx, data, getResponseData)

}
