package botpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/bot"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &BotpolicyResource{}
var _ resource.ResourceWithConfigure = (*BotpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*BotpolicyResource)(nil)

func NewBotpolicyResource() resource.Resource {
	return &BotpolicyResource{}
}

// BotpolicyResource defines the resource implementation.
type BotpolicyResource struct {
	client *service.NitroClient
}

func (r *BotpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botpolicy"
}

func (r *BotpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botpolicy resource")

	botpolicy := botpolicyGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (POST)
	botpolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Botpolicy.Type(), botpolicyName, &botpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created botpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(botpolicyName)

	// Read the updated state back
	if !r.readBotpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "botpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botpolicy resource")

	found := r.readBotpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *BotpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state BotpolicyResourceModel

	// Read Terraform prior state to preserve the live object ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state; it holds the current live object name.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating botpolicy resource")

	// Handle in-place rename via the NITRO rename action. The rename source must
	// be the CURRENT LIVE name (state.Id), not the configured name attribute.
	if !data.Newname.IsNull() && !data.Newname.IsUnknown() && data.Newname.ValueString() != "" && !data.Newname.Equal(state.Newname) {
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("newname has changed for botpolicy, renaming %s -> %s", state.Id.ValueString(), newName))
		renamePayload := bot.Botpolicy{
			Name:    state.Id.ValueString(),
			Newname: newName,
		}
		err := r.client.ActOnResource(service.Botpolicy.Type(), &renamePayload, "rename")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename botpolicy, got error: %s", err))
			return
		}
		// The live object is now named newName; the ID must track it.
		data.Id = types.StringValue(newName)
	}

	// Detect changes to in-place updatable attributes (name, rule and profilename
	// are RequiresReplace and never reach Update).
	hasChange := false
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for botpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for botpolicy")
		hasChange = true
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for botpolicy")
		hasChange = true
	}

	if hasChange {
		botpolicy := botpolicyGetThePayloadFromthePlan(ctx, &data)
		// Target the live object name (post-rename if applicable). NITRO's update
		// PUT requires the name field in the payload.
		botpolicy.Name = data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Botpolicy.Type(), data.Id.ValueString(), &botpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated botpolicy resource")
	} else {
		tflog.Debug(ctx, "No in-place changes detected for botpolicy resource")
	}

	// Read the updated state back
	if !r.readBotpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "botpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botpolicy resource")

	// Named resource - delete using DeleteResource keyed on the live name (ID).
	botpolicyName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Botpolicy.Type(), botpolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete botpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted botpolicy resource")
}

// Helper function to read botpolicy data from API
func (r *BotpolicyResource) readBotpolicyFromApi(ctx context.Context, data *BotpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the live object name.
	botpolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Botpolicy.Type(), botpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botpolicy, got error: %s", err))
		return false
	}

	botpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
