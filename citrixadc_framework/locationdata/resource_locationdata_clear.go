package locationdata

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// locationdata_clear is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the clear action:
//     POST /nitro/v1/config/locationdata?action=clear, which clears the static
//     location (GSLB geo) database from memory.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the clear action, Read/Update are no-ops (there is nothing
//     to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for locationdata_clear.
var _ resource.Resource = &LocationdataClearResource{}
var _ resource.ResourceWithConfigure = (*LocationdataClearResource)(nil)

func NewLocationdataClearResource() resource.Resource {
	return &LocationdataClearResource{}
}

// LocationdataClearResource defines the resource implementation.
type LocationdataClearResource struct {
	client *service.NitroClient
}

// LocationdataClearResourceModel describes the resource data model.
//
// locationdata_clear is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// "locationdata" object exposes no read/write properties and only the clear
// action (POST /nitro/v1/config/locationdata?action=clear). The model therefore
// carries only the synthetic id.
type LocationdataClearResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *LocationdataClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_locationdata_clear"
}

func (r *LocationdataClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LocationdataClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the locationdata_clear resource.",
			},
		},
	}
}

func (r *LocationdataClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LocationdataClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating locationdata_clear resource (clear action)")
	locationdata := locationdata_clearGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=clear
	err := r.client.ActOnResource(service.Locationdata.Type(), &locationdata, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear locationdata, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared locationdata")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("locationdata_clear")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. locationdata_clear has no GET endpoint; there is nothing to reconcile.
func (r *LocationdataClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LocationdataClearResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for locationdata_clear; NITRO exposes no GET endpoint (action=clear only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. locationdata_clear has no attributes and no set endpoint.
func (r *LocationdataClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LocationdataClearResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for locationdata_clear; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. locationdata_clear has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *LocationdataClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for locationdata_clear; NITRO has no delete endpoint")
}

// locationdata_clearGetThePayloadFromthePlan builds the (empty) NITRO payload for
// the clear action. locationdata_clear has no read/write attributes, so the
// payload is an empty basic.Locationdata struct.
func locationdata_clearGetThePayloadFromthePlan(ctx context.Context, data *LocationdataClearResourceModel) basic.Locationdata {
	tflog.Debug(ctx, "In locationdata_clearGetThePayloadFromthePlan Function")
	locationdata := basic.Locationdata{}
	return locationdata
}
