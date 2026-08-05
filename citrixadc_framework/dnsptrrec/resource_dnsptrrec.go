package dnsptrrec

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
var _ resource.Resource = &DnsptrrecResource{}
var _ resource.ResourceWithConfigure = (*DnsptrrecResource)(nil)
var _ resource.ResourceWithImportState = (*DnsptrrecResource)(nil)

func NewDnsptrrecResource() resource.Resource {
	return &DnsptrrecResource{}
}

// DnsptrrecResource defines the resource implementation.
type DnsptrrecResource struct {
	client *service.NitroClient
}

func (r *DnsptrrecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsptrrecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsptrrec"
}

func (r *DnsptrrecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsptrrecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsptrrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsptrrec resource")

	dnsptrrec := dnsptrrecGetThePayloadFromtheConfig(ctx, &data)

	// Named resource keyed on reversedomain - use AddResource
	dnsptrrecName := data.Reversedomain.ValueString()

	_, err := r.client.AddResource(service.Dnsptrrec.Type(), dnsptrrecName, &dnsptrrec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsptrrec, got error: %s", err))
		return
	}

	// Set ID to reversedomain before reading state back
	data.Id = types.StringValue(dnsptrrecName)

	tflog.Trace(ctx, "Created dnsptrrec resource")

	// Read the updated state back
	if !r.readDnsptrrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsptrrec not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsptrrecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsptrrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsptrrec resource")

	found := r.readDnsptrrecFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnsptrrecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsptrrecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnsptrrec resource")

	// NITRO exposes no set/update operation for dnsptrrec and every attribute is
	// ForceNew (RequiresReplace), so any attribute change is handled by Terraform via
	// destroy+recreate and this Update path is never reached with a real change.
	// Simply refresh state from the ADC.
	if !r.readDnsptrrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsptrrec not found during update")
		}
		return
	}

	tflog.Trace(ctx, "Updated dnsptrrec resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsptrrecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsptrrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsptrrec resource")

	// Named resource - delete by reversedomain (the resource ID)
	dnsptrrecName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Dnsptrrec.Type(), dnsptrrecName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsptrrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsptrrec resource")
}

// Helper function to read dnsptrrec data from API.
// Returns false (without adding an error) when the resource no longer exists.
func (r *DnsptrrecResource) readDnsptrrecFromApi(ctx context.Context, data *DnsptrrecResourceModel, diags *diag.Diagnostics) bool {
	dnsptrrecName := data.Id.ValueString()
	getResponseData, err := r.client.FindResource(service.Dnsptrrec.Type(), dnsptrrecName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsptrrec, got error: %s", err))
		return false
	}

	dnsptrrecSetAttrFromGet(ctx, data, getResponseData)

	return true
}
