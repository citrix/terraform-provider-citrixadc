package appfwlearningsettings

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwlearningsettingsResourceModel describes the resource data model.
type AppfwlearningsettingsResourceModel struct {
	Id                                      types.String `tfsdk:"id"`
	Contenttypeautodeploygraceperiod        types.Int64  `tfsdk:"contenttypeautodeploygraceperiod"`
	Contenttypeminthreshold                 types.Int64  `tfsdk:"contenttypeminthreshold"`
	Contenttypepercentthreshold             types.Int64  `tfsdk:"contenttypepercentthreshold"`
	Cookieconsistencyautodeploygraceperiod  types.Int64  `tfsdk:"cookieconsistencyautodeploygraceperiod"`
	Cookieconsistencyminthreshold           types.Int64  `tfsdk:"cookieconsistencyminthreshold"`
	Cookieconsistencypercentthreshold       types.Int64  `tfsdk:"cookieconsistencypercentthreshold"`
	Creditcardnumberminthreshold            types.Int64  `tfsdk:"creditcardnumberminthreshold"`
	Creditcardnumberpercentthreshold        types.Int64  `tfsdk:"creditcardnumberpercentthreshold"`
	Crosssitescriptingautodeploygraceperiod types.Int64  `tfsdk:"crosssitescriptingautodeploygraceperiod"`
	Crosssitescriptingminthreshold          types.Int64  `tfsdk:"crosssitescriptingminthreshold"`
	Crosssitescriptingpercentthreshold      types.Int64  `tfsdk:"crosssitescriptingpercentthreshold"`
	Csrftagautodeploygraceperiod            types.Int64  `tfsdk:"csrftagautodeploygraceperiod"`
	Csrftagminthreshold                     types.Int64  `tfsdk:"csrftagminthreshold"`
	Csrftagpercentthreshold                 types.Int64  `tfsdk:"csrftagpercentthreshold"`
	Fieldconsistencyautodeploygraceperiod   types.Int64  `tfsdk:"fieldconsistencyautodeploygraceperiod"`
	Fieldconsistencyminthreshold            types.Int64  `tfsdk:"fieldconsistencyminthreshold"`
	Fieldconsistencypercentthreshold        types.Int64  `tfsdk:"fieldconsistencypercentthreshold"`
	Fieldformatautodeploygraceperiod        types.Int64  `tfsdk:"fieldformatautodeploygraceperiod"`
	Fieldformatminthreshold                 types.Int64  `tfsdk:"fieldformatminthreshold"`
	Fieldformatpercentthreshold             types.Int64  `tfsdk:"fieldformatpercentthreshold"`
	Profilename                             types.String `tfsdk:"profilename"`
	Sqlinjectionautodeploygraceperiod       types.Int64  `tfsdk:"sqlinjectionautodeploygraceperiod"`
	Sqlinjectionminthreshold                types.Int64  `tfsdk:"sqlinjectionminthreshold"`
	Sqlinjectionpercentthreshold            types.Int64  `tfsdk:"sqlinjectionpercentthreshold"`
	Starturlautodeploygraceperiod           types.Int64  `tfsdk:"starturlautodeploygraceperiod"`
	Starturlminthreshold                    types.Int64  `tfsdk:"starturlminthreshold"`
	Starturlpercentthreshold                types.Int64  `tfsdk:"starturlpercentthreshold"`
	Xmlattachmentminthreshold               types.Int64  `tfsdk:"xmlattachmentminthreshold"`
	Xmlattachmentpercentthreshold           types.Int64  `tfsdk:"xmlattachmentpercentthreshold"`
	Xmlwsiminthreshold                      types.Int64  `tfsdk:"xmlwsiminthreshold"`
	Xmlwsipercentthreshold                  types.Int64  `tfsdk:"xmlwsipercentthreshold"`
}

