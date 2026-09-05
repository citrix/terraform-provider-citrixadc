package auditsyslogpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// auditsyslogpolicyGlobalbindingDSAttrs returns the nested attributes for the
// inline globalbinding block on the data source. It is defined here (not in
// datasource_schema.go) so the merged schema file stays a flat top-level
// attribute/block map matching the AuditsyslogpolicyDataSourceModel.
func auditsyslogpolicyGlobalbindingDSAttrs() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"feature": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The feature to be checked while applying this config.",
		},
		"globalbindtype": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The bind type for the global binding.",
		},
		"gotopriorityexpression": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.",
		},
		"nextfactor": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "On success invoke label. Applicable for advanced authentication policy binding.",
		},
		"priority": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The priority of the policy binding.",
		},
	}
}

var _ datasource.DataSource = (*AuditsyslogpolicyDataSource)(nil)

func AUditsyslogpolicyDataSource() datasource.DataSource {
	return &AuditsyslogpolicyDataSource{}
}

type AuditsyslogpolicyDataSource struct {
	client *service.NitroClient
}

func (d *AuditsyslogpolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditsyslogpolicy"
}

func (d *AuditsyslogpolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AuditsyslogpolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AuditsyslogpolicyDataSourceSchema()
}

func (d *AuditsyslogpolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AuditsyslogpolicyDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	name_Name := data.Name.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Auditsyslogpolicy.Type(), name_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read auditsyslogpolicy, got error: %s", err))
		return
	}

	auditsyslogpolicyDataSourceSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
