package nat64

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
var _ resource.Resource = &Nat64Resource{}
var _ resource.ResourceWithConfigure = (*Nat64Resource)(nil)
var _ resource.ResourceWithImportState = (*Nat64Resource)(nil)

func NewNat64Resource() resource.Resource {
	return &Nat64Resource{}
}

// Nat64Resource defines the resource implementation.
type Nat64Resource struct {
	client *service.NitroClient
}

func (r *Nat64Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nat64Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nat64"
}

func (r *Nat64Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nat64Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nat64ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nat64 resource")

	nat64 := nat64GetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	nat64Name := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nat64.Type(), nat64Name, &nat64)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nat64, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nat64 resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(nat64Name)

	// Read the updated state back
	if !r.readNat64FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nat64 not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nat64Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nat64ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nat64 resource")

	found := r.readNat64FromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Nat64Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state Nat64ResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read raw config to detect attributes removed from configuration (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nat64 resource")

	// Check if there are any changes in updateable attributes (name is RequiresReplace)
	hasChange := false
	attributesToUnset := []string{}
	if !data.Acl6name.Equal(state.Acl6name) {
		tflog.Debug(ctx, "acl6name has changed for nat64")
		hasChange = true
	}
	if !data.Netprofile.Equal(state.Netprofile) {
		tflog.Debug(ctx, "netprofile has changed for nat64")
		if config.Netprofile.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "netprofile")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		nat64 := nat64GetThePayloadFromtheConfig(ctx, &data)
		// Named resource - use UpdateResource
		nat64Name := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Nat64.Type(), nat64Name, &nat64)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nat64, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated nat64 resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nat64 resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Nat64.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nat64 attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNat64FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nat64 not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nat64Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nat64ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nat64 resource")

	// Named resource - delete using DeleteResource (matches SDK v2 DeleteResource(type, d.Id()))
	nat64Name := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nat64.Type(), nat64Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nat64, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nat64 resource")
}

// Helper function to read nat64 data from API. Returns false when the resource no longer exists.
func (r *Nat64Resource) readNat64FromApi(ctx context.Context, data *Nat64ResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (name)
	nat64Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nat64.Type(), nat64Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nat64, got error: %s", err))
		return false
	}

	nat64SetAttrFromGet(ctx, data, getResponseData)

	return true
}
