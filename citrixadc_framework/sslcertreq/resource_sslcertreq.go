package sslcertreq

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslcertreqResource{}
var _ resource.ResourceWithConfigure = (*SslcertreqResource)(nil)
var _ resource.ResourceWithImportState = (*SslcertreqResource)(nil)

func NewSslcertreqResource() resource.Resource {
	return &SslcertreqResource{}
}

// SslcertreqResource defines the resource implementation.
type SslcertreqResource struct {
	client *service.NitroClient
}

func (r *SslcertreqResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcertreqResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertreq"
}

func (r *SslcertreqResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcertreqResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslcertreqResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcertreq resource")
	// Get payload from plan (regular attributes)
	sslcertreq := sslcertreqGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslcertreqGetThePayloadFromtheConfig(ctx, &config, &sslcertreq)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.ActOnResource(service.Sslcertreq.Type(), &sslcertreq, "create")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcertreq, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcertreq resource")

	// Set ID for the resource
	data.Id = types.StringValue("sslcertreq-config")

	// Resolve unknown Optional+Computed attributes to null since Read is a noop
	// (no server-side state to read back for this action-only resource)
	if data.Commonname.IsUnknown() {
		data.Commonname = types.StringNull()
	}
	if data.Companyname.IsUnknown() {
		data.Companyname = types.StringNull()
	}
	if data.Countryname.IsUnknown() {
		data.Countryname = types.StringNull()
	}
	if data.Digestmethod.IsUnknown() {
		data.Digestmethod = types.StringNull()
	}
	if data.Emailaddress.IsUnknown() {
		data.Emailaddress = types.StringNull()
	}
	if data.Fipskeyname.IsUnknown() {
		data.Fipskeyname = types.StringNull()
	}
	if data.Keyfile.IsUnknown() {
		data.Keyfile = types.StringNull()
	}
	if data.Keyform.IsUnknown() {
		data.Keyform = types.StringNull()
	}
	if data.Localityname.IsUnknown() {
		data.Localityname = types.StringNull()
	}
	if data.Organizationname.IsUnknown() {
		data.Organizationname = types.StringNull()
	}
	if data.Organizationunitname.IsUnknown() {
		data.Organizationunitname = types.StringNull()
	}
	if data.Reqfile.IsUnknown() {
		data.Reqfile = types.StringNull()
	}
	if data.Statename.IsUnknown() {
		data.Statename = types.StringNull()
	}
	if data.Subjectaltname.IsUnknown() {
		data.Subjectaltname = types.StringNull()
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertreqResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcertreqResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcertreq resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertreqResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslcertreqResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslcertreq resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertreqResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcertreqResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcertreq resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed sslcertreq from Terraform state")
}
