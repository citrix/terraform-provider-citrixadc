package sslcertfile

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
var _ resource.Resource = &SslcertfileResource{}
var _ resource.ResourceWithConfigure = (*SslcertfileResource)(nil)
var _ resource.ResourceWithImportState = (*SslcertfileResource)(nil)

func NewSslcertfileResource() resource.Resource {
	return &SslcertfileResource{}
}

// SslcertfileResource defines the resource implementation.
type SslcertfileResource struct {
	client *service.NitroClient
}

func (r *SslcertfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcertfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertfile"
}

func (r *SslcertfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcertfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcertfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcertfile resource")

	sslcertfile := sslcertfileGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes sslcertfile create only via POST ?action=import (no `add`).
	// This matches the legacy SDK v2 resource which used
	// client.ActOnResource(..., "import"). Verb casing is case-sensitive.
	err := r.client.ActOnResource(service.Sslcertfile.Type(), &sslcertfile, "import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcertfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcertfile resource")

	// Set ID for the resource before reading state (plain value = name), matching
	// the legacy SDK v2 ID scheme d.SetId(name).
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	r.readSslcertfileFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcertfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcertfile resource")

	r.readSslcertfileFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Object is gone out-of-band; remove from state.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for sslcertfile (only import, delete,
	// get (all)). Every schema attribute (name, src) is marked RequiresReplace,
	// exactly matching the legacy SDK v2 ForceNew flags, so Terraform will never
	// actually invoke Update with field changes. This body is a documented no-op
	// that preserves the prior ID and re-reads state.
	var data, state SslcertfileResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslcertfile; NITRO has no update endpoint and all attributes are RequiresReplace")

	r.readSslcertfileFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcertfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcertfile resource")
	// NITRO delete is keyless (DELETE /sslcertfile?args=name:<name>), not a
	// URL-path key. This matches the legacy SDK v2
	// DeleteResourceWithArgs(type, "", ["name:<name>"]).
	args := []string{"name:" + utils.UrlEncode(data.Name.ValueString())}
	err := r.client.DeleteResourceWithArgs(service.Sslcertfile.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcertfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcertfile resource")
}

// Helper function to read sslcertfile data from API
func (r *SslcertfileResource) readSslcertfileFromApi(ctx context.Context, data *SslcertfileResourceModel, diags *diag.Diagnostics) {

	// sslcertfile has no reliable get-by-name endpoint; the legacy SDK v2
	// resource used FindAllResources + filter by name. Mirror that here.
	name := data.Id.ValueString()

	allResources, err := r.client.FindAllResources(service.Sslcertfile.Type())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcertfile, got error: %s", err))
		return
	}

	var getResponseData map[string]interface{}
	for _, v := range allResources {
		if n, ok := v["name"].(string); ok && n == name {
			getResponseData = v
			break
		}
	}

	if getResponseData == nil {
		// Object is gone out-of-band; signal removal via null Id.
		data.Id = types.StringNull()
		return
	}

	sslcertfileSetAttrFromGet(ctx, data, getResponseData)
}
