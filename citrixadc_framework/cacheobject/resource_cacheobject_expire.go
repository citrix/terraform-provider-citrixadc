package cacheobject

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cache"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
var _ resource.Resource = &CacheobjectExpireResource{}
var _ resource.ResourceWithConfigure = (*CacheobjectExpireResource)(nil)
var _ resource.ResourceWithImportState = (*CacheobjectExpireResource)(nil)

// ValidateConfig enforces the CLI-mandatory choice for the expire action:
// exactly one of (locator) OR (url + host). NITRO/tfdata mark none required, so
// the check must live in the schema (Pattern 8).
var _ resource.ResourceWithValidateConfig = (*CacheobjectExpireResource)(nil)

func NewCacheobjectExpireResource() resource.Resource {
	return &CacheobjectExpireResource{}
}

// CacheobjectExpireResource models the NITRO cacheobject `?action=expire`
// action. expire is a one-shot side-effect action with no GET endpoint and no
// inverse API, so Read/Update/Delete are no-ops. Its payload accepts
// locator | (url,host[,port,groupname,httpmethod]).
type CacheobjectExpireResource struct {
	client *service.NitroClient
}

// CacheobjectExpireResourceModel describes the resource data model.
type CacheobjectExpireResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Groupname  types.String `tfsdk:"groupname"`
	Host       types.String `tfsdk:"host"`
	Httpmethod types.String `tfsdk:"httpmethod"`
	Locator    types.Int64  `tfsdk:"locator"`
	Port       types.Int64  `tfsdk:"port"`
	Url        types.String `tfsdk:"url"`
}

func (r *CacheobjectExpireResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CacheobjectExpireResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cacheobject_expire"
}

func (r *CacheobjectExpireResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CacheobjectExpireResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cacheobject_expire resource.",
			},
			"groupname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the content group to which the object belongs. It will display only the objects belonging to the specified content group. You must also set the Host parameter.",
			},
			"host": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Host name of the object. Parameter \"url\" must be specified.",
			},
			"httpmethod": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "HTTP request method that caused the object to be stored.",
			},
			"locator": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the cached object.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Host port of the object. You must also set the Host parameter.",
			},
			"url": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL of the particular object whose details is required. Parameter \"host\" must be specified along with the URL.",
			},
		},
	}
}

// ValidateConfig enforces the NITRO expire mandatory choice: (locator) XOR
// (url + host). NITRO marks none red and tfdata says is_required:false, but the
// appliance CLI rejects a call with neither (Required argument missing
// [url, locator]) and requires host when url is given.
func (r *CacheobjectExpireResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data CacheobjectExpireResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	hasLocator := !data.Locator.IsNull() && !data.Locator.IsUnknown()
	hasUrl := !data.Url.IsNull() && !data.Url.IsUnknown()
	hasHost := !data.Host.IsNull() && !data.Host.IsUnknown()

	if hasLocator && (hasUrl || hasHost) {
		resp.Diagnostics.AddError(
			"Invalid cacheobject_expire configuration",
			"Specify either \"locator\" OR (\"url\" + \"host\"), not both.",
		)
		return
	}
	if !hasLocator && !(hasUrl && hasHost) {
		resp.Diagnostics.AddError(
			"Invalid cacheobject_expire configuration",
			"You must specify either \"locator\" OR both \"url\" and \"host\".",
		)
		return
	}
}

func (r *CacheobjectExpireResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CacheobjectExpireResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Firing cacheobject expire action")
	payload := cacheobject_expireGetThePayloadFromthePlan(ctx, &data)

	// cacheobject has no add/set/delete; the only writes are POST ?action=<verb>.
	// Verb casing per the NITRO URL: action=expire.
	err := r.client.ActOnResource(service.Cacheobject.Type(), &payload, "expire")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to expire cacheobject, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Performed cacheobject expire action")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("cacheobject_expire")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheobjectExpireResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// expire is a one-shot action. NITRO exposes no GET endpoint that reports
	// expire-state, so Read is a pure preserve-state no-op.
	var data CacheobjectExpireResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for cacheobject_expire; NITRO has no query endpoint for expire state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheobjectExpireResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for expire; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state CacheobjectExpireResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for cacheobject_expire; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheobjectExpireResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// expire is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for cacheobject_expire; NITRO has no inverse of the expire action")
}

// cacheobject_expireGetThePayloadFromthePlan builds the expire payload. Per the
// NITRO doc the expire action accepts locator | (url,host[,port,groupname,httpmethod]).
func cacheobject_expireGetThePayloadFromthePlan(ctx context.Context, data *CacheobjectExpireResourceModel) cache.Cacheobject {
	tflog.Debug(ctx, "In cacheobject_expireGetThePayloadFromthePlan Function")

	cacheobject := cache.Cacheobject{}
	if !data.Locator.IsNull() && !data.Locator.IsUnknown() {
		cacheobject.Locator = utils.IntPtr(int(data.Locator.ValueInt64()))
	}
	if !data.Url.IsNull() && !data.Url.IsUnknown() {
		cacheobject.Url = data.Url.ValueString()
	}
	if !data.Host.IsNull() && !data.Host.IsUnknown() {
		cacheobject.Host = data.Host.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		cacheobject.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		cacheobject.Groupname = data.Groupname.ValueString()
	}
	if !data.Httpmethod.IsNull() && !data.Httpmethod.IsUnknown() {
		cacheobject.Httpmethod = data.Httpmethod.ValueString()
	}

	return cacheobject
}
