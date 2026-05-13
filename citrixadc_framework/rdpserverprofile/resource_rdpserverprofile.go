package rdpserverprofile

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
var _ resource.Resource = &RdpserverprofileResource{}
var _ resource.ResourceWithConfigure = (*RdpserverprofileResource)(nil)
var _ resource.ResourceWithImportState = (*RdpserverprofileResource)(nil)
var _ resource.ResourceWithValidateConfig = (*RdpserverprofileResource)(nil)

func NewRdpserverprofileResource() resource.Resource {
	return &RdpserverprofileResource{}
}

// RdpserverprofileResource defines the resource implementation.
type RdpserverprofileResource struct {
	client *service.NitroClient
}

func (r *RdpserverprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RdpserverprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rdpserverprofile"
}

func (r *RdpserverprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RdpserverprofileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data RdpserverprofileResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either psk or psk_wo is specified
	if data.Psk.IsNull() && data.PskWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("psk"),
			"Missing Required Attribute",
			"Either \"psk\" or \"psk_wo\" must be specified.",
		)
	}
}

func (r *RdpserverprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config RdpserverprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rdpserverprofile resource")
	// Get payload from plan (regular attributes)
	rdpserverprofile := rdpserverprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	rdpserverprofileGetThePayloadFromtheConfig(ctx, &config, &rdpserverprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Rdpserverprofile.Type(), name_value, &rdpserverprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rdpserverprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rdpserverprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readRdpserverprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RdpserverprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RdpserverprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rdpserverprofile resource")

	r.readRdpserverprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RdpserverprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state RdpserverprofileResourceModel

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

	tflog.Debug(ctx, "Updating rdpserverprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute psk or its version tracker
	if !data.Psk.Equal(state.Psk) {
		tflog.Debug(ctx, fmt.Sprintf("psk has changed for rdpserverprofile"))
		hasChange = true
	} else if !data.PskWoVersion.Equal(state.PskWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("psk_wo_version has changed for rdpserverprofile"))
		hasChange = true
	}
	if !data.Rdpip.Equal(state.Rdpip) {
		tflog.Debug(ctx, fmt.Sprintf("rdpip has changed for rdpserverprofile"))
		hasChange = true
	}
	if !data.Rdpport.Equal(state.Rdpport) {
		tflog.Debug(ctx, fmt.Sprintf("rdpport has changed for rdpserverprofile"))
		hasChange = true
	}
	if !data.Rdpredirection.Equal(state.Rdpredirection) {
		tflog.Debug(ctx, fmt.Sprintf("rdpredirection has changed for rdpserverprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		rdpserverprofile := rdpserverprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		rdpserverprofileGetThePayloadFromtheConfig(ctx, &config, &rdpserverprofile)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Rdpserverprofile.Type(), name_value, &rdpserverprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rdpserverprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated rdpserverprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for rdpserverprofile resource, skipping update")
	}

	// Read the updated state back
	r.readRdpserverprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RdpserverprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RdpserverprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rdpserverprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Rdpserverprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete rdpserverprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted rdpserverprofile resource")
}

// Helper function to read rdpserverprofile data from API
func (r *RdpserverprofileResource) readRdpserverprofileFromApi(ctx context.Context, data *RdpserverprofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Rdpserverprofile.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rdpserverprofile, got error: %s", err))
		return
	}

	rdpserverprofileSetAttrFromGet(ctx, data, getResponseData)

}
