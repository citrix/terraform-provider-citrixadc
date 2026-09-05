package nsservicefunction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsservicefunctionResource{}
var _ resource.ResourceWithConfigure = (*NsservicefunctionResource)(nil)
var _ resource.ResourceWithImportState = (*NsservicefunctionResource)(nil)

func NewNsservicefunctionResource() resource.Resource {
	return &NsservicefunctionResource{}
}

// NsservicefunctionResource defines the resource implementation.
type NsservicefunctionResource struct {
	client *service.NitroClient
}

func (r *NsservicefunctionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsservicefunctionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsservicefunction"
}

func (r *NsservicefunctionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsservicefunctionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsservicefunctionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsservicefunction resource")

	// Create API request body from the model
	nsservicefunction := nsservicefunctionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	servicefunctionname_value := data.Servicefunctionname.ValueString()
	_, err := r.client.AddResource(service.Nsservicefunction.Type(), servicefunctionname_value, &nsservicefunction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsservicefunction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsservicefunction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Servicefunctionname.ValueString()))

	// Read the updated state back
	if !r.readNsservicefunctionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsservicefunction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicefunctionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsservicefunctionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsservicefunction resource")

	found := r.readNsservicefunctionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicefunctionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsservicefunctionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsservicefunction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Ingressvlan.Equal(state.Ingressvlan) {
		tflog.Debug(ctx, "ingressvlan has changed for nsservicefunction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		nsservicefunction := nsservicefunctionGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		servicefunctionname_value := data.Servicefunctionname.ValueString()
		_, err := r.client.UpdateResource(service.Nsservicefunction.Type(), servicefunctionname_value, &nsservicefunction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsservicefunction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsservicefunction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsservicefunction resource, skipping update")
	}

	// Read the updated state back
	if !r.readNsservicefunctionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsservicefunction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicefunctionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsservicefunctionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsservicefunction resource")

	// Named resource - delete using DeleteResource
	servicefunctionname_value := data.Servicefunctionname.ValueString()
	err := r.client.DeleteResource(service.Nsservicefunction.Type(), servicefunctionname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsservicefunction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsservicefunction resource")
}

// Helper function to read nsservicefunction data from API
func (r *NsservicefunctionResource) readNsservicefunctionFromApi(ctx context.Context, data *NsservicefunctionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	servicefunctionname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nsservicefunction.Type(), servicefunctionname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsservicefunction, got error: %s", err))
		return false
	}

	nsservicefunctionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
