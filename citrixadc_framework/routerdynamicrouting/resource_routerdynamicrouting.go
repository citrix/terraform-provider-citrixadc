package routerdynamicrouting

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkv2resource "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RouterdynamicroutingResource{}
var _ resource.ResourceWithConfigure = (*RouterdynamicroutingResource)(nil)
var _ resource.ResourceWithImportState = (*RouterdynamicroutingResource)(nil)

func NewRouterdynamicroutingResource() resource.Resource {
	return &RouterdynamicroutingResource{}
}

// RouterdynamicroutingResource defines the resource implementation.
type RouterdynamicroutingResource struct {
	client *service.NitroClient
}

func (r *RouterdynamicroutingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RouterdynamicroutingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routerdynamicrouting"
}

func (r *RouterdynamicroutingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RouterdynamicroutingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RouterdynamicroutingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating routerdynamicrouting resource")

	// Build the apply-action payload from the plan (command lines joined by newlines)
	routerdynamicrouting := routerdynamicroutingGetThePayloadFromthePlan(ctx, &data)

	// routerdynamicrouting is an action-only resource: push the configuration via
	// the NITRO "apply" action (mirrors SDK v2 ActOnResource(..., "apply")).
	err := r.client.ActOnResource(service.Routerdynamicrouting.Type(), &routerdynamicrouting, "apply")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to apply routerdynamicrouting, got error: %s", err))
		return
	}

	// Generate a unique ID for this configuration resource (matches SDK v2 id format:
	// resource.PrefixedUniqueId("tf-routerdynamicrouting-")).
	data.Id = types.StringValue(sdkv2resource.PrefixedUniqueId("tf-routerdynamicrouting-"))

	// commandlines is Optional+Computed; if it was not configured, give it a
	// concrete (empty) value so the saved state is never unknown.
	if data.Commandlines.IsNull() || data.Commandlines.IsUnknown() {
		data.Commandlines = types.ListValueMust(types.StringType, []attr.Value{})
	}

	tflog.Trace(ctx, "Created routerdynamicrouting resource")

	// Save data into Terraform state. There is no GET for the applied config, so
	// the configured command lines are the source of truth (SDK v2 Read is a no-op).
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouterdynamicroutingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RouterdynamicroutingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading routerdynamicrouting resource")

	// routerdynamicrouting exposes no GET for the applied configuration; drift
	// detection is impossible by definition. Preserve the prior state unchanged
	// (mirrors SDK v2 Read = schema.Noop).
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouterdynamicroutingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// commandlines is RequiresReplaceIfConfigured, so Terraform never reaches
	// Update with a real change. This preserves identity and state defensively.
	var data, state RouterdynamicroutingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating routerdynamicrouting resource")

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Trace(ctx, "Updated routerdynamicrouting resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouterdynamicroutingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RouterdynamicroutingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting routerdynamicrouting resource")

	// routerdynamicrouting is a one-shot action resource with no NITRO delete
	// verb; removing it from state (done automatically by the framework) is the
	// correct behavior, mirroring the SDK v2 no-op delete.
	tflog.Trace(ctx, "Deleted routerdynamicrouting resource from state")
}
