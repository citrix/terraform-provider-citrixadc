package nshmackey

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
var _ resource.Resource = &NshmackeyResource{}
var _ resource.ResourceWithConfigure = (*NshmackeyResource)(nil)
var _ resource.ResourceWithImportState = (*NshmackeyResource)(nil)

func NewNshmackeyResource() resource.Resource {
	return &NshmackeyResource{}
}

// NshmackeyResource defines the resource implementation.
type NshmackeyResource struct {
	client *service.NitroClient
}

func (r *NshmackeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NshmackeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nshmackey"
}

func (r *NshmackeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NshmackeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config NshmackeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nshmackey resource")
	// Get payload from plan (regular attributes)
	nshmackey := nshmackeyGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	nshmackeyGetThePayloadFromtheConfig(ctx, &config, &nshmackey)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nshmackey.Type(), name_value, &nshmackey)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nshmackey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nshmackey resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readNshmackeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshmackeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NshmackeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nshmackey resource")

	r.readNshmackeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshmackeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NshmackeyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nshmackey resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, fmt.Sprintf("comment has changed for nshmackey"))
		hasChange = true
	}
	if !data.Digest.Equal(state.Digest) {
		tflog.Debug(ctx, fmt.Sprintf("digest has changed for nshmackey"))
		hasChange = true
	}
	// Check secret attribute keyvalue or its version tracker
	if !data.Keyvalue.Equal(state.Keyvalue) {
		tflog.Debug(ctx, fmt.Sprintf("keyvalue has changed for nshmackey"))
		hasChange = true
	} else if !data.KeyvalueWoVersion.Equal(state.KeyvalueWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("keyvalue_wo_version has changed for nshmackey"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		nshmackey := nshmackeyGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		nshmackeyGetThePayloadFromtheConfig(ctx, &config, &nshmackey)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Nshmackey.Type(), name_value, &nshmackey)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nshmackey, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nshmackey resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nshmackey resource, skipping update")
	}

	// Read the updated state back
	r.readNshmackeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshmackeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NshmackeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nshmackey resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Nshmackey.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nshmackey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nshmackey resource")
}

// Helper function to read nshmackey data from API
func (r *NshmackeyResource) readNshmackeyFromApi(ctx context.Context, data *NshmackeyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nshmackey.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nshmackey, got error: %s", err))
		return
	}

	nshmackeySetAttrFromGet(ctx, data, getResponseData)

}
