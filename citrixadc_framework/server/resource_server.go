package server

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkv2resource "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ServerResource{}
var _ resource.ResourceWithConfigure = (*ServerResource)(nil)
var _ resource.ResourceWithImportState = (*ServerResource)(nil)

func NewServerResource() resource.Resource {
	return &ServerResource{}
}

// ServerResource defines the resource implementation.
type ServerResource struct {
	client *service.NitroClient
}

func (r *ServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (r *ServerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServerResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating server resource")

	// Resolve the server name. SDK v2 auto-generates a name when it is not set.
	var serverName string
	if !data.Name.IsNull() && !data.Name.IsUnknown() && data.Name.ValueString() != "" {
		serverName = data.Name.ValueString()
	} else {
		serverName = sdkv2resource.PrefixedUniqueId("tf-server-")
	}
	data.Name = types.StringValue(serverName)

	// Named resource - use AddResource
	server := serverGetThePayloadFromthePlan(ctx, &data)
	server.Name = serverName

	_, err := r.client.AddResource(service.Server.Type(), serverName, &server)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create server, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created server resource")

	// Set ID for the resource before reading state (SDK v2: d.SetId(serverName)).
	data.Id = types.StringValue(serverName)

	// Read the updated state back
	if !r.readServerFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "server not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ServerResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading server resource")

	found := r.readServerFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ServerResourceModel

	// Read Terraform prior state to preserve ID / live name
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (name is ForceNew, so the live name never changes).
	data.Id = state.Id

	tflog.Debug(ctx, "Updating server resource")

	serverName := data.Id.ValueString()

	// Build the update payload with only the changed, updateable attributes,
	// mirroring the SDK v2 updateServerFunc.
	server := basic.Server{
		Name: serverName,
	}
	hasChange := false
	stateChange := false

	if !data.Comment.Equal(state.Comment) {
		server.Comment = data.Comment.ValueString()
		hasChange = true
	}
	if !data.Domainresolvenow.Equal(state.Domainresolvenow) {
		server.Domainresolvenow = data.Domainresolvenow.ValueBool()
		hasChange = true
	}
	if !data.Domainresolveretry.Equal(state.Domainresolveretry) {
		server.Domainresolveretry = utils.IntPtr(int(data.Domainresolveretry.ValueInt64()))
		hasChange = true
	}
	if !data.Internal.Equal(state.Internal) {
		server.Internal = data.Internal.ValueBool()
		hasChange = true
	}
	if !data.Ipaddress.Equal(state.Ipaddress) {
		server.Ipaddress = data.Ipaddress.ValueString()
		hasChange = true
	}
	if !data.Querytype.Equal(state.Querytype) {
		server.Querytype = data.Querytype.ValueString()
		hasChange = true
	}
	if !data.Translationip.Equal(state.Translationip) {
		server.Translationip = data.Translationip.ValueString()
		hasChange = true
	}
	if !data.Translationmask.Equal(state.Translationmask) {
		server.Translationmask = data.Translationmask.ValueString()
		hasChange = true
	}
	if !data.State.Equal(state.State) {
		stateChange = true
	}

	if hasChange {
		_, err := r.client.UpdateResource(service.Server.Type(), serverName, &server)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update server %s, got error: %s", serverName, err))
			return
		}
		tflog.Trace(ctx, "Updated server resource")
	}
	if stateChange {
		if err := r.doServerStateChange(&data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error enabling/disabling server %s, got error: %s", serverName, err))
			return
		}
	}
	if !hasChange && !stateChange {
		tflog.Debug(ctx, "No changes detected for server resource, skipping update")
	}

	// Read the updated state back
	if !r.readServerFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "server not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ServerResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting server resource")

	// Named resource - delete using DeleteResource
	serverName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Server.Type(), serverName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete server, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted server resource")
}

// doServerStateChange mirrors the SDK v2 doServerStateChange: it enables or
// disables the server via the NITRO enable/disable actions.
func (r *ServerResource) doServerStateChange(data *ServerResourceModel) error {
	// A fresh struct is required - ActOnResource fails on superfluous attributes.
	server := basic.Server{
		Name: data.Id.ValueString(),
	}

	newstate := data.State.ValueString()

	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Server.Type(), server, "enable")
	case "DISABLED":
		// Add attributes relevant to the disable operation.
		if !data.Delay.IsNull() && !data.Delay.IsUnknown() {
			server.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
		}
		if !data.Graceful.IsNull() && !data.Graceful.IsUnknown() {
			server.Graceful = data.Graceful.ValueString()
		}
		return r.client.ActOnResource(service.Server.Type(), server, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// Helper function to read server data from API
func (r *ServerResource) readServerFromApi(ctx context.Context, data *ServerResourceModel, diags *diag.Diagnostics) bool {
	serverName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Server.Type(), serverName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read server, got error: %s", err))
		return false
	}

	serverSetAttrFromGet(ctx, data, getResponseData)

	return true
}
