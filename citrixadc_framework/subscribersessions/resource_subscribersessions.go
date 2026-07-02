package subscribersessions

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SubscribersessionsResource{}
var _ resource.ResourceWithConfigure = (*SubscribersessionsResource)(nil)
var _ resource.ResourceWithImportState = (*SubscribersessionsResource)(nil)

func NewSubscribersessionsResource() resource.Resource {
	return &SubscribersessionsResource{}
}

// SubscribersessionsResource defines the resource implementation.
//
// subscribersessions is an action-only resource. NITRO exposes only get(all),
// count, and ?action=clear (POST). There is NO add, NO set/update, NO delete
// endpoint. The resource therefore models the "clear" action (Pattern 13):
// Create fires ?action=clear; Read/Update/Delete are no-ops. The read-only
// telemetry (subscriptionidtype, subscriptionidvalue, subscriberrules, flags,
// ttl, idlettl, avpdisplaybuffer, servicepath) is exposed by the datasource,
// not the resource.
type SubscribersessionsResource struct {
	client *service.NitroClient
}

func (r *SubscribersessionsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SubscribersessionsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscribersessions"
}

func (r *SubscribersessionsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SubscribersessionsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SubscribersessionsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating subscribersessions resource (firing ?action=clear)")
	payload := subscribersessionsGetThePayloadFromthePlan(ctx, &data)

	// subscribersessions has no add/set endpoint. Fire the clear action.
	// NITRO verb is case-sensitive: "clear".
	err := r.client.ActOnResource(service.Subscribersessions.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear subscribersessions, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared subscribersessions")

	// Synthetic ID: there is no persistent object to key off. Derive from the
	// clear selectors so distinct clears are distinguishable, but the resource
	// is otherwise transient.
	switch {
	case !data.Ip.IsNull() && !data.Ip.IsUnknown() && data.Ip.ValueString() != "":
		if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
			data.Id = types.StringValue(fmt.Sprintf("subscribersessions-clear-%s-%d", data.Ip.ValueString(), data.Vlan.ValueInt64()))
		} else {
			data.Id = types.StringValue(fmt.Sprintf("subscribersessions-clear-%s", data.Ip.ValueString()))
		}
	case !data.Vlan.IsNull() && !data.Vlan.IsUnknown():
		data.Id = types.StringValue(fmt.Sprintf("subscribersessions-clear-vlan-%d", data.Vlan.ValueInt64()))
	default:
		data.Id = types.StringValue("subscribersessions-clear-all")
	}

	// Save data into Terraform state (no Read: transient table, no GET-by-key).
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. subscribersessions is a transient diagnostics table with no
// GET-by-key endpoint; the cleared sessions are not a persistent object, so
// there is nothing to reconcile. Preserve prior state unchanged.
func (r *SubscribersessionsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SubscribersessionsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for subscribersessions (action-only clear, transient table)")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. All attributes (ip, vlan) are RequiresReplace, so
// Terraform never routes a change through Update. There is no set endpoint.
func (r *SubscribersessionsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SubscribersessionsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for subscribersessions; all attributes are RequiresReplace and there is no set endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. subscribersessions has no delete endpoint; the clear action
// is fire-and-forget. Removing the resource simply drops it from state.
func (r *SubscribersessionsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for subscribersessions (no delete endpoint on NITRO side)")
}
