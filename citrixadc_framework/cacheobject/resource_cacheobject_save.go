package cacheobject

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cache"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CacheobjectSaveResource{}
var _ resource.ResourceWithConfigure = (*CacheobjectSaveResource)(nil)

func NewCacheobjectSaveResource() resource.Resource {
	return &CacheobjectSaveResource{}
}

// CacheobjectSaveResource models the NITRO cacheobject `?action=save` action.
// save is a one-shot side-effect action with no GET endpoint and no inverse
// API, so Read/Update/Delete are no-ops. Its payload accepts [locator]
// [tosecondary]; per the CLI both args are optional (save has no mandatory
// choice, unlike expire/flush).
type CacheobjectSaveResource struct {
	client *service.NitroClient
}

// CacheobjectSaveResourceModel describes the resource data model.
type CacheobjectSaveResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Locator     types.Int64  `tfsdk:"locator"`
	Tosecondary types.String `tfsdk:"tosecondary"`
}

func (r *CacheobjectSaveResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cacheobject_save"
}

func (r *CacheobjectSaveResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CacheobjectSaveResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cacheobject_save resource.",
			},
			"locator": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the cached object.",
			},
			"tosecondary": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Object will be saved onto Secondary. Applies only to the save action.",
			},
		},
	}
}

func (r *CacheobjectSaveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CacheobjectSaveResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Firing cacheobject save action")
	payload := cacheobject_saveGetThePayloadFromthePlan(ctx, &data)

	// cacheobject has no add/set/delete; the only writes are POST ?action=<verb>.
	// Verb casing per the NITRO URL: action=save.
	err := r.client.ActOnResource(service.Cacheobject.Type(), &payload, "save")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to save cacheobject, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Performed cacheobject save action")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("cacheobject_save")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheobjectSaveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// save is a one-shot action. NITRO exposes no GET endpoint that reports
	// save-state, so Read is a pure preserve-state no-op.
	var data CacheobjectSaveResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for cacheobject_save; NITRO has no query endpoint for save state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheobjectSaveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for save; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state CacheobjectSaveResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for cacheobject_save; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheobjectSaveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// save is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for cacheobject_save; NITRO has no inverse of the save action")
}

// cacheobject_saveGetThePayloadFromthePlan builds the save payload. Per the
// NITRO doc the save action accepts [locator] [tosecondary].
func cacheobject_saveGetThePayloadFromthePlan(ctx context.Context, data *CacheobjectSaveResourceModel) cache.Cacheobject {
	tflog.Debug(ctx, "In cacheobject_saveGetThePayloadFromthePlan Function")

	cacheobject := cache.Cacheobject{}
	if !data.Locator.IsNull() && !data.Locator.IsUnknown() {
		cacheobject.Locator = utils.IntPtr(int(data.Locator.ValueInt64()))
	}
	if !data.Tosecondary.IsNull() && !data.Tosecondary.IsUnknown() {
		cacheobject.Tosecondary = data.Tosecondary.ValueString()
	}

	return cacheobject
}
