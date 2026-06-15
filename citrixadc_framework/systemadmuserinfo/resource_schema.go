package systemadmuserinfo

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemadmuserinfoResourceModel describes the resource data model.
type SystemadmuserinfoResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Username types.String `tfsdk:"username"`
}

func (r *SystemadmuserinfoResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemadmuserinfo resource.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of adm-user to log in syslogs.",
			},
		},
	}
}

func systemadmuserinfoGetThePayloadFromthePlan(ctx context.Context, data *SystemadmuserinfoResourceModel) system.Systemadmuserinfo {
	tflog.Debug(ctx, "In systemadmuserinfoGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemadmuserinfo := system.Systemadmuserinfo{}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		systemadmuserinfo.Username = data.Username.ValueString()
	}

	return systemadmuserinfo
}
