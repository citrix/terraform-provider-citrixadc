package policyexpression

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
var _ resource.Resource = &PolicyexpressionResource{}
var _ resource.ResourceWithConfigure = (*PolicyexpressionResource)(nil)
var _ resource.ResourceWithImportState = (*PolicyexpressionResource)(nil)

func NewPolicyexpressionResource() resource.Resource {
	return &PolicyexpressionResource{}
}

// PolicyexpressionResource defines the resource implementation.
type PolicyexpressionResource struct {
	client *service.NitroClient
}

func (r *PolicyexpressionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicyexpressionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policyexpression"
}

func (r *PolicyexpressionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicyexpressionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicyexpressionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policyexpression resource")

	// policyexpression := policyexpressionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Policyexpression.Type(), &policyexpression)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policyexpression, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("policyexpression-config")

	tflog.Trace(ctx, "Created policyexpression resource")

	// Read the updated state back
	r.readPolicyexpressionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyexpressionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicyexpressionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policyexpression resource")

	r.readPolicyexpressionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyexpressionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data PolicyexpressionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating policyexpression resource")

	// Create API request body from the model
	// policyexpression := policyexpressionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Policyexpression.Type(), &policyexpression)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policyexpression, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated policyexpression resource")

	// Read the updated state back
	r.readPolicyexpressionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyexpressionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicyexpressionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policyexpression resource")

	// For policyexpression, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted policyexpression resource from state")
}

// Helper function to read policyexpression data from API
func (r *PolicyexpressionResource) readPolicyexpressionFromApi(ctx context.Context, data *PolicyexpressionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Policyexpression.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policyexpression, got error: %s", err))
		return
	}

	policyexpressionSetAttrFromGet(ctx, data, getResponseData)

}
