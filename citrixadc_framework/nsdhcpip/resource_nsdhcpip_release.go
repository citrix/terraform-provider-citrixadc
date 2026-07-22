package nsdhcpip

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// nsdhcpip_release is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the release action:
//     POST /nitro/v1/config/nsdhcpip?action=release, which releases the DHCP
//     lease for the appliance management IP.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the release action, Read/Update are no-ops (there is
//     nothing to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for nsdhcpip_release.
//
// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsdhcpipReleaseResource{}
var _ resource.ResourceWithConfigure = (*NsdhcpipReleaseResource)(nil)

func NewNsdhcpipReleaseResource() resource.Resource {
	return &NsdhcpipReleaseResource{}
}

// NsdhcpipReleaseResource defines the resource implementation.
type NsdhcpipReleaseResource struct {
	client *service.NitroClient
}

// NsdhcpipReleaseResourceModel describes the resource data model.
//
// nsdhcpip_release is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// "nsdhcpip" object exposes no read/write properties and only the release
// action (POST /nitro/v1/config/nsdhcpip?action=release). The model therefore
// carries only the synthetic id.
type NsdhcpipReleaseResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *NsdhcpipReleaseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsdhcpip_release"
}

func (r *NsdhcpipReleaseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsdhcpipReleaseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsdhcpip_release resource.",
			},
		},
	}
}

func (r *NsdhcpipReleaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsdhcpipReleaseResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsdhcpip_release resource (release action)")
	nsdhcpip := nsdhcpip_releaseGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=release
	err := r.client.ActOnResource(service.Nsdhcpip.Type(), &nsdhcpip, "release")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to release nsdhcpip, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Released nsdhcpip")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("nsdhcpip_release")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. nsdhcpip_release has no GET endpoint; there is nothing to reconcile.
func (r *NsdhcpipReleaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsdhcpipReleaseResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nsdhcpip_release; NITRO exposes no GET endpoint (action=release only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. nsdhcpip_release has no attributes and no set endpoint.
func (r *NsdhcpipReleaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsdhcpipReleaseResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for nsdhcpip_release; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. nsdhcpip_release has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *NsdhcpipReleaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for nsdhcpip_release; NITRO has no delete endpoint")
}

// nsdhcpip_releaseGetThePayloadFromthePlan builds the (empty) NITRO payload for
// the release action. nsdhcpip_release has no read/write attributes, so the
// payload is an empty ns.Nsdhcpip struct.
func nsdhcpip_releaseGetThePayloadFromthePlan(ctx context.Context, data *NsdhcpipReleaseResourceModel) ns.Nsdhcpip {
	tflog.Debug(ctx, "In nsdhcpip_releaseGetThePayloadFromthePlan Function")
	nsdhcpip := ns.Nsdhcpip{}
	return nsdhcpip
}
