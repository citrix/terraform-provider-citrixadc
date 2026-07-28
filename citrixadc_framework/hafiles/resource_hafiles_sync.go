package hafiles

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &HafilesSyncResource{}
var _ resource.ResourceWithConfigure = (*HafilesSyncResource)(nil)

func NewHafilesSyncResource() resource.Resource {
	return &HafilesSyncResource{}
}

// HafilesSyncResource defines the resource implementation.
type HafilesSyncResource struct {
	client *service.NitroClient
}

// HafilesSyncResourceModel describes the resource data model.
type HafilesSyncResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Mode types.List   `tfsdk:"mode"`
}

func (r *HafilesSyncResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hafiles_sync"
}

func (r *HafilesSyncResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HafilesSyncResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the hafiles_sync resource.",
			},
			"mode": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "Specify one of the following modes of synchronization.\n* all - Synchronize files related to system configuration, Access Gateway bookmarks, SSL certificates, SSL CRL lists,  and Application Firewall XML objects.\n* bookmarks - Synchronize all Access Gateway bookmarks.\n* ssl - Synchronize all certificates, keys, and CRLs for the SSL feature.\n* imports. Synchronize all XML objects (for example, WSDLs, schemas, error pages) configured for the application firewall.\n* misc - Synchronize all license files and the rc.conf file.\n* all_plus_misc - Synchronize files related to system configuration, Access Gateway bookmarks, SSL certificates, SSL CRL lists, application firewall XML objects, licenses, and the rc.conf file.",
			},
		},
	}
}

func (r *HafilesSyncResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HafilesSyncResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hafiles_sync resource")

	// hafiles exposes only the POST ?action=sync action on NITRO.
	// There is no add/get/update/delete endpoint. Use ActOnResource with
	// the "sync" verb.
	payload := hafiles_syncGetThePayloadFromthePlan(ctx, &data)

	err := r.client.ActOnResource(service.Hafiles.Type(), payload, "sync")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to sync hafiles, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Synced hafiles_sync resource")

	// Synthetic constant ID - there is no NITRO identity for this action resource.
	data.Id = types.StringValue("hafiles_sync")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafilesSyncResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HafilesSyncResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for hafiles_sync: NITRO exposes no GET endpoint for this
	// action-only resource, so drift detection is impossible. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for hafiles_sync; no GET endpoint on NITRO side")

	// Save (unchanged) data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafilesSyncResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state HafilesSyncResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for hafiles_sync; the only attribute (mode) is
	// RequiresReplace, so Terraform re-creates (re-syncs) on change instead.
	tflog.Debug(ctx, "Update is a no-op for hafiles_sync; all attributes are RequiresReplace")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafilesSyncResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HafilesSyncResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op for hafiles_sync: NITRO exposes no DELETE endpoint for
	// this action-only resource. Just remove from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for hafiles_sync; no DELETE endpoint on NITRO side")
	tflog.Trace(ctx, "Removed hafiles_sync from Terraform state")
}

func hafiles_syncGetThePayloadFromthePlan(ctx context.Context, data *HafilesSyncResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In hafiles_syncGetThePayloadFromthePlan Function")

	// Build the sync action payload. mode is included only when set.
	hafiles := make(map[string]interface{})
	if !data.Mode.IsNull() && !data.Mode.IsUnknown() {
		var modeList []string
		data.Mode.ElementsAs(ctx, &modeList, false)
		hafiles["mode"] = modeList
	}

	return hafiles
}
