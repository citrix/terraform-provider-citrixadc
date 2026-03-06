package sslcertkey

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslCertKeyUpdateResource{}
var _ resource.ResourceWithConfigure = (*SslCertKeyUpdateResource)(nil)
var _ resource.ResourceWithImportState = (*SslCertKeyUpdateResource)(nil)

func NewSslCertKeyUpdateResource() resource.Resource {
	return &SslCertKeyUpdateResource{}
}

// SslCertKeyUpdateResource defines the resource implementation.
type SslCertKeyUpdateResource struct {
	client *service.NitroClient
}

// SslCertKeyUpdateResourceModel describes the resource data model.
type SslCertKeyUpdateResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Certkey         types.String `tfsdk:"certkey"`
	Cert            types.String `tfsdk:"cert"`
	Fipskey         types.String `tfsdk:"fipskey"`
	Inform          types.String `tfsdk:"inform"`
	Key             types.String `tfsdk:"key"`
	NoDomainCheck   types.Bool   `tfsdk:"nodomaincheck"`
	Passplain       types.String `tfsdk:"passplain"`
	Password        types.Bool   `tfsdk:"password"`
	LinkCertKeyName types.String `tfsdk:"linkcertkeyname"`
}

func (r *SslCertKeyUpdateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertkey_update"
}

func (r *SslCertKeyUpdateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     1,
		Description: "Resource to update an existing SSL certificate key pair. This resource performs an update operation on an existing certificate.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the SSL certificate key pair. This is the same as the certkey attribute.",
			},
			"certkey": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the certificate and private-key pair to update.",
			},
			"cert": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the X509 certificate file that is used to form the certificate-key pair.",
			},
			"fipskey": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the FIPS key that was created inside the Hardware Security Module (HSM) of a FIPS appliance, or a key that was imported into the HSM.",
			},
			"inform": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Input format of the certificate and the private-key files. The three formats supported by the appliance are: PEM - Privacy Enhanced Mail, DER - Distinguished Encoding Rule, PFX - Personal Information Exchange.",
			},
			"key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the private-key file that is used to form the certificate-key pair.",
			},
			"nodomaincheck": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Override the check for matching domain names during a certificate update operation.",
			},
			"passplain": schema.StringAttribute{
				Optional:  true,
				Computed:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Pass phrase used to encrypt the private-key. Required when updating an encrypted private-key in PEM format.",
			},
			"password": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Passphrase that was used to encrypt the private-key. Use this option to load encrypted private-keys in PEM format.",
			},
			"linkcertkeyname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.",
			},
		},
	}
}

