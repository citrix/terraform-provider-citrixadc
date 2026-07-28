package reputationsettings

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
var _ resource.Resource = &ReputationsettingsResource{}
var _ resource.ResourceWithConfigure = (*ReputationsettingsResource)(nil)
var _ resource.ResourceWithImportState = (*ReputationsettingsResource)(nil)

func NewReputationsettingsResource() resource.Resource {
	return &ReputationsettingsResource{}
}

// ReputationsettingsResource defines the resource implementation.
type ReputationsettingsResource struct {
	client *service.NitroClient
}

func (r *ReputationsettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ReputationsettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_reputationsettings"
}

func (r *ReputationsettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ReputationsettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config ReputationsettingsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating reputationsettings resource")
	// Get payload from plan (regular attributes)
	reputationsettings := reputationsettingsGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	reputationsettingsGetThePayloadFromtheConfig(ctx, &config, &reputationsettings)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Reputationsettings.Type(), &reputationsettings)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create reputationsettings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created reputationsettings resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("reputationsettings-config")

	// Read the updated state back
	if !r.readReputationsettingsFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "reputationsettings not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ReputationsettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ReputationsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading reputationsettings resource")

	found := r.readReputationsettingsFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ReputationsettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ReputationsettingsResourceModel

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

	tflog.Debug(ctx, "Updating reputationsettings resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute proxypassword or its version tracker
	if !data.Proxypassword.Equal(state.Proxypassword) {
		tflog.Debug(ctx, fmt.Sprintf("proxypassword has changed for reputationsettings"))
		hasChange = true
	} else if !data.ProxypasswordWoVersion.Equal(state.ProxypasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("proxypassword_wo_version has changed for reputationsettings"))
		hasChange = true
	}
	if !data.Proxyport.Equal(state.Proxyport) {
		tflog.Debug(ctx, fmt.Sprintf("proxyport has changed for reputationsettings"))
		hasChange = true
	}
	if !data.Proxyserver.Equal(state.Proxyserver) {
		tflog.Debug(ctx, fmt.Sprintf("proxyserver has changed for reputationsettings"))
		hasChange = true
	}
	if !data.Proxyusername.Equal(state.Proxyusername) {
		tflog.Debug(ctx, fmt.Sprintf("proxyusername has changed for reputationsettings"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		reputationsettings := reputationsettingsGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		reputationsettingsGetThePayloadFromtheConfig(ctx, &config, &reputationsettings)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Reputationsettings.Type(), &reputationsettings)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update reputationsettings, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated reputationsettings resource")
	} else {
		tflog.Debug(ctx, "No changes detected for reputationsettings resource, skipping update")
	}

	// Read the updated state back
	if !r.readReputationsettingsFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "reputationsettings not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ReputationsettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ReputationsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting reputationsettings resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed reputationsettings from Terraform state")
}

// Helper function to read reputationsettings data from API
func (r *ReputationsettingsResource) readReputationsettingsFromApi(ctx context.Context, data *ReputationsettingsResourceModel, diags *diag.Diagnostics) bool {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Reputationsettings.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read reputationsettings, got error: %s", err))
		return false
	}

	reputationsettingsSetAttrFromGet(ctx, data, getResponseData)

	return true
}
