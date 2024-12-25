package yoomoney

type Client struct {
	ClientId    string `json:"client_id" url:"client_id"` // ClientId который выдали после регистрации приложения
	AccessToken string `url:"access_token"`               // AccessToken живет 3 года
	RedirectUri string `url:"redirect_uri"`               // RedirectUri  который указывали при регистрации приложения
}

func NewClient(clientId string, accessToken string, redirectUri string) *Client {
	return &Client{
		ClientId:    clientId,
		AccessToken: accessToken,
		RedirectUri: redirectUri,
	}
}
