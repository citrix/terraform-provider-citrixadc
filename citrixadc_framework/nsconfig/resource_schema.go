package nsconfig

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsconfigResourceModel describes the data model shared by the nsconfig datasource.
type NsconfigResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Async                   types.Bool   `tfsdk:"async"`
	All                     types.Bool   `tfsdk:"all"`
	Changedpassword         types.Bool   `tfsdk:"changedpassword"`
	Cip                     types.String `tfsdk:"cip"`
	Cipheader               types.String `tfsdk:"cipheader"`
	Config                  types.String `tfsdk:"config"`
	Config1                 types.String `tfsdk:"config1"`
	Config2                 types.String `tfsdk:"config2"`
	Configfile              types.String `tfsdk:"configfile"`
	Cookieversion           types.String `tfsdk:"cookieversion"`
	Crportrange             types.String `tfsdk:"crportrange"`
	Exclusivequotamaxclient types.Int64  `tfsdk:"exclusivequotamaxclient"`
	Exclusivequotaspillover types.Int64  `tfsdk:"exclusivequotaspillover"`
	Force                   types.Bool   `tfsdk:"force"`
	Ftpportrange            types.String `tfsdk:"ftpportrange"`
	Grantquotamaxclient     types.Int64  `tfsdk:"grantquotamaxclient"`
	Grantquotaspillover     types.Int64  `tfsdk:"grantquotaspillover"`
	Httpport                types.List   `tfsdk:"httpport"`
	Ifnum                   types.List   `tfsdk:"ifnum"`
	Ignoredevicespecific    types.Bool   `tfsdk:"ignoredevicespecific"`
	Ipaddress               types.String `tfsdk:"ipaddress"`
	Level                   types.String `tfsdk:"level"`
	Maxconn                 types.Int64  `tfsdk:"maxconn"`
	Maxreq                  types.Int64  `tfsdk:"maxreq"`
	Netmask                 types.String `tfsdk:"netmask"`
	Nsvlan                  types.Int64  `tfsdk:"nsvlan"`
	Outtype                 types.String `tfsdk:"outtype"`
	Pmtumin                 types.Int64  `tfsdk:"pmtumin"`
	Pmtutimeout             types.Int64  `tfsdk:"pmtutimeout"`
	Rbaconfig               types.String `tfsdk:"rbaconfig"`
	Responsefile            types.String `tfsdk:"responsefile"`
	Securecookie            types.String `tfsdk:"securecookie"`
	Securemanagementtd      types.Int64  `tfsdk:"securemanagementtd"`
	Securemanagementtraffic types.String `tfsdk:"securemanagementtraffic"`
	Tagged                  types.String `tfsdk:"tagged"`
	Template                types.Bool   `tfsdk:"template"`
	Timezone                types.String `tfsdk:"timezone"`
	Weakpassword            types.Bool   `tfsdk:"weakpassword"`
}

