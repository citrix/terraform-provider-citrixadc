package ip6tunnel

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
var _ resource.Resource = &Ip6tunnelResource{}
var _ resource.ResourceWithConfigure = (*Ip6tunnelResource)(nil)
var _ resource.ResourceWithImportState = (*Ip6tunnelResource)(nil)

func NewIp6tunnelResource() resource.Resource {
	return &Ip6tunnelResource{}
}

// Ip6tunnelResource defines the resource implementation.
type Ip6tunnelResource struct {
	client *service.NitroClient
}

func (r *Ip6tunnelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Ip6tunnelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip6tunnel"
}

func (r *Ip6tunnelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Ip6tunnelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Ip6tunnelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ip6tunnel resource")

	ip6tunnel := ip6tunnelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource keyed on "name" - use AddResource
	ip6tunnelName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Ip6tunnel.Type(), ip6tunnelName, &ip6tunnel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ip6tunnel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ip6tunnel resource")

	// Set ID for the resource before reading state.
	// ID is the resource name, matching the SDK v2 behavior (d.SetId(name)).
	data.Id = types.StringValue(ip6tunnelName)

	// Read the updated state back
	if !r.readIp6tunnelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ip6tunnel not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ip6tunnelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Ip6tunnelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ip6tunnel resource")

	found := r.readIp6tunnelFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Ip6tunnelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Ip6tunnelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// ip6tunnel exposes no NITRO "update"/"set" operation and every attribute is
	// ForceNew/RequiresReplace (matching the SDK v2 resource, which had no Update
	// function). There is therefore nothing to update in place - just re-read the
	// current state from the ADC.
	tflog.Debug(ctx, "Updating ip6tunnel resource (no updatable attributes; re-reading state)")

	if !r.readIp6tunnelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ip6tunnel not found during update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ip6tunnelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ip6tunnelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ip6tunnel resource")

	// Delete the resource by name (the ID is the resource name).
	ip6tunnelName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Ip6tunnel.Type(), ip6tunnelName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ip6tunnel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted ip6tunnel resource from state")
}

// Helper function to read ip6tunnel data from API.
// Returns false (without adding an error) when the resource no longer exists.
func (r *Ip6tunnelResource) readIp6tunnelFromApi(ctx context.Context, data *Ip6tunnelResourceModel, diags *diag.Diagnostics) bool {
	ip6tunnelName := data.Id.ValueString()
	getResponseData, err := r.client.FindResource(service.Ip6tunnel.Type(), ip6tunnelName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ip6tunnel, got error: %s", err))
		return false
	}

	ip6tunnelSetAttrFromGet(ctx, data, getResponseData)

	return true
}
