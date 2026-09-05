package icapolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ica"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IcapolicyResource{}
var _ resource.ResourceWithConfigure = (*IcapolicyResource)(nil)
var _ resource.ResourceWithImportState = (*IcapolicyResource)(nil)

func NewIcapolicyResource() resource.Resource {
	return &IcapolicyResource{}
}

// IcapolicyResource defines the resource implementation.
type IcapolicyResource struct {
	client *service.NitroClient
}

func (r *IcapolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IcapolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_icapolicy"
}

func (r *IcapolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IcapolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IcapolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating icapolicy resource")

	icapolicy := icapolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Icapolicy.Type(), name_value, &icapolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create icapolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created icapolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readIcapolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcapolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IcapolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading icapolicy resource")

	r.readIcapolicyFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcapolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state IcapolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating icapolicy resource")

	// Rename support: icapolicy exposes a NITRO `rename` action. `name` uses
	// RequiresReplace, so a name change recreates the resource and never reaches here.
	// The ONLY name-related change that lands in Update is `newname`, which drives an
	// in-place rename (mirrors SDK v2 appfwpolicy convention).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, which is tracked by the ID - NOT
		// state.Name. state.Name stays pinned to the originally configured value, so on
		// a SECOND rename it would point at the wrong (no longer live) name.
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming icapolicy from %q to %q", oldName, newName))

		renamePayload := ica.Icapolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Icapolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename icapolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update/read
		// below (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Handle in-place attribute updates via NITRO update (PUT).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for icapolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for icapolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for icapolicy")
		if config.Logaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logaction")
		} else {
			hasChange = true
		}
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for icapolicy")
		hasChange = true
	}

	if hasChange {
		icapolicy := icapolicyGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// The update must target the CURRENT LIVE name (post-rename), which is data.Id -
		// NOT data.Name (which stays pinned to the originally configured value).
		liveName := data.Id.ValueString()
		icapolicy.Name = liveName
		_, err := r.client.UpdateResource(service.Icapolicy.Type(), liveName, &icapolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update icapolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated icapolicy resource")
	} else {
		tflog.Debug(ctx, "No attribute changes detected for icapolicy resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them
	// to their defaults. The unset must target the CURRENT LIVE name (post-rename),
	// which is tracked by data.Id.
	unsetIdPayload := map[string]interface{}{
		"name": data.Id.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Icapolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset icapolicy attributes, got error: %s", err))
		return
	}

	// Read the current state back. The resource may now be physically named newName,
	// so preserve the user-facing name/newname across the read-back to avoid an
	// inconsistent-result / perpetual diff.
	planName := data.Name
	planNewname := data.Newname
	r.readIcapolicyFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcapolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IcapolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting icapolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== name at create, == newname after a rename), so we must delete by data.Id, NOT
	// data.Name (which stays at the originally configured value and would target a
	// non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Icapolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete icapolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted icapolicy resource")
}

// Helper function to read icapolicy data from API
func (r *IcapolicyResource) readIcapolicyFromApi(ctx context.Context, data *IcapolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value (live name)
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Icapolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read icapolicy, got error: %s", err))
		return
	}

	icapolicySetAttrFromGet(ctx, data, getResponseData)

}
