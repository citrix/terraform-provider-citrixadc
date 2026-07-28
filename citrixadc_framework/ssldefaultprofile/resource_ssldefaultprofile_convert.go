package ssldefaultprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ssldefaultprofile_convert is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the convert action:
//     POST /nitro/v1/config/ssldefaultprofile?action=convert, which converts the
//     appliance to the SSL default profile mode.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the convert action, Read/Update are no-ops (there is
//     nothing to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for
//     ssldefaultprofile_convert.

// ssldefaultprofileResourceType is the NITRO resource-type string. There is no
// service.Ssldefaultprofile enum in the vendored adc-nitro-go, so the raw type
// string is used with a map payload. NO vendor/ edits.
const ssldefaultprofileResourceType = "ssldefaultprofile"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SsldefaultprofileConvertResource{}
var _ resource.ResourceWithConfigure = (*SsldefaultprofileConvertResource)(nil)

func NewSsldefaultprofileConvertResource() resource.Resource {
	return &SsldefaultprofileConvertResource{}
}

// SsldefaultprofileConvertResource defines the resource implementation.
type SsldefaultprofileConvertResource struct {
	client *service.NitroClient
}

// SsldefaultprofileConvertResourceModel describes the resource data model.
//
// ssldefaultprofile_convert is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// "ssldefaultprofile" object exposes no read/write properties and only the
// convert action (POST /nitro/v1/config/ssldefaultprofile?action=convert). The
// model therefore carries only the synthetic id.
type SsldefaultprofileConvertResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SsldefaultprofileConvertResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssldefaultprofile_convert"
}

func (r *SsldefaultprofileConvertResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SsldefaultprofileConvertResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ssldefaultprofile_convert resource.",
			},
		},
	}
}

func (r *SsldefaultprofileConvertResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SsldefaultprofileConvertResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ssldefaultprofile_convert resource (convert action)")
	ssldefaultprofile := ssldefaultprofile_convertGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=convert
	err := r.client.ActOnResource(ssldefaultprofileResourceType, ssldefaultprofile, "convert")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to convert ssldefaultprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Converted ssldefaultprofile")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("ssldefaultprofile_convert")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. ssldefaultprofile_convert has no GET endpoint; nothing to reconcile.
func (r *SsldefaultprofileConvertResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SsldefaultprofileConvertResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for ssldefaultprofile_convert; NITRO exposes no GET endpoint (action=convert only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. ssldefaultprofile_convert has no attributes and no set endpoint.
func (r *SsldefaultprofileConvertResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SsldefaultprofileConvertResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for ssldefaultprofile_convert; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. ssldefaultprofile_convert has no delete endpoint; the action is
// not reversible and there is no persistent object to remove.
func (r *SsldefaultprofileConvertResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for ssldefaultprofile_convert; NITRO has no delete endpoint")
}

// ssldefaultprofile_convertGetThePayloadFromthePlan builds the (empty) NITRO
// payload for the convert action. ssldefaultprofile_convert has no read/write
// attributes, so the payload is an empty map.
func ssldefaultprofile_convertGetThePayloadFromthePlan(ctx context.Context, data *SsldefaultprofileConvertResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In ssldefaultprofile_convertGetThePayloadFromthePlan Function")
	ssldefaultprofile := make(map[string]interface{})
	return ssldefaultprofile
}
