package rdpclientprofile

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
var _ resource.Resource = &RdpclientprofileResource{}
var _ resource.ResourceWithConfigure = (*RdpclientprofileResource)(nil)
var _ resource.ResourceWithImportState = (*RdpclientprofileResource)(nil)

func NewRdpclientprofileResource() resource.Resource {
	return &RdpclientprofileResource{}
}

// RdpclientprofileResource defines the resource implementation.
type RdpclientprofileResource struct {
	client *service.NitroClient
}

func (r *RdpclientprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RdpclientprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rdpclientprofile"
}

func (r *RdpclientprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RdpclientprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config RdpclientprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rdpclientprofile resource")
	// Get payload from plan (regular attributes)
	rdpclientprofile := rdpclientprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	rdpclientprofileGetThePayloadFromtheConfig(ctx, &config, &rdpclientprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Rdpclientprofile.Type(), name_value, &rdpclientprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rdpclientprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rdpclientprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readRdpclientprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RdpclientprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RdpclientprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rdpclientprofile resource")

	r.readRdpclientprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RdpclientprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state RdpclientprofileResourceModel

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

	tflog.Debug(ctx, "Updating rdpclientprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Addusernameinrdpfile.Equal(state.Addusernameinrdpfile) {
		tflog.Debug(ctx, fmt.Sprintf("addusernameinrdpfile has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Audiocapturemode.Equal(state.Audiocapturemode) {
		tflog.Debug(ctx, fmt.Sprintf("audiocapturemode has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Keyboardhook.Equal(state.Keyboardhook) {
		tflog.Debug(ctx, fmt.Sprintf("keyboardhook has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Multimonitorsupport.Equal(state.Multimonitorsupport) {
		tflog.Debug(ctx, fmt.Sprintf("multimonitorsupport has changed for rdpclientprofile"))
		hasChange = true
	}
	// Check secret attribute psk or its version tracker
	if !data.Psk.Equal(state.Psk) {
		tflog.Debug(ctx, fmt.Sprintf("psk has changed for rdpclientprofile"))
		hasChange = true
	} else if !data.PskWoVersion.Equal(state.PskWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("psk_wo_version has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Randomizerdpfilename.Equal(state.Randomizerdpfilename) {
		tflog.Debug(ctx, fmt.Sprintf("randomizerdpfilename has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Rdpcookievalidity.Equal(state.Rdpcookievalidity) {
		tflog.Debug(ctx, fmt.Sprintf("rdpcookievalidity has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Rdpcustomparams.Equal(state.Rdpcustomparams) {
		tflog.Debug(ctx, fmt.Sprintf("rdpcustomparams has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Rdpfilename.Equal(state.Rdpfilename) {
		tflog.Debug(ctx, fmt.Sprintf("rdpfilename has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Rdphost.Equal(state.Rdphost) {
		tflog.Debug(ctx, fmt.Sprintf("rdphost has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Rdplinkattribute.Equal(state.Rdplinkattribute) {
		tflog.Debug(ctx, fmt.Sprintf("rdplinkattribute has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Rdplistener.Equal(state.Rdplistener) {
		tflog.Debug(ctx, fmt.Sprintf("rdplistener has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Rdpurloverride.Equal(state.Rdpurloverride) {
		tflog.Debug(ctx, fmt.Sprintf("rdpurloverride has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Rdpvalidateclientip.Equal(state.Rdpvalidateclientip) {
		tflog.Debug(ctx, fmt.Sprintf("rdpvalidateclientip has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Redirectclipboard.Equal(state.Redirectclipboard) {
		tflog.Debug(ctx, fmt.Sprintf("redirectclipboard has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Redirectcomports.Equal(state.Redirectcomports) {
		tflog.Debug(ctx, fmt.Sprintf("redirectcomports has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Redirectdrives.Equal(state.Redirectdrives) {
		tflog.Debug(ctx, fmt.Sprintf("redirectdrives has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Redirectpnpdevices.Equal(state.Redirectpnpdevices) {
		tflog.Debug(ctx, fmt.Sprintf("redirectpnpdevices has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Redirectprinters.Equal(state.Redirectprinters) {
		tflog.Debug(ctx, fmt.Sprintf("redirectprinters has changed for rdpclientprofile"))
		hasChange = true
	}
	if !data.Videoplaybackmode.Equal(state.Videoplaybackmode) {
		tflog.Debug(ctx, fmt.Sprintf("videoplaybackmode has changed for rdpclientprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		rdpclientprofile := rdpclientprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		rdpclientprofileGetThePayloadFromtheConfig(ctx, &config, &rdpclientprofile)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Rdpclientprofile.Type(), name_value, &rdpclientprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rdpclientprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated rdpclientprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for rdpclientprofile resource, skipping update")
	}

	// Read the updated state back
	r.readRdpclientprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RdpclientprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RdpclientprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rdpclientprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Rdpclientprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete rdpclientprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted rdpclientprofile resource")
}

// Helper function to read rdpclientprofile data from API
func (r *RdpclientprofileResource) readRdpclientprofileFromApi(ctx context.Context, data *RdpclientprofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Rdpclientprofile.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rdpclientprofile, got error: %s", err))
		return
	}

	rdpclientprofileSetAttrFromGet(ctx, data, getResponseData)

}
