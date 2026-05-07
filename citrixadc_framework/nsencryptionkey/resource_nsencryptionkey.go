package nsencryptionkey

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
var _ resource.Resource = &NsencryptionkeyResource{}
var _ resource.ResourceWithConfigure = (*NsencryptionkeyResource)(nil)
var _ resource.ResourceWithImportState = (*NsencryptionkeyResource)(nil)

func NewNsencryptionkeyResource() resource.Resource {
	return &NsencryptionkeyResource{}
}

// NsencryptionkeyResource defines the resource implementation.
type NsencryptionkeyResource struct {
	client *service.NitroClient
}

func (r *NsencryptionkeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsencryptionkeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsencryptionkey"
}

func (r *NsencryptionkeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsencryptionkeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config NsencryptionkeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsencryptionkey resource")
	// Get payload from plan (regular attributes)
	nsencryptionkey := nsencryptionkeyGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	nsencryptionkeyGetThePayloadFromtheConfig(ctx, &config, &nsencryptionkey)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nsencryptionkey.Type(), name_value, &nsencryptionkey)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsencryptionkey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsencryptionkey resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readNsencryptionkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsencryptionkeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsencryptionkeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsencryptionkey resource")

	r.readNsencryptionkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsencryptionkeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NsencryptionkeyResourceModel

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

	tflog.Debug(ctx, "Updating nsencryptionkey resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, fmt.Sprintf("comment has changed for nsencryptionkey"))
		hasChange = true
	}
	if !data.Iv.Equal(state.Iv) {
		tflog.Debug(ctx, fmt.Sprintf("iv has changed for nsencryptionkey"))
		hasChange = true
	}
	// Check secret attribute keyvalue or its version tracker
	if !data.Keyvalue.Equal(state.Keyvalue) {
		tflog.Debug(ctx, fmt.Sprintf("keyvalue has changed for nsencryptionkey"))
		hasChange = true
	} else if !data.KeyvalueWoVersion.Equal(state.KeyvalueWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("keyvalue_wo_version has changed for nsencryptionkey"))
		hasChange = true
	}
	if !data.Method.Equal(state.Method) {
		tflog.Debug(ctx, fmt.Sprintf("method has changed for nsencryptionkey"))
		hasChange = true
	}
	if !data.Padding.Equal(state.Padding) {
		tflog.Debug(ctx, fmt.Sprintf("padding has changed for nsencryptionkey"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		nsencryptionkey := nsencryptionkeyGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		nsencryptionkeyGetThePayloadFromtheConfig(ctx, &config, &nsencryptionkey)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Nsencryptionkey.Type(), name_value, &nsencryptionkey)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsencryptionkey, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsencryptionkey resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsencryptionkey resource, skipping update")
	}

	// Read the updated state back
	r.readNsencryptionkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsencryptionkeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsencryptionkeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsencryptionkey resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Nsencryptionkey.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsencryptionkey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsencryptionkey resource")
}

// Helper function to read nsencryptionkey data from API
func (r *NsencryptionkeyResource) readNsencryptionkeyFromApi(ctx context.Context, data *NsencryptionkeyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nsencryptionkey.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsencryptionkey, got error: %s", err))
		return
	}

	nsencryptionkeySetAttrFromGet(ctx, data, getResponseData)

}
