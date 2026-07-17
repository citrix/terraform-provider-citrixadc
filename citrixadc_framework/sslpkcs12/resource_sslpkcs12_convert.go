package sslpkcs12

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Sslpkcs12ConvertResource{}
var _ resource.ResourceWithConfigure = (*Sslpkcs12ConvertResource)(nil)
var _ resource.ResourceWithImportState = (*Sslpkcs12ConvertResource)(nil)
var _ resource.ResourceWithValidateConfig = (*Sslpkcs12ConvertResource)(nil)

func NewSslpkcs12ConvertResource() resource.Resource {
	return &Sslpkcs12ConvertResource{}
}

// Sslpkcs12ConvertResource defines the resource implementation.
type Sslpkcs12ConvertResource struct {
	client *service.NitroClient
}

// Sslpkcs12ConvertResourceModel describes the resource data model.
//
// This resource models the NITRO sslpkcs12 `?action=convert` action. convert is
// a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops.
type Sslpkcs12ConvertResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Import                 types.Bool   `tfsdk:"import"`
	Aes256                 types.Bool   `tfsdk:"aes256"`
	Certfile               types.String `tfsdk:"certfile"`
	Des                    types.Bool   `tfsdk:"des"`
	Des3                   types.Bool   `tfsdk:"des3"`
	Export                 types.Bool   `tfsdk:"export"`
	Keyfile                types.String `tfsdk:"keyfile"`
	Outfile                types.String `tfsdk:"outfile"`
	Password               types.String `tfsdk:"password"`
	PasswordWo             types.String `tfsdk:"password_wo"`
	PasswordWoVersion      types.Int64  `tfsdk:"password_wo_version"`
	Pempassphrase          types.String `tfsdk:"pempassphrase"`
	PempassphraseWo        types.String `tfsdk:"pempassphrase_wo"`
	PempassphraseWoVersion types.Int64  `tfsdk:"pempassphrase_wo_version"`
	Pkcs12file             types.String `tfsdk:"pkcs12file"`
}

// ValidateConfig enforces that the mandatory secret attribute password is
// supplied via one of its value/write-only variants (Pattern 17). password
// protects the PKCS#12 material and is required for the convert action.
//
// pempassphrase is intentionally NOT required: it is only needed to unlock an
// ENCRYPTED input PEM key when exporting. Exporting an unencrypted key needs no
// pass phrase, so requiring it unconditionally would reject a valid config. When
// it is actually needed but missing/wrong, the NITRO convert call returns a
// clear error.
func (r *Sslpkcs12ConvertResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Sslpkcs12ConvertResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Password.IsNull() && data.PasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing required attribute",
			"One of \"password\" or \"password_wo\" must be set for sslpkcs12_convert.",
		)
	}
}

func (r *Sslpkcs12ConvertResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Sslpkcs12ConvertResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslpkcs12_convert"
}

func (r *Sslpkcs12ConvertResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Sslpkcs12ConvertResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslpkcs12_convert resource.",
			},
			"import": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Convert the certificate and private-key from PKCS#12 format to PEM format.",
			},
			"aes256": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the private key by using the AES algorithm (256-bit key) during the import operation. On the command line, you are prompted to enter the pass phrase.",
			},
			"certfile": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Certificate file to be converted from PEM to PKCS#12 format.",
			},
			"des": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the private key by using the DES algorithm in CBC mode during the import operation. On the command line, you are prompted to enter the pass phrase.",
			},
			"des3": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the private key by using the Triple-DES algorithm in EDE CBC mode (168-bit key) during the import operation. On the command line, you are prompted to enter the pass phrase.",
			},
			"export": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Convert the certificate and private key from PEM format to PKCS#12 format. On the command line, you are prompted to enter the pass phrase.",
			},
			"keyfile": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the private key file to be converted from PEM to PKCS#12 format. If the key file is encrypted, you are prompted to enter the pass phrase used for encrypting the key.",
			},
			"outfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to, the output file that contains the certificate and the private key after converting from PKCS#12 to PEM format. /nsconfig/ssl/ is the default path.\nIf importing, the certificate-key pair is stored in PEM format. If exporting, the certificate-key pair is stored in PKCS#12 format.",
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
			"pempassphrase": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"pempassphrase_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"pempassphrase_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a pempassphrase_wo update.",
			},
			"pkcs12file": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to, the PKCS#12 file. If importing, specify the input file name that contains the certificate and the private key in PKCS#12 format. If exporting, specify the output file name that contains the certificate and the private key after converting from PEM to\nPKCS#12 format. /nsconfig/ssl/ is the default path.\nDuring the import operation, if the key is encrypted, you are prompted to enter the pass phrase used for encrypting the key.",
			},
		},
	}
}

