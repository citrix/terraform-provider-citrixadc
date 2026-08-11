package dnsmxrec

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnsmxrecResource{}
var _ resource.ResourceWithConfigure = (*DnsmxrecResource)(nil)
var _ resource.ResourceWithImportState = (*DnsmxrecResource)(nil)

func NewDnsmxrecResource() resource.Resource {
	return &DnsmxrecResource{}
}

// DnsmxrecResource defines the resource implementation.
type DnsmxrecResource struct {
	client *service.NitroClient
}

func (r *DnsmxrecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsmxrecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsmxrec"
}

func (r *DnsmxrecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsmxrecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsmxrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsmxrec resource")

	dnsmxrec := dnsmxrecGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource (NITRO add is POST)
	dnsmxrecName := data.Domain.ValueString()
	_, err := r.client.AddResource(service.Dnsmxrec.Type(), dnsmxrecName, &dnsmxrec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsmxrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnsmxrec resource")

	// Set ID for the resource before reading state (plain value = domain, matches SDK v2)
	data.Id = types.StringValue(dnsmxrecName)

	// Read the updated state back
	if !r.readDnsmxrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsmxrec not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsmxrecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsmxrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsmxrec resource")

	found := r.readDnsmxrecFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnsmxrecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state DnsmxrecResourceModel

	// Read Terraform prior state to detect changes and preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read Terraform config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnsmxrec resource")

	dnsmxrec, hasChange := dnsmxrecGetTheUpdatablePayloadFromThePlan(ctx, &data, &state)

	// Detect optional attributes removed from config so they can be unset
	// (reverted to their NITRO defaults) after any update.
	attributesToUnset := []string{}
	if !data.Ttl.Equal(state.Ttl) && config.Ttl.IsNull() {
		attributesToUnset = append(attributesToUnset, "ttl")
	}

	if hasChange {
		// Named resource - use UpdateResource (domain is the resource name)
		dnsmxrecName := data.Domain.ValueString()
		_, err := r.client.UpdateResource(service.Dnsmxrec.Type(), dnsmxrecName, &dnsmxrec)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsmxrec, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated dnsmxrec resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnsmxrec resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. The unset payload keys on domain and mx (both mandatory in
	// the NITRO dnsmxrec unset operation).
	unsetIdPayload := map[string]interface{}{
		"domain": data.Domain.ValueString(),
		"mx":     data.Mx.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Dnsmxrec.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset dnsmxrec attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readDnsmxrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsmxrec not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsmxrecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsmxrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsmxrec resource")

	// Named resource - delete keyed on domain, with mx (required) and ecssubnet
	// (optional) as delete args, matching the NITRO delete endpoint and SDK v2.
	dnsmxrecName := data.Id.ValueString()
	argsMap := make(map[string]string)
	argsMap["mx"] = url.QueryEscape(data.Mx.ValueString())
	if !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown() && data.Ecssubnet.ValueString() != "" {
		argsMap["ecssubnet"] = url.QueryEscape(data.Ecssubnet.ValueString())
	}

	err := r.client.DeleteResourceWithArgsMap(service.Dnsmxrec.Type(), dnsmxrecName, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsmxrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsmxrec resource")
}

// Helper function to read dnsmxrec data from API.
// Returns false (without adding an error) when the resource no longer exists.
func (r *DnsmxrecResource) readDnsmxrecFromApi(ctx context.Context, data *DnsmxrecResourceModel, diags *diag.Diagnostics) bool {
	// ID is the plain domain value (matches SDK v2)
	dnsmxrecName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Dnsmxrec.Type(), dnsmxrecName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsmxrec, got error: %s", err))
		return false
	}

	dnsmxrecSetAttrFromGet(ctx, data, getResponseData)

	return true
}
