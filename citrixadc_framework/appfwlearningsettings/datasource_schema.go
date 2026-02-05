package appfwlearningsettings

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwlearningsettingsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"contenttypeautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"contenttypeminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum threshold to learn Content Type information.",
			},
			"contenttypepercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum threshold in percent to learn Content Type information.",
			},
			"cookieconsistencyautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"cookieconsistencyminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn cookies.",
			},
			"cookieconsistencypercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular cookie pattern for the learning engine to learn that cookie.",
			},
			"creditcardnumberminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum threshold to learn Credit Card information.",
			},
			"creditcardnumberpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum threshold in percent to learn Credit Card information.",
			},
			"crosssitescriptingautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"crosssitescriptingminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn HTML cross-site scripting patterns.",
			},
			"crosssitescriptingpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular cross-site scripting pattern for the learning engine to learn that cross-site scripting pattern.",
			},
			"csrftagautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"csrftagminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn cross-site request forgery (CSRF) tags.",
			},
			"csrftagpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular CSRF tag for the learning engine to learn that CSRF tag.",
			},
			"fieldconsistencyautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"fieldconsistencyminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn field consistency information.",
			},
			"fieldconsistencypercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular field consistency pattern for the learning engine to learn that field consistency pattern.",
			},
			"fieldformatautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"fieldformatminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn field formats.",
			},
			"fieldformatpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular web form field pattern for the learning engine to recommend a field format for that form field.",
			},
			"profilename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile.",
			},
			"sqlinjectionautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"sqlinjectionminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn HTML SQL injection patterns.",
			},
			"sqlinjectionpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular HTML SQL injection pattern for the learning engine to learn that HTML SQL injection pattern.",
			},
			"starturlautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"starturlminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn start URLs.",
			},
			"starturlpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular start URL pattern for the learning engine to learn that start URL.",
			},
			"xmlattachmentminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn XML attachment patterns.",
			},
			"xmlattachmentpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular XML attachment pattern for the learning engine to learn that XML attachment pattern.",
			},
			"xmlwsiminthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn web services interoperability (WSI) information.",
			},
			"xmlwsipercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular pattern for the learning engine to learn a web services interoperability (WSI) pattern.",
			},
		},
	}
}
