package model

type AdCoount struct {
	ID                       string                `json:"id"`
	AccountID                string                `json:"account_id"`
	AccountGroups            []AccountGroup        `json:"account_groups"`
	AccountStatus            int                   `json:"account_status"` //1 = Active	2 = Disabled	3 = Unsettled 7 = Pending Review	9 = In Grace Period	101 = temporarily unavailable	100 = pending closure.
	Age                      float64               `json:"age"`
	AgencyClientDeclaration  interface{}           `json:"agency_client_declaration"`
	AmountSpent              string                `json:"amount_spent"`
	Balance                  string                `json:"balance"`
	Business                 Business              `json:"business"`
	BusinessCity             string                `json:"business_city"`
	BusinessCountryCode      string                `json:"business_country_code"`
	BusinessName             string                `json:"business_name"`
	BusinessState            string                `json:"business_state"`
	BusinessStreet           string                `json:"business_street"`
	BusinessStreet2          string                `json:"business_street2"`
	BusinessZip              string                `json:"business_zip"`
	Capabilities             interface{}           `json:"capabilities"`
	CreatedTime              string                `json:"created_time"`
	Currency                 string                `json:"currency"`
	FundingSource            string                `json:"funding_source"`
	FundingSourceDetails     []FundingSourceDetail `json:"funding_source_details"`
	IsPersonal               string                `json:"is_personal"`
	MediaAgency              string                `json:"media_agency"`
	Name                     string                `json:"name"`
	OffsitePixelsTosAccepted bool                  `json:"offsite_pixels_tos_accepted"`
	Partner                  string                `json:"partner"`
	RfSpec                   interface{}           `json:"rf_spec"`
	SpendCap                 string                `json:"spend_cap"`
	TaxIdStatus              int                   `json:"tax_id_status"`
	TimezoneId               int                   `json:"timezone_id"`
	TimezoneName             string                `json:"timezone_name"`
	TimezoneOffsetHoursUtc   float64               `json:"timezone_offset_hours_utc"`
	TosAccepted              interface{}           `json:"tos_accepted"`
	Users                    []User                `json:"users"`
}

type AccountGroup struct {
	Name           string `json:"name"`
	Status         string `json:"status"`
	AccountGroupId string `json:"account_group_id "`
}
type User struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Role        interface{} `json:"role"`
	Permissions interface{} `json:"permissions"`
}
type FundingSourceDetail struct {
	ID             string `json:"ID"`
	COUPON         string `json:"COUPON"`
	AMOUNT         string `json:"AMOUNT"`
	CURRENCY       string `json:"CURRENCY"`
	DISPLAY_AMOUNT string `json:"DISPLAY_AMOUNT"`
	EXPIRATION     string `json:"EXPIRATION"`
	DISPLAY_STRING string `json:"DISPLAY_STRING"`
}
