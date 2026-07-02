package sslcertbundle

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslcertbundleResource{}
var _ resource.ResourceWithConfigure = (*SslcertbundleResource)(nil)
var _ resource.ResourceWithImportState = (*SslcertbundleResource)(nil)

func NewSslcertbundleResource() resource.Resource {
	return &SslcertbundleResource{}
}

// SslcertbundleResource defines the resource implementation.
type SslcertbundleResource struct {
	client *service.NitroClient
}

func (r *SslcertbundleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcertbundleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertbundle"
}

func (r *SslcertbundleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcertbundleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcertbundleResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcertbundle resource")
	sslcertbundle := sslcertbundleGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Action-based resource - create via ?action=Import (capital I)
	err := r.client.ActOnResource(service.Sslcertbundle.Type(), &sslcertbundle, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcertbundle, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcertbundle resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readSslcertbundleFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertbundleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcertbundleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcertbundle resource")

	r.readSslcertbundleFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertbundleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslcertbundleResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// NITRO has no update endpoint for sslcertbundle; all attributes are RequiresReplace.
	tflog.Debug(ctx, "Update is a no-op for sslcertbundle; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSslcertbundleFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertbundleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcertbundleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcertbundle resource")
	// sslcertbundle has NO get-by-name endpoint and its DELETE does NOT accept the
	// resource name in the URL path (DELETE /sslcertbundle/<name> => errorcode 1090
	// "No such argument [arguid]"). The working delete form is
	// DELETE /sslcertbundle?args=name:<name> with an empty path name, which
	// DeleteResourceWithArgs produces. Its list pre-check (GET ?args=name:..) 400s
	// with errorcode 278 but transparently falls back to GET ?filter=name:.. which
	// succeeds, so the DELETE is correctly issued.
	name_value := data.Name.ValueString()
	args := []string{fmt.Sprintf("name:%s", name_value)}
	err := r.client.DeleteResourceWithArgs(service.Sslcertbundle.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcertbundle, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcertbundle resource")
}

// Helper function to read sslcertbundle data from API
func (r *SslcertbundleResource) readSslcertbundleFromApi(ctx context.Context, data *SslcertbundleResourceModel, diags *diag.Diagnostics) {

	// No get-byname endpoint - get all and filter by name
	name_value := data.Id.ValueString()

	allResources, err := r.client.FindAllResources(service.Sslcertbundle.Type())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcertbundle, got error: %s", err))
		return
	}

	var getResponseData map[string]interface{}
	for _, v := range allResources {
		if n, ok := v["name"].(string); ok && n == name_value {
			getResponseData = v
			break
		}
	}

	if getResponseData == nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcertbundle: no record found with name %s", name_value))
		return
	}

	sslcertbundleSetAttrFromGet(ctx, data, getResponseData)

}
