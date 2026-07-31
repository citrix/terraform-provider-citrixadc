package locationfile6

import (
	"context"
	"fmt"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Locationfile6ImportResource{}
var _ resource.ResourceWithConfigure = (*Locationfile6ImportResource)(nil)

func NewLocationfile6ImportResource() resource.Resource {
	return &Locationfile6ImportResource{}
}

// Locationfile6ImportResource defines the resource implementation.
//
// This resource models the NITRO locationfile6 `?action=import` action. import
// is a one-shot side-effect action: it imports an IPv6 location file from a
// source URL. The SDK v2 resource (resource_citrixadc_locationfile6_import.go)
// declared Read and Delete as schema.Noop and had no Update, so this is a PURE
// ACTION resource with no-op Read/Update/Delete. To preserve backward
// compatibility the id is a freshly generated prefixed unique string (SDK v2
// used resource.PrefixedUniqueId("tf-locationfile6-")).
type Locationfile6ImportResource struct {
	client *service.NitroClient
}

// Locationfile6ImportResourceModel describes the resource data model.
type Locationfile6ImportResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Format       types.String `tfsdk:"format"`
	Locationfile types.String `tfsdk:"locationfile"`
	Src          types.String `tfsdk:"src"`
}

func (r *Locationfile6ImportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_locationfile6_import"
}

func (r *Locationfile6ImportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Locationfile6ImportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the locationfile6_import resource.",
			},
			// SDK v2: format is Optional + ForceNew (no Default, not Computed).
			// Mirror exactly. import is a one-shot action with a no-op Read, so
			// this must NOT be Computed.
			"format": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Format of the IPv6 location file. Required for the NetScaler to identify how to read the location file.",
			},
			// SDK v2: locationfile is Required + ForceNew.
			"locationfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the IPv6 location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.",
			},
			// SDK v2: src is Required + ForceNew.
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path, and file name) from where the location file will be imported. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func (r *Locationfile6ImportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Locationfile6ImportResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Importing locationfile6 (action-only resource)")
	locationfile := locationfile6ImportGetThePayloadFromthePlan(ctx, &data)

	// import is a POST ?action=import action (no add endpoint for this
	// workflow). The verb casing is lower-case "import" to match SDK v2.
	err := r.client.ActOnResource(service.Locationfile6.Type(), &locationfile, "import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to import locationfile6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Imported locationfile6")

	// SDK v2 used resource.PrefixedUniqueId("tf-locationfile6-"), a freshly
	// generated unique prefixed handle. Reproduce a prefixed unique id here so
	// the id scheme stays backward compatible; the imported location file is
	// not a queryable managed object, so the id is purely a state handle.
	data.Id = types.StringValue(fmt.Sprintf("tf-locationfile6-%d", time.Now().UnixNano()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Locationfile6ImportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// SDK v2 Read was schema.Noop. import is a one-shot action; NITRO has no
	// GET endpoint that reports import-state, so Read is a preserve-state no-op.
	var data Locationfile6ImportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for locationfile6_import; import has no stable GET-backed object")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Locationfile6ImportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for the import action; every schema attribute
	// is RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state Locationfile6ImportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for locationfile6_import; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Locationfile6ImportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// SDK v2 Delete was schema.Noop. import is a one-shot side-effect action
	// with no inverse NITRO API, so Delete simply removes the resource from
	// Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for locationfile6_import; NITRO has no inverse of the import action")
}

// locationfile6ImportGetThePayloadFromthePlan builds the body for the import
// action. Mirroring SDK v2, only Locationfile and Src are sent (these are the
// fields the import action accepts per the NITRO metadata); format is a schema
// attribute but is not part of the import payload.
func locationfile6ImportGetThePayloadFromthePlan(ctx context.Context, data *Locationfile6ImportResourceModel) basic.Locationfile6 {
	tflog.Debug(ctx, "In locationfile6ImportGetThePayloadFromthePlan Function")

	locationfile := basic.Locationfile6{}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		locationfile.Src = data.Src.ValueString()
	}
	if !data.Locationfile.IsNull() && !data.Locationfile.IsUnknown() {
		locationfile.Locationfile = data.Locationfile.ValueString()
	}

	return locationfile
}
