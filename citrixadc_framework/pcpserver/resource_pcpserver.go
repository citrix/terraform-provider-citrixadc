package pcpserver

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
var _ resource.Resource = &PcpserverResource{}
var _ resource.ResourceWithConfigure = (*PcpserverResource)(nil)
var _ resource.ResourceWithImportState = (*PcpserverResource)(nil)

func NewPcpserverResource() resource.Resource {
	return &PcpserverResource{}
}

// PcpserverResource defines the resource implementation.
type PcpserverResource struct {
	client *service.NitroClient
}

func (r *PcpserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PcpserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pcpserver"
}

func (r *PcpserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PcpserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PcpserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating pcpserver resource")

	pcpserver := pcpserverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource (NITRO add is POST /config/pcpserver)
	pcpserverName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Pcpserver.Type(), pcpserverName, &pcpserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create pcpserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created pcpserver resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(pcpserverName)

	// Read the updated state back
	if !r.readPcpserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "pcpserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PcpserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PcpserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading pcpserver resource")

	found := r.readPcpserverFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *PcpserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state PcpserverResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating pcpserver resource")

	// Check if there are any changes in updateable attributes
	// (ipaddress and name are RequiresReplace and never reach Update)
	hasChange := false
	if !data.Pcpprofile.Equal(state.Pcpprofile) {
		tflog.Debug(ctx, "pcpprofile has changed for pcpserver")
		hasChange = true
	}
	if !data.Port.Equal(state.Port) {
		tflog.Debug(ctx, "port has changed for pcpserver")
		hasChange = true
	}

	if hasChange {
		// NITRO update is PUT /config/pcpserver (unnamed) with name in the body
		pcpserver := pcpserverGetTheUpdatablePayloadFromThePlan(ctx, &data)
		err := r.client.UpdateUnnamedResource(service.Pcpserver.Type(), &pcpserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update pcpserver, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated pcpserver resource")
	} else {
		tflog.Debug(ctx, "No changes detected for pcpserver resource, skipping update")
	}

	// Read the updated state back
	if !r.readPcpserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "pcpserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PcpserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PcpserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting pcpserver resource")

	// Named resource - delete using DeleteResource (DELETE /config/pcpserver/{name})
	pcpserverName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Pcpserver.Type(), pcpserverName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete pcpserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted pcpserver resource")
}

// Helper function to read pcpserver data from API. Returns false when the
// resource no longer exists on the ADC.
func (r *PcpserverResource) readPcpserverFromApi(ctx context.Context, data *PcpserverResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	pcpserverName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Pcpserver.Type(), pcpserverName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read pcpserver, got error: %s", err))
		return false
	}

	pcpserverSetAttrFromGet(ctx, data, getResponseData)

	return true
}
