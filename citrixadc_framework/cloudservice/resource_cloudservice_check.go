package cloudservice

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// cloudservice_check is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the check action:
//     POST /nitro/v1/config/cloudservice?action=check, which checks the cloud
//     service configuration.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the check action, Read/Update are no-ops (there is nothing
//     to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for cloudservice_check.

// cloudserviceResourceType is the NITRO resource-type string. There is no
// service.Cloudservice enum in the vendored adc-nitro-go, so the raw type string
// is used with a map payload. NO vendor/ edits.
const cloudserviceResourceType = "cloudservice"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CloudserviceCheckResource{}
var _ resource.ResourceWithConfigure = (*CloudserviceCheckResource)(nil)
var _ resource.ResourceWithImportState = (*CloudserviceCheckResource)(nil)

func NewCloudserviceCheckResource() resource.Resource {
	return &CloudserviceCheckResource{}
}

// CloudserviceCheckResource defines the resource implementation.
type CloudserviceCheckResource struct {
	client *service.NitroClient
}

// CloudserviceCheckResourceModel describes the resource data model.
//
// cloudservice_check is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// "cloudservice" object exposes no read/write properties and only the check
// action (POST /nitro/v1/config/cloudservice?action=check). The model therefore
// carries only the synthetic id.
type CloudserviceCheckResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *CloudserviceCheckResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudserviceCheckResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudservice_check"
}

func (r *CloudserviceCheckResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudserviceCheckResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudservice_check resource.",
			},
		},
	}
}

func (r *CloudserviceCheckResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudserviceCheckResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudservice_check resource (check action)")
	cloudservice := cloudservice_checkGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=check
	err := r.client.ActOnResource(cloudserviceResourceType, cloudservice, "check")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to check cloudservice, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Checked cloudservice")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("cloudservice_check")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. cloudservice_check has no GET endpoint; there is nothing to reconcile.
func (r *CloudserviceCheckResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudserviceCheckResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for cloudservice_check; NITRO exposes no GET endpoint (action=check only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. cloudservice_check has no attributes and no set endpoint.
func (r *CloudserviceCheckResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CloudserviceCheckResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for cloudservice_check; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. cloudservice_check has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *CloudserviceCheckResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for cloudservice_check; NITRO has no delete endpoint")
}

// cloudservice_checkGetThePayloadFromthePlan builds the (empty) NITRO payload for
// the check action. cloudservice_check has no read/write attributes, so the
// payload is an empty map.
func cloudservice_checkGetThePayloadFromthePlan(ctx context.Context, data *CloudserviceCheckResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In cloudservice_checkGetThePayloadFromthePlan Function")
	cloudservice := make(map[string]interface{})
	return cloudservice
}
