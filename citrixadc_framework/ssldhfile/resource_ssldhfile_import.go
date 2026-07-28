package ssldhfile

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

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SsldhfileImportResource{}
var _ resource.ResourceWithConfigure = (*SsldhfileImportResource)(nil)
var _ resource.ResourceWithImportState = (*SsldhfileImportResource)(nil)

func NewSsldhfileImportResource() resource.Resource {
	return &SsldhfileImportResource{}
}

// SsldhfileImportResource defines the resource implementation.
type SsldhfileImportResource struct {
	client *service.NitroClient
}

// SsldhfileImportResourceModel describes the resource data model.
type SsldhfileImportResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Src  types.String `tfsdk:"src"`
}

func (r *SsldhfileImportResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SsldhfileImportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssldhfile_import"
}

func (r *SsldhfileImportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SsldhfileImportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ssldhfile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported DH file.  Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL specifying the protocol, host, and path, including file name, to the DH file to be imported. For example, http://www.example.com/dh_file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
		},
	}
}

func (r *SsldhfileImportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SsldhfileImportResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ssldhfile resource")
	ssldhfile := ssldhfileImportGetThePayloadFromthePlan(ctx, &data)

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
	r.readSsldhfileImportFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldhfileImportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SsldhfileImportResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ssldhfile resource")

	r.readSsldhfileImportFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// If the object was deleted out-of-band, remove it from state so a
	// subsequent apply re-creates it instead of erroring.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldhfileImportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for ssldhfile (only Import, delete,
	// get (all)). Every schema attribute is marked RequiresReplace, so Terraform
	// will never actually invoke Update with field changes. This body is a
	// documented no-op that preserves the prior ID and re-reads state.
	var data, state SsldhfileImportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for ssldhfile; NITRO has no update endpoint and all attributes are RequiresReplace")

	r.readSsldhfileImportFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldhfileImportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SsldhfileImportResourceModel

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
func (r *SsldhfileImportResource) readSsldhfileImportFromApi(ctx context.Context, data *SsldhfileImportResourceModel, diags *diag.Diagnostics) {

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
		// Object is gone out-of-band; signal removal via null Id.
		data.Id = types.StringNull()
		return
	}

	ssldhfileImportSetAttrFromGet(ctx, data, getResponseData)

}

func ssldhfileImportGetThePayloadFromthePlan(ctx context.Context, data *SsldhfileImportResourceModel) ssl.Ssldhfile {
	tflog.Debug(ctx, "In ssldhfileImportGetThePayloadFromthePlan Function")

	// Create API request body from the model
	ssldhfile := ssl.Ssldhfile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		ssldhfile.Name = data.Name.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		ssldhfile.Src = data.Src.ValueString()
	}

	return ssldhfile
}

func ssldhfileImportSetAttrFromGet(ctx context.Context, data *SsldhfileImportResourceModel, getResponseData map[string]interface{}) *SsldhfileImportResourceModel {
	tflog.Debug(ctx, "In ssldhfileImportSetAttrFromGet Function")

	// Resource setter: preserves plan/state values. `name` is the key and `src`
	// is a write-only Import input that NITRO does not faithfully echo back, so
	// do not overwrite either from the GET response. ID is set once in Create.
	if data.Name.IsNull() || data.Name.IsUnknown() {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}

	return data
}
