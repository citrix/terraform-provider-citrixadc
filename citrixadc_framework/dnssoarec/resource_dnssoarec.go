package dnssoarec

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
var _ resource.Resource = &DnssoarecResource{}
var _ resource.ResourceWithConfigure = (*DnssoarecResource)(nil)
var _ resource.ResourceWithImportState = (*DnssoarecResource)(nil)

func NewDnssoarecResource() resource.Resource {
	return &DnssoarecResource{}
}

// DnssoarecResource defines the resource implementation.
type DnssoarecResource struct {
	client *service.NitroClient
}

func (r *DnssoarecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnssoarecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssoarec"
}

func (r *DnssoarecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnssoarecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnssoarecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnssoarec resource")

	dnssoarec := dnssoarecGetThePayloadFromthePlan(ctx, &data)

	// Named resource keyed on domain - use AddResource
	dnssoarecId := data.Domain.ValueString()
	_, err := r.client.AddResource(service.Dnssoarec.Type(), dnssoarecId, &dnssoarec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnssoarec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnssoarec resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(dnssoarecId)

	// Read the updated state back
	if !r.readDnssoarecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnssoarec not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssoarecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnssoarecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnssoarec resource")

	found := r.readDnssoarecFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnssoarecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state DnssoarecResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnssoarec resource")

	// Check if there are any changes in updateable attributes
	// (domain is the primary key and is RequiresReplace, so it never reaches Update)
	hasChange := false
	attributesToUnset := []string{}
	if !data.Contact.Equal(state.Contact) {
		tflog.Debug(ctx, "contact has changed for dnssoarec")
		hasChange = true
	}
	if !data.Ecssubnet.Equal(state.Ecssubnet) {
		tflog.Debug(ctx, "ecssubnet has changed for dnssoarec")
		hasChange = true
	}
	if !data.Expire.Equal(state.Expire) {
		tflog.Debug(ctx, "expire has changed for dnssoarec")
		if config.Expire.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "expire")
		} else {
			hasChange = true
		}
	}
	if !data.Minimum.Equal(state.Minimum) {
		tflog.Debug(ctx, "minimum has changed for dnssoarec")
		if config.Minimum.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "minimum")
		} else {
			hasChange = true
		}
	}
	if !data.Nodeid.Equal(state.Nodeid) {
		tflog.Debug(ctx, "nodeid has changed for dnssoarec")
		hasChange = true
	}
	if !data.Originserver.Equal(state.Originserver) {
		tflog.Debug(ctx, "originserver has changed for dnssoarec")
		hasChange = true
	}
	if !data.Refresh.Equal(state.Refresh) {
		tflog.Debug(ctx, "refresh has changed for dnssoarec")
		if config.Refresh.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "refresh")
		} else {
			hasChange = true
		}
	}
	if !data.Retry.Equal(state.Retry) {
		tflog.Debug(ctx, "retry has changed for dnssoarec")
		if config.Retry.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "retry")
		} else {
			hasChange = true
		}
	}
	if !data.Serial.Equal(state.Serial) {
		tflog.Debug(ctx, "serial has changed for dnssoarec")
		if config.Serial.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "serial")
		} else {
			hasChange = true
		}
	}
	if !data.Ttl.Equal(state.Ttl) {
		tflog.Debug(ctx, "ttl has changed for dnssoarec")
		if config.Ttl.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ttl")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		dnssoarec := dnssoarecGetThePayloadFromthePlan(ctx, &data)
		// Named resource keyed on domain - use UpdateResource
		dnssoarecId := data.Domain.ValueString()
		_, err := r.client.UpdateResource(service.Dnssoarec.Type(), dnssoarecId, &dnssoarec)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnssoarec, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dnssoarec resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnssoarec resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"domain": data.Domain.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Dnssoarec.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset dnssoarec attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readDnssoarecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnssoarec not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssoarecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnssoarecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnssoarec resource")
	// Named resource keyed on domain - delete using DeleteResource
	dnssoarecId := data.Id.ValueString()
	err := r.client.DeleteResource(service.Dnssoarec.Type(), dnssoarecId)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnssoarec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnssoarec resource")
}

// Helper function to read dnssoarec data from API. Returns false when the
// resource is not found on the ADC.
func (r *DnssoarecResource) readDnssoarecFromApi(ctx context.Context, data *DnssoarecResourceModel, diags *diag.Diagnostics) bool {
	// Named resource keyed on domain - the ID is the plain domain value
	dnssoarecName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Dnssoarec.Type(), dnssoarecName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnssoarec, got error: %s", err))
		return false
	}

	dnssoarecSetAttrFromGet(ctx, data, getResponseData)

	return true
}
