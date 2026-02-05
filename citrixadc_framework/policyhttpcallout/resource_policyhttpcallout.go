package policyhttpcallout

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
var _ resource.Resource = &PolicyhttpcalloutResource{}
var _ resource.ResourceWithConfigure = (*PolicyhttpcalloutResource)(nil)
var _ resource.ResourceWithImportState = (*PolicyhttpcalloutResource)(nil)

func NewPolicyhttpcalloutResource() resource.Resource {
	return &PolicyhttpcalloutResource{}
}

// PolicyhttpcalloutResource defines the resource implementation.
type PolicyhttpcalloutResource struct {
	client *service.NitroClient
}

func (r *PolicyhttpcalloutResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicyhttpcalloutResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policyhttpcallout"
}

func (r *PolicyhttpcalloutResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicyhttpcalloutResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicyhttpcalloutResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policyhttpcallout resource")

	// policyhttpcallout := policyhttpcalloutGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Policyhttpcallout.Type(), &policyhttpcallout)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policyhttpcallout, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("policyhttpcallout-config")

	tflog.Trace(ctx, "Created policyhttpcallout resource")

	// Read the updated state back
	r.readPolicyhttpcalloutFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyhttpcalloutResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicyhttpcalloutResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policyhttpcallout resource")

	r.readPolicyhttpcalloutFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyhttpcalloutResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data PolicyhttpcalloutResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating policyhttpcallout resource")

	// Create API request body from the model
	// policyhttpcallout := policyhttpcalloutGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Policyhttpcallout.Type(), &policyhttpcallout)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policyhttpcallout, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated policyhttpcallout resource")

	// Read the updated state back
	r.readPolicyhttpcalloutFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyhttpcalloutResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicyhttpcalloutResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policyhttpcallout resource")

	// For policyhttpcallout, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted policyhttpcallout resource from state")
}

// Helper function to read policyhttpcallout data from API
func (r *PolicyhttpcalloutResource) readPolicyhttpcalloutFromApi(ctx context.Context, data *PolicyhttpcalloutResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Policyhttpcallout.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policyhttpcallout, got error: %s", err))
		return
	}

	policyhttpcalloutSetAttrFromGet(ctx, data, getResponseData)

}
