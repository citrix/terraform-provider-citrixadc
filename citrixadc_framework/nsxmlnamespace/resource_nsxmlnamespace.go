package nsxmlnamespace

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
var _ resource.Resource = &NsxmlnamespaceResource{}
var _ resource.ResourceWithConfigure = (*NsxmlnamespaceResource)(nil)
var _ resource.ResourceWithImportState = (*NsxmlnamespaceResource)(nil)

func NewNsxmlnamespaceResource() resource.Resource {
	return &NsxmlnamespaceResource{}
}

// NsxmlnamespaceResource defines the resource implementation.
type NsxmlnamespaceResource struct {
	client *service.NitroClient
}

func (r *NsxmlnamespaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsxmlnamespaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsxmlnamespace"
}

func (r *NsxmlnamespaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsxmlnamespaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsxmlnamespaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsxmlnamespace resource")

	nsxmlnamespace := nsxmlnamespaceGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	prefix := data.Prefix.ValueString()
	_, err := r.client.AddResource(service.Nsxmlnamespace.Type(), prefix, &nsxmlnamespace)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsxmlnamespace, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsxmlnamespace resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(prefix)

	// Read the updated state back
	if !r.readNsxmlnamespaceFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsxmlnamespace not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsxmlnamespaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsxmlnamespaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsxmlnamespace resource")

	found := r.readNsxmlnamespaceFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NsxmlnamespaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NsxmlnamespaceResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsxmlnamespace resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Namespace.Equal(state.Namespace) {
		tflog.Debug(ctx, "namespace has changed for nsxmlnamespace")
		hasChange = true
	}
	if !data.Description.Equal(state.Description) {
		tflog.Debug(ctx, "description has changed for nsxmlnamespace")
		if config.Description.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "description")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		nsxmlnamespace := nsxmlnamespaceGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		prefix := data.Prefix.ValueString()
		_, err := r.client.UpdateResource(service.Nsxmlnamespace.Type(), prefix, &nsxmlnamespace)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsxmlnamespace, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated nsxmlnamespace resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsxmlnamespace resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"prefix": data.Prefix.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Nsxmlnamespace.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nsxmlnamespace attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNsxmlnamespaceFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsxmlnamespace not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsxmlnamespaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsxmlnamespaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsxmlnamespace resource")

	// Named resource - delete using DeleteResource
	prefix := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nsxmlnamespace.Type(), prefix)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsxmlnamespace, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsxmlnamespace resource")
}

// Helper function to read nsxmlnamespace data from API
func (r *NsxmlnamespaceResource) readNsxmlnamespaceFromApi(ctx context.Context, data *NsxmlnamespaceResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (prefix)
	prefixName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nsxmlnamespace.Type(), prefixName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsxmlnamespace, got error: %s", err))
		return false
	}

	nsxmlnamespaceSetAttrFromGet(ctx, data, getResponseData)

	return true
}
