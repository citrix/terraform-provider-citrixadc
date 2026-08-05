package autoscalepolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/autoscale"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AutoscalepolicyResource{}
var _ resource.ResourceWithConfigure = (*AutoscalepolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AutoscalepolicyResource)(nil)

func NewAutoscalepolicyResource() resource.Resource {
	return &AutoscalepolicyResource{}
}

// AutoscalepolicyResource defines the resource implementation.
type AutoscalepolicyResource struct {
	client *service.NitroClient
}

func (r *AutoscalepolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AutoscalepolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_autoscalepolicy"
}

func (r *AutoscalepolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AutoscalepolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AutoscalepolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating autoscalepolicy resource")

	autoscalepolicy := autoscalepolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Autoscalepolicy.Type(), name_value, &autoscalepolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create autoscalepolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created autoscalepolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAutoscalepolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscalepolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AutoscalepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading autoscalepolicy resource")

	r.readAutoscalepolicyFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource no longer exists on the ADC - remove it from state.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscalepolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AutoscalepolicyResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (holds the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating autoscalepolicy resource")

	// Rename support (NITRO ?action=rename). name is RequiresReplace, so a name
	// change never reaches Update - only a newname change does. Mirrors the SDK v2
	// rename convention (citrixadc/resource_citrixadc_appfwpolicy.go).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Name (which stays pinned to the originally configured value and
		// would point at a stale name on a second rename).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming autoscalepolicy from %q to %q", oldName, newName))

		renamePayload := autoscale.Autoscalepolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Autoscalepolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename autoscalepolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so all
		// subsequent calls (update below, read-back, future reads) address it.
		data.Id = types.StringValue(newName)
	}

	// In-place update of mutable attributes. NITRO's update verb is PUT to the
	// unnamed URL (payload carries the name), matching the SDK v2 implementation's
	// use of UpdateUnnamedResource.
	hasChange := false
	autoscalepolicy := autoscale.Autoscalepolicy{
		Name: data.Id.ValueString(),
	}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for autoscalepolicy")
		autoscalepolicy.Action = data.Action.ValueString()
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for autoscalepolicy")
		autoscalepolicy.Comment = data.Comment.ValueString()
		hasChange = true
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for autoscalepolicy")
		autoscalepolicy.Logaction = data.Logaction.ValueString()
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for autoscalepolicy")
		autoscalepolicy.Rule = data.Rule.ValueString()
		hasChange = true
	}

	if hasChange {
		err := r.client.UpdateUnnamedResource(service.Autoscalepolicy.Type(), &autoscalepolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update autoscalepolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated autoscalepolicy resource")
	} else {
		tflog.Debug(ctx, "No mutable changes detected for autoscalepolicy resource, skipping update")
	}

	// Read the current state back. Capture the user-facing name and newname
	// (which GET does not faithfully echo after a rename) and restore them after
	// the read to avoid an inconsistent-result / perpetual diff.
	planName := data.Name
	planNewname := data.Newname
	r.readAutoscalepolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscalepolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AutoscalepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting autoscalepolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id,
	// NOT data.Name (which stays at the originally configured value and would
	// target a non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Autoscalepolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete autoscalepolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted autoscalepolicy resource")
}

// Helper function to read autoscalepolicy data from API
func (r *AutoscalepolicyResource) readAutoscalepolicyFromApi(ctx context.Context, data *AutoscalepolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain (live) name value
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Autoscalepolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read autoscalepolicy, got error: %s", err))
		return
	}

	autoscalepolicySetAttrFromGet(ctx, data, getResponseData)
}
