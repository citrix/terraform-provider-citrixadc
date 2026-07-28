package sslwrapkey

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslwrapkeyCreateResource{}
var _ resource.ResourceWithConfigure = (*SslwrapkeyCreateResource)(nil)
var _ resource.ResourceWithImportState = (*SslwrapkeyCreateResource)(nil)
var _ resource.ResourceWithValidateConfig = (*SslwrapkeyCreateResource)(nil)

func NewSslwrapkeyCreateResource() resource.Resource {
	return &SslwrapkeyCreateResource{}
}

// SslwrapkeyCreateResource defines the resource implementation.
type SslwrapkeyCreateResource struct {
	client *service.NitroClient
}

// SslwrapkeyCreateResourceModel describes the resource data model.
type SslwrapkeyCreateResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Salt              types.String `tfsdk:"salt"`
	SaltWo            types.String `tfsdk:"salt_wo"`
	SaltWoVersion     types.Int64  `tfsdk:"salt_wo_version"`
	Wrapkeyname       types.String `tfsdk:"wrapkeyname"`
}

func (r *SslwrapkeyCreateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslwrapkeyCreateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslwrapkey_create"
}

func (r *SslwrapkeyCreateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslwrapkeyCreateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslwrapkey resource.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password string for the wrap key.",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password string for the wrap key.",
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
			"salt": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Salt string for the wrap key.",
			},
			"salt_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Salt string for the wrap key.",
			},
			"salt_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a salt_wo update.",
			},
			"wrapkeyname": schema.StringAttribute{
				// Key attribute - mandatory (Pattern 8: tfdata wrongly marks optional).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the wrap key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the wrap key is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my key\" or 'my key').",
			},
		},
	}
}

func (r *SslwrapkeyCreateResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data SslwrapkeyCreateResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// password and salt are mandatory for the create action (Pattern 8 / Pattern 17).
	// Each is expanded into a secret triple whose value attributes are both Optional,
	// so enforce at-least-one-of at plan time.
	if data.Password.IsNull() && data.PasswordWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Required Attribute",
			"Either \"password\" or \"password_wo\" must be specified.",
		)
	}
	if data.Salt.IsNull() && data.SaltWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("salt"),
			"Missing Required Attribute",
			"Either \"salt\" or \"salt_wo\" must be specified.",
		)
	}
}

func (r *SslwrapkeyCreateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslwrapkeyCreateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslwrapkey resource")
	// Get payload from plan (regular attributes)
	sslwrapkey := sslwrapkeyGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslwrapkeyGetThePayloadFromtheConfig(ctx, &config, &sslwrapkey)

	// Make API call
	// sslwrapkey is created via the NITRO `create` action (?action=create, POST),
	// NOT a plain add. There is no update endpoint.
	// NOTE: may require the FIPS/crypto subsystem to be available - gate the
	// acceptance test accordingly.
	err := r.client.ActOnResource(service.Sslwrapkey.Type(), &sslwrapkey, "create")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslwrapkey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslwrapkey resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(data.Wrapkeyname.ValueString())

	// Read the updated state back
	r.readSslwrapkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslwrapkeyCreateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslwrapkeyCreateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslwrapkey resource")

	r.readSslwrapkeyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Resource is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslwrapkeyCreateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslwrapkeyCreateResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for sslwrapkey; the NITRO doc exposes no update endpoint and
	// all attributes are RequiresReplace (Pattern 5).
	tflog.Debug(ctx, "Update is a no-op for sslwrapkey; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSslwrapkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslwrapkeyCreateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslwrapkeyCreateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslwrapkey resource")
	// NITRO supports DELETE /config/sslwrapkey/<wrapkeyname> (Pattern 4).
	err := r.client.DeleteResource(service.Sslwrapkey.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslwrapkey, got error: %s", err))
		return
	}
	tflog.Trace(ctx, "Deleted sslwrapkey resource")
}

// Helper function to read sslwrapkey data from API
func (r *SslwrapkeyCreateResource) readSslwrapkeyFromApi(ctx context.Context, data *SslwrapkeyCreateResourceModel, diags *diag.Diagnostics) {

	// Named resource: read by wrapkeyname (the ID holds the plain key value).
	getResponseData, err := r.client.FindResource(service.Sslwrapkey.Type(), data.Id.ValueString())
	if err != nil {
		if utils.IsNotFoundError(err) {
			// Resource no longer exists on the ADC. Signal removal via a null Id so the
			// Read caller drops it from state instead of erroring.
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslwrapkey, got error: %s", err))
		return
	}

	sslwrapkeySetAttrFromGet(ctx, data, getResponseData)
}

func sslwrapkeyGetThePayloadFromthePlan(ctx context.Context, data *SslwrapkeyCreateResourceModel) ssl.Sslwrapkey {
	tflog.Debug(ctx, "In sslwrapkeyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslwrapkey := ssl.Sslwrapkey{}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		sslwrapkey.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Salt.IsNull() && !data.Salt.IsUnknown() {
		sslwrapkey.Salt = data.Salt.ValueString()
	}
	// Skip write-only attribute: salt_wo
	// Skip version tracker attribute: salt_wo_version
	if !data.Wrapkeyname.IsNull() && !data.Wrapkeyname.IsUnknown() {
		sslwrapkey.Wrapkeyname = data.Wrapkeyname.ValueString()
	}

	return sslwrapkey
}

func sslwrapkeyGetThePayloadFromtheConfig(ctx context.Context, data *SslwrapkeyCreateResourceModel, payload *ssl.Sslwrapkey) {
	tflog.Debug(ctx, "In sslwrapkeyGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
	// Handle write-only secret attribute: salt_wo -> salt
	if !data.SaltWo.IsNull() {
		saltWo := data.SaltWo.ValueString()
		if saltWo != "" {
			payload.Salt = saltWo
		}
	}
}

func sslwrapkeySetAttrFromGet(ctx context.Context, data *SslwrapkeyCreateResourceModel, getResponseData map[string]interface{}) *SslwrapkeyCreateResourceModel {
	tflog.Debug(ctx, "In sslwrapkeySetAttrFromGet Function")

	// Convert API response to model.
	// password and salt are write-only secrets - the GET response returns only
	// wrapkeyname, so the secrets are preserved from plan/state and never touched.
	if val, ok := getResponseData["wrapkeyname"]; ok && val != nil {
		data.Wrapkeyname = types.StringValue(val.(string))
	}
	// ID is set once in Create; do not recompute it here (Pattern 6).

	return data
}
