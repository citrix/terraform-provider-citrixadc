package systemcollectionparam

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemcollectionparamResource{}
var _ resource.ResourceWithConfigure = (*SystemcollectionparamResource)(nil)
var _ resource.ResourceWithImportState = (*SystemcollectionparamResource)(nil)

func NewSystemcollectionparamResource() resource.Resource {
	return &SystemcollectionparamResource{}
}

// SystemcollectionparamResource defines the resource implementation.
type SystemcollectionparamResource struct {
	client *service.NitroClient
}

// SystemcollectionparamResourceModel describes the resource data model.
//
// Mirrors the SDK v2 `citrixadc_systemcollectionparam` resource: systemcollectionparam
// is an unnamed singleton config object with no ADD/DELETE NITRO endpoint, so
// Create/Update both issue `set systemcollectionparam` (UpdateUnnamedResource / PUT),
// Read fetches the singleton via FindResource, and Delete is a no-op (the object has
// no delete API). The ID is a synthetic constant because the object has no key.
//
// Every schema attribute has a matching tfsdk field. All three config attributes
// (communityname, datapath, loglevel) are TypeString and Optional+Computed, exactly
// as in the SDK v2 schema.
type SystemcollectionparamResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Communityname types.String `tfsdk:"communityname"`
	Datapath      types.String `tfsdk:"datapath"`
	Loglevel      types.String `tfsdk:"loglevel"`
}

func (r *SystemcollectionparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemcollectionparam"
}

func (r *SystemcollectionparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The synthetic ID of the systemcollectionparam resource.",
			},
			"communityname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SNMPv1 community name for authentication.",
			},
			"datapath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "specify the data path to the database.",
			},
			"loglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "specify the log level. Possible values CRITICAL,WARNING,INFO,DEBUG1,DEBUG2",
			},
		},
	}
}

func (r *SystemcollectionparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemcollectionparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// getPayload builds the system.Systemcollectionparam PUT body from the model.
func (r *SystemcollectionparamResource) getPayload(ctx context.Context, data *SystemcollectionparamResourceModel) system.Systemcollectionparam {
	tflog.Debug(ctx, "In SystemcollectionparamResource getPayload Function")

	systemcollectionparam := system.Systemcollectionparam{}
	if !data.Communityname.IsNull() && !data.Communityname.IsUnknown() {
		systemcollectionparam.Communityname = data.Communityname.ValueString()
	}
	if !data.Datapath.IsNull() && !data.Datapath.IsUnknown() {
		systemcollectionparam.Datapath = data.Datapath.ValueString()
	}
	if !data.Loglevel.IsNull() && !data.Loglevel.IsUnknown() {
		systemcollectionparam.Loglevel = data.Loglevel.ValueString()
	}
	return systemcollectionparam
}

func (r *SystemcollectionparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemcollectionparamResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemcollectionparam resource (set systemcollectionparam)")
	systemcollectionparam := r.getPayload(ctx, &data)
	if err := r.client.UpdateUnnamedResource(service.Systemcollectionparam.Type(), &systemcollectionparam); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemcollectionparam, got error: %s", err))
		return
	}

	// Synthetic ID for the unnamed singleton object (mirrors the opaque ID the
	// SDK v2 resource assigned; Read/Delete never depend on its value).
	data.Id = types.StringValue("systemcollectionparam-config")

	r.readFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemcollectionparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemcollectionparamResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Reading systemcollectionparam resource")
	r.readFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemcollectionparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemcollectionparamResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemcollectionparam resource")

	// Mirror SDK v2 updateSystemcollectionparamFunc: only issue the set call when
	// one of the config attributes actually changed.
	hasChange := false
	if !data.Communityname.Equal(state.Communityname) {
		hasChange = true
	}
	if !data.Datapath.Equal(state.Datapath) {
		hasChange = true
	}
	if !data.Loglevel.Equal(state.Loglevel) {
		hasChange = true
	}

	if hasChange {
		systemcollectionparam := r.getPayload(ctx, &data)
		if err := r.client.UpdateUnnamedResource(service.Systemcollectionparam.Type(), &systemcollectionparam); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemcollectionparam, got error: %s", err))
			return
		}
	} else {
		tflog.Debug(ctx, "No changes detected for systemcollectionparam resource, skipping update")
	}

	r.readFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemcollectionparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Mirrors SDK v2 deleteSystemcollectionparamFunc: systemcollectionparam does not
	// support a delete operation, so Delete only removes the resource from state.
	tflog.Debug(ctx, "Deleting systemcollectionparam: no delete API, removing from state only")
}

// readFromApi reads the live systemcollectionparam singleton and populates the model.
// On read failure it clears the ID (mirrors SDK v2 readSystemcollectionparamFunc,
// which does `d.SetId("")` when FindResource fails).
//
// Note: communityname is intentionally NOT read back from the API, exactly as in the
// SDK v2 Read (where `d.Set("communityname", ...)` is commented out) — the SNMPv1
// community name is not reliably returned and reading it would cause perpetual diffs.
// The plan/state value is preserved instead; any Unknown (Optional+Computed with no
// config value) is resolved to Null to avoid an "inconsistent result after apply".
func (r *SystemcollectionparamResource) readFromApi(ctx context.Context, data *SystemcollectionparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemcollectionparam.Type(), "")
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Clearing systemcollectionparam state, got error: %s", err))
		data.Id = types.StringNull()
		return
	}

	// communityname: preserve the configured value (SDK v2 does not read it back);
	// resolve Unknown -> Null so a Computed attribute is always known after apply.
	if data.Communityname.IsUnknown() {
		data.Communityname = types.StringNull()
	}

	if val, ok := getResponseData["datapath"]; ok && val != nil {
		data.Datapath = types.StringValue(val.(string))
	} else {
		data.Datapath = types.StringNull()
	}
	if val, ok := getResponseData["loglevel"]; ok && val != nil {
		data.Loglevel = types.StringValue(val.(string))
	} else {
		data.Loglevel = types.StringNull()
	}
}
