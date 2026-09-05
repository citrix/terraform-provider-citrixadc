package aaagroup

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
var _ resource.Resource = &AaagroupResource{}
var _ resource.ResourceWithConfigure = (*AaagroupResource)(nil)
var _ resource.ResourceWithImportState = (*AaagroupResource)(nil)

func NewAaagroupResource() resource.Resource {
	return &AaagroupResource{}
}

// AaagroupResource defines the resource implementation.
type AaagroupResource struct {
	client *service.NitroClient
}

func (r *AaagroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaagroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup"
}

func (r *AaagroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaagroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaagroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaagroup resource")
	// Get payload from plan
	aaagroup := aaagroupGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	groupname_value := data.Groupname.ValueString()
	_, err := r.client.AddResource(service.Aaagroup.Type(), groupname_value, &aaagroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaagroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaagroup resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Groupname.ValueString()))

	// Read the updated state back
	if !r.readAaagroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "aaagroup not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaagroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaagroup resource")

	found := r.readAaagroupFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AaagroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AaagroupResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaagroup resource")

	// aaagroup has no NITRO-updatable attributes: groupname, loggedin and weight
	// are all ForceNew/RequiresReplace (matching the SDK v2 resource, which has no
	// update path). Any attribute change triggers a replace, so there is nothing to
	// push here - just refresh the state from the appliance.
	tflog.Trace(ctx, "No updatable attributes for aaagroup; refreshing state")

	// Read the updated state back
	if !r.readAaagroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "aaagroup not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaagroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaagroup resource")
	// Named resource - delete using DeleteResource
	groupname_value := data.Groupname.ValueString()
	err := r.client.DeleteResource(service.Aaagroup.Type(), groupname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaagroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaagroup resource")
}

// Helper function to read aaagroup data from API
func (r *AaagroupResource) readAaagroupFromApi(ctx context.Context, data *AaagroupResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	groupname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Aaagroup.Type(), groupname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup, got error: %s", err))
		return false
	}

	aaagroupSetAttrFromGet(ctx, data, getResponseData)

	return true
}