func (r *Sslpkcs12ConvertResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config Sslpkcs12ConvertResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Converting sslpkcs12 (action-only resource)")
	// Get payload from plan (regular attributes)
	sslpkcs12 := sslpkcs12_convertGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslpkcs12_convertGetThePayloadFromtheConfig(ctx, &config, &sslpkcs12)

	// Make API call
	// Action-only resource: sslpkcs12 exposes ONLY ?action=convert (no add/get/delete).
	// NOTE: this operation is DISRUPTIVE and NON-IDEMPOTENT - it requires the source
	// certificate/key files to be present on the appliance filesystem and writes the
	// converted output file. There is no GET endpoint to read back the result.
	err := r.client.ActOnResource(service.Sslpkcs12.Type(), &sslpkcs12, "convert")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to convert sslpkcs12, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Converted sslpkcs12 resource")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("sslpkcs12_convert")

	// No Read-back: sslpkcs12 has no GET endpoint.

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs12ConvertResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// convert is a one-shot action. NITRO has no GET endpoint that reports
	// convert-state, so Read is a pure preserve-state no-op.
	var data Sslpkcs12ConvertResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for sslpkcs12_convert; NITRO has no query endpoint for convert state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs12ConvertResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for convert; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state Sslpkcs12ConvertResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslpkcs12_convert; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs12ConvertResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// convert is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-convert"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for sslpkcs12_convert; NITRO has no inverse of the convert action")
}

func sslpkcs12_convertGetThePayloadFromthePlan(ctx context.Context, data *Sslpkcs12ConvertResourceModel) ssl.Sslpkcs12 {
	tflog.Debug(ctx, "In sslpkcs12_convertGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslpkcs12 := ssl.Sslpkcs12{}
	if !data.Import.IsNull() && !data.Import.IsUnknown() {
		sslpkcs12.Import = data.Import.ValueBool()
	}
	if !data.Aes256.IsNull() && !data.Aes256.IsUnknown() {
		sslpkcs12.Aes256 = data.Aes256.ValueBool()
	}
	if !data.Certfile.IsNull() && !data.Certfile.IsUnknown() {
		sslpkcs12.Certfile = data.Certfile.ValueString()
	}
	if !data.Des.IsNull() && !data.Des.IsUnknown() {
		sslpkcs12.Des = data.Des.ValueBool()
	}
	if !data.Des3.IsNull() && !data.Des3.IsUnknown() {
		sslpkcs12.Des3 = data.Des3.ValueBool()
	}
	if !data.Export.IsNull() && !data.Export.IsUnknown() {
		sslpkcs12.Export = data.Export.ValueBool()
	}
	if !data.Keyfile.IsNull() && !data.Keyfile.IsUnknown() {
		sslpkcs12.Keyfile = data.Keyfile.ValueString()
	}
	if !data.Outfile.IsNull() && !data.Outfile.IsUnknown() {
		sslpkcs12.Outfile = data.Outfile.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		sslpkcs12.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Pempassphrase.IsNull() && !data.Pempassphrase.IsUnknown() {
		sslpkcs12.Pempassphrase = data.Pempassphrase.ValueString()
	}
	// Skip write-only attribute: pempassphrase_wo
	// Skip version tracker attribute: pempassphrase_wo_version
	if !data.Pkcs12file.IsNull() && !data.Pkcs12file.IsUnknown() {
		sslpkcs12.Pkcs12file = data.Pkcs12file.ValueString()
	}

	return sslpkcs12
}

func sslpkcs12_convertGetThePayloadFromtheConfig(ctx context.Context, data *Sslpkcs12ConvertResourceModel, payload *ssl.Sslpkcs12) {
	tflog.Debug(ctx, "In sslpkcs12_convertGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
	// Handle write-only secret attribute: pempassphrase_wo -> pempassphrase
	if !data.PempassphraseWo.IsNull() {
		pempassphraseWo := data.PempassphraseWo.ValueString()
		if pempassphraseWo != "" {
			payload.Pempassphrase = pempassphraseWo
		}
	}
}
