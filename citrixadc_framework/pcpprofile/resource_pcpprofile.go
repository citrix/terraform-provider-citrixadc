package pcpprofile

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
var _ resource.Resource = &PcpprofileResource{}
var _ resource.ResourceWithConfigure = (*PcpprofileResource)(nil)
var _ resource.ResourceWithImportState = (*PcpprofileResource)(nil)

func NewPcpprofileResource() resource.Resource {
	return &PcpprofileResource{}
}

// PcpprofileResource defines the resource implementation.
type PcpprofileResource struct {
	client *service.NitroClient
}

func (r *PcpprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PcpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pcpprofile"
}

func (r *PcpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PcpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PcpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating pcpprofile resource")

	pcpprofile := pcpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	pcpprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Pcpprofile.Type(), pcpprofileName, &pcpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create pcpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created pcpprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(pcpprofileName)

	// Read the updated state back
	if !r.readPcpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "pcpprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PcpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PcpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading pcpprofile resource")

	found := r.readPcpprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *PcpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state PcpprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating pcpprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Announcemulticount.Equal(state.Announcemulticount) {
		tflog.Debug(ctx, "announcemulticount has changed for pcpprofile")
		hasChange = true
	}
	if !data.Mapping.Equal(state.Mapping) {
		tflog.Debug(ctx, "mapping has changed for pcpprofile")
		hasChange = true
	}
	if !data.Maxmaplife.Equal(state.Maxmaplife) {
		tflog.Debug(ctx, "maxmaplife has changed for pcpprofile")
		hasChange = true
	}
	if !data.Minmaplife.Equal(state.Minmaplife) {
		tflog.Debug(ctx, "minmaplife has changed for pcpprofile")
		hasChange = true
	}
	if !data.Peer.Equal(state.Peer) {
		tflog.Debug(ctx, "peer has changed for pcpprofile")
		hasChange = true
	}
	if !data.Thirdparty.Equal(state.Thirdparty) {
		tflog.Debug(ctx, "thirdparty has changed for pcpprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		pcpprofile := pcpprofileGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		pcpprofileName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Pcpprofile.Type(), pcpprofileName, &pcpprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update pcpprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated pcpprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for pcpprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readPcpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "pcpprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PcpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PcpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting pcpprofile resource")
	// Named resource - delete using DeleteResource
	pcpprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Pcpprofile.Type(), pcpprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete pcpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted pcpprofile resource")
}

// Helper function to read pcpprofile data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *PcpprofileResource) readPcpprofileFromApi(ctx context.Context, data *PcpprofileResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	pcpprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Pcpprofile.Type(), pcpprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read pcpprofile, got error: %s", err))
		return false
	}

	pcpprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