func (r *AppfwlearningsettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwlearningsettings resource.",
			},
			"contenttypeautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10080),
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"contenttypeminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Minimum threshold to learn Content Type information.",
			},
			"contenttypepercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum threshold in percent to learn Content Type information.",
			},
			"cookieconsistencyautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10080),
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"cookieconsistencyminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn cookies.",
			},
			"cookieconsistencypercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular cookie pattern for the learning engine to learn that cookie.",
			},
			"creditcardnumberminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Minimum threshold to learn Credit Card information.",
			},
			"creditcardnumberpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum threshold in percent to learn Credit Card information.",
			},
			"crosssitescriptingautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10080),
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"crosssitescriptingminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn HTML cross-site scripting patterns.",
			},
			"crosssitescriptingpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular cross-site scripting pattern for the learning engine to learn that cross-site scripting pattern.",
			},
			"csrftagautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10080),
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"csrftagminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn cross-site request forgery (CSRF) tags.",
			},
			"csrftagpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular CSRF tag for the learning engine to learn that CSRF tag.",
			},
			"fieldconsistencyautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10080),
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"fieldconsistencyminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn field consistency information.",
			},
			"fieldconsistencypercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular field consistency pattern for the learning engine to learn that field consistency pattern.",
			},
			"fieldformatautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10080),
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"fieldformatminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
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
				Default:     int64default.StaticInt64(10080),
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"sqlinjectionminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn HTML SQL injection patterns.",
			},
			"sqlinjectionpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular HTML SQL injection pattern for the learning engine to learn that HTML SQL injection pattern.",
			},
			"starturlautodeploygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10080),
				Description: "The number of minutes after the threshold hit alert the learned rule will be deployed",
			},
			"starturlminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn start URLs.",
			},
			"starturlpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular start URL pattern for the learning engine to learn that start URL.",
			},
			"xmlattachmentminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Minimum number of application firewall sessions that the learning engine must observe to learn XML attachment patterns.",
			},
			"xmlattachmentpercentthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum percentage of application firewall sessions that must contain a particular XML attachment pattern for the learning engine to learn that XML attachment pattern.",
			},
			"xmlwsiminthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
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

