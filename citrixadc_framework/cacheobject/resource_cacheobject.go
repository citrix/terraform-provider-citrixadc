package cacheobject

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
var _ resource.Resource = &CacheobjectResource{}
var _ resource.ResourceWithConfigure = (*CacheobjectResource)(nil)
var _ resource.ResourceWithImportState = (*CacheobjectResource)(nil)
var _ resource.ResourceWithValidateConfig = (*CacheobjectResource)(nil)

func NewCacheobjectResource() resource.Resource {
	return &CacheobjectResource{}
}

// CacheobjectResource defines the resource implementation.
//
// cacheobject is an ACTION-ONLY runtime object of the integrated cache.
// NITRO exposes ONLY get(all), count, and the POST actions expire/flush/save.
// There is NO add, NO update/set, NO delete. Cached objects are created by the
// traffic engine, not the config API. This resource therefore fires the chosen
// action (expire|flush|save) on Create and treats Read/Update/Delete as no-ops.
type CacheobjectResource struct {
	client *service.NitroClient
}

func (r *CacheobjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CacheobjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cacheobject"
}

func (r *CacheobjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the NITRO expire/flush mutually-exclusive mandatory
// choice: (locator) XOR (url + host). save has no such requirement (save-all is
// allowed with no args).
func (r *CacheobjectResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data CacheobjectResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	action := "flush"
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		action = data.Action.ValueString()
	}

	if action == "expire" || action == "flush" {
		hasLocator := !data.Locator.IsNull() && !data.Locator.IsUnknown()
		hasUrl := !data.Url.IsNull() && !data.Url.IsUnknown()
		hasHost := !data.Host.IsNull() && !data.Host.IsUnknown()

		if hasLocator && (hasUrl || hasHost) {
			resp.Diagnostics.AddError(
				"Invalid cacheobject configuration",
				fmt.Sprintf("For action %q specify either \"locator\" OR (\"url\" + \"host\"), not both.", action),
			)
			return
		}
		if !hasLocator {
			if !(hasUrl && hasHost) {
				resp.Diagnostics.AddError(
					"Invalid cacheobject configuration",
					fmt.Sprintf("For action %q you must specify either \"locator\" OR both \"url\" and \"host\".", action),
				)
				return
			}
		}
	}
}

func (r *CacheobjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CacheobjectResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	action := "flush"
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		action = data.Action.ValueString()
	}

	tflog.Debug(ctx, fmt.Sprintf("Firing cacheobject action %q", action))

	// Build the action payload. expire/flush accept locator | (url,host[,port,groupname,httpmethod]).
	// save accepts [locator] [tosecondary]. We include only the fields relevant to each action.
	cacheobject := cacheobjectGetActionPayload(ctx, &data, action)

	// cacheobject has no add/set/delete; the only writes are POST ?action=<verb>.
	err := r.client.ActOnResource(service.Cacheobject.Type(), &cacheobject, action)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to perform %q on cacheobject, got error: %s", action, err))
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("Performed cacheobject action %q", action))

	// Synthetic ID: this is an action, not a persistent object. Use action plus the
	// primary identity the caller supplied so repeated applies with different args
	// produce distinct IDs.
	id := action
	if !data.Locator.IsNull() && !data.Locator.IsUnknown() {
		id = fmt.Sprintf("%s:locator:%d", action, data.Locator.ValueInt64())
	} else if !data.Url.IsNull() && !data.Url.IsUnknown() {
		id = fmt.Sprintf("%s:url:%s", action, data.Url.ValueString())
	}
	data.Id = types.StringValue(id)

	// Read is a no-op (cached objects are ephemeral and not reliably re-findable);
	// persist the plan as-is.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheobjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CacheobjectResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op. cacheobject is an action-only runtime object: the traffic
	// engine owns its lifecycle, and a fired action (expire/flush/save) leaves no
	// re-findable object keyed to our synthetic ID. Preserve prior state unchanged.
	tflog.Debug(ctx, "Read is a no-op for cacheobject (action-only resource, no re-findable object)")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheobjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CacheobjectResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op: cacheobject has no set endpoint and every input attribute
	// is RequiresReplace, so a meaningful change forces recreation and never reaches
	// here.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for cacheobject; all attributes are RequiresReplace and there is no set endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheobjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CacheobjectResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op: cacheobject exposes no delete operation. The prior action
	// (expire/flush/save) is not reversible via the config API; removing the
	// resource simply drops it from state.
	tflog.Debug(ctx, "Delete is a no-op for cacheobject (no delete operation on NITRO side)")
}
