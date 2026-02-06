/*
Copyright 2025 Citrix Systems, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"
	"os"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcertkey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_appfwpolicy_binding"
)

// Ensure CitrixAdcFrameworkProvider satisfies various provider interfaces.
var _ provider.Provider = &CitrixAdcFrameworkProvider{}

// CitrixAdcFrameworkProvider defines the provider implementation.
type CitrixAdcFrameworkProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CitrixAdcFrameworkProviderModel describes the provider data model.
type CitrixAdcFrameworkProviderModel struct {
	Username           types.String `tfsdk:"username"`
	Password           types.String `tfsdk:"password"`
	Endpoint           types.String `tfsdk:"endpoint"`
	InsecureSkipVerify types.Bool   `tfsdk:"insecure_skip_verify"`
	ProxiedNs          types.String `tfsdk:"proxied_ns"`
	Partition          types.String `tfsdk:"partition"`
	DoLogin            types.Bool   `tfsdk:"do_login"`
}

// ProviderData contains the configured client for data sources and resources.
type ProviderData struct {
	Client   *service.NitroClient
	Username string
	Password string
	Endpoint string
}

func (p *CitrixAdcFrameworkProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "citrixadc"
	resp.Version = p.version
}

func (p *CitrixAdcFrameworkProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Description: "Username to login to the NetScaler",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Password to login to the NetScaler",
				Optional:    true,
			},
			"endpoint": schema.StringAttribute{
				Description: "The URL to the API",
				Optional:    true,
			},
			"insecure_skip_verify": schema.BoolAttribute{
				Description: "Ignore validity of endpoint TLS certificate if true",
				Optional:    true,
			},
			"proxied_ns": schema.StringAttribute{
				Description: "Target NS ip. When defined username, password and endpoint must refer to MAS.",
				Optional:    true,
			},
			"partition": schema.StringAttribute{
				Description: "Partition to target",
				Optional:    true,
			},
			"do_login": schema.BoolAttribute{
				Description: "Perform login to NetScaler",
				Optional:    true,
			},
		},
	}
}

func (p *CitrixAdcFrameworkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data CitrixAdcFrameworkProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// Validate required parameters
	username := data.Username.ValueString()
	if username == "" {
		username = os.Getenv("NS_LOGIN")
	}

	password := data.Password.ValueString()
	if password == "" {
		password = os.Getenv("NS_PASSWORD")
	}

	endpoint := data.Endpoint.ValueString()
	if endpoint == "" {
		endpoint = os.Getenv("NS_URL")
	}

	// Check if required parameters are empty and add errors
	if username == "" {
		resp.Diagnostics.AddError(
			"Missing required parameter",
			"The 'username' parameter is required. It can be set via the provider configuration or the NS_LOGIN environment variable.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddError(
			"Missing required parameter",
			"The 'password' parameter is required. It can be set via the provider configuration or the NS_PASSWORD environment variable.",
		)
	}

	if endpoint == "" {
		resp.Diagnostics.AddError(
			"Missing required parameter",
			"The 'endpoint' parameter is required. It can be set via the provider configuration or the NS_URL environment variable.",
		)
	}

	// Return early if any required parameters are missing
	if resp.Diagnostics.HasError() {
		return
	}

	proxiedNs := os.Getenv("_MPS_API_PROXY_MANAGED_INSTANCE_IP")
	if !data.ProxiedNs.IsNull() {
		proxiedNs = data.ProxiedNs.ValueString()
	}

	partition := os.Getenv("NS_PARTITION")
	if !data.Partition.IsNull() {
		partition = data.Partition.ValueString()
	}

	insecureSkipVerify := false
	if !data.InsecureSkipVerify.IsNull() {
		insecureSkipVerify = data.InsecureSkipVerify.ValueBool()
	}

	doLogin := false
	if !data.DoLogin.IsNull() {
		doLogin = data.DoLogin.ValueBool()
	}

	userHeaders := map[string]string{
		"User-Agent": "terraform-ctxadc-framework",
	}

	params := service.NitroParams{
		Url:       endpoint,
		Username:  username,
		Password:  password,
		ProxiedNs: proxiedNs,
		SslVerify: !insecureSkipVerify,
		Headers:   userHeaders,
	}

	client, err := service.NewNitroClientFromParams(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create Nitro client",
			"Unable to create client:\n\n"+err.Error(),
		)
		return
	}

	if doLogin {
		client.Login()
	}

	if partition != "" {
		nspartition := make(map[string]interface{})
		nspartition["partitionname"] = partition
		err := client.ActOnResource("nspartition", &nspartition, "Switch")
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to switch partition",
				"Unable to switch partition:\n\n"+err.Error(),
			)
			return
		}
	}

	// providerData := &ProviderData{
	// 	Client:   client,
	// 	Username: username,
	// 	Password: password,
	// 	Endpoint: endpoint,
	// }

	resp.DataSourceData = &client
	resp.ResourceData = &client

	tflog.Info(ctx, "Configured CitrixADC Framework Provider", map[string]any{"success": true})
}

func (p *CitrixAdcFrameworkProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		lbparameter.NewLbParameterResource,
		sslcertkey.NewSslCertKeyResource,
		vpnvserver_appfwpolicy_binding.NewVpnvserverAppfwpolicyBindingResource,
	}
}

func (p *CitrixAdcFrameworkProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		lbparameter.LBParameterDataSource,
		sslcertkey.SslCertKeyDataSource,
		vpnvserver_appfwpolicy_binding.VpnvserverAppfwpolicyBindingDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CitrixAdcFrameworkProvider{
			version: version,
		}
	}
}
