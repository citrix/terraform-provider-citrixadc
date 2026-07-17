package gslbconfig

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

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GslbconfigSyncResource{}
var _ resource.ResourceWithConfigure = (*GslbconfigSyncResource)(nil)
var _ resource.ResourceWithImportState = (*GslbconfigSyncResource)(nil)

func NewGslbconfigSyncResource() resource.Resource {
	return &GslbconfigSyncResource{}
}

// GslbconfigSyncResource defines the resource implementation.
type GslbconfigSyncResource struct {
	client *service.NitroClient
}

// GslbconfigSyncResourceModel describes the resource data model.
type GslbconfigSyncResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Command    types.String `tfsdk:"command"`
	Debug      types.Bool   `tfsdk:"debug"`
	Forcesync  types.String `tfsdk:"forcesync"`
	Nowarn     types.Bool   `tfsdk:"nowarn"`
	Preview    types.Bool   `tfsdk:"preview"`
	Saveconfig types.Bool   `tfsdk:"saveconfig"`
}

func (r *GslbconfigSyncResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbconfigSyncResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbconfig_sync"
}

func (r *GslbconfigSyncResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbconfigSyncResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbconfig_sync resource.",
			},
			"command": schema.StringAttribute{
				Optional:    true,
				Description: "Run the specified command on the master node and then on all the slave nodes. You cannot use this option with the force sync and preview options.",
			},
			"debug": schema.BoolAttribute{
				Optional:    true,
				Description: "Generate verbose output when synchronizing the GSLB sites. The Debug option generates more verbose output than the sync gslb config command in which the option is not used, and is useful for analyzing synchronization issues.",
			},
			"forcesync": schema.StringAttribute{
				Optional:    true,
				Description: "Force synchronization of the specified site even if a dependent configuration on the remote site is preventing synchronization or if one or more GSLB entities on the remote site have the same name but are of a different type. You can specify either the name of the remote site that you want to synchronize with the local site, or you can specify All Sites in the configuration utility (the string all-sites in the CLI). If you specify All Sites, all the sites in the GSLB setup are synchronized with the site on the master node.\nNote: If you select the Force Sync option, the synchronization starts without displaying the commands that are going to be executed.",
			},
			"nowarn": schema.BoolAttribute{
				Optional:    true,
				Description: "Suppress the warning and the confirmation prompt that are displayed before site synchronization begins. This option can be used in automation scripts that must not be interrupted by a prompt.",
			},
			"preview": schema.BoolAttribute{
				Optional:    true,
				Description: "Do not synchronize the GSLB sites, but display the commands that would be applied on the slave node upon synchronization. Mutually exclusive with the Save Configuration option.",
			},
			"saveconfig": schema.BoolAttribute{
				Optional:    true,
				Description: "Save the configuration on all the nodes participating in the synchronization process, automatically. The master saves its configuration immediately before synchronization begins. Slave nodes save their configurations after the process of synchronization is complete. A slave node saves its configuration only if the configuration difference was successfully applied to it. Mutually exclusive with the Preview option.",
			},
		},
	}
}

func (r *GslbconfigSyncResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbconfigSyncResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbconfig_sync resource")

	// gslbconfig exposes only the POST ?action=sync action on NITRO.
	// There is no add/get/update/delete endpoint. Use ActOnResource with
	// the "sync" verb.
	payload := gslbconfig_syncGetThePayloadFromthePlan(ctx, &data)

	err := r.client.ActOnResource(service.Gslbconfig.Type(), payload, "sync")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to sync gslbconfig, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Synced gslbconfig_sync resource")

	// Synthetic constant ID - there is no NITRO identity for this action resource.
	data.Id = types.StringValue("gslbconfig_sync")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbconfigSyncResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbconfigSyncResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for gslbconfig_sync: NITRO exposes no GET endpoint for this
	// action-only resource, so drift detection is impossible. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for gslbconfig_sync; no GET endpoint on NITRO side")

	// Save (unchanged) data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbconfigSyncResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state GslbconfigSyncResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for gslbconfig_sync; NITRO exposes no update endpoint for
	// this action-only resource (only ?action=sync). Changes re-run sync via
	// Create on the next apply.
	tflog.Debug(ctx, "Update is a no-op for gslbconfig_sync; no update endpoint on NITRO side")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbconfigSyncResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbconfigSyncResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op for gslbconfig_sync: NITRO exposes no DELETE endpoint for
	// this action-only resource. Just remove from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for gslbconfig_sync; no DELETE endpoint on NITRO side")
	tflog.Trace(ctx, "Removed gslbconfig_sync from Terraform state")
}

func gslbconfig_syncGetThePayloadFromthePlan(ctx context.Context, data *GslbconfigSyncResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In gslbconfig_syncGetThePayloadFromthePlan Function")

	// Build the sync action payload. Only the attributes the user set are included.
	gslbconfig := make(map[string]interface{})
	if !data.Command.IsNull() && !data.Command.IsUnknown() {
		gslbconfig["command"] = data.Command.ValueString()
	}
	if !data.Debug.IsNull() && !data.Debug.IsUnknown() {
		gslbconfig["debug"] = data.Debug.ValueBool()
	}
	if !data.Forcesync.IsNull() && !data.Forcesync.IsUnknown() {
		gslbconfig["forcesync"] = data.Forcesync.ValueString()
	}
	if !data.Nowarn.IsNull() && !data.Nowarn.IsUnknown() {
		gslbconfig["nowarn"] = data.Nowarn.ValueBool()
	}
	if !data.Preview.IsNull() && !data.Preview.IsUnknown() {
		gslbconfig["preview"] = data.Preview.ValueBool()
	}
	if !data.Saveconfig.IsNull() && !data.Saveconfig.IsUnknown() {
		gslbconfig["saveconfig"] = data.Saveconfig.ValueBool()
	}

	return gslbconfig
}
