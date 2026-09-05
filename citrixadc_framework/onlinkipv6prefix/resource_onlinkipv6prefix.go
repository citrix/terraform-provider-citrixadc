package onlinkipv6prefix

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
var _ resource.Resource = &Onlinkipv6prefixResource{}
var _ resource.ResourceWithConfigure = (*Onlinkipv6prefixResource)(nil)
var _ resource.ResourceWithImportState = (*Onlinkipv6prefixResource)(nil)

func NewOnlinkipv6prefixResource() resource.Resource {
	return &Onlinkipv6prefixResource{}
}

// Onlinkipv6prefixResource defines the resource implementation.
type Onlinkipv6prefixResource struct {
	client *service.NitroClient
}

func (r *Onlinkipv6prefixResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Onlinkipv6prefixResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_onlinkipv6prefix"
}

func (r *Onlinkipv6prefixResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Onlinkipv6prefixResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Onlinkipv6prefixResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating onlinkipv6prefix resource")

	onlinkipv6prefix := onlinkipv6prefixGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource (matches SDK v2 client.AddResource keyed on ipv6prefix)
	ipv6prefixName := data.Ipv6prefix.ValueString()
	_, err := r.client.AddResource(service.Onlinkipv6prefix.Type(), ipv6prefixName, &onlinkipv6prefix)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create onlinkipv6prefix, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created onlinkipv6prefix resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(ipv6prefix))
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Ipv6prefix.ValueString()))

	// Read the updated state back
	if !r.readOnlinkipv6prefixFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "onlinkipv6prefix not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Onlinkipv6prefixResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Onlinkipv6prefixResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading onlinkipv6prefix resource")

	found := r.readOnlinkipv6prefixFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Onlinkipv6prefixResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state Onlinkipv6prefixResourceModel

	// Read Terraform prior state to preserve ID
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

	tflog.Debug(ctx, "Updating onlinkipv6prefix resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Autonomusprefix.Equal(state.Autonomusprefix) {
		tflog.Debug(ctx, "autonomusprefix has changed for onlinkipv6prefix")
		if config.Autonomusprefix.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "autonomusprefix")
		} else {
			hasChange = true
		}
	}
	if !data.Decrementprefixlifetimes.Equal(state.Decrementprefixlifetimes) {
		tflog.Debug(ctx, "decrementprefixlifetimes has changed for onlinkipv6prefix")
		if config.Decrementprefixlifetimes.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "decrementprefixlifetimes")
		} else {
			hasChange = true
		}
	}
	if !data.Depricateprefix.Equal(state.Depricateprefix) {
		tflog.Debug(ctx, "depricateprefix has changed for onlinkipv6prefix")
		if config.Depricateprefix.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "depricateprefix")
		} else {
			hasChange = true
		}
	}
	if !data.Onlinkprefix.Equal(state.Onlinkprefix) {
		tflog.Debug(ctx, "onlinkprefix has changed for onlinkipv6prefix")
		if config.Onlinkprefix.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "onlinkprefix")
		} else {
			hasChange = true
		}
	}
	if !data.Prefixpreferredlifetime.Equal(state.Prefixpreferredlifetime) {
		tflog.Debug(ctx, "prefixpreferredlifetime has changed for onlinkipv6prefix")
		if config.Prefixpreferredlifetime.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "prefixpreferredlifetime")
		} else {
			hasChange = true
		}
	}
	if !data.Prefixvalidelifetime.Equal(state.Prefixvalidelifetime) {
		tflog.Debug(ctx, "prefixvalidelifetime has changed for onlinkipv6prefix")
		if config.Prefixvalidelifetime.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "prefixvalidelifetime")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model. ipv6prefix (the key) is always
		// included so NITRO knows which prefix to update.
		onlinkipv6prefix := onlinkipv6prefixGetThePayloadFromtheConfig(ctx, &data)
		// Named resource updated via UpdateUnnamedResource (matches SDK v2 semantics:
		// PUT /config/onlinkipv6prefix with ipv6prefix carried in the body).
		err := r.client.UpdateUnnamedResource(service.Onlinkipv6prefix.Type(), &onlinkipv6prefix)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update onlinkipv6prefix, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated onlinkipv6prefix resource")
	} else {
		tflog.Debug(ctx, "No changes detected for onlinkipv6prefix resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. Keyed on ipv6prefix (carried in the unset body).
	unsetIdPayload := map[string]interface{}{
		"ipv6prefix": data.Ipv6prefix.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Onlinkipv6prefix.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset onlinkipv6prefix attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readOnlinkipv6prefixFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "onlinkipv6prefix not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Onlinkipv6prefixResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Onlinkipv6prefixResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting onlinkipv6prefix resource")

	// Named resource - delete using DeleteResource keyed on the ID (ipv6prefix value)
	ipv6prefixName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Onlinkipv6prefix.Type(), ipv6prefixName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete onlinkipv6prefix, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted onlinkipv6prefix resource")
}

// Helper function to read onlinkipv6prefix data from API
func (r *Onlinkipv6prefixResource) readOnlinkipv6prefixFromApi(ctx context.Context, data *Onlinkipv6prefixResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (ipv6prefix)
	ipv6prefixName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Onlinkipv6prefix.Type(), ipv6prefixName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read onlinkipv6prefix, got error: %s", err))
		return false
	}

	onlinkipv6prefixSetAttrFromGet(ctx, data, getResponseData)

	return true
}
