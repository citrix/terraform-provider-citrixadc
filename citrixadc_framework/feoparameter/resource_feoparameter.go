package feoparameter

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
var _ resource.Resource = &FeoparameterResource{}
var _ resource.ResourceWithConfigure = (*FeoparameterResource)(nil)
var _ resource.ResourceWithImportState = (*FeoparameterResource)(nil)

func NewFeoparameterResource() resource.Resource {
	return &FeoparameterResource{}
}

// FeoparameterResource defines the resource implementation.
type FeoparameterResource struct {
	client *service.NitroClient
}

func (r *FeoparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *FeoparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_feoparameter"
}

func (r *FeoparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FeoparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FeoparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating feoparameter resource")

	feoparameter := feoparameterGetThePayloadFromthePlan(ctx, &data)

	// Singleton (unnamed) resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Feoparameter.Type(), &feoparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create feoparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("feoparameter-config")

	tflog.Trace(ctx, "Created feoparameter resource")

	// Read the updated state back
	r.readFeoparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FeoparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading feoparameter resource")

	r.readFeoparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state FeoparameterResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating feoparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Cssinlinethressize.Equal(state.Cssinlinethressize) {
		tflog.Debug(ctx, "cssinlinethressize has changed for feoparameter")
		hasChange = true
	}
	if !data.Imginlinethressize.Equal(state.Imginlinethressize) {
		tflog.Debug(ctx, "imginlinethressize has changed for feoparameter")
		hasChange = true
	}
	if !data.Jpegqualitypercent.Equal(state.Jpegqualitypercent) {
		tflog.Debug(ctx, "jpegqualitypercent has changed for feoparameter")
		hasChange = true
	}
	if !data.Jsinlinethressize.Equal(state.Jsinlinethressize) {
		tflog.Debug(ctx, "jsinlinethressize has changed for feoparameter")
		hasChange = true
	}

	if hasChange {
		feoparameter := feoparameterGetThePayloadFromthePlan(ctx, &data)

		// Singleton (unnamed) resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Feoparameter.Type(), &feoparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update feoparameter, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated feoparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for feoparameter resource, skipping update")
	}

	// Read the updated state back
	r.readFeoparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FeoparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting feoparameter resource")

	// feoparameter is a global configuration singleton and does not support a
	// NITRO DELETE operation (matches SDK v2). Just remove it from state.
	tflog.Trace(ctx, "Deleted feoparameter resource from state")
}

// Helper function to read feoparameter data from API
func (r *FeoparameterResource) readFeoparameterFromApi(ctx context.Context, data *FeoparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Feoparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read feoparameter, got error: %s", err))
		return
	}

	feoparameterSetAttrFromGet(ctx, data, getResponseData)

}
