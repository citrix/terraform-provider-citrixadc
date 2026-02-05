package nslimitidentifier

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
var _ resource.Resource = &NslimitidentifierResource{}
var _ resource.ResourceWithConfigure = (*NslimitidentifierResource)(nil)
var _ resource.ResourceWithImportState = (*NslimitidentifierResource)(nil)

func NewNslimitidentifierResource() resource.Resource {
	return &NslimitidentifierResource{}
}

// NslimitidentifierResource defines the resource implementation.
type NslimitidentifierResource struct {
	client *service.NitroClient
}

func (r *NslimitidentifierResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslimitidentifierResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslimitidentifier"
}

func (r *NslimitidentifierResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslimitidentifierResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslimitidentifierResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslimitidentifier resource")

	// nslimitidentifier := nslimitidentifierGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nslimitidentifier.Type(), &nslimitidentifier)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nslimitidentifier, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nslimitidentifier-config")

	tflog.Trace(ctx, "Created nslimitidentifier resource")

	// Read the updated state back
	r.readNslimitidentifierFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitidentifierResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslimitidentifierResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nslimitidentifier resource")

	r.readNslimitidentifierFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitidentifierResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NslimitidentifierResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nslimitidentifier resource")

	// Create API request body from the model
	// nslimitidentifier := nslimitidentifierGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nslimitidentifier.Type(), &nslimitidentifier)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nslimitidentifier, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nslimitidentifier resource")

	// Read the updated state back
	r.readNslimitidentifierFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitidentifierResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslimitidentifierResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nslimitidentifier resource")

	// For nslimitidentifier, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nslimitidentifier resource from state")
}

// Helper function to read nslimitidentifier data from API
func (r *NslimitidentifierResource) readNslimitidentifierFromApi(ctx context.Context, data *NslimitidentifierResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nslimitidentifier.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nslimitidentifier, got error: %s", err))
		return
	}

	nslimitidentifierSetAttrFromGet(ctx, data, getResponseData)

}
