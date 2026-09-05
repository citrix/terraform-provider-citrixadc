package nsassignment

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
var _ resource.Resource = &NsassignmentResource{}
var _ resource.ResourceWithConfigure = (*NsassignmentResource)(nil)
var _ resource.ResourceWithImportState = (*NsassignmentResource)(nil)

func NewNsassignmentResource() resource.Resource {
	return &NsassignmentResource{}
}

// NsassignmentResource defines the resource implementation.
type NsassignmentResource struct {
	client *service.NitroClient
}

func (r *NsassignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsassignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsassignment"
}

func (r *NsassignmentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsassignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsassignmentResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsassignment resource")

	// Create API request body from the model
	nsassignment := nsassignmentGetThePayloadFromtheConfig(ctx, &data)

	// Make API call - named resource, use AddResource
	nsassignmentName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nsassignment.Type(), nsassignmentName, &nsassignment)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsassignment, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsassignment resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(nsassignmentName)

	// Read the updated state back
	if !r.readNsassignmentFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsassignment not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsassignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsassignmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsassignment resource")

	found := r.readNsassignmentFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NsassignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NsassignmentResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to be unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsassignment resource")

	// Check if there are any changes in the updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Add.Equal(state.Add) {
		tflog.Debug(ctx, "add has changed for nsassignment")
		hasChange = true
	}
	if !data.Append.Equal(state.Append) {
		tflog.Debug(ctx, "append has changed for nsassignment")
		hasChange = true
	}
	if !data.Clear.Equal(state.Clear) {
		tflog.Debug(ctx, "clear has changed for nsassignment")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for nsassignment")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Set.Equal(state.Set) {
		tflog.Debug(ctx, "set has changed for nsassignment")
		hasChange = true
	}
	if !data.Sub.Equal(state.Sub) {
		tflog.Debug(ctx, "sub has changed for nsassignment")
		hasChange = true
	}
	if !data.Variable.Equal(state.Variable) {
		tflog.Debug(ctx, "variable has changed for nsassignment")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		nsassignment := nsassignmentGetThePayloadFromtheConfig(ctx, &data)
		// Make API call - named resource, use UpdateResource
		nsassignmentName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Nsassignment.Type(), nsassignmentName, &nsassignment)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsassignment %s, got error: %s", nsassignmentName, err))
			return
		}

		tflog.Trace(ctx, "Updated nsassignment resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsassignment resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Nsassignment.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nsassignment attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNsassignmentFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsassignment not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsassignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsassignmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsassignment resource")

	// Named resource - delete using DeleteResource
	nsassignmentName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nsassignment.Type(), nsassignmentName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsassignment %s, got error: %s", nsassignmentName, err))
		return
	}

	tflog.Trace(ctx, "Deleted nsassignment resource")
}

// Helper function to read nsassignment data from API.
// Returns false (without adding a diagnostic) when the resource does not exist,
// so callers can remove it from state.
func (r *NsassignmentResource) readNsassignmentFromApi(ctx context.Context, data *NsassignmentResourceModel, diags *diag.Diagnostics) bool {
	// Named resource - ID is the plain name value
	nsassignmentName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nsassignment.Type(), nsassignmentName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsassignment %s, got error: %s", nsassignmentName, err))
		return false
	}

	nsassignmentSetAttrFromGet(ctx, data, getResponseData)

	return true
}
