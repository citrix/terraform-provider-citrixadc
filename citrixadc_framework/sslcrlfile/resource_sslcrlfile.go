package sslcrlfile

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
var _ resource.Resource = &SslcrlfileResource{}
var _ resource.ResourceWithConfigure = (*SslcrlfileResource)(nil)
var _ resource.ResourceWithImportState = (*SslcrlfileResource)(nil)

func NewSslcrlfileResource() resource.Resource {
	return &SslcrlfileResource{}
}

// SslcrlfileResource defines the resource implementation.
type SslcrlfileResource struct {
	client *service.NitroClient
}

func (r *SslcrlfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcrlfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcrlfile"
}

func (r *SslcrlfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcrlfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcrlfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcrlfile resource")
	sslcrlfile := sslcrlfileGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes sslcrlfile create only via POST ?action=Import (no `add`).
	// Use ActOnResource with the case-sensitive "Import" verb.
	err := r.client.ActOnResource(service.Sslcrlfile.Type(), &sslcrlfile, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcrlfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcrlfile resource")

	// Set ID for the resource before reading state (plain value = name)
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	r.readSslcrlfileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcrlfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcrlfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcrlfile resource")

	r.readSslcrlfileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcrlfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for sslcrlfile (only Import, delete,
	// get (all)). Every schema attribute is marked RequiresReplace, so Terraform
	// will never actually invoke Update with field changes. This body is a
	// documented no-op that preserves the prior ID and re-reads state.
	var data, state SslcrlfileResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslcrlfile; NITRO has no update endpoint and all attributes are RequiresReplace")

	r.readSslcrlfileFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcrlfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcrlfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcrlfile resource")
	// NITRO delete is keyless (DELETE /sslcrlfile?args=name:<name>), not a URL-path key.
	args := []string{"name:" + utils.UrlEncode(data.Name.ValueString())}
	err := r.client.DeleteResourceWithArgs(service.Sslcrlfile.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcrlfile, got error: %s", err))
		return
	}
	tflog.Trace(ctx, "Deleted sslcrlfile resource")
}

// Helper function to read sslcrlfile data from API
func (r *SslcrlfileResource) readSslcrlfileFromApi(ctx context.Context, data *SslcrlfileResourceModel, diags *diag.Diagnostics) {

	// sslcrlfile has NO get-by-name endpoint (GET /sslcrlfile/<name> => errorcode
	// 1090 "No such argument [arguid]"). Get all records and filter by name.
	name := data.Id.ValueString()

	allResources, err := r.client.FindAllResources(service.Sslcrlfile.Type())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcrlfile, got error: %s", err))
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
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcrlfile: no record found with name %s", name))
		return
	}

	sslcrlfileSetAttrFromGet(ctx, data, getResponseData)

}
