package nskeymanagerproxy

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
var _ resource.Resource = &NskeymanagerproxyResource{}
var _ resource.ResourceWithConfigure = (*NskeymanagerproxyResource)(nil)
var _ resource.ResourceWithImportState = (*NskeymanagerproxyResource)(nil)
var _ resource.ResourceWithValidateConfig = (*NskeymanagerproxyResource)(nil)

func NewNskeymanagerproxyResource() resource.Resource {
	return &NskeymanagerproxyResource{}
}

// NskeymanagerproxyResource defines the resource implementation.
type NskeymanagerproxyResource struct {
	client *service.NitroClient
}

func (r *NskeymanagerproxyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NskeymanagerproxyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nskeymanagerproxy"
}

func (r *NskeymanagerproxyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NskeymanagerproxyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data NskeymanagerproxyResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// NITRO CLI requires exactly one of serverip / servername to identify the proxy.
	// tfdata marks both is_required:false, so enforce here.
	ipSet := !data.Serverip.IsNull() && !data.Serverip.IsUnknown() && data.Serverip.ValueString() != ""
	nameSet := !data.Servername.IsNull() && !data.Servername.IsUnknown() && data.Servername.ValueString() != ""

	if !ipSet && !nameSet {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"Exactly one of \"serverip\" or \"servername\" must be set for nskeymanagerproxy.",
		)
	} else if ipSet && nameSet {
		resp.Diagnostics.AddError(
			"Conflicting Attributes",
			"Only one of \"serverip\" or \"servername\" may be set for nskeymanagerproxy, not both.",
		)
	}
}

func (r *NskeymanagerproxyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NskeymanagerproxyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nskeymanagerproxy resource")
	nskeymanagerproxy := nskeymanagerproxyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - add via POST (NITRO doc: add HTTP Method POST)
	_, err := r.client.AddResource(service.Nskeymanagerproxy.Type(), "", &nskeymanagerproxy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nskeymanagerproxy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nskeymanagerproxy resource")

	// Set ID for the resource before reading state.
	// serverip is the primary key (delete URL path); fall back to servername when unset.
	data.Id = types.StringValue(nskeymanagerproxyComputeId(&data))

	// Read the updated state back
	r.readNskeymanagerproxyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NskeymanagerproxyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NskeymanagerproxyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nskeymanagerproxy resource")

	r.readNskeymanagerproxyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Object deleted out-of-band: remove from state so a later apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NskeymanagerproxyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for nskeymanagerproxy (only add/delete/get).
	// Every schema attribute uses RequiresReplace, so Terraform never reaches Update with a
	// real change; this body is a documented no-op that just re-reads live state.
	var data, state NskeymanagerproxyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for nskeymanagerproxy; all attributes are RequiresReplace")

	// Read the updated state back
	r.readNskeymanagerproxyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NskeymanagerproxyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NskeymanagerproxyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nskeymanagerproxy resource")
	// NITRO doc: delete URL /nskeymanagerproxy/serverip_value with Query-parameters
	// args=servername:<value>. Path key is serverip; servername is passed as an arg.
	serverip_value := data.Serverip.ValueString()
	args := []string{}
	if !data.Servername.IsNull() && data.Servername.ValueString() != "" {
		args = append(args, fmt.Sprintf("servername:%s", utils.UrlEncode(data.Servername.ValueString())))
	}

	err := r.client.DeleteResourceWithArgs(service.Nskeymanagerproxy.Type(), serverip_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nskeymanagerproxy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nskeymanagerproxy resource")
}

// nskeymanagerproxyComputeId returns the resource identifier: serverip when set,
// otherwise servername (both are x-unique-attr). It is also the GET/DELETE path key.
func nskeymanagerproxyComputeId(data *NskeymanagerproxyResourceModel) string {
	if !data.Serverip.IsNull() && !data.Serverip.IsUnknown() && data.Serverip.ValueString() != "" {
		return data.Serverip.ValueString()
	}
	return data.Servername.ValueString()
}

// Helper function to read nskeymanagerproxy data from API
func (r *NskeymanagerproxyResource) readNskeymanagerproxyFromApi(ctx context.Context, data *NskeymanagerproxyResourceModel, diags *diag.Diagnostics) {

	// ID is a plain value (serverip, or servername fallback) - read directly by name.
	name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nskeymanagerproxy.Type(), name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nskeymanagerproxy, got error: %s", err))
		return
	}

	nskeymanagerproxySetAttrFromGet(ctx, data, getResponseData)
}
