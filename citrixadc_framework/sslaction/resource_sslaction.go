package sslaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslactionResource{}
var _ resource.ResourceWithConfigure = (*SslactionResource)(nil)
var _ resource.ResourceWithImportState = (*SslactionResource)(nil)

func NewSslactionResource() resource.Resource {
	return &SslactionResource{}
}

// SslactionResource defines the resource implementation.
type SslactionResource struct {
	client *service.NitroClient
}

func (r *SslactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslaction"
}

func (r *SslactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslactionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslaction resource")

	sslactionName := data.Name.ValueString()
	sslaction := sslactionGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - NITRO add (HTTP POST). sslaction has no set/update command.
	_, err := r.client.AddResource(service.Sslaction.Type(), sslactionName, &sslaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslaction, got error: %s", err))
		return
	}

	data.Id = types.StringValue(sslactionName)

	tflog.Trace(ctx, "Created sslaction resource")

	if !r.readSslactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslaction not found immediately after create")
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslactionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslaction resource")

	found := r.readSslactionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update: sslaction is immutable on the appliance (all attributes are ForceNew /
// RequiresReplace and NITRO exposes no set command). Terraform therefore replaces
// the resource on any change and never actually invokes Update in practice; this
// implementation simply refreshes state without issuing a NITRO write.
func (r *SslactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslactionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslaction resource (no-op: immutable resource)")

	if !r.readSslactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslaction not found during update")
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslactionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslaction resource")

	err := r.client.DeleteResource(service.Sslaction.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslaction %s, got error: %s", data.Id.ValueString(), err))
		return
	}

	tflog.Trace(ctx, "Deleted sslaction resource")
}

// readSslactionFromApi reads the sslaction into the model. Returns false when the
// resource no longer exists on the appliance.
func (r *SslactionResource) readSslactionFromApi(ctx context.Context, data *SslactionResourceModel, diags *diag.Diagnostics) bool {
	sslactionName := data.Id.ValueString()
	if sslactionName == "" {
		sslactionName = data.Name.ValueString()
	}

	getResponseData, err := r.client.FindResource(service.Sslaction.Type(), sslactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslaction %s, got error: %s", sslactionName, err))
		return false
	}
	if getResponseData == nil {
		return false
	}

	sslactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
