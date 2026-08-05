package channel

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
var _ resource.Resource = &ChannelResource{}
var _ resource.ResourceWithConfigure = (*ChannelResource)(nil)
var _ resource.ResourceWithImportState = (*ChannelResource)(nil)

func NewChannelResource() resource.Resource {
	return &ChannelResource{}
}

// ChannelResource defines the resource implementation.
type ChannelResource struct {
	client *service.NitroClient
}

func (r *ChannelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ChannelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_channel"
}

func (r *ChannelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ChannelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ChannelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating channel resource")

	channel := channelGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource keyed on channel_id
	channelName := data.ChannelId.ValueString()
	_, err := r.client.AddResource(service.Channel.Type(), channelName, &channel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create channel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created channel resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(channelName)

	// Read the updated state back
	if !r.readChannelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "channel not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChannelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ChannelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading channel resource")

	found := r.readChannelFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ChannelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ChannelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating channel resource")

	// Create API request body from the model.
	// channel is updated via a PUT to /config/channel with the "id" in the body
	// (no name in the URL), matching the SDK v2 behavior.
	channel := channelGetThePayloadFromtheConfig(ctx, &data)

	err := r.client.UpdateUnnamedResource(service.Channel.Type(), &channel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update channel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated channel resource")

	// Read the updated state back
	if !r.readChannelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "channel not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChannelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ChannelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting channel resource")

	// Named resource - delete using DeleteResource keyed on the channel id
	channelName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Channel.Type(), channelName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete channel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted channel resource")
}

// Helper function to read channel data from API.
// Returns false (without an error diagnostic) when the resource no longer exists.
func (r *ChannelResource) readChannelFromApi(ctx context.Context, data *ChannelResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain channel id value
	channelName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Channel.Type(), channelName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read channel, got error: %s", err))
		return false
	}

	channelSetAttrFromGet(ctx, data, getResponseData)

	return true
}