func appfwlearningsettingsGetThePayloadFromtheConfig(ctx context.Context, data *AppfwlearningsettingsResourceModel) appfw.Appfwlearningsettings {
	tflog.Debug(ctx, "In appfwlearningsettingsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwlearningsettings := appfw.Appfwlearningsettings{}
	if !data.Contenttypeautodeploygraceperiod.IsNull() {
		appfwlearningsettings.Contenttypeautodeploygraceperiod = utils.IntPtr(int(data.Contenttypeautodeploygraceperiod.ValueInt64()))
	}
	if !data.Contenttypeminthreshold.IsNull() {
		appfwlearningsettings.Contenttypeminthreshold = utils.IntPtr(int(data.Contenttypeminthreshold.ValueInt64()))
	}
	if !data.Contenttypepercentthreshold.IsNull() {
		appfwlearningsettings.Contenttypepercentthreshold = utils.IntPtr(int(data.Contenttypepercentthreshold.ValueInt64()))
	}
	if !data.Cookieconsistencyautodeploygraceperiod.IsNull() {
		appfwlearningsettings.Cookieconsistencyautodeploygraceperiod = utils.IntPtr(int(data.Cookieconsistencyautodeploygraceperiod.ValueInt64()))
	}
	if !data.Cookieconsistencyminthreshold.IsNull() {
		appfwlearningsettings.Cookieconsistencyminthreshold = utils.IntPtr(int(data.Cookieconsistencyminthreshold.ValueInt64()))
	}
	if !data.Cookieconsistencypercentthreshold.IsNull() {
		appfwlearningsettings.Cookieconsistencypercentthreshold = utils.IntPtr(int(data.Cookieconsistencypercentthreshold.ValueInt64()))
	}
	if !data.Creditcardnumberminthreshold.IsNull() {
		appfwlearningsettings.Creditcardnumberminthreshold = utils.IntPtr(int(data.Creditcardnumberminthreshold.ValueInt64()))
	}
	if !data.Creditcardnumberpercentthreshold.IsNull() {
		appfwlearningsettings.Creditcardnumberpercentthreshold = utils.IntPtr(int(data.Creditcardnumberpercentthreshold.ValueInt64()))
	}
	if !data.Crosssitescriptingautodeploygraceperiod.IsNull() {
		appfwlearningsettings.Crosssitescriptingautodeploygraceperiod = utils.IntPtr(int(data.Crosssitescriptingautodeploygraceperiod.ValueInt64()))
	}
	if !data.Crosssitescriptingminthreshold.IsNull() {
		appfwlearningsettings.Crosssitescriptingminthreshold = utils.IntPtr(int(data.Crosssitescriptingminthreshold.ValueInt64()))
	}
	if !data.Crosssitescriptingpercentthreshold.IsNull() {
		appfwlearningsettings.Crosssitescriptingpercentthreshold = utils.IntPtr(int(data.Crosssitescriptingpercentthreshold.ValueInt64()))
	}
	if !data.Csrftagautodeploygraceperiod.IsNull() {
		appfwlearningsettings.Csrftagautodeploygraceperiod = utils.IntPtr(int(data.Csrftagautodeploygraceperiod.ValueInt64()))
	}
	if !data.Csrftagminthreshold.IsNull() {
		appfwlearningsettings.Csrftagminthreshold = utils.IntPtr(int(data.Csrftagminthreshold.ValueInt64()))
	}
	if !data.Csrftagpercentthreshold.IsNull() {
		appfwlearningsettings.Csrftagpercentthreshold = utils.IntPtr(int(data.Csrftagpercentthreshold.ValueInt64()))
	}
	if !data.Fieldconsistencyautodeploygraceperiod.IsNull() {
		appfwlearningsettings.Fieldconsistencyautodeploygraceperiod = utils.IntPtr(int(data.Fieldconsistencyautodeploygraceperiod.ValueInt64()))
	}
	if !data.Fieldconsistencyminthreshold.IsNull() {
		appfwlearningsettings.Fieldconsistencyminthreshold = utils.IntPtr(int(data.Fieldconsistencyminthreshold.ValueInt64()))
	}
	if !data.Fieldconsistencypercentthreshold.IsNull() {
		appfwlearningsettings.Fieldconsistencypercentthreshold = utils.IntPtr(int(data.Fieldconsistencypercentthreshold.ValueInt64()))
	}
	if !data.Fieldformatautodeploygraceperiod.IsNull() {
		appfwlearningsettings.Fieldformatautodeploygraceperiod = utils.IntPtr(int(data.Fieldformatautodeploygraceperiod.ValueInt64()))
	}
	if !data.Fieldformatminthreshold.IsNull() {
		appfwlearningsettings.Fieldformatminthreshold = utils.IntPtr(int(data.Fieldformatminthreshold.ValueInt64()))
	}
	if !data.Fieldformatpercentthreshold.IsNull() {
		appfwlearningsettings.Fieldformatpercentthreshold = utils.IntPtr(int(data.Fieldformatpercentthreshold.ValueInt64()))
	}
	if !data.Profilename.IsNull() {
		appfwlearningsettings.Profilename = data.Profilename.ValueString()
	}
	if !data.Sqlinjectionautodeploygraceperiod.IsNull() {
		appfwlearningsettings.Sqlinjectionautodeploygraceperiod = utils.IntPtr(int(data.Sqlinjectionautodeploygraceperiod.ValueInt64()))
	}
	if !data.Sqlinjectionminthreshold.IsNull() {
		appfwlearningsettings.Sqlinjectionminthreshold = utils.IntPtr(int(data.Sqlinjectionminthreshold.ValueInt64()))
	}
	if !data.Sqlinjectionpercentthreshold.IsNull() {
		appfwlearningsettings.Sqlinjectionpercentthreshold = utils.IntPtr(int(data.Sqlinjectionpercentthreshold.ValueInt64()))
	}
	if !data.Starturlautodeploygraceperiod.IsNull() {
		appfwlearningsettings.Starturlautodeploygraceperiod = utils.IntPtr(int(data.Starturlautodeploygraceperiod.ValueInt64()))
	}
	if !data.Starturlminthreshold.IsNull() {
		appfwlearningsettings.Starturlminthreshold = utils.IntPtr(int(data.Starturlminthreshold.ValueInt64()))
	}
	if !data.Starturlpercentthreshold.IsNull() {
		appfwlearningsettings.Starturlpercentthreshold = utils.IntPtr(int(data.Starturlpercentthreshold.ValueInt64()))
	}
	if !data.Xmlattachmentminthreshold.IsNull() {
		appfwlearningsettings.Xmlattachmentminthreshold = utils.IntPtr(int(data.Xmlattachmentminthreshold.ValueInt64()))
	}
	if !data.Xmlattachmentpercentthreshold.IsNull() {
		appfwlearningsettings.Xmlattachmentpercentthreshold = utils.IntPtr(int(data.Xmlattachmentpercentthreshold.ValueInt64()))
	}
	if !data.Xmlwsiminthreshold.IsNull() {
		appfwlearningsettings.Xmlwsiminthreshold = utils.IntPtr(int(data.Xmlwsiminthreshold.ValueInt64()))
	}
	if !data.Xmlwsipercentthreshold.IsNull() {
		appfwlearningsettings.Xmlwsipercentthreshold = utils.IntPtr(int(data.Xmlwsipercentthreshold.ValueInt64()))
	}

	return appfwlearningsettings
}

func appfwlearningsettingsSetAttrFromGet(ctx context.Context, data *AppfwlearningsettingsResourceModel, getResponseData map[string]interface{}) *AppfwlearningsettingsResourceModel {
	tflog.Debug(ctx, "In appfwlearningsettingsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["contenttypeautodeploygraceperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Contenttypeautodeploygraceperiod = types.Int64Value(intVal)
		}
	} else {
		data.Contenttypeautodeploygraceperiod = types.Int64Null()
	}
	if val, ok := getResponseData["contenttypeminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Contenttypeminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Contenttypeminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["contenttypepercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Contenttypepercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Contenttypepercentthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["cookieconsistencyautodeploygraceperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cookieconsistencyautodeploygraceperiod = types.Int64Value(intVal)
		}
	} else {
		data.Cookieconsistencyautodeploygraceperiod = types.Int64Null()
	}
	if val, ok := getResponseData["cookieconsistencyminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cookieconsistencyminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Cookieconsistencyminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["cookieconsistencypercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cookieconsistencypercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Cookieconsistencypercentthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["creditcardnumberminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Creditcardnumberminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Creditcardnumberminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["creditcardnumberpercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Creditcardnumberpercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Creditcardnumberpercentthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["crosssitescriptingautodeploygraceperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Crosssitescriptingautodeploygraceperiod = types.Int64Value(intVal)
		}
	} else {
		data.Crosssitescriptingautodeploygraceperiod = types.Int64Null()
	}
	if val, ok := getResponseData["crosssitescriptingminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Crosssitescriptingminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Crosssitescriptingminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["crosssitescriptingpercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Crosssitescriptingpercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Crosssitescriptingpercentthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["csrftagautodeploygraceperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Csrftagautodeploygraceperiod = types.Int64Value(intVal)
		}
	} else {
		data.Csrftagautodeploygraceperiod = types.Int64Null()
	}
	if val, ok := getResponseData["csrftagminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Csrftagminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Csrftagminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["csrftagpercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Csrftagpercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Csrftagpercentthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["fieldconsistencyautodeploygraceperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldconsistencyautodeploygraceperiod = types.Int64Value(intVal)
		}
	} else {
		data.Fieldconsistencyautodeploygraceperiod = types.Int64Null()
	}
	if val, ok := getResponseData["fieldconsistencyminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldconsistencyminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Fieldconsistencyminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["fieldconsistencypercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldconsistencypercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Fieldconsistencypercentthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["fieldformatautodeploygraceperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldformatautodeploygraceperiod = types.Int64Value(intVal)
		}
	} else {
		data.Fieldformatautodeploygraceperiod = types.Int64Null()
	}
	if val, ok := getResponseData["fieldformatminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldformatminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Fieldformatminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["fieldformatpercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldformatpercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Fieldformatpercentthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["profilename"]; ok && val != nil {
		data.Profilename = types.StringValue(val.(string))
	} else {
		data.Profilename = types.StringNull()
	}
	if val, ok := getResponseData["sqlinjectionautodeploygraceperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sqlinjectionautodeploygraceperiod = types.Int64Value(intVal)
		}
	} else {
		data.Sqlinjectionautodeploygraceperiod = types.Int64Null()
	}
	if val, ok := getResponseData["sqlinjectionminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sqlinjectionminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Sqlinjectionminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["sqlinjectionpercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sqlinjectionpercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Sqlinjectionpercentthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["starturlautodeploygraceperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Starturlautodeploygraceperiod = types.Int64Value(intVal)
		}
	} else {
		data.Starturlautodeploygraceperiod = types.Int64Null()
	}
	if val, ok := getResponseData["starturlminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Starturlminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Starturlminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["starturlpercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Starturlpercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Starturlpercentthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["xmlattachmentminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlattachmentminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Xmlattachmentminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["xmlattachmentpercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlattachmentpercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Xmlattachmentpercentthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["xmlwsiminthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlwsiminthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Xmlwsiminthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["xmlwsipercentthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlwsipercentthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Xmlwsipercentthreshold = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Profilename.ValueString())

	return data
}
