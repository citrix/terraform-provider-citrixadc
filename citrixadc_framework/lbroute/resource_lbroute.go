package lbroute

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LbrouteResource{}
var _ resource.ResourceWithConfigure = (*LbrouteResource)(nil)
var _ resource.ResourceWithImportState = (*LbrouteResource)(nil)

func NewLbrouteResource() resource.Resource {
	return &LbrouteResource{}
}

// LbrouteResource defines the resource implementation.
type LbrouteResource struct {
	client *service.NitroClient
}

func (r *LbrouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// lbroute has no single-key GET endpoint; Read matches the enumerated array on
	// (network, netmask, gatewayname). A bare passthrough would populate only id,
	// leaving those key attributes null so Read finds nothing and drops the
	// resource. Parse the composite id "network,netmask,gatewayname" (the same
	// format Create/Read set as the canonical id) into the key attributes so Read
	// can locate the route. SplitN keeps a comma-bearing gatewayname intact.
	parts := strings.SplitN(req.ID, ",", 3)
	if len(parts) != 3 {
		resp.Diagnostics.AddError(
			"Invalid import ID for lbroute",
			fmt.Sprintf("Expected import ID in the format \"network,netmask,gatewayname\", got %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network"), types.StringValue(parts[0]))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("netmask"), types.StringValue(parts[1]))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("gatewayname"), types.StringValue(parts[2]))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(req.ID))...)
}

func (r *LbrouteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbroute"
}

func (r *LbrouteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbrouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbrouteResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbroute resource")

	// Build payload from the plan (td only when explicitly configured)
	lbroute := lbrouteGetThePayloadFromtheConfig(ctx, &data)

	// The SDK v2 resource used "network,netmask,gatewayname" as the NITRO
	// resource name for the AddResource call. Preserve that behavior.
	lbrouteName := fmt.Sprintf("%s,%s,%s", data.Network.ValueString(), data.Netmask.ValueString(), data.Gatewayname.ValueString())

	_, err := r.client.AddResource(service.Lbroute.Type(), lbrouteName, &lbroute)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbroute, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbroute resource")

	// Read the created state back (also sets the ID to the SDK v2 composite)
	if !r.readLbrouteFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbroute not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbrouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbrouteResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbroute resource")

	found := r.readLbrouteFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LbrouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbrouteResourceModel

	// Read Terraform prior state to preserve the ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// lbroute has no NITRO-updatable attributes (every attribute is ForceNew in
	// SDK v2 and RequiresReplace here), so Update never performs a write. It is
	// only present to satisfy the resource.Resource interface. Read state back.
	tflog.Debug(ctx, "Updating lbroute resource (read-back only; no updatable attributes)")

	if !r.readLbrouteFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbroute not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbrouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbrouteResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbroute resource")

	// Only the network and netmask properties are required for deletion - not
	// gatewayname. This mirrors the SDK v2 resource exactly.
	argsMap := make(map[string]string)
	argsMap["network"] = url.QueryEscape(data.Network.ValueString())
	argsMap["netmask"] = url.QueryEscape(data.Netmask.ValueString())

	err := r.client.DeleteResourceWithArgsMap(service.Lbroute.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbroute, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbroute resource")
}

// Helper function to read lbroute data from API.
// Returns true if the route was found, false otherwise.
func (r *LbrouteResource) readLbrouteFromApi(ctx context.Context, data *LbrouteResourceModel, diags *diag.Diagnostics) bool {
	// lbroute has no GET-by-name endpoint; enumerate the array and filter on the
	// identity attributes (network, netmask, gatewayname) exactly as SDK v2 did.
	findParams := service.FindParams{
		ResourceType: service.Lbroute.Type(),
	}
	dataArray, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Clearing lbroute state, got error: %s", err))
		return false
	}
	if len(dataArray) == 0 {
		tflog.Warn(ctx, "lbroute does not exist. Clearing state.")
		return false
	}

	foundIndex := -1
	for i, v := range dataArray {
		match := true
		if v["network"] != data.Network.ValueString() {
			match = false
		}
		if v["netmask"] != data.Netmask.ValueString() {
			match = false
		}
		if v["gatewayname"] != data.Gatewayname.ValueString() {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		tflog.Warn(ctx, "lbroute not found in array. Clearing state.")
		return false
	}

	lbrouteSetAttrFromGet(ctx, data, dataArray[foundIndex])

	return true
}
