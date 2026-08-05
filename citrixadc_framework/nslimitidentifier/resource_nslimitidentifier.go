package nslimitidentifier

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
var _ resource.Resource = &NslimitidentifierResource{}
var _ resource.ResourceWithConfigure = (*NslimitidentifierResource)(nil)
var _ resource.ResourceWithImportState = (*NslimitidentifierResource)(nil)

func NewNslimitidentifierResource() resource.Resource {
	return &NslimitidentifierResource{}
}

// NslimitidentifierResource defines the resource implementation.
type NslimitidentifierResource struct {
	client *service.NitroClient
}

func (r *NslimitidentifierResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslimitidentifierResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslimitidentifier"
}

func (r *NslimitidentifierResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslimitidentifierResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslimitidentifierResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslimitidentifier resource")

	nslimitidentifier := nslimitidentifierGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	limitidentifier_value := data.Limitidentifier.ValueString()
	_, err := r.client.AddResource(service.Nslimitidentifier.Type(), limitidentifier_value, &nslimitidentifier)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nslimitidentifier, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nslimitidentifier resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(limitidentifier))
	data.Id = types.StringValue(limitidentifier_value)

	// Read the updated state back
	if !r.readNslimitidentifierFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nslimitidentifier not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitidentifierResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslimitidentifierResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nslimitidentifier resource")

	found := r.readNslimitidentifierFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NslimitidentifierResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NslimitidentifierResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nslimitidentifier resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Limittype.Equal(state.Limittype) {
		tflog.Debug(ctx, "limittype has changed for nslimitidentifier")
		hasChange = true
	}
	if !data.Maxbandwidth.Equal(state.Maxbandwidth) {
		tflog.Debug(ctx, "maxbandwidth has changed for nslimitidentifier")
		hasChange = true
	}
	if !data.Mode.Equal(state.Mode) {
		tflog.Debug(ctx, "mode has changed for nslimitidentifier")
		hasChange = true
	}
	if !data.Selectorname.Equal(state.Selectorname) {
		tflog.Debug(ctx, "selectorname has changed for nslimitidentifier")
		hasChange = true
	}
	if !data.Threshold.Equal(state.Threshold) {
		tflog.Debug(ctx, "threshold has changed for nslimitidentifier")
		hasChange = true
	}
	if !data.Timeslice.Equal(state.Timeslice) {
		tflog.Debug(ctx, "timeslice has changed for nslimitidentifier")
		hasChange = true
	}
	if !data.Trapsintimeslice.Equal(state.Trapsintimeslice) {
		tflog.Debug(ctx, "trapsintimeslice has changed for nslimitidentifier")
		hasChange = true
	}

	if hasChange {
		nslimitidentifier := nslimitidentifierGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		limitidentifier_value := data.Limitidentifier.ValueString()
		_, err := r.client.UpdateResource(service.Nslimitidentifier.Type(), limitidentifier_value, &nslimitidentifier)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nslimitidentifier, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated nslimitidentifier resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nslimitidentifier resource, skipping update")
	}

	// Read the updated state back
	if !r.readNslimitidentifierFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nslimitidentifier not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitidentifierResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslimitidentifierResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nslimitidentifier resource")

	// Named resource - delete using DeleteResource (key off the live ID)
	limitidentifier_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nslimitidentifier.Type(), limitidentifier_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nslimitidentifier, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nslimitidentifier resource")
}

// Helper function to read nslimitidentifier data from API
func (r *NslimitidentifierResource) readNslimitidentifierFromApi(ctx context.Context, data *NslimitidentifierResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	limitidentifier_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nslimitidentifier.Type(), limitidentifier_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nslimitidentifier, got error: %s", err))
		return false
	}

	nslimitidentifierSetAttrFromGet(ctx, data, getResponseData)

	return true
}
