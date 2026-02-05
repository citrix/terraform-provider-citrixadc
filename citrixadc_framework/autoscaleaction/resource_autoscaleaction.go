package autoscaleaction

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
var _ resource.Resource = &AutoscaleactionResource{}
var _ resource.ResourceWithConfigure = (*AutoscaleactionResource)(nil)
var _ resource.ResourceWithImportState = (*AutoscaleactionResource)(nil)

func NewAutoscaleactionResource() resource.Resource {
	return &AutoscaleactionResource{}
}

// AutoscaleactionResource defines the resource implementation.
type AutoscaleactionResource struct {
	client *service.NitroClient
}

func (r *AutoscaleactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AutoscaleactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_autoscaleaction"
}

func (r *AutoscaleactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AutoscaleactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AutoscaleactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating autoscaleaction resource")

	// autoscaleaction := autoscaleactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Autoscaleaction.Type(), &autoscaleaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create autoscaleaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("autoscaleaction-config")

	tflog.Trace(ctx, "Created autoscaleaction resource")

	// Read the updated state back
	r.readAutoscaleactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscaleactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AutoscaleactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading autoscaleaction resource")

	r.readAutoscaleactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscaleactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AutoscaleactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating autoscaleaction resource")

	// Create API request body from the model
	// autoscaleaction := autoscaleactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Autoscaleaction.Type(), &autoscaleaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update autoscaleaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated autoscaleaction resource")

	// Read the updated state back
	r.readAutoscaleactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscaleactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AutoscaleactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting autoscaleaction resource")

	// For autoscaleaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted autoscaleaction resource from state")
}

// Helper function to read autoscaleaction data from API
func (r *AutoscaleactionResource) readAutoscaleactionFromApi(ctx context.Context, data *AutoscaleactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Autoscaleaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read autoscaleaction, got error: %s", err))
		return
	}

	autoscaleactionSetAttrFromGet(ctx, data, getResponseData)

}
