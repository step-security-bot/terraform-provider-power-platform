package powerplatform_bapi

var (
	//https://api.bap.microsoft.com/providers/Microsoft.BusinessAppPlatform/locations/unitedstates/environmentCurrencies?api-version=2023-06-01
	EnvironmentCurrencyCodes = []string{"DJF", "ZAR", "ETB", "AED", "BHD", "DZD", "EGP", "IQD", "JOD", "KWD",
		"LBP", "LYD", "MAD", "OMR", "QAR", "SAR", "SYP", "TND", "YER", "CLP",
		"INR", "AZN", "RUB", "BYN", "BGN", "NGN", "BDT", "CNY", "EUR", "BAM",
		"USD", "CZK", "GBP", "DKK", "CHF", "MVR", "BTN", "XCD", "AUD", "BZD",
		"CAD", "HKD", "IDR", "JMD", "MYR", "NZD", "PHP", "SGD", "TTD", "XDR",
		"ARS", "BOB", "COP", "CRC", "CUP", "DOP", "GTQ", "HNL", "MXN", "NIO",
		"PAB", "PEN", "PYG", "UYU", "VES", "IRR", "XOF", "CDF", "XAF", "HTG",
		"ILS", "HRK", "HUF", "AMD", "ISK", "JPY", "GEL", "KZT", "KHR", "KRW",
		"KGS", "LAK", "MKD", "MNT", "BND", "MMK", "NOK", "NPR", "PKR", "PLN",
		"AFN", "BRL", "MDL", "RON", "RWF", "SEK", "LKR", "SOS", "ALL", "RSD",
		"KES", "TJS", "THB", "ERN", "TMT", "BWP", "TRY", "UAH", "UZS", "VND",
		"MOP", "TWD"}

	//https://api.bap.microsoft.com/providers/Microsoft.BusinessAppPlatform/locations?api-version=2023-06-01
	EnvironmentLocations = []string{"unitedstates",
		"europe", "asia", "australia", "india", "japan", "canada",
		"unitedkingdom", "southamerica", "france", "unitedarabemirates", "germany",
		"switzerland", "norway", "korea", "southafrica"}

	//https://api.bap.microsoft.com/providers/Microsoft.BusinessAppPlatform/locations/unitedstates/environmentLanguages?api-version=2023-06-01
	EnvironmentLanguages = []int64{1033, 1025, 1026, 1069, 1027, 3076, 2052, 1028, 1050, 1029, 1030, 1043, 1061,
		1035, 1036, 1110, 1031, 1032, 1037, 1081, 1038, 1040, 1041, 1087, 1042, 1062,
		1063, 1044, 1045, 1046, 2070, 1048, 1049, 2074, 1051, 1060, 3082, 1053, 1054,
		1055, 1058, 1066, 3098, 1086, 1057}

	EnvironmentTypes = []string{"Sandbox", "Production", "Trial", "Developer"}
)

type EnvironmentDto struct {
	Id         string                   `json:"id"`
	Type       string                   `json:"type"`
	Location   string                   `json:"location"`
	Name       string                   `json:"name"`
	Properties EnvironmentPropertiesDto `json:"properties"`
}

type EnvironmentPropertiesDto struct {
	TenantID                  string                       `json:"tenantId"`
	DisplayName               string                       `json:"displayName"`
	EnvironmentSku            string                       `json:"environmentSku"`
	DatabaseType              string                       `json:"databaseType"`
	LinkedEnvironmentMetadata LinkedEnvironmentMetadataDto `json:"linkedEnvironmentMetadata"`
	States                    StatesEnvironmentDto         `json:"states"`
}

type LinkedEnvironmentMetadataDto struct {
	DomainName      string `json:"domainName,omitempty"`
	InstanceURL     string `json:"instanceUrl"`
	BaseLanguage    int    `json:"baseLanguage"`
	SecurityGroupId string `json:"securityGroupId"`
	ResourceId      string `json:"resourceId"`
	Version         string `json:"version"`
	Currency        EnvironmentCreateCurrency
}

type StatesEnvironmentDto struct {
	Management StatesManagementEnvironmentDto `json:"management"`
}

type StatesManagementEnvironmentDto struct {
	Id string `json:"id"`
}

type EnvironmentDtoArray struct {
	Value []EnvironmentDto `json:"value"`
}

type EnvironmentCreateDto struct {
	Location   string                         `json:"location"`
	Properties EnvironmentCreatePropertiesDto `json:"properties"`
}

type EnvironmentCreatePropertiesDto struct {
	DisplayName               string                                      `json:"displayName"`
	DataBaseType              string                                      `json:"databaseType,omitempty"`
	BillingPolicy             string                                      `json:"billingPolicy,omitempty"`
	EnvironmentSku            string                                      `json:"environmentSku"`
	LinkedEnvironmentMetadata EnvironmentCreateLinkEnvironmentMetadataDto `json:"linkedEnvironmentMetadata"`
}

type EnvironmentCreateLinkEnvironmentMetadataDto struct {
	BaseLanguage    int                       `json:"baseLanguage"`
	DomainName      string                    `json:"domainName,omitempty"`
	Currency        EnvironmentCreateCurrency `json:"currency"`
	Templates       []string                  `json:"templates,omitempty"`
	SecurityGroupId string                    `json:"securityGroupId,omitempty"`
}
type EnvironmentCreateCurrency struct {
	Code string `json:"code"`
}

type EnvironmentDeleteDto struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
