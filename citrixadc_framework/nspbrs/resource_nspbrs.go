package nspbrs

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkid "github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NspbrsResource{}
var _ resource.ResourceWithConfigure = (*NspbrsResource)(nil)
var _ resource.ResourceWithImportState = (*NspbrsResource)(nil)

func NewNspbrsResource() resource.Resource {
	return &NspbrsResource{}
}

// NspbrsResource defines the resource implementation.
type NspbrsResource struct {
	client *service.NitroClient
}

func (r *NspbrsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NspbrsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nspbrs"
}

func (r *NspbrsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NspbrsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NspbrsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nspbrs resource")

	// nspbrs is action-only. The payload is an empty struct; ActOnResource wraps
	// it as {"nspbrs":{}} and POSTs to ?action=<action>, matching both the NITRO
	// doc and the SDK v2 behavior.
	nspbrs := ns.Nspbrs{}

	var err error
	action := data.Action.ValueString()
	switch action {
	case "apply":
		err = r.client.ActOnResource(service.Nspbrs.Type(), &nspbrs, "apply")
	case "clear":
		err = r.client.ActOnResource(service.Nspbrs.Type(), &nspbrs, "clear")
	case "renumber":
		err = r.client.ActOnResource(service.Nspbrs.Type(), &nspbrs, "renumber")
	default:
		// Mirror the SDK v2 validation/error message exactly.
		resp.Diagnostics.AddError(
			"Configuration Error",
			fmt.Sprintf("Invalid value for action %s. Supported values of action are `apply`, `clear` or `renumber`", action),
		)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to perform nspbrs %s action, got error: %s", action, err))
		return
	}

	// Generate a unique ID for this resource (matches SDK v2 id format:
	// resource.PrefixedUniqueId("tf-nspbrs-")).
	data.Id = types.StringValue(sdkid.PrefixedUniqueId("tf-nspbrs-"))

	tflog.Trace(ctx, "Created nspbrs resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspbrsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// nspbrs exposes no NITRO GET verb (action-only). Read is a no-op that
	// preserves the prior state unchanged, mirroring the SDK v2 schema.Noop read.
	var data NspbrsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nspbrs resource (no-op: nspbrs has no GET verb)")

	// Save state back unchanged
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspbrsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// action is RequiresReplace (SDK v2 ForceNew), so Terraform never reaches
	// Update with a real change. This preserves the identity defensively.
	var data, state NspbrsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nspbrs resource (no-op)")

	// Preserve ID from prior state
	data.Id = state.Id

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspbrsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// nspbrs is a one-shot action resource with no NITRO delete verb. Removing it
	// from state (done automatically by the framework once Delete returns without
	// error) is the correct behavior, mirroring the SDK v2 schema.Noop delete.
	var data NspbrsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "Deleted nspbrs resource from state (no-op)")
}