// nsconfigSetAttrFromGetForDatasource faithfully copies every field from the GET
// response (Pattern 7 datasource split) and sets the synthetic ID, since the
// datasource has no Create. It exposes the rich read-only status fields returned
// by the nsconfig GET.
func nsconfigSetAttrFromGetForDatasource(ctx context.Context, data *NsconfigResourceModel, getResponseData map[string]interface{}) *NsconfigResourceModel {
	tflog.Debug(ctx, "In nsconfigSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["Async"]; ok && val != nil {
		data.Async = types.BoolValue(val.(bool))
	} else {
		data.Async = types.BoolNull()
	}
	if val, ok := getResponseData["all"]; ok && val != nil {
		data.All = types.BoolValue(val.(bool))
	} else {
		data.All = types.BoolNull()
	}
	if val, ok := getResponseData["changedpassword"]; ok && val != nil {
		data.Changedpassword = types.BoolValue(val.(bool))
	} else {
		data.Changedpassword = types.BoolNull()
	}
	if val, ok := getResponseData["cip"]; ok && val != nil {
		data.Cip = types.StringValue(val.(string))
	} else {
		data.Cip = types.StringNull()
	}
	if val, ok := getResponseData["cipheader"]; ok && val != nil {
		data.Cipheader = types.StringValue(val.(string))
	} else {
		data.Cipheader = types.StringNull()
	}
	if val, ok := getResponseData["config"]; ok && val != nil {
		data.Config = types.StringValue(val.(string))
	} else {
		data.Config = types.StringNull()
	}
	if val, ok := getResponseData["config1"]; ok && val != nil {
		data.Config1 = types.StringValue(val.(string))
	} else {
		data.Config1 = types.StringNull()
	}
	if val, ok := getResponseData["config2"]; ok && val != nil {
		data.Config2 = types.StringValue(val.(string))
	} else {
		data.Config2 = types.StringNull()
	}
	if val, ok := getResponseData["configfile"]; ok && val != nil {
		data.Configfile = types.StringValue(val.(string))
	} else {
		data.Configfile = types.StringNull()
	}
	if val, ok := getResponseData["cookieversion"]; ok && val != nil {
		data.Cookieversion = types.StringValue(val.(string))
	} else {
		data.Cookieversion = types.StringNull()
	}
	if val, ok := getResponseData["crportrange"]; ok && val != nil {
		data.Crportrange = types.StringValue(val.(string))
	} else {
		data.Crportrange = types.StringNull()
	}
	if val, ok := getResponseData["exclusivequotamaxclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Exclusivequotamaxclient = types.Int64Value(intVal)
		}
	} else {
		data.Exclusivequotamaxclient = types.Int64Null()
	}
	if val, ok := getResponseData["exclusivequotaspillover"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Exclusivequotaspillover = types.Int64Value(intVal)
		}
	} else {
		data.Exclusivequotaspillover = types.Int64Null()
	}
	if val, ok := getResponseData["force"]; ok && val != nil {
		data.Force = types.BoolValue(val.(bool))
	} else {
		data.Force = types.BoolNull()
	}
	if val, ok := getResponseData["ftpportrange"]; ok && val != nil {
		data.Ftpportrange = types.StringValue(val.(string))
	} else {
		data.Ftpportrange = types.StringNull()
	}
	if val, ok := getResponseData["grantquotamaxclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Grantquotamaxclient = types.Int64Value(intVal)
		}
	} else {
		data.Grantquotamaxclient = types.Int64Null()
	}
	if val, ok := getResponseData["grantquotaspillover"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Grantquotaspillover = types.Int64Value(intVal)
		}
	} else {
		data.Grantquotaspillover = types.Int64Null()
	}
	if val, ok := getResponseData["httpport"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			intList := utils.StringListToIntList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.Int64Type, intList)
			data.Httpport = listValue
		} else {
			data.Httpport = types.ListNull(types.Int64Type)
		}
	} else {
		data.Httpport = types.ListNull(types.Int64Type)
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Ifnum = listValue
		} else {
			data.Ifnum = types.ListNull(types.StringType)
		}
	} else {
		data.Ifnum = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["ignoredevicespecific"]; ok && val != nil {
		data.Ignoredevicespecific = types.BoolValue(val.(bool))
	} else {
		data.Ignoredevicespecific = types.BoolNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["level"]; ok && val != nil {
		data.Level = types.StringValue(val.(string))
	} else {
		data.Level = types.StringNull()
	}
	if val, ok := getResponseData["maxconn"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxconn = types.Int64Value(intVal)
		}
	} else {
		data.Maxconn = types.Int64Null()
	}
	if val, ok := getResponseData["maxreq"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxreq = types.Int64Value(intVal)
		}
	} else {
		data.Maxreq = types.Int64Null()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["nsvlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nsvlan = types.Int64Value(intVal)
		}
	} else {
		data.Nsvlan = types.Int64Null()
	}
	if val, ok := getResponseData["outtype"]; ok && val != nil {
		data.Outtype = types.StringValue(val.(string))
	} else {
		data.Outtype = types.StringNull()
	}
	if val, ok := getResponseData["pmtumin"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pmtumin = types.Int64Value(intVal)
		}
	} else {
		data.Pmtumin = types.Int64Null()
	}
	if val, ok := getResponseData["pmtutimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pmtutimeout = types.Int64Value(intVal)
		}
	} else {
		data.Pmtutimeout = types.Int64Null()
	}
	if val, ok := getResponseData["rbaconfig"]; ok && val != nil {
		data.Rbaconfig = types.StringValue(val.(string))
	} else {
		data.Rbaconfig = types.StringNull()
	}
	if val, ok := getResponseData["responsefile"]; ok && val != nil {
		data.Responsefile = types.StringValue(val.(string))
	} else {
		data.Responsefile = types.StringNull()
	}
	if val, ok := getResponseData["securecookie"]; ok && val != nil {
		data.Securecookie = types.StringValue(val.(string))
	} else {
		data.Securecookie = types.StringNull()
	}
	if val, ok := getResponseData["securemanagementtd"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Securemanagementtd = types.Int64Value(intVal)
		}
	} else {
		data.Securemanagementtd = types.Int64Null()
	}
	if val, ok := getResponseData["securemanagementtraffic"]; ok && val != nil {
		data.Securemanagementtraffic = types.StringValue(val.(string))
	} else {
		data.Securemanagementtraffic = types.StringNull()
	}
	if val, ok := getResponseData["tagged"]; ok && val != nil {
		data.Tagged = types.StringValue(val.(string))
	} else {
		data.Tagged = types.StringNull()
	}
	if val, ok := getResponseData["template"]; ok && val != nil {
		data.Template = types.BoolValue(val.(bool))
	} else {
		data.Template = types.BoolNull()
	}
	if val, ok := getResponseData["timezone"]; ok && val != nil {
		data.Timezone = types.StringValue(val.(string))
	} else {
		data.Timezone = types.StringNull()
	}
	if val, ok := getResponseData["weakpassword"]; ok && val != nil {
		data.Weakpassword = types.BoolValue(val.(bool))
	} else {
		data.Weakpassword = types.BoolNull()
	}

	// Set ID for the datasource (no Create to set it).
	data.Id = types.StringValue("nsconfig-config")

	return data
}
