package sslpkcs8

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Sslpkcs8ConvertResource{}
var _ resource.ResourceWithConfigure = (*Sslpkcs8ConvertResource)(nil)

func NewSslpkcs8ConvertResource() resource.Resource {
	return &Sslpkcs8ConvertResource{}
}

// Sslpkcs8ConvertResource defines the resource implementation.
type Sslpkcs8ConvertResource struct {
	client *service.NitroClient
}

// Sslpkcs8ConvertResourceModel describes the resource data model.
//
// This resource models the NITRO sslpkcs8 `?action=convert` action. convert is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The convert payload carries the key file,
// key format, optional password (write-only) and the output PKCS#8 file.
type Sslpkcs8ConvertResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Keyfile           types.String `tfsdk:"keyfile"`
	Keyform           types.String `tfsdk:"keyform"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Pkcs8file         types.String `tfsdk:"pkcs8file"`
}

func (r *Sslpkcs8ConvertResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslpkcs8_convert"
}

func (r *Sslpkcs8ConvertResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Sslpkcs8ConvertResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslpkcs8_convert resource.",
			},
			"keyfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the input key file to be converted from PEM or DER format to PKCS#8 format. /nsconfig/ssl/ is the default path.",
			},
			"keyform": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("PEM"),
				Description: "Format in which the key file is stored on the appliance.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password to assign to the file if the key is encrypted. Applies only for PEM format files.",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password to assign to the file if the key is encrypted. Applies only for PEM format files.",
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
			"pkcs8file": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to, the output file where the PKCS#8 format key file is stored. /nsconfig/ssl/ is the default path.",
			},
		},
	}
}

func (r *Sslpkcs8ConvertResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config Sslpkcs8ConvertResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslpkcs8_convert resource")
	// Get payload from plan (regular attributes)
	sslpkcs8 := sslpkcs8_convertGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslpkcs8_convertGetThePayloadFromtheConfig(ctx, &config, &sslpkcs8)

	// Make API call
	// Action-only resource: sslpkcs8 exposes ONLY ?action=convert (no add/get/delete).
	// NOTE: this operation requires the source key file to be present on the appliance
	// filesystem and is NON-IDEMPOTENT - it writes the converted PKCS#8 output file.
	// There is no GET endpoint to read back the result.
	err := r.client.ActOnResource(service.Sslpkcs8.Type(), &sslpkcs8, "convert")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to convert sslpkcs8, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Converted sslpkcs8_convert resource")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("sslpkcs8_convert")

	// No Read-back: sslpkcs8 has no GET endpoint.

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs8ConvertResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// convert is a one-shot action. NITRO has no GET endpoint that reports
	// convert-state, so Read is a pure preserve-state no-op.
	var data Sslpkcs8ConvertResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for sslpkcs8_convert; NITRO has no GET endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs8ConvertResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for convert; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state Sslpkcs8ConvertResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for sslpkcs8_convert; NITRO has no update endpoint and all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs8ConvertResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// convert is a one-shot side-effect action. There is no inverse NITRO API.
	// The converted output file persists on the appliance; Delete only removes
	// Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for sslpkcs8_convert; NITRO has no inverse of the convert action")
}

func sslpkcs8_convertGetThePayloadFromthePlan(ctx context.Context, data *Sslpkcs8ConvertResourceModel) ssl.Sslpkcs8 {
	tflog.Debug(ctx, "In sslpkcs8_convertGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslpkcs8 := ssl.Sslpkcs8{}
	if !data.Keyfile.IsNull() && !data.Keyfile.IsUnknown() {
		sslpkcs8.Keyfile = data.Keyfile.ValueString()
	}
	if !data.Keyform.IsNull() && !data.Keyform.IsUnknown() {
		sslpkcs8.Keyform = data.Keyform.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		sslpkcs8.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Pkcs8file.IsNull() && !data.Pkcs8file.IsUnknown() {
		sslpkcs8.Pkcs8file = data.Pkcs8file.ValueString()
	}

	return sslpkcs8
}

func sslpkcs8_convertGetThePayloadFromtheConfig(ctx context.Context, data *Sslpkcs8ConvertResourceModel, payload *ssl.Sslpkcs8) {
	tflog.Debug(ctx, "In sslpkcs8_convertGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}
