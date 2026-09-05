package nsservicepath

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
var _ resource.Resource = &NsservicepathResource{}
var _ resource.ResourceWithConfigure = (*NsservicepathResource)(nil)
var _ resource.ResourceWithImportState = (*NsservicepathResource)(nil)

func NewNsservicepathResource() resource.Resource {
	return &NsservicepathResource{}
}

// NsservicepathResource defines the resource implementation.
type NsservicepathResource struct {
	client *service.NitroClient
}

func (r *NsservicepathResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsservicepathResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsservicepath"
}

func (r *NsservicepathResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsservicepathResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsservicepathResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsservicepath resource")

	nsservicepath := nsservicepathGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	servicepathname_value := data.Servicepathname.ValueString()
	_, err := r.client.AddResource(service.Nsservicepath.Type(), servicepathname_value, &nsservicepath)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsservicepath, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsservicepath resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Servicepathname.ValueString()))

	// Read the updated state back
	if !r.readNsservicepathFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsservicepath not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicepathResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsservicepathResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsservicepath resource")

	found := r.readNsservicepathFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NsservicepathResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsservicepathResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsservicepath resource")

	// nsservicepath has a single writable attribute (servicepathname) which is
	// the key and is ForceNew (RequiresReplace). NITRO exposes no update
	// operation for this resource, so there is nothing to push here; any change
	// to servicepathname triggers a destroy/create instead. Re-read to keep
	// state consistent.
	if !r.readNsservicepathFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsservicepath not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicepathResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsservicepathResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsservicepath resource")
	// Named resource - delete using DeleteResource
	servicepathname_value := data.Servicepathname.ValueString()
	err := r.client.DeleteResource(service.Nsservicepath.Type(), servicepathname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsservicepath, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsservicepath resource")
}

// Helper function to read nsservicepath data from API
func (r *NsservicepathResource) readNsservicepathFromApi(ctx context.Context, data *NsservicepathResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	servicepathname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nsservicepath.Type(), servicepathname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsservicepath, got error: %s", err))
		return false
	}

	nsservicepathSetAttrFromGet(ctx, data, getResponseData)

	return true
}
