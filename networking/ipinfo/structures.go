package github.com/thomas-osgood/OGOR/networking/ipinfo

// structure defining the main IPInfo object that the
// user will create to query IPInfo.io
type IPInfoQuery struct {
	Options ConfigStruct
}

// structure defining the configuration settings for the
// IPInfoQuery object.
type ConfigStruct struct {
	Token string
}

// structure designed to hold the return value from
// IPInfo.io's query.
type IPInfoStruct struct {

	// the following objects are available to non-authenticated
	// users of IPInfo.io's API.

	IP           string `json:"ip"`
	Hostname     string `json:"hostname"`
	Anycast      bool   `json:"anycast"`
	City         string `json:"city"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	Location     string `json:"loc"`
	Organization string `json:"org"`
	ZipCode      string `json:"postal"`
	Timezone     string `json:"timezone"`
	Readme       string `json:"readme"`
	Bogon        bool   `json:"bogon"`

	// the following objects are only returned when the
	// user has been authenticated and has an account
	// with IPInfo.io.

	Asn     ASNStruct
	Company CompanyStruct
	Privacy PrivacyStruct
	Abuse   AbuseStruct
}

// structure designed to hold the Abuse info pice
// of the IPInfo return JSON.
type AbuseStruct struct {
	Address string `json:"address"`
	Country string `json:"country"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Network string `json:"network"`
	Phone   string `json:"phone"`
}

// structure designed to hold the ASN info piece
// of the IPInfo return JSON.
type ASNStruct struct {
	Asn    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

// structure designed to hold the Company info piece
// of the IPInfo return JSON.
type CompanyStruct struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

// structure designed to hold the Privacy info piece
// of the IPInfo return JSON.
type PrivacyStruct struct {
	Vpn     bool   `json:"vpn"`
	Proxy   bool   `json:"proxy"`
	Tor     bool   `json:"tor"`
	Relay   bool   `json:"relay"`
	Hosting bool   `json:"hosting"`
	Service string `json:"service"`
}

