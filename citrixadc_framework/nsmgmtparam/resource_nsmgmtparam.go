package nsmgmtparam

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
var _ resource.Resource = &NsmgmtparamResource{}
var _ resource.ResourceWithConfigure = (*NsmgmtparamResource)(nil)
var _ resource.ResourceWithImportState = (*NsmgmtparamResource)(nil)

func NewNsmgmtparamResource() resource.Resource {
	return &NsmgmtparamResource{}
}

// NsmgmtparamResource defines the resource implementation.
type NsmgmtparamResource struct {
	client *service.NitroClient
}

func (r *NsmgmtparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsmgmtparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsmgmtparam"
}

func (r *NsmgmtparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsmgmtparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsmgmtparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsmgmtparam resource")
	nsmgmtparam := nsmgmtparamGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Parameter singleton - NITRO has no "add" verb; create is the "set" verb (PUT).
	_, err := r.client.UpdateResource(service.Nsmgmtparam.Type(), "", &nsmgmtparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsmgmtparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsmgmtparam resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("nsmgmtparam-config")

	// Read the updated state back
	r.readNsmgmtparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmgmtparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsmgmtparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsmgmtparam resource")

	r.readNsmgmtparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmgmtparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsmgmtparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsmgmtparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Httpdmaxclients.Equal(state.Httpdmaxclients) {
		tflog.Debug(ctx, fmt.Sprintf("httpdmaxclients has changed for nsmgmtparam"))
		hasChange = true
	}
	if !data.Httpdmaxreqworkers.Equal(state.Httpdmaxreqworkers) {
		tflog.Debug(ctx, fmt.Sprintf("httpdmaxreqworkers has changed for nsmgmtparam"))
		hasChange = true
	}
	if !data.Mgmthttpport.Equal(state.Mgmthttpport) {
		tflog.Debug(ctx, fmt.Sprintf("mgmthttpport has changed for nsmgmtparam"))
		hasChange = true
	}
	if !data.Mgmthttpsport.Equal(state.Mgmthttpsport) {
		tflog.Debug(ctx, fmt.Sprintf("mgmthttpsport has changed for nsmgmtparam"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		nsmgmtparam := nsmgmtparamGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Parameter singleton - update uses the same "set" verb (PUT).
		_, err := r.client.UpdateResource(service.Nsmgmtparam.Type(), "", &nsmgmtparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsmgmtparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsmgmtparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsmgmtparam resource, skipping update")
	}

	// Read the updated state back
	r.readNsmgmtparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmgmtparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsmgmtparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsmgmtparam resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed nsmgmtparam from Terraform state")
}

// Helper function to read nsmgmtparam data from API
func (r *NsmgmtparamResource) readNsmgmtparamFromApi(ctx context.Context, data *NsmgmtparamResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nsmgmtparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsmgmtparam, got error: %s", err))
		return
	}

	nsmgmtparamSetAttrFromGet(ctx, data, getResponseData)

}
