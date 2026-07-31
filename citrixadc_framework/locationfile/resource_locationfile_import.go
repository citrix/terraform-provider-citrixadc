package locationfile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	// Aliased to avoid a name collision with the Plugin Framework `resource`
	// package. Used only for PrefixedUniqueId so the synthetic ID scheme is
	// byte-for-byte identical to the SDK v2 resource (backward compatibility).
	sdkresource "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LocationfileImportResource{}
var _ resource.ResourceWithConfigure = (*LocationfileImportResource)(nil)

func NewLocationfileImportResource() resource.Resource {
	return &LocationfileImportResource{}
}

// LocationfileImportResource defines the resource implementation.
//
// This resource models the NITRO locationfile `import` action. It is a
// one-shot side-effect action: the SDK v2 resource used Create only, with
// Read and Delete declared as schema.Noop and no Update. There is no inverse
// NITRO API and no stable GET-backed object keyed by the synthetic Terraform
// ID, so Read/Update/Delete are pure no-ops here. The import payload carries
// only Locationfile and src; `format` is present in the schema for backward
// compatibility but (mirroring SDK v2) is NOT sent in the import payload.
type LocationfileImportResource struct {
	client *service.NitroClient
}

// LocationfileImportResourceModel describes the resource data model.
type LocationfileImportResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Format       types.String `tfsdk:"format"`
	Locationfile types.String `tfsdk:"locationfile"`
	Src          types.String `tfsdk:"src"`
}

func (r *LocationfileImportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_locationfile_import"
}

func (r *LocationfileImportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LocationfileImportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the locationfile_import resource.",
			},
			// SDK v2: Optional + ForceNew -> RequiresReplace. Must NOT be
			// Computed because Read is a no-op (an action, not a managed
			// object), otherwise Terraform reports an unknown value after apply.
			"format": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Format of the location file. Required for the NetScaler to identify how to read the location file.",
			},
			// SDK v2: Required + ForceNew -> RequiresReplace.
			"locationfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.",
			},
			// SDK v2: Required + ForceNew -> RequiresReplace.
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

func (r *LocationfileImportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LocationfileImportResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Importing locationfile (action-only resource)")
	payload := locationfileImportGetThePayloadFromthePlan(ctx, &data)

	// import is a POST ?action=import action (there is no `add` used here in
	// SDK v2). Mirror the SDK v2 verb casing exactly: lower-case "import".
	err := r.client.ActOnResource(service.Locationfile.Type(), &payload, "import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to import locationfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Imported locationfile")

	// Synthetic, per-apply unique ID. Identical scheme to SDK v2
	// (resource.PrefixedUniqueId("tf-locationfile-")); the imported file is not
	// a queryable managed object, so this ID is purely a Terraform state handle.
	data.Id = types.StringValue(sdkresource.PrefixedUniqueId("tf-locationfile-"))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LocationfileImportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// SDK v2 declared Read as schema.Noop. import is a one-shot action with no
	// GET endpoint keyed by the synthetic ID, so Read preserves state as-is.
	var data LocationfileImportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for locationfile_import; import has no stable GET-backed object")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LocationfileImportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for the import action; every schema attribute
	// is RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state LocationfileImportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for locationfile_import; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LocationfileImportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// SDK v2 declared Delete as schema.Noop. import is a one-shot side-effect
	// action with no inverse NITRO API; Delete simply drops the resource from
	// Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for locationfile_import; NITRO has no inverse of the import action")
}

// locationfileImportGetThePayloadFromthePlan builds the body for the import
// action. Mirroring SDK v2, ONLY Locationfile and Src are included; `format` is
// intentionally excluded from the import payload (the NITRO import action
// accepts only Locationfile and src).
func locationfileImportGetThePayloadFromthePlan(ctx context.Context, data *LocationfileImportResourceModel) basic.Locationfile {
	tflog.Debug(ctx, "In locationfileImportGetThePayloadFromthePlan Function")

	locationfile := basic.Locationfile{}
	if !data.Locationfile.IsNull() && !data.Locationfile.IsUnknown() {
		locationfile.Locationfile = data.Locationfile.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		locationfile.Src = data.Src.ValueString()
	}

	return locationfile
}
