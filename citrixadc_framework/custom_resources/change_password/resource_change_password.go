package change_password

import (
	"context"
	"fmt"

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
var _ resource.Resource = &ChangePasswordResource{}
var _ resource.ResourceWithConfigure = (*ChangePasswordResource)(nil)
var _ resource.ResourceWithValidateConfig = (*ChangePasswordResource)(nil)

func NewChangePasswordResource() resource.Resource {
	return &ChangePasswordResource{}
}

// ChangePasswordResource models the NetScaler password-change operation.
//
// It is a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops and every input attribute is RequiresReplace. The
// first_time_password_reset flag selects the operation:
//   - true  -> POST {username, password, new_password} to the `login` endpoint
//     (the default/first-login password reset).
//   - false -> update `systemuser` to set the user's password to new_password
//     (equivalent to the CLI `set system user <username> <new_password>`); this
//     changes the password of an already-active user at any later time.
//
// Both secrets support the write-only "triple": password / password_wo /
// password_wo_version and new_password / new_password_wo / new_password_wo_version.
// The _wo variants are WriteOnly (read from config, never persisted to state) so
// ephemeral/Vault-sourced secrets never land in Terraform state. Bumping a
// _wo_version forces a replace, re-running the operation with the rotated secret.
type ChangePasswordResource struct {
	client *service.NitroClient
}

// ChangePasswordResourceModel describes the resource data model. Every schema
// attribute has a matching tfsdk field.
type ChangePasswordResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Username               types.String `tfsdk:"username"`
	Password               types.String `tfsdk:"password"`
	PasswordWo             types.String `tfsdk:"password_wo"`
	PasswordWoVersion      types.Int64  `tfsdk:"password_wo_version"`
	NewPassword            types.String `tfsdk:"new_password"`
	NewPasswordWo          types.String `tfsdk:"new_password_wo"`
	NewPasswordWoVersion   types.Int64  `tfsdk:"new_password_wo_version"`
	FirstTimePasswordReset types.Bool   `tfsdk:"first_time_password_reset"`
}

// changePasswordPayload is the request body for the login / systemuser endpoints.
// Neither login has a vendored adc-nitro-go struct, so a local struct is used
// (mirrors the legacy SDKv2 changePasswordPayload). Its json field names match the
// systemuser NITRO object, so the same struct serves both operations.
type changePasswordPayload struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	New_password string `json:"new_password,omitempty"`
}

func (r *ChangePasswordResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_change_password"
}

func (r *ChangePasswordResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the mandatory secrets (Pattern 17): password and
// new_password were both Required in the legacy resource. Each is expanded into a
// value/_wo/_wo_version triple whose value attributes are Optional, so at least
// one of the pair must be set.
func (r *ChangePasswordResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data ChangePasswordResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Password.IsNull() && data.PasswordWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Required Attribute",
			"Either \"password\" or \"password_wo\" must be specified.",
		)
	}

	if data.NewPassword.IsNull() && data.NewPasswordWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("new_password"),
			"Missing Required Attribute",
			"Either \"new_password\" or \"new_password_wo\" must be specified.",
		)
	}
}

func (r *ChangePasswordResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the change_password resource.",
			},
			"username": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "User name for the operation.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The default (current) password.",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The default (current) password. Write-only: read from configuration and never persisted to state.",
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
			"new_password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new password.",
			},
			"new_password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new password. Write-only: read from configuration and never persisted to state.",
			},
			"new_password_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a new_password_wo update.",
			},
			"first_time_password_reset": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Set to `true` to perform the default password reset at first login (via the login endpoint). Set to `false` to change the password of an existing user at any later time (via a systemuser update).",
			},
		},
	}
}

func (r *ChangePasswordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config ChangePasswordResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating change_password resource")
	// Get payload from plan (regular attributes)
	payload := changePasswordGetThePayloadFromthePlan(ctx, &data)
	// Overlay write-only secrets from config (they take precedence over the legacy attrs)
	changePasswordGetThePayloadFromtheConfig(ctx, &config, &payload)

	if data.FirstTimePasswordReset.ValueBool() {
		// First-time default password reset: authenticate with the current
		// password and set the new one via the login endpoint.
		_, err := r.client.AddResource("login", "", &payload)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to reset default password, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Reset default password via login action")
	} else {
		// Change the password of an existing user at any later time via a
		// systemuser update (equivalent to `set system user <username> <new_password>`).
		systemuserPayload := changePasswordPayload{
			Username: payload.Username,
			Password: payload.New_password,
		}
		err := r.client.UpdateUnnamedResource("systemuser", &systemuserPayload)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to change password, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Changed password via systemuser update")
	}

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue(fmt.Sprintf("change_password-%v", data.Username.ValueString()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChangePasswordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// The password change is a one-shot action. NITRO has no GET endpoint that
	// reports its state, and secrets are never read back, so Read is a pure
	// preserve-state no-op.
	var data ChangePasswordResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for change_password; NITRO has no GET endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChangePasswordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for this action; every schema attribute
	// (including the _wo_version trackers) is RequiresReplace, so Terraform never
	// invokes Update for a real change.
	var data, state ChangePasswordResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for change_password; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChangePasswordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// The password change is a one-shot side-effect action. There is no inverse
	// NITRO API. Delete only removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for change_password; NITRO has no inverse of the password-change action")
}

// changePasswordGetThePayloadFromthePlan builds the body from the plan (regular
// attributes). The write-only attributes and version trackers are skipped here and
// applied from config in changePasswordGetThePayloadFromtheConfig.
func changePasswordGetThePayloadFromthePlan(ctx context.Context, data *ChangePasswordResourceModel) changePasswordPayload {
	tflog.Debug(ctx, "In changePasswordGetThePayloadFromthePlan Function")

	payload := changePasswordPayload{}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		payload.Username = data.Username.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		payload.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.NewPassword.IsNull() && !data.NewPassword.IsUnknown() {
		payload.New_password = data.NewPassword.ValueString()
	}
	// Skip write-only attribute: new_password_wo
	// Skip version tracker attribute: new_password_wo_version

	return payload
}

// changePasswordGetThePayloadFromtheConfig overlays the write-only secrets read
// from configuration onto the payload. Applied after the plan helper so the _wo
// values win when both the legacy attribute and its _wo counterpart are set.
func changePasswordGetThePayloadFromtheConfig(ctx context.Context, data *ChangePasswordResourceModel, payload *changePasswordPayload) {
	tflog.Debug(ctx, "In changePasswordGetThePayloadFromtheConfig Function")

	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
	// Handle write-only secret attribute: new_password_wo -> new_password
	if !data.NewPasswordWo.IsNull() {
		newPasswordWo := data.NewPasswordWo.ValueString()
		if newPasswordWo != "" {
			payload.New_password = newPasswordWo
		}
	}
}
