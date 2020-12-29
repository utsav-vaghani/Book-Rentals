package models

//Token data
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUUID   string `json:"access_uuid"`
	RefreshUUID  string `json:"refresh_uuid"`
	AtExp        int64  `json:"at_exp"`
	RtExp        int64  `json:"rt_exp"`
}
