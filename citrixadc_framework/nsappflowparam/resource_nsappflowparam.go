package nsappflowparam

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
var _ resource.Resource = &NsappflowparamResource{}
var _ resource.ResourceWithConfigure = (*NsappflowparamResource)(nil)
var _ resource.ResourceWithImportState = (*NsappflowparamResource)(nil)

func NewNsappflowparamResource() resource.Resource {
	return &NsappflowparamResource{}
}

// NsappflowparamResource defines the resource implementation.
type NsappflowparamResource struct {
	client *service.NitroClient
}

func (r *NsappflowparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsappflowparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsappflowparam"
}

func (r *NsappflowparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsappflowparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsappflowparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsappflowparam resource")
	nsappflowparam := nsappflowparamGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Parameter singleton - NITRO has no "add" verb; create is the "set" verb (PUT).
	_, err := r.client.UpdateResource(service.Nsappflowparam.Type(), "", &nsappflowparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsappflowparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsappflowparam resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("nsappflowparam-config")

	// Read the updated state back
	r.readNsappflowparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsappflowparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsappflowparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsappflowparam resource")

	r.readNsappflowparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsappflowparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsappflowparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsappflowparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Clienttrafficonly.Equal(state.Clienttrafficonly) {
		tflog.Debug(ctx, fmt.Sprintf("clienttrafficonly has changed for nsappflowparam"))
		hasChange = true
	}
	if !data.Httpcookie.Equal(state.Httpcookie) {
		tflog.Debug(ctx, fmt.Sprintf("httpcookie has changed for nsappflowparam"))
		hasChange = true
	}
	if !data.Httphost.Equal(state.Httphost) {
		tflog.Debug(ctx, fmt.Sprintf("httphost has changed for nsappflowparam"))
		hasChange = true
	}
	if !data.Httpmethod.Equal(state.Httpmethod) {
		tflog.Debug(ctx, fmt.Sprintf("httpmethod has changed for nsappflowparam"))
		hasChange = true
	}
	if !data.Httpreferer.Equal(state.Httpreferer) {
		tflog.Debug(ctx, fmt.Sprintf("httpreferer has changed for nsappflowparam"))
		hasChange = true
	}
	if !data.Httpurl.Equal(state.Httpurl) {
		tflog.Debug(ctx, fmt.Sprintf("httpurl has changed for nsappflowparam"))
		hasChange = true
	}
	if !data.Httpuseragent.Equal(state.Httpuseragent) {
		tflog.Debug(ctx, fmt.Sprintf("httpuseragent has changed for nsappflowparam"))
		hasChange = true
	}
	if !data.Templaterefresh.Equal(state.Templaterefresh) {
		tflog.Debug(ctx, fmt.Sprintf("templaterefresh has changed for nsappflowparam"))
		hasChange = true
	}
	if !data.Udppmtu.Equal(state.Udppmtu) {
		tflog.Debug(ctx, fmt.Sprintf("udppmtu has changed for nsappflowparam"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		nsappflowparam := nsappflowparamGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Parameter singleton - update uses the same "set" verb (PUT).
		_, err := r.client.UpdateResource(service.Nsappflowparam.Type(), "", &nsappflowparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsappflowparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsappflowparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsappflowparam resource, skipping update")
	}

	// Read the updated state back
	r.readNsappflowparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsappflowparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsappflowparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsappflowparam resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed nsappflowparam from Terraform state")
}

// Helper function to read nsappflowparam data from API
func (r *NsappflowparamResource) readNsappflowparamFromApi(ctx context.Context, data *NsappflowparamResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nsappflowparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsappflowparam, got error: %s", err))
		return
	}

	nsappflowparamSetAttrFromGet(ctx, data, getResponseData)

}
