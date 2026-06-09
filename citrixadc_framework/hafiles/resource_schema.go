package hafiles

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// HafilesResourceModel describes the resource data model.
type HafilesResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Mode types.List   `tfsdk:"mode"`
}

func (r *HafilesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the hafiles resource.",
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

func hafilesGetThePayloadFromthePlan(ctx context.Context, data *HafilesResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In hafilesGetThePayloadFromthePlan Function")

	// Build the sync action payload. mode is included only when set.
	hafiles := make(map[string]interface{})
	if !data.Mode.IsNull() && !data.Mode.IsUnknown() {
		var modeList []string
		data.Mode.ElementsAs(ctx, &modeList, false)
		hafiles["mode"] = modeList
	}

	return hafiles
}
