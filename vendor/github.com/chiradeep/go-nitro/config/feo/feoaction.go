package feo

type Feoaction struct {
	Builtin                interface{} `json:"builtin,omitempty"`
	Cachemaxage            int         `json:"cachemaxage,omitempty"`
	Clientsidemeasurements bool        `json:"clientsidemeasurements,omitempty"`
	Convertimporttolink    bool        `json:"convertimporttolink,omitempty"`
	Csscombine             bool        `json:"csscombine,omitempty"`
	Cssflattenimports      bool        `json:"cssflattenimports,omitempty"`
	Cssimginline           bool        `json:"cssimginline,omitempty"`
	Cssinline              bool        `json:"cssinline,omitempty"`
	Cssminify              bool        `json:"cssminify,omitempty"`
	Cssmovetohead          bool        `json:"cssmovetohead,omitempty"`
	Dnsshards              interface{} `json:"dnsshards,omitempty"`
	Domainsharding         string      `json:"domainsharding,omitempty"`
	Hits                   int         `json:"hits,omitempty"`
	Htmlminify             bool        `json:"htmlminify,omitempty"`
	Htmlrmattribquotes     bool        `json:"htmlrmattribquotes,omitempty"`
	Htmlrmdefaultattribs   bool        `json:"htmlrmdefaultattribs,omitempty"`
	Htmltrimurls           bool        `json:"htmltrimurls,omitempty"`
	Imgadddimensions       bool        `json:"imgadddimensions,omitempty"`
	Imggiftopng            bool        `json:"imggiftopng,omitempty"`
	Imginline              bool        `json:"imginline,omitempty"`
	Imglazyload            bool        `json:"imglazyload,omitempty"`
	Imgshrinkformobile     bool        `json:"imgshrinkformobile,omitempty"`
	Imgshrinktoattrib      bool        `json:"imgshrinktoattrib,omitempty"`
	Imgtojpegxr            bool        `json:"imgtojpegxr,omitempty"`
	Imgtowebp              bool        `json:"imgtowebp,omitempty"`
	Imgweaken              bool        `json:"imgweaken,omitempty"`
	Jpgoptimize            bool        `json:"jpgoptimize,omitempty"`
	Jpgprogressive         bool        `json:"jpgprogressive,omitempty"`
	Jscombine              bool        `json:"jscombine,omitempty"`
	Jsinline               bool        `json:"jsinline,omitempty"`
	Jsminify               bool        `json:"jsminify,omitempty"`
	Jsmovetoend            bool        `json:"jsmovetoend,omitempty"`
	Name                   string      `json:"name,omitempty"`
	Pageextendcache        bool        `json:"pageextendcache,omitempty"`
	Undefhits              int         `json:"undefhits,omitempty"`
}
