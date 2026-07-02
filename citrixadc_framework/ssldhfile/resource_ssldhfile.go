package ssldhfile

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
var _ resource.Resource = &SsldhfileResource{}
var _ resource.ResourceWithConfigure = (*SsldhfileResource)(nil)
var _ resource.ResourceWithImportState = (*SsldhfileResource)(nil)

func NewSsldhfileResource() resource.Resource {
	return &SsldhfileResource{}
}

// SsldhfileResource defines the resource implementation.
type SsldhfileResource struct {
	client *service.NitroClient
}

func (r *SsldhfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SsldhfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssldhfile"
}

func (r *SsldhfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SsldhfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SsldhfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ssldhfile resource")
	ssldhfile := ssldhfileGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes ssldhfile create only via POST ?action=Import (no `add`).
	// Use ActOnResource with the case-sensitive "Import" verb.
	err := r.client.ActOnResource(service.Ssldhfile.Type(), &ssldhfile, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ssldhfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ssldhfile resource")

	// Set ID for the resource before reading state (plain value = name)
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	r.readSsldhfileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldhfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SsldhfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ssldhfile resource")

	r.readSsldhfileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldhfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for ssldhfile (only Import, delete,
	// get (all)). Every schema attribute is marked RequiresReplace, so Terraform
	// will never actually invoke Update with field changes. This body is a
	// documented no-op that preserves the prior ID and re-reads state.
	var data, state SsldhfileResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for ssldhfile; NITRO has no update endpoint and all attributes are RequiresReplace")

	r.readSsldhfileFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldhfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SsldhfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ssldhfile resource")
	// NITRO delete is keyless (DELETE /ssldhfile?args=name:<name>), not a URL-path key.
	args := []string{"name:" + utils.UrlEncode(data.Name.ValueString())}
	err := r.client.DeleteResourceWithArgs(service.Ssldhfile.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ssldhfile, got error: %s", err))
		return
	}
	tflog.Trace(ctx, "Deleted ssldhfile resource")
}

// Helper function to read ssldhfile data from API
func (r *SsldhfileResource) readSsldhfileFromApi(ctx context.Context, data *SsldhfileResourceModel, diags *diag.Diagnostics) {

	// ssldhfile has NO get-by-name endpoint (GET /ssldhfile/<name> => errorcode
	// 1090 "No such argument [arguid]"). Get all records and filter by name.
	name := data.Id.ValueString()

	allResources, err := r.client.FindAllResources(service.Ssldhfile.Type())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ssldhfile, got error: %s", err))
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
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ssldhfile: no record found with name %s", name))
		return
	}

	ssldhfileSetAttrFromGet(ctx, data, getResponseData)

}
