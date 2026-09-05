package ssldhparam

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
var _ resource.Resource = &SsldhparamResource{}
var _ resource.ResourceWithConfigure = (*SsldhparamResource)(nil)
var _ resource.ResourceWithImportState = (*SsldhparamResource)(nil)

func NewSsldhparamResource() resource.Resource {
	return &SsldhparamResource{}
}

// SsldhparamResource defines the resource implementation.
type SsldhparamResource struct {
	client *service.NitroClient
}

func (r *SsldhparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SsldhparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssldhparam"
}

func (r *SsldhparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SsldhparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SsldhparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ssldhparam resource")

	ssldhparam := ssldhparamGetThePayloadFromtheConfig(ctx, &data)

	// ssldhparam is an action-only resource: NITRO only exposes the "create"
	// action (DH key file generation). There is no GET / update / delete
	// operation, matching the SDK v2 behaviour (ActOnResource "create").
	err := r.client.ActOnResource(service.Ssldhparam.Type(), &ssldhparam, "create")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ssldhparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ssldhparam resource")

	// gen is Optional+Computed but there is no GET to read it back. Resolve the
	// computed/unknown value to the NITRO default ("2") so the state is fully
	// known and Terraform does not report an inconsistent result after apply.
	if data.Gen.IsNull() || data.Gen.IsUnknown() {
		data.Gen = types.StringValue("2")
	}

	// ID is the dhfile value, matching SDK v2 d.SetId(dhfile).
	data.Id = types.StringValue(data.Dhfile.ValueString())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldhparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SsldhparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ssldhparam resource")

	// ssldhparam has no NITRO GET operation (action-only DH generation).
	// Mirror the SDK v2 schema.Noop Read: preserve prior state as-is without
	// clobbering configured values.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldhparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SsldhparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read Terraform prior state to preserve ID / computed values
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating ssldhparam resource")

	// All ssldhparam attributes are ForceNew and there is no NITRO update
	// operation, so any real change is handled via replace. This body only runs
	// for no-op plan differences; preserve prior ID and computed values.
	data.Id = state.Id
	if data.Gen.IsNull() || data.Gen.IsUnknown() {
		data.Gen = state.Gen
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldhparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// ssldhparam has no NITRO delete operation. Matching the SDK v2 no-op
	// delete, we simply let the framework drop the resource from state.
	tflog.Debug(ctx, "Deleting ssldhparam resource (state-only, no NITRO delete)")
}
