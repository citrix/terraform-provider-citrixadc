package gslbconfig

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbconfigResourceModel describes the resource data model.
type GslbconfigResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Command    types.String `tfsdk:"command"`
	Debug      types.Bool   `tfsdk:"debug"`
	Forcesync  types.String `tfsdk:"forcesync"`
	Nowarn     types.Bool   `tfsdk:"nowarn"`
	Preview    types.Bool   `tfsdk:"preview"`
	Saveconfig types.Bool   `tfsdk:"saveconfig"`
}

func (r *GslbconfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbconfig resource.",
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

func gslbconfigGetThePayloadFromthePlan(ctx context.Context, data *GslbconfigResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In gslbconfigGetThePayloadFromthePlan Function")

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
