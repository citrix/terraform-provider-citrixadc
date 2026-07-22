package callhome

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
var _ resource.Resource = &CallhomeResource{}
var _ resource.ResourceWithConfigure = (*CallhomeResource)(nil)
var _ resource.ResourceWithImportState = (*CallhomeResource)(nil)

func NewCallhomeResource() resource.Resource {
	return &CallhomeResource{}
}

// CallhomeResource defines the resource implementation.
type CallhomeResource struct {
	client *service.NitroClient
}

func (r *CallhomeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CallhomeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_callhome"
}

func (r *CallhomeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CallhomeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CallhomeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating callhome resource")
	callhome := callhomeGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Callhome.Type(), &callhome)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create callhome, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created callhome resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("callhome")

	// Read the updated state back
	r.readCallhomeFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CallhomeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CallhomeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading callhome resource")

	r.readCallhomeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CallhomeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CallhomeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating callhome resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Emailaddress.Equal(state.Emailaddress) {
		tflog.Debug(ctx, fmt.Sprintf("emailaddress has changed for callhome"))
		hasChange = true
	}
	if !data.Hbcustominterval.Equal(state.Hbcustominterval) {
		tflog.Debug(ctx, fmt.Sprintf("hbcustominterval has changed for callhome"))
		hasChange = true
	}
	if !data.Ipaddress.Equal(state.Ipaddress) {
		tflog.Debug(ctx, fmt.Sprintf("ipaddress has changed for callhome"))
		hasChange = true
	}
	if !data.Mode.Equal(state.Mode) {
		tflog.Debug(ctx, fmt.Sprintf("mode has changed for callhome"))
		hasChange = true
	}
	if !data.Port.Equal(state.Port) {
		tflog.Debug(ctx, fmt.Sprintf("port has changed for callhome"))
		hasChange = true
	}
	if !data.Proxyauthservice.Equal(state.Proxyauthservice) {
		tflog.Debug(ctx, fmt.Sprintf("proxyauthservice has changed for callhome"))
		hasChange = true
	}
	if !data.Proxymode.Equal(state.Proxymode) {
		tflog.Debug(ctx, fmt.Sprintf("proxymode has changed for callhome"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		callhome := callhomeGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Callhome.Type(), &callhome)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update callhome, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated callhome resource")
	} else {
		tflog.Debug(ctx, "No changes detected for callhome resource, skipping update")
	}

	// Read the updated state back
	r.readCallhomeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CallhomeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CallhomeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting callhome resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed callhome from Terraform state")
}

// Helper function to read callhome data from API
func (r *CallhomeResource) readCallhomeFromApi(ctx context.Context, data *CallhomeResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Callhome.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read callhome, got error: %s", err))
		return
	}

	callhomeSetAttrFromGet(ctx, data, getResponseData)

}
