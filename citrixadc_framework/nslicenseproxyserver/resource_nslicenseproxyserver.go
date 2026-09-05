package nslicenseproxyserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NslicenseproxyserverResource{}
var _ resource.ResourceWithConfigure = (*NslicenseproxyserverResource)(nil)
var _ resource.ResourceWithImportState = (*NslicenseproxyserverResource)(nil)

func NewNslicenseproxyserverResource() resource.Resource {
	return &NslicenseproxyserverResource{}
}

// NslicenseproxyserverResource defines the resource implementation.
type NslicenseproxyserverResource struct {
	client *service.NitroClient
}

func (r *NslicenseproxyserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslicenseproxyserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslicenseproxyserver"
}

func (r *NslicenseproxyserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// resourceName derives the SDK v2-compatible identifier: serverip if configured,
// otherwise servername (serverip precedence).
func nslicenseproxyserverResourceName(data *NslicenseproxyserverResourceModel) string {
	if !data.Serverip.IsNull() && !data.Serverip.IsUnknown() && data.Serverip.ValueString() != "" {
		return data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() && data.Servername.ValueString() != "" {
		return data.Servername.ValueString()
	}
	return ""
}

func (r *NslicenseproxyserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslicenseproxyserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslicenseproxyserver resource")

	nslicenseproxyserver := nslicenseproxyserverGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource (POST /config/nslicenseproxyserver)
	name := nslicenseproxyserverResourceName(&data)
	_, err := r.client.AddResource(service.Nslicenseproxyserver.Type(), name, &nslicenseproxyserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nslicenseproxyserver, got error: %s", err))
		return
	}

	// ID is the plain serverip/servername value (SDK v2 backward-compat)
	data.Id = types.StringValue(name)

	tflog.Trace(ctx, "Created nslicenseproxyserver resource")

	// Read the updated state back
	if !r.readNslicenseproxyserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nslicenseproxyserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseproxyserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslicenseproxyserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nslicenseproxyserver resource")

	found := r.readNslicenseproxyserverFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NslicenseproxyserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NslicenseproxyserverResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (serverip/servername are ForceNew, only port updates)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nslicenseproxyserver resource")

	// Only port is updateable in SDK v2 (serverip/servername are ForceNew).
	if !data.Port.Equal(state.Port) {
		tflog.Debug(ctx, "port has changed for nslicenseproxyserver, starting update")
		nslicenseproxyserver := nslicenseproxyserverGetThePayloadFromtheConfig(ctx, &data)
		// Singleton-style update: PUT /config/nslicenseproxyserver
		err := r.client.UpdateUnnamedResource(service.Nslicenseproxyserver.Type(), &nslicenseproxyserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nslicenseproxyserver, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated nslicenseproxyserver resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nslicenseproxyserver resource, skipping update")
	}

	// Read the updated state back
	if !r.readNslicenseproxyserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nslicenseproxyserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseproxyserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslicenseproxyserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nslicenseproxyserver resource")

	// Named resource - delete using DeleteResource (DELETE /config/nslicenseproxyserver/<id>)
	name := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nslicenseproxyserver.Type(), name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nslicenseproxyserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nslicenseproxyserver resource")
}

// Helper function to read nslicenseproxyserver data from API.
// Returns false if the resource no longer exists on the ADC.
func (r *NslicenseproxyserverResource) readNslicenseproxyserverFromApi(ctx context.Context, data *NslicenseproxyserverResourceModel, diags *diag.Diagnostics) bool {
	// nslicenseproxyserver only supports "get (all)" - read via empty name (SDK v2 parity).
	getResponseData, err := r.client.FindResource(service.Nslicenseproxyserver.Type(), "")
	if err != nil {
		// Treat any read failure as "not found" (matches SDK v2 which clears state on read error).
		return false
	}

	nslicenseproxyserverSetAttrFromGet(ctx, data, getResponseData)

	return true
}
