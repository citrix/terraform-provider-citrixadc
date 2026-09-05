package uservserver

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/user"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &UservserverResource{}
var _ resource.ResourceWithConfigure = (*UservserverResource)(nil)
var _ resource.ResourceWithImportState = (*UservserverResource)(nil)

func NewUservserverResource() resource.Resource {
	return &UservserverResource{}
}

// UservserverResource defines the resource implementation.
type UservserverResource struct {
	client *service.NitroClient
}

func (r *UservserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *UservserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_uservserver"
}

func (r *UservserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *UservserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data UservserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating uservserver resource")

	uservserver := uservserverGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource. Mirrors SDK v2 client.AddResource("uservserver", name, ...).
	uservserverName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Uservserver.Type(), uservserverName, &uservserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create uservserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created uservserver resource")

	// Set ID for the resource before reading state (SDK v2 d.SetId(name)).
	data.Id = types.StringValue(uservserverName)

	// Read the updated state back
	if !r.readUservserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "uservserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UservserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data UservserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading uservserver resource")

	found := r.readUservserverFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// Mirrors SDK v2 read which clears the ID when the resource is gone.
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UservserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state UservserverResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating uservserver resource")

	uservserverName := data.Name.ValueString()

	// Detect changes on updateable, non-state attributes (SDK v2 updateUservserverFunc).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for uservserver")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Defaultlb.Equal(state.Defaultlb) {
		tflog.Debug(ctx, "defaultlb has changed for uservserver")
		hasChange = true
	}
	if !data.Ipaddress.Equal(state.Ipaddress) {
		tflog.Debug(ctx, "ipaddress has changed for uservserver")
		hasChange = true
	}
	if !data.Params.Equal(state.Params) {
		tflog.Debug(ctx, "params has changed for uservserver")
		if config.Params.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "Params")
		} else {
			hasChange = true
		}
	}

	// state is toggled via the enable/disable action, not UpdateResource (SDK v2 doUservserverStateChange).
	stateChange := !data.State.Equal(state.State)

	if stateChange {
		if err := r.doUservserverStateChange(ctx, uservserverName, data.State.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error enabling/disabling uservserver %s, got error: %s", uservserverName, err))
			return
		}
	}

	if hasChange {
		uservserver := uservserverGetTheUpdatablePayloadFromThePlan(ctx, &data)
		_, err := r.client.UpdateResource(service.Uservserver.Type(), uservserverName, &uservserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update uservserver %s, got error: %s", uservserverName, err))
			return
		}
		tflog.Trace(ctx, "Updated uservserver resource")
	} else {
		tflog.Debug(ctx, "No in-place attribute changes detected for uservserver resource")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Uservserver.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset uservserver attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readUservserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "uservserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UservserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data UservserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting uservserver resource")

	// Named resource - delete using DeleteResource (SDK v2 client.DeleteResource("uservserver", id)).
	err := r.client.DeleteResource(service.Uservserver.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete uservserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted uservserver resource")
}

// doUservserverStateChange enables or disables the user vserver via the NITRO
// enable/disable action. Mirrors SDK v2 doUservserverStateChange.
func (r *UservserverResource) doUservserverStateChange(ctx context.Context, name string, newState string) error {
	tflog.Debug(ctx, "In doUservserverStateChange Function")

	uservserver := user.Uservserver{
		Name: name,
	}

	switch strings.ToLower(newState) {
	case "enabled":
		return r.client.ActOnResource(service.Uservserver.Type(), &uservserver, "enable")
	case "disabled":
		return r.client.ActOnResource(service.Uservserver.Type(), &uservserver, "disable")
	case "":
		// No explicit state requested; nothing to toggle.
		return nil
	default:
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newState)
	}
}

// Helper function to read uservserver data from API.
// Returns false (without error) when the resource no longer exists on the ADC.
func (r *UservserverResource) readUservserverFromApi(ctx context.Context, data *UservserverResourceModel, diags *diag.Diagnostics) bool {
	// Named resource - find by ID (== name). SDK v2 read used client.FindResource("uservserver", id).
	uservserverName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Uservserver.Type(), uservserverName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read uservserver, got error: %s", err))
		return false
	}

	uservserverSetAttrFromGet(ctx, data, getResponseData)

	return true
}
