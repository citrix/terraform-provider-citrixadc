package lbroute6

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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Lbroute6Resource{}
var _ resource.ResourceWithConfigure = (*Lbroute6Resource)(nil)
var _ resource.ResourceWithImportState = (*Lbroute6Resource)(nil)

func NewLbroute6Resource() resource.Resource {
	return &Lbroute6Resource{}
}

// Lbroute6Resource defines the resource implementation.
type Lbroute6Resource struct {
	client *service.NitroClient
}

func (r *Lbroute6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// SDK v2 ID scheme is the plain network value; passthrough it into the id
	// attribute so Read can resolve the resource by network.
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Lbroute6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbroute6"
}

func (r *Lbroute6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Lbroute6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Lbroute6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbroute6 resource")

	lbroute6 := lbroute6GetThePayloadFromtheConfig(ctx, &data)

	// Named resource identified by network — SDK v2 used AddResource with an
	// empty resource name.
	_, err := r.client.AddResource(service.Lbroute6.Type(), "", &lbroute6)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbroute6, got error: %s", err))
		return
	}

	// SDK v2 ID scheme: d.SetId(network)
	data.Id = types.StringValue(data.Network.ValueString())

	tflog.Trace(ctx, "Created lbroute6 resource")

	// Read the updated state back
	found := r.readLbroute6FromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// The resource was created but could not be read back; make sure no
		// computed attribute is left unknown to avoid an inconsistent-result
		// error, and preserve the plan values.
		if data.Td.IsUnknown() {
			data.Td = types.Int64Value(0)
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Lbroute6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Lbroute6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbroute6 resource")

	found := r.readLbroute6FromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// SDK v2 cleared state (d.SetId("")) when the route could not be found.
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Lbroute6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Lbroute6ResourceModel

	// Read Terraform prior state to preserve the ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbroute6 resource")

	// SDK v2 lbroute6 has no update path — every configurable attribute is
	// ForceNew, so any real change forces recreation via RequiresReplace. This
	// Update simply refreshes state from the ADC without issuing a NITRO write.
	found := r.readLbroute6FromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	tflog.Trace(ctx, "Updated lbroute6 resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Lbroute6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Lbroute6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbroute6 resource")

	// SDK v2 deleted via DeleteResourceWithArgsMap with the (url-escaped)
	// network as the disambiguating argument. The internal client does not
	// escape the args, so escape here to match SDK v2 exactly.
	argsMap := make(map[string]string)
	argsMap["network"] = url.QueryEscape(data.Network.ValueString())

	err := r.client.DeleteResourceWithArgsMap(service.Lbroute6.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbroute6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbroute6 resource")
}

// Helper function to read lbroute6 data from API.
// Returns true when the route is found, false when it does not exist (so the
// caller can remove it from state). Only genuine parse/response errors are
// surfaced via diags.
func (r *Lbroute6Resource) readLbroute6FromApi(ctx context.Context, data *Lbroute6ResourceModel, diags *diag.Diagnostics) bool {
	// SDK v2 ID scheme: the ID is the plain network value.
	network := data.Id.ValueString()

	findParams := service.FindParams{
		ResourceType: service.Lbroute6.Type(),
	}
	dataArray, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		// SDK v2 cleared state on any find error; treat as not-found.
		tflog.Warn(ctx, fmt.Sprintf("Unable to list lbroute6, clearing state for %s: %s", network, err.Error()))
		return false
	}

	if len(dataArray) == 0 {
		tflog.Warn(ctx, "lbroute6 does not exist; clearing state")
		return false
	}

	foundIndex := -1
	for i, v := range dataArray {
		if n, ok := v["network"].(string); ok && n == network {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		tflog.Warn(ctx, fmt.Sprintf("lbroute6 with network %s not found; clearing state", network))
		return false
	}

	lbroute6SetAttrFromGet(ctx, data, dataArray[foundIndex])

	return true
}
