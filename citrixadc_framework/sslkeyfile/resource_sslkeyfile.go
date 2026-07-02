package sslkeyfile

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
var _ resource.Resource = &SslkeyfileResource{}
var _ resource.ResourceWithConfigure = (*SslkeyfileResource)(nil)
var _ resource.ResourceWithImportState = (*SslkeyfileResource)(nil)

func NewSslkeyfileResource() resource.Resource {
	return &SslkeyfileResource{}
}

// SslkeyfileResource defines the resource implementation.
type SslkeyfileResource struct {
	client *service.NitroClient
}

func (r *SslkeyfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslkeyfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslkeyfile"
}

func (r *SslkeyfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslkeyfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslkeyfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslkeyfile resource")
	// Get payload from plan (regular attributes)
	sslkeyfile := sslkeyfileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslkeyfileGetThePayloadFromtheConfig(ctx, &config, &sslkeyfile)

	// NITRO exposes sslkeyfile create only via POST ?action=Import (no `add`).
	// Use ActOnResource with the case-sensitive "Import" verb.
	err := r.client.ActOnResource(service.Sslkeyfile.Type(), &sslkeyfile, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslkeyfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslkeyfile resource")

	// Set ID for the resource before reading state (plain value = name)
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	r.readSslkeyfileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslkeyfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslkeyfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslkeyfile resource")

	r.readSslkeyfileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslkeyfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for sslkeyfile (only Import, delete,
	// get (all)). Every schema attribute is marked RequiresReplace, so Terraform
	// will never actually invoke Update with field changes. This body is a
	// documented no-op that preserves the prior ID and re-reads state.
	var data, state SslkeyfileResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslkeyfile; NITRO has no update endpoint and all attributes are RequiresReplace")

	r.readSslkeyfileFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslkeyfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslkeyfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslkeyfile resource")
	// NITRO delete is keyless (DELETE /sslkeyfile?args=name:<name>), not a URL-path key.
	args := []string{"name:" + utils.UrlEncode(data.Name.ValueString())}
	err := r.client.DeleteResourceWithArgs(service.Sslkeyfile.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslkeyfile, got error: %s", err))
		return
	}
	tflog.Trace(ctx, "Deleted sslkeyfile resource")
}

// Helper function to read sslkeyfile data from API
func (r *SslkeyfileResource) readSslkeyfileFromApi(ctx context.Context, data *SslkeyfileResourceModel, diags *diag.Diagnostics) {

	// sslkeyfile has NO get-by-name endpoint (GET /sslkeyfile/<name> => errorcode
	// 1090 "No such argument [arguid]"). Get all records and filter by name.
	name := data.Id.ValueString()

	allResources, err := r.client.FindAllResources(service.Sslkeyfile.Type())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslkeyfile, got error: %s", err))
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
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslkeyfile: no record found with name %s", name))
		return
	}

	sslkeyfileSetAttrFromGet(ctx, data, getResponseData)

}
