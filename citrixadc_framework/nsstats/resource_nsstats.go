package nsstats

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
var _ resource.Resource = &NsstatsResource{}
var _ resource.ResourceWithConfigure = (*NsstatsResource)(nil)
var _ resource.ResourceWithImportState = (*NsstatsResource)(nil)
var _ resource.ResourceWithValidateConfig = (*NsstatsResource)(nil)

func NewNsstatsResource() resource.Resource {
	return &NsstatsResource{}
}

// NsstatsResource defines the resource implementation.
type NsstatsResource struct {
	client *service.NitroClient
}

func (r *NsstatsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsstatsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsstats"
}

func (r *NsstatsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the cleanuplevel enum (global | all).
func (r *NsstatsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data NsstatsResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.Cleanuplevel.IsNull() && !data.Cleanuplevel.IsUnknown() {
		v := data.Cleanuplevel.ValueString()
		if v != "global" && v != "all" {
			resp.Diagnostics.AddAttributeError(
				path.Root("cleanuplevel"),
				"Invalid cleanuplevel",
				fmt.Sprintf("cleanuplevel must be one of [global, all], got: %q", v),
			)
		}
	}
}

func (r *NsstatsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsstatsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing nsstats (action-only resource)")
	payload := nsstatsGetThePayloadFromthePlan(ctx, &data)

	// nsstats is an action-only resource (clear, no add/get).
	// The Create maps to the "clear" action.
	err := r.client.ActOnResource(service.Nsstats.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear nsstats, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared nsstats resource")

	// Generate a synthetic ID; nsstats has no GET endpoint.
	data.Id = types.StringValue("nsstats-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsstatsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsstatsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op: nsstats is an action-only resource with no GET endpoint.
	tflog.Debug(ctx, "Read is a no-op for nsstats (no GET endpoint); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsstatsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsstatsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op for nsstats; it is an action-only resource.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsstats; action-only resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsstatsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsstatsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op: nsstats is action-only with no inverse API; just remove from state.
	tflog.Debug(ctx, "Delete is a no-op for nsstats; removing from Terraform state")
}
