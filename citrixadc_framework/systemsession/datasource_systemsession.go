package systemsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*SystemsessionDataSource)(nil)

func SYstemsessionDataSource() datasource.DataSource {
	return &SystemsessionDataSource{}
}

type SystemsessionDataSource struct {
	client *service.NitroClient
}

// SystemsessionDataSourceModel is a dedicated read-only model exposing the GET-only
// session fields. It is intentionally separate from SystemsessionResourceModel, which
// is kept minimal (id, sid, all) for the kill action.
type SystemsessionDataSourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Sid                   types.Int64  `tfsdk:"sid"`
	Username              types.String `tfsdk:"username"`
	Logintime             types.String `tfsdk:"logintime"`
	Logintimelocal        types.String `tfsdk:"logintimelocal"`
	Lastactivitytime      types.String `tfsdk:"lastactivitytime"`
	Lastactivitytimelocal types.String `tfsdk:"lastactivitytimelocal"`
	Expirytime            types.String `tfsdk:"expirytime"`
	Numofconnections      types.String `tfsdk:"numofconnections"`
	Currentconn           types.String `tfsdk:"currentconn"`
	Clienttype            types.String `tfsdk:"clienttype"`
	Partitionname         types.String `tfsdk:"partitionname"`
	Clientipaddress       types.String `tfsdk:"clientipaddress"`
}

func (d *SystemsessionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsession"
}

func (d *SystemsessionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SystemsessionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SystemsessionDataSourceSchema()
}

func (d *SystemsessionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SystemsessionDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sidName := fmt.Sprintf("%v", data.Sid.ValueInt64())

	getResponseData, err := d.client.FindResource(service.Systemsession.Type(), sidName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read systemsession, got error: %s", err))
		return
	}

	systemsessionSetAttrFromGetForDatasource(ctx, &data, getResponseData)

	// Datasource has no Create — set its ID here.
	data.Id = types.StringValue(sidName)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// systemsessionSetAttrFromGetForDatasource faithfully copies the GET response into the
// datasource model (Pattern 7 split: the resource model is action-only and does not use
// this setter).
func systemsessionSetAttrFromGetForDatasource(ctx context.Context, data *SystemsessionDataSourceModel, getResponseData map[string]interface{}) {
	tflog.Debug(ctx, "In systemsessionSetAttrFromGetForDatasource Function")

	setString := func(key string, target *types.String) {
		if val, ok := getResponseData[key]; ok && val != nil {
			*target = types.StringValue(fmt.Sprintf("%v", val))
		} else {
			*target = types.StringNull()
		}
	}

	setString("username", &data.Username)
	setString("logintime", &data.Logintime)
	setString("logintimelocal", &data.Logintimelocal)
	setString("lastactivitytime", &data.Lastactivitytime)
	setString("lastactivitytimelocal", &data.Lastactivitytimelocal)
	setString("expirytime", &data.Expirytime)
	setString("numofconnections", &data.Numofconnections)
	setString("currentconn", &data.Currentconn)
	setString("clienttype", &data.Clienttype)
	setString("partitionname", &data.Partitionname)
	setString("clientipaddress", &data.Clientipaddress)
}
