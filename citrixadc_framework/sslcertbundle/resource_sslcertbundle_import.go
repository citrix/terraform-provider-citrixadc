package sslcertbundle

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslcertbundleImportResource{}
var _ resource.ResourceWithConfigure = (*SslcertbundleImportResource)(nil)
var _ resource.ResourceWithImportState = (*SslcertbundleImportResource)(nil)

func NewSslcertbundleImportResource() resource.Resource {
	return &SslcertbundleImportResource{}
}

// SslcertbundleImportResource defines the resource implementation.
//
// This is a FULL managed resource modelling the NITRO sslcertbundle object,
// created via ?action=Import. It preserves real CRUD: Create imports the bundle,
// Read reflects it from GET (all + filter by name), Update is a no-op (all
// attributes RequiresReplace), and Delete removes the bundle from the appliance.
type SslcertbundleImportResource struct {
	client *service.NitroClient
}

// SslcertbundleImportResourceModel describes the resource data model.
type SslcertbundleImportResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Src  types.String `tfsdk:"src"`
}

func (r *SslcertbundleImportResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcertbundleImportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertbundle_import"
}

func (r *SslcertbundleImportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcertbundleImportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcertbundle resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported certificate bundle. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL specifying the protocol, host, and path, including file name, to the certificate bundle to be imported or exported. For example, http://www.example.com/cert_bundle_file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
		},
	}
}

func (r *SslcertbundleImportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcertbundleImportResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcertbundle resource")
	sslcertbundle := sslcertbundleImportGetThePayloadFromthePlan(ctx, &data)

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
	r.readSslcertbundleImportFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertbundleImportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcertbundleImportResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcertbundle resource")

	r.readSslcertbundleImportFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertbundleImportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslcertbundleImportResourceModel

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
	r.readSslcertbundleImportFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertbundleImportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcertbundleImportResourceModel

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
func (r *SslcertbundleImportResource) readSslcertbundleImportFromApi(ctx context.Context, data *SslcertbundleImportResourceModel, diags *diag.Diagnostics) {

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
		// Object is gone out-of-band; signal removal via null Id.
		data.Id = types.StringNull()
		return
	}

	sslcertbundleImportSetAttrFromGet(ctx, data, getResponseData)

}

func sslcertbundleImportGetThePayloadFromthePlan(ctx context.Context, data *SslcertbundleImportResourceModel) ssl.Sslcertbundle {
	tflog.Debug(ctx, "In sslcertbundleImportGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslcertbundle := ssl.Sslcertbundle{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslcertbundle.Name = data.Name.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		sslcertbundle.Src = data.Src.ValueString()
	}

	return sslcertbundle
}

func sslcertbundleImportSetAttrFromGet(ctx context.Context, data *SslcertbundleImportResourceModel, getResponseData map[string]interface{}) *SslcertbundleImportResourceModel {
	tflog.Debug(ctx, "In sslcertbundleImportSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// src is an Import-only input; NITRO GET does not echo it back. Preserve plan/state value (Pattern 7).

	// ID is set once in Create; do not recompute here (Pattern 6).

	return data
}
