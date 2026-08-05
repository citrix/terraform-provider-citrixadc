package icaaccessprofile

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
var _ resource.Resource = &IcaaccessprofileResource{}
var _ resource.ResourceWithConfigure = (*IcaaccessprofileResource)(nil)
var _ resource.ResourceWithImportState = (*IcaaccessprofileResource)(nil)

func NewIcaaccessprofileResource() resource.Resource {
	return &IcaaccessprofileResource{}
}

// IcaaccessprofileResource defines the resource implementation.
type IcaaccessprofileResource struct {
	client *service.NitroClient
}

func (r *IcaaccessprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IcaaccessprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_icaaccessprofile"
}

func (r *IcaaccessprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IcaaccessprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IcaaccessprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating icaaccessprofile resource")

	icaaccessprofile := icaaccessprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource (POST /nitro/v1/config/icaaccessprofile)
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Icaaccessprofile.Type(), name_value, &icaaccessprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create icaaccessprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created icaaccessprofile resource")

	// Set ID for the resource before reading state (single unique attribute)
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	if !r.readIcaaccessprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "icaaccessprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcaaccessprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IcaaccessprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading icaaccessprofile resource")

	found := r.readIcaaccessprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *IcaaccessprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state IcaaccessprofileResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating icaaccessprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Clientaudioredirection.Equal(state.Clientaudioredirection) {
		tflog.Debug(ctx, "clientaudioredirection has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Clientclipboardredirection.Equal(state.Clientclipboardredirection) {
		tflog.Debug(ctx, "clientclipboardredirection has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Clientcomportredirection.Equal(state.Clientcomportredirection) {
		tflog.Debug(ctx, "clientcomportredirection has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Clientdriveredirection.Equal(state.Clientdriveredirection) {
		tflog.Debug(ctx, "clientdriveredirection has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Clientprinterredirection.Equal(state.Clientprinterredirection) {
		tflog.Debug(ctx, "clientprinterredirection has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Clienttwaindeviceredirection.Equal(state.Clienttwaindeviceredirection) {
		tflog.Debug(ctx, "clienttwaindeviceredirection has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Clientusbdriveredirection.Equal(state.Clientusbdriveredirection) {
		tflog.Debug(ctx, "clientusbdriveredirection has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Connectclientlptports.Equal(state.Connectclientlptports) {
		tflog.Debug(ctx, "connectclientlptports has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Draganddrop.Equal(state.Draganddrop) {
		tflog.Debug(ctx, "draganddrop has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Fido2redirection.Equal(state.Fido2redirection) {
		tflog.Debug(ctx, "fido2redirection has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Localremotedatasharing.Equal(state.Localremotedatasharing) {
		tflog.Debug(ctx, "localremotedatasharing has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Multistream.Equal(state.Multistream) {
		tflog.Debug(ctx, "multistream has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Smartcardredirection.Equal(state.Smartcardredirection) {
		tflog.Debug(ctx, "smartcardredirection has changed for icaaccessprofile")
		hasChange = true
	}
	if !data.Wiaredirection.Equal(state.Wiaredirection) {
		tflog.Debug(ctx, "wiaredirection has changed for icaaccessprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		icaaccessprofile := icaaccessprofileGetThePayloadFromthePlan(ctx, &data)
		// Named resource update is a PUT to the collection URL with name in the body
		err := r.client.UpdateUnnamedResource(service.Icaaccessprofile.Type(), &icaaccessprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update icaaccessprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated icaaccessprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for icaaccessprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readIcaaccessprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "icaaccessprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcaaccessprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IcaaccessprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting icaaccessprofile resource")
	// Named resource - delete using DeleteResource (DELETE /nitro/v1/config/icaaccessprofile/<name>)
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Icaaccessprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete icaaccessprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted icaaccessprofile resource")
}

// Helper function to read icaaccessprofile data from API
func (r *IcaaccessprofileResource) readIcaaccessprofileFromApi(ctx context.Context, data *IcaaccessprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	name_value := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Icaaccessprofile.Type(), name_value)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read icaaccessprofile, got error: %s", err))
		return false
	}

	icaaccessprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
