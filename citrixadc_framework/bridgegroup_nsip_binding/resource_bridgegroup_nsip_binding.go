package bridgegroup_nsip_binding

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
var _ resource.Resource = &BridgegroupNsipBindingResource{}
var _ resource.ResourceWithConfigure = (*BridgegroupNsipBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BridgegroupNsipBindingResource)(nil)

func NewBridgegroupNsipBindingResource() resource.Resource {
	return &BridgegroupNsipBindingResource{}
}

// BridgegroupNsipBindingResource defines the resource implementation.
type BridgegroupNsipBindingResource struct {
	client *service.NitroClient
}

func (r *BridgegroupNsipBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BridgegroupNsipBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgegroup_nsip_binding"
}

func (r *BridgegroupNsipBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BridgegroupNsipBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BridgegroupNsipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating bridgegroup_nsip_binding resource")
	bridgegroup_nsip_binding := bridgegroup_nsip_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Bridgegroup_nsip_binding.Type(), &bridgegroup_nsip_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create bridgegroup_nsip_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created bridgegroup_nsip_binding resource")

	// Set ID for the resource before reading state.
	// Composite ID uses the legacy attribute order (bridgegroup_id,ipaddress).
	data.Id = types.StringValue(bridgegroupNsipBindingComputeId(&data))

	// Read the updated state back
	r.readBridgegroupNsipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupNsipBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BridgegroupNsipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading bridgegroup_nsip_binding resource")

	r.readBridgegroupNsipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupNsipBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state BridgegroupNsipBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state. All attributes are RequiresReplace and this
	// binding has no NITRO update endpoint, so Update is effectively a no-op
	// read-back (Terraform forces recreation on any attribute change).
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for bridgegroup_nsip_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readBridgegroupNsipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupNsipBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BridgegroupNsipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting bridgegroup_nsip_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// The parent (bridge group integer key) is the NITRO resource name; the
	// bound nsip attributes are delete query args. ParseIdString handles both
	// the new key:value and the legacy "bridgegroup_id,ipaddress" forms.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"bridgegroup_id", "ipaddress"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	bridgegroupId, ok := idMap["bridgegroup_id"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'bridgegroup_id' not found in ID")
		return
	}

	// Build delete args. ipaddress comes from the ID; netmask/ownergroup/td come
	// from state (they are not encoded in the ID). DeleteResourceWithArgs
	// URL-encodes the args for the NITRO request.
	args := make([]string, 0)
	if val, ok := idMap["ipaddress"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ipaddress:%s", val))
	} else if !data.Ipaddress.IsNull() {
		args = append(args, fmt.Sprintf("ipaddress:%s", data.Ipaddress.ValueString()))
	}
	if !data.Netmask.IsNull() && data.Netmask.ValueString() != "" {
		args = append(args, fmt.Sprintf("netmask:%s", data.Netmask.ValueString()))
	}
	if !data.Ownergroup.IsNull() && data.Ownergroup.ValueString() != "" {
		args = append(args, fmt.Sprintf("ownergroup:%s", data.Ownergroup.ValueString()))
	}
	if !data.Td.IsNull() {
		args = append(args, fmt.Sprintf("td:%d", data.Td.ValueInt64()))
	}

	err = r.client.DeleteResourceWithArgs(service.Bridgegroup_nsip_binding.Type(), bridgegroupId, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete bridgegroup_nsip_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted bridgegroup_nsip_binding binding")
}

// Helper function to read bridgegroup_nsip_binding data from API
func (r *BridgegroupNsipBindingResource) readBridgegroupNsipBindingFromApi(ctx context.Context, data *BridgegroupNsipBindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse from ID (handles new + legacy formats).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"bridgegroup_id", "ipaddress"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	bridgegroupId, ok := idMap["bridgegroup_id"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'bridgegroup_id' not found in ID string")
		return
	}

	ipaddress, ok := idMap["ipaddress"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'ipaddress' not found in ID string")
		return
	}

	findParams := service.FindParams{
		ResourceType:             service.Bridgegroup_nsip_binding.Type(),
		ResourceName:             bridgegroupId,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read bridgegroup_nsip_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "bridgegroup_nsip_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the matching ipaddress.
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ipaddress"].(string); ok && val == ipaddress {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "bridgegroup_nsip_binding not found with the provided ID attributes")
		return
	}

	bridgegroup_nsip_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
