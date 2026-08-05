package l2param

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
var _ resource.Resource = &L2paramResource{}
var _ resource.ResourceWithConfigure = (*L2paramResource)(nil)
var _ resource.ResourceWithImportState = (*L2paramResource)(nil)

func NewL2paramResource() resource.Resource {
	return &L2paramResource{}
}

// L2paramResource defines the resource implementation.
type L2paramResource struct {
	client *service.NitroClient
}

func (r *L2paramResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *L2paramResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l2param"
}

func (r *L2paramResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *L2paramResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data L2paramResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating l2param resource")

	// Create API request body from the model
	l2param := l2paramGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.L2param.Type(), &l2param)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create l2param, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created l2param resource")

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("l2param-config")

	// Read the updated state back
	if !r.readL2paramFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "l2param not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L2paramResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data L2paramResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading l2param resource")

	found := r.readL2paramFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *L2paramResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state L2paramResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating l2param resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Bdggrpproxyarp.Equal(state.Bdggrpproxyarp) {
		tflog.Debug(ctx, "bdggrpproxyarp has changed for l2param")
		hasChange = true
	}
	if !data.Bdgsetting.Equal(state.Bdgsetting) {
		tflog.Debug(ctx, "bdgsetting has changed for l2param")
		hasChange = true
	}
	if !data.Bridgeagetimeout.Equal(state.Bridgeagetimeout) {
		tflog.Debug(ctx, "bridgeagetimeout has changed for l2param")
		hasChange = true
	}
	if !data.Garponvridintf.Equal(state.Garponvridintf) {
		tflog.Debug(ctx, "garponvridintf has changed for l2param")
		hasChange = true
	}
	if !data.Garpreply.Equal(state.Garpreply) {
		tflog.Debug(ctx, "garpreply has changed for l2param")
		hasChange = true
	}
	if !data.Macmodefwdmypkt.Equal(state.Macmodefwdmypkt) {
		tflog.Debug(ctx, "macmodefwdmypkt has changed for l2param")
		hasChange = true
	}
	if !data.Maxbridgecollision.Equal(state.Maxbridgecollision) {
		tflog.Debug(ctx, "maxbridgecollision has changed for l2param")
		hasChange = true
	}
	if !data.Mbfinstlearning.Equal(state.Mbfinstlearning) {
		tflog.Debug(ctx, "mbfinstlearning has changed for l2param")
		hasChange = true
	}
	if !data.Mbfpeermacupdate.Equal(state.Mbfpeermacupdate) {
		tflog.Debug(ctx, "mbfpeermacupdate has changed for l2param")
		hasChange = true
	}
	if !data.Proxyarp.Equal(state.Proxyarp) {
		tflog.Debug(ctx, "proxyarp has changed for l2param")
		hasChange = true
	}
	if !data.Returntoethernetsender.Equal(state.Returntoethernetsender) {
		tflog.Debug(ctx, "returntoethernetsender has changed for l2param")
		hasChange = true
	}
	if !data.Rstintfonhafo.Equal(state.Rstintfonhafo) {
		tflog.Debug(ctx, "rstintfonhafo has changed for l2param")
		hasChange = true
	}
	if !data.Skipproxyingbsdtraffic.Equal(state.Skipproxyingbsdtraffic) {
		tflog.Debug(ctx, "skipproxyingbsdtraffic has changed for l2param")
		hasChange = true
	}
	if !data.Stopmacmoveupdate.Equal(state.Stopmacmoveupdate) {
		tflog.Debug(ctx, "stopmacmoveupdate has changed for l2param")
		hasChange = true
	}
	if !data.Usemymac.Equal(state.Usemymac) {
		tflog.Debug(ctx, "usemymac has changed for l2param")
		hasChange = true
	}
	if !data.Usenetprofilebsdtraffic.Equal(state.Usenetprofilebsdtraffic) {
		tflog.Debug(ctx, "usenetprofilebsdtraffic has changed for l2param")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		l2param := l2paramGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.L2param.Type(), &l2param)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update l2param, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated l2param resource")
	} else {
		tflog.Debug(ctx, "No changes detected for l2param resource, skipping update")
	}

	// Read the updated state back
	if !r.readL2paramFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "l2param not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L2paramResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data L2paramResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting l2param resource")

	// l2param does not support a DELETE operation (singleton global configuration).
	// Mirror the SDK v2 behavior: simply remove the resource from Terraform state.
	tflog.Trace(ctx, "Deleted l2param resource from state")
}

// Helper function to read l2param data from API
func (r *L2paramResource) readL2paramFromApi(ctx context.Context, data *L2paramResourceModel, diags *diag.Diagnostics) bool {
	// Case 1: Simple find without ID (singleton)
	getResponseData, err := r.client.FindResource(service.L2param.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read l2param, got error: %s", err))
		return false
	}

	l2paramSetAttrFromGet(ctx, data, getResponseData)

	return true
}
