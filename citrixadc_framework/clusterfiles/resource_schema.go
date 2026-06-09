package clusterfiles

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ClusterfilesResourceModel describes the resource data model.
type ClusterfilesResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Mode types.List   `tfsdk:"mode"`
}

func (r *ClusterfilesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusterfiles resource.",
			},
			"mode": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "The directories and files to be synchronized. The available settings function as follows:\n Mode    Paths\n all           /nsconfig/ssl/\n                /var/netscaler/ssl/\n                /var/vpn/bookmark/\n                /nsconfig/dns/\n                /nsconfig/monitors/\n                /nsconfig/nstemplates/\n                /nsconfig/ssh/\n                /nsconfig/rc.netscaler\n                /nsconfig/resolv.conf\n                /nsconfig/inetd.conf\n                /nsconfig/syslog.conf\n                /nsconfig/ntp.conf\n                /nsconfig/httpd.conf\n                /nsconfig/sshd_config\n                /nsconfig/hosts\n                /nsconfig/enckey\n                /var/nslw.bin/etc/krb5.conf\n                /var/nslw.bin/etc/krb5.keytab\n                /var/lib/likewise/db/\n                /var/download/\n                /var/wi/tomcat/webapps/\n                /var/wi/tomcat/conf/Catalina/localhost/\n                /var/wi/java_home/lib/security/cacerts\n                /var/wi/java_home/jre/lib/security/cacerts\n                /var/netscaler/locdb/\nssl            /nsconfig/ssl/\n                 /var/netscaler/ssl/\nbookmarks     /var/vpn/bookmark/\ndns                  /nsconfig/dns/\nimports          /var/download/\nmisc               /nsconfig/license/\n                       /nsconfig/rc.conf\nall_plus_misc    Includes *all* files and /nsconfig/license/ and /nsconfig/rc.conf.\nDefault value: all\nPossible values = all, bookmarks, ssl, imports, misc, dns, krb, AAA, app_catalog, all_plus_misc, all_minus_misc",
			},
		},
	}
}

func clusterfilesGetThePayloadFromthePlan(ctx context.Context, data *ClusterfilesResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In clusterfilesGetThePayloadFromthePlan Function")

	// Build the sync action payload. mode is included only when set.
	clusterfiles := make(map[string]interface{})
	if !data.Mode.IsNull() && !data.Mode.IsUnknown() {
		var modeList []string
		data.Mode.ElementsAs(ctx, &modeList, false)
		clusterfiles["mode"] = modeList
	}

	return clusterfiles
}
