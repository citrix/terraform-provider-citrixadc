package sslkeyfile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslkeyfileImportResource{}
var _ resource.ResourceWithConfigure = (*SslkeyfileImportResource)(nil)
var _ resource.ResourceWithImportState = (*SslkeyfileImportResource)(nil)

func NewSslkeyfileImportResource() resource.Resource {
	return &SslkeyfileImportResource{}
}

// SslkeyfileImportResource defines the resource implementation.
type SslkeyfileImportResource struct {
	client *service.NitroClient
}

// SslkeyfileImportResourceModel describes the resource data model.
type SslkeyfileImportResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Src               types.String `tfsdk:"src"`
}

func (r *SslkeyfileImportResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslkeyfileImportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslkeyfile_import"
}

func (r *SslkeyfileImportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslkeyfileImportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslkeyfile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported key file. Must begin with an ASCII alphanumeric or underscore(_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a password_wo update.",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL specifying the protocol, host, and path, including file name, to the key file to be imported. For example, http://www.example.com/key_file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
		},
	}
}

func (r *SslkeyfileImportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslkeyfileImportResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslkeyfile resource")
	// Get payload from plan (regular attributes)
	sslkeyfile := sslkeyfileImportGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslkeyfileImportGetThePayloadFromtheConfig(ctx, &config, &sslkeyfile)

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
	r.readSslkeyfileImportFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslkeyfileImportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslkeyfileImportResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslkeyfile resource")

	r.readSslkeyfileImportFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslkeyfileImportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for sslkeyfile (only Import, delete,
	// get (all)). Every schema attribute is marked RequiresReplace, so Terraform
	// will never actually invoke Update with field changes. This body is a
	// documented no-op that preserves the prior ID and re-reads state.
	var data, state SslkeyfileImportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslkeyfile; NITRO has no update endpoint and all attributes are RequiresReplace")

	r.readSslkeyfileImportFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslkeyfileImportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslkeyfileImportResourceModel

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
func (r *SslkeyfileImportResource) readSslkeyfileImportFromApi(ctx context.Context, data *SslkeyfileImportResourceModel, diags *diag.Diagnostics) {

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
		// Object is gone out-of-band; signal removal via null Id.
		data.Id = types.StringNull()
		return
	}

	sslkeyfileImportSetAttrFromGet(ctx, data, getResponseData)

}

func sslkeyfileImportGetThePayloadFromthePlan(ctx context.Context, data *SslkeyfileImportResourceModel) ssl.Sslkeyfile {
	tflog.Debug(ctx, "In sslkeyfileImportGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslkeyfile := ssl.Sslkeyfile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslkeyfile.Name = data.Name.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		sslkeyfile.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		sslkeyfile.Src = data.Src.ValueString()
	}

	return sslkeyfile
}

func sslkeyfileImportGetThePayloadFromtheConfig(ctx context.Context, data *SslkeyfileImportResourceModel, payload *ssl.Sslkeyfile) {
	tflog.Debug(ctx, "In sslkeyfileImportGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func sslkeyfileImportSetAttrFromGet(ctx context.Context, data *SslkeyfileImportResourceModel, getResponseData map[string]interface{}) *SslkeyfileImportResourceModel {
	tflog.Debug(ctx, "In sslkeyfileImportSetAttrFromGet Function")

	// Resource setter: preserves plan/state values. `name` is the key and `src`
	// is a write-only Import input that NITRO does not faithfully echo back, so
	// do not overwrite either from the GET response. The password secret triple
	// is never returned by NITRO and is retained from config. ID is set in Create.
	if data.Name.IsNull() || data.Name.IsUnknown() {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}

	return data
}
