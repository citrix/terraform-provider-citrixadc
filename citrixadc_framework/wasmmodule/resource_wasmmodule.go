package wasmmodule

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
var _ resource.Resource = &WasmmoduleResource{}
var _ resource.ResourceWithConfigure = (*WasmmoduleResource)(nil)
var _ resource.ResourceWithImportState = (*WasmmoduleResource)(nil)

func NewWasmmoduleResource() resource.Resource {
	return &WasmmoduleResource{}
}

// WasmmoduleResource defines the resource implementation.
type WasmmoduleResource struct {
	client *service.NitroClient
}

func (r *WasmmoduleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *WasmmoduleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wasmmodule"
}

func (r *WasmmoduleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *WasmmoduleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data WasmmoduleResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating wasmmodule resource")

	// Build untyped NITRO add payload (no SDK struct exists for wasmmodule).
	wasmmodule := wasmmoduleGetTheCreatePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Wasmmodule.Type(), name_value, &wasmmodule)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create wasmmodule, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created wasmmodule resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readWasmmoduleFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "wasmmodule not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WasmmoduleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WasmmoduleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading wasmmodule resource")

	found := r.readWasmmoduleFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *WasmmoduleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state WasmmoduleResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (-> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating wasmmodule resource")

	// Only settingfile and comment are updatable/unsettable per the NITRO doc;
	// modulefile and signaturefile are create-only (RequiresReplace).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Settingfile.Equal(state.Settingfile) {
		tflog.Debug(ctx, "settingfile has changed for wasmmodule")
		if config.Settingfile.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "settingfile")
		} else {
			hasChange = true
		}
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for wasmmodule")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Build untyped NITRO update payload.
		wasmmodule := wasmmoduleGetTheUpdatePayloadFromthePlan(ctx, &data)
		// Make API call
		// Update is an unnamed PUT (name is carried in the payload)
		err := r.client.UpdateUnnamedResource(service.Wasmmodule.Type(), &wasmmodule)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update wasmmodule, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated wasmmodule resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for wasmmodule resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Wasmmodule.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset wasmmodule attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readWasmmoduleFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "wasmmodule not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WasmmoduleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data WasmmoduleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting wasmmodule resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Wasmmodule.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete wasmmodule, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted wasmmodule resource")
}

// Helper function to read wasmmodule data from API
func (r *WasmmoduleResource) readWasmmoduleFromApi(ctx context.Context, data *WasmmoduleResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Wasmmodule.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read wasmmodule, got error: %s", err))
		return false
	}

	wasmmoduleSetAttrFromGet(ctx, data, getResponseData)

	return true
}
