package ssldtlsprofile

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
var _ resource.Resource = &SsldtlsprofileResource{}
var _ resource.ResourceWithConfigure = (*SsldtlsprofileResource)(nil)
var _ resource.ResourceWithImportState = (*SsldtlsprofileResource)(nil)

func NewSsldtlsprofileResource() resource.Resource {
	return &SsldtlsprofileResource{}
}

// SsldtlsprofileResource defines the resource implementation.
type SsldtlsprofileResource struct {
	client *service.NitroClient
}

func (r *SsldtlsprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SsldtlsprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssldtlsprofile"
}

func (r *SsldtlsprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SsldtlsprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SsldtlsprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ssldtlsprofile resource")

	// Create API request body from the model
	ssldtlsprofile := ssldtlsprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	ssldtlsprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Ssldtlsprofile.Type(), ssldtlsprofileName, &ssldtlsprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ssldtlsprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ssldtlsprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", ssldtlsprofileName))

	// Read the updated state back
	if !r.readSsldtlsprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ssldtlsprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldtlsprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SsldtlsprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ssldtlsprofile resource")

	found := r.readSsldtlsprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SsldtlsprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SsldtlsprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating ssldtlsprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Helloverifyrequest.Equal(state.Helloverifyrequest) {
		tflog.Debug(ctx, "helloverifyrequest has changed for ssldtlsprofile")
		hasChange = true
	}
	if !data.Initialretrytimeout.Equal(state.Initialretrytimeout) {
		tflog.Debug(ctx, "initialretrytimeout has changed for ssldtlsprofile")
		hasChange = true
	}
	if !data.Maxbadmacignorecount.Equal(state.Maxbadmacignorecount) {
		tflog.Debug(ctx, "maxbadmacignorecount has changed for ssldtlsprofile")
		hasChange = true
	}
	if !data.Maxholdqlen.Equal(state.Maxholdqlen) {
		tflog.Debug(ctx, "maxholdqlen has changed for ssldtlsprofile")
		hasChange = true
	}
	if !data.Maxpacketsize.Equal(state.Maxpacketsize) {
		tflog.Debug(ctx, "maxpacketsize has changed for ssldtlsprofile")
		hasChange = true
	}
	if !data.Maxrecordsize.Equal(state.Maxrecordsize) {
		tflog.Debug(ctx, "maxrecordsize has changed for ssldtlsprofile")
		hasChange = true
	}
	if !data.Maxretrytime.Equal(state.Maxretrytime) {
		tflog.Debug(ctx, "maxretrytime has changed for ssldtlsprofile")
		hasChange = true
	}
	if !data.Pmtudiscovery.Equal(state.Pmtudiscovery) {
		tflog.Debug(ctx, "pmtudiscovery has changed for ssldtlsprofile")
		hasChange = true
	}
	if !data.Terminatesession.Equal(state.Terminatesession) {
		tflog.Debug(ctx, "terminatesession has changed for ssldtlsprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		ssldtlsprofile := ssldtlsprofileGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		ssldtlsprofileName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Ssldtlsprofile.Type(), ssldtlsprofileName, &ssldtlsprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ssldtlsprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated ssldtlsprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for ssldtlsprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readSsldtlsprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ssldtlsprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldtlsprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SsldtlsprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ssldtlsprofile resource")
	// Named resource - delete using DeleteResource
	ssldtlsprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Ssldtlsprofile.Type(), ssldtlsprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ssldtlsprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted ssldtlsprofile resource")
}

// Helper function to read ssldtlsprofile data from API
func (r *SsldtlsprofileResource) readSsldtlsprofileFromApi(ctx context.Context, data *SsldtlsprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	ssldtlsprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Ssldtlsprofile.Type(), ssldtlsprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ssldtlsprofile, got error: %s", err))
		return false
	}

	ssldtlsprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
