package nslicenseserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NslicenseserverResource{}
var _ resource.ResourceWithConfigure = (*NslicenseserverResource)(nil)
var _ resource.ResourceWithImportState = (*NslicenseserverResource)(nil)

func NewNslicenseserverResource() resource.Resource {
	return &NslicenseserverResource{}
}

// NslicenseserverResource defines the resource implementation.
type NslicenseserverResource struct {
	client *service.NitroClient
}

func (r *NslicenseserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslicenseserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslicenseserver"
}

func (r *NslicenseserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslicenseserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslicenseserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslicenseserver resource")

	// Build the payload and add the license server (SDK v2 used AddResource).
	nslicenseserver := nslicenseserverGetThePayloadFromthePlan(ctx, &data)
	_, err := r.client.AddResource(service.Nslicenseserver.Type(), "", &nslicenseserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nslicenseserver, got error: %s", err))
		return
	}

	// ID matches SDK v2: d.SetId(servername)
	data.Id = types.StringValue(data.Servername.ValueString())

	tflog.Trace(ctx, "Created nslicenseserver resource")

	// Read the updated state back
	if !r.readNslicenseserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nslicenseserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslicenseserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nslicenseserver resource")

	found := r.readNslicenseserverFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NslicenseserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NslicenseserverResourceModel

	// Read Terraform prior state to detect changes and preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nslicenseserver resource")

	// Only licensemode and port are updateable in SDK v2; everything else is
	// RequiresReplace and never reaches Update.
	nslicenseserver := ns.Nslicenseserver{Servername: data.Servername.ValueString()}
	hasChange := false
	if !data.Licensemode.Equal(state.Licensemode) {
		tflog.Debug(ctx, "licensemode has changed for nslicenseserver, starting update")
		nslicenseserver.Licensemode = data.Licensemode.ValueString()
		hasChange = true
	}
	if !data.Port.Equal(state.Port) {
		tflog.Debug(ctx, "port has changed for nslicenseserver, starting update")
		if !data.Port.IsNull() && !data.Port.IsUnknown() {
			nslicenseserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
		}
		hasChange = true
	}

	if hasChange {
		_, err := r.client.UpdateResource(service.Nslicenseserver.Type(), "", &nslicenseserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nslicenseserver, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated nslicenseserver resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nslicenseserver resource, skipping update")
	}

	// Read the updated state back
	if !r.readNslicenseserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nslicenseserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslicenseserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nslicenseserver resource")

	// SDK v2 deleted with args: servername:<value>
	args := []string{fmt.Sprintf("servername:%s", data.Servername.ValueString())}
	err := r.client.DeleteResourceWithArgs(service.Nslicenseserver.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nslicenseserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nslicenseserver resource")
}

// readNslicenseserverFromApi reads the license server via the array-filter GET
// (matching SDK v2 FindResourceArrayWithParams). Returns false when no license
// server is configured so the caller can remove it from state.
func (r *NslicenseserverResource) readNslicenseserverFromApi(ctx context.Context, data *NslicenseserverResourceModel, diags *diag.Diagnostics) bool {
	findParams := service.FindParams{
		ResourceType:             "nslicenseserver",
		ResourceMissingErrorCode: 258,
	}

	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nslicenseserver, got error: %s", err))
		return false
	}

	// No license server configured.
	if len(dataArr) == 0 {
		return false
	}

	// License server returns at most one element.
	nslicenseserverSetAttrFromGet(ctx, data, dataArr[0])

	return true
}