func (r *SslCertKeyUpdateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslCertKeyUpdateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslCertKeyUpdateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslcertkey resource via update action")

	sslcertkeyName := data.Certkey.ValueString()

	// Build the update payload
	sslcertkey := ssl.Sslcertkey{
		Certkey:       sslcertkeyName,
		Nodomaincheck: true,
	}

	if !data.Cert.IsNull() {
		sslcertkey.Cert = data.Cert.ValueString()
	}
	if !data.Fipskey.IsNull() {
		sslcertkey.Fipskey = data.Fipskey.ValueString()
	}
	if !data.Inform.IsNull() {
		sslcertkey.Inform = data.Inform.ValueString()
	}
	if !data.Key.IsNull() {
		sslcertkey.Key = data.Key.ValueString()
	}
	if !data.Passplain.IsNull() {
		sslcertkey.Passplain = data.Passplain.ValueString()
	}
	if !data.Password.IsNull() {
		sslcertkey.Password = data.Password.ValueBool()
	}

	// Perform the update action
	err := r.client.ActOnResource(service.Sslcertkey.Type(), &sslcertkey, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcertkey, got error: %s", err))
		return
	}

	data.Id = types.StringValue(sslcertkeyName)

	tflog.Trace(ctx, "Updated sslcertkey resource")

	// Handle linked certificate if configured
	if !data.LinkCertKeyName.IsNull() && data.LinkCertKeyName.ValueString() != "" {
		if err := r.handleLinkedCertificateUpdate(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to handle linked certificate for %s, got error: %s", sslcertkeyName, err))
			return
		}
	} else if data.LinkCertKeyName.IsUnknown() {
		// Set to null if not configured
		data.LinkCertKeyName = types.StringNull()
	}

	// Ensure all optional+computed attributes have known values
	if data.Cert.IsUnknown() {
		data.Cert = types.StringNull()
	}
	if data.Fipskey.IsUnknown() {
		data.Fipskey = types.StringNull()
	}
	if data.Inform.IsUnknown() {
		data.Inform = types.StringNull()
	}
	if data.Key.IsUnknown() {
		data.Key = types.StringNull()
	}
	if data.NoDomainCheck.IsUnknown() {
		data.NoDomainCheck = types.BoolNull()
	}
	if data.Passplain.IsUnknown() {
		data.Passplain = types.StringNull()
	}
	if data.Password.IsUnknown() {
		data.Password = types.BoolNull()
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslCertKeyUpdateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// This is a no-op for update-only resources
	var data SslCertKeyUpdateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslCertKeyUpdateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// This is a no-op since all attributes have RequiresReplace modifier
	var data SslCertKeyUpdateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslCertKeyUpdateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// This is a no-op for update-only resources
	tflog.Debug(ctx, "Delete called for sslcertkey_update resource - no-op")
}

func (r *SslCertKeyUpdateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Helper function to handle linked certificate during update
func (r *SslCertKeyUpdateResource) handleLinkedCertificateUpdate(ctx context.Context, data *SslCertKeyUpdateResourceModel) error {
	tflog.Debug(ctx, "In handleLinkedCertificateUpdate")

	sslcertkeyName := data.Certkey.ValueString()

	// Get current state from API
	apiData, err := r.client.FindResource(service.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Clearing sslcertkey state %s", sslcertkeyName))
		data.Id = types.StringNull()
		return err
	}

	actualLinkedCertKeyname := apiData["linkcertkeyname"]
	configuredLinkedCertKeyname := data.LinkCertKeyName.ValueString()

	// Check for noop conditions
	if actualLinkedCertKeyname == configuredLinkedCertKeyname {
		tflog.Debug(ctx, fmt.Sprintf("actual and configured linked certificates identical: %v", actualLinkedCertKeyname))
		return nil
	}

	if actualLinkedCertKeyname == nil && configuredLinkedCertKeyname == "" {
		tflog.Debug(ctx, "actual and configured linked certificates both empty")
		return nil
	}

	// Unlink existing certificate if present
	if err := r.unlinkCertificateUpdate(ctx, data); err != nil {
		return err
	}

	// Link new certificate if configured
	if configuredLinkedCertKeyname != "" {
		tflog.Debug(ctx, fmt.Sprintf("Linking certkey: %s", configuredLinkedCertKeyname))
		sslCertkey := ssl.Sslcertkey{
			Certkey:         apiData["certkey"].(string),
			Linkcertkeyname: configuredLinkedCertKeyname,
		}
		if err := r.client.ActOnResource(service.Sslcertkey.Type(), &sslCertkey, "link"); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error linking certificate: %v", err))
			return err
		}
	} else {
		tflog.Debug(ctx, "configured linked certkey is empty, nothing to do")
	}

	return nil
}

// Helper function to unlink certificate during update
func (r *SslCertKeyUpdateResource) unlinkCertificateUpdate(ctx context.Context, data *SslCertKeyUpdateResourceModel) error {
	sslcertkeyName := data.Certkey.ValueString()

	apiData, err := r.client.FindResource(service.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Clearing sslcertkey state %s", sslcertkeyName))
		data.Id = types.StringNull()
		return err
	}

	actualLinkedCertKeyname := apiData["linkcertkeyname"]

	if actualLinkedCertKeyname != nil {
		tflog.Debug(ctx, fmt.Sprintf("Unlinking certkey: %s", actualLinkedCertKeyname))

		sslCertkey := ssl.Sslcertkey{
			Certkey: apiData["certkey"].(string),
		}
		if err := r.client.ActOnResource(service.Sslcertkey.Type(), &sslCertkey, "unlink"); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error unlinking certificate: %v", err))
			return err
		}
	} else {
		tflog.Debug(ctx, "actual linked certkey is nil, nothing to do")
	}

	return nil
}
