package nspartition

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
var _ resource.Resource = &NspartitionResource{}
var _ resource.ResourceWithConfigure = (*NspartitionResource)(nil)
var _ resource.ResourceWithImportState = (*NspartitionResource)(nil)

func NewNspartitionResource() resource.Resource {
	return &NspartitionResource{}
}

// NspartitionResource defines the resource implementation.
type NspartitionResource struct {
	client *service.NitroClient
}

func (r *NspartitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NspartitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nspartition"
}

func (r *NspartitionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NspartitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NspartitionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nspartition resource")

	nspartition := nspartitionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	partitionname_value := data.Partitionname.ValueString()
	_, err := r.client.AddResource(service.Nspartition.Type(), partitionname_value, &nspartition)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nspartition, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nspartition resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Partitionname.ValueString()))

	// Read the updated state back
	if !r.readNspartitionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nspartition not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NspartitionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nspartition resource")

	found := r.readNspartitionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NspartitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NspartitionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nspartition resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Force.Equal(state.Force) {
		tflog.Debug(ctx, "force has changed for nspartition")
		hasChange = true
	}
	if !data.Maxbandwidth.Equal(state.Maxbandwidth) {
		tflog.Debug(ctx, "maxbandwidth has changed for nspartition")
		hasChange = true
	}
	if !data.Maxconn.Equal(state.Maxconn) {
		tflog.Debug(ctx, "maxconn has changed for nspartition")
		hasChange = true
	}
	if !data.Maxmemlimit.Equal(state.Maxmemlimit) {
		tflog.Debug(ctx, "maxmemlimit has changed for nspartition")
		hasChange = true
	}
	if !data.Minbandwidth.Equal(state.Minbandwidth) {
		tflog.Debug(ctx, "minbandwidth has changed for nspartition")
		hasChange = true
	}
	if !data.Partitionmac.Equal(state.Partitionmac) {
		tflog.Debug(ctx, "partitionmac has changed for nspartition")
		hasChange = true
	}
	if !data.Save.Equal(state.Save) {
		tflog.Debug(ctx, "save has changed for nspartition")
		hasChange = true
	}

	if hasChange {
		nspartition := nspartitionGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		partitionname_value := data.Partitionname.ValueString()
		_, err := r.client.UpdateResource(service.Nspartition.Type(), partitionname_value, &nspartition)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nspartition, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nspartition resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nspartition resource, skipping update")
	}

	// Read the updated state back
	if !r.readNspartitionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nspartition not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NspartitionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nspartition resource")

	// Named resource - delete using DeleteResource
	partitionname_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nspartition.Type(), partitionname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nspartition, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nspartition resource")
}

// Helper function to read nspartition data from API.
// Returns false when the resource no longer exists on the appliance.
func (r *NspartitionResource) readNspartitionFromApi(ctx context.Context, data *NspartitionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain partitionname value
	partitionname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nspartition.Type(), partitionname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nspartition, got error: %s", err))
		return false
	}

	nspartitionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
