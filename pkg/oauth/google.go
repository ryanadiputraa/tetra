package oauth

import (
	"context"
	"encoding/json"

	"github.com/ryanadiputraa/inventra/internal/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	googleAPIURL     = "https://www.googleapis.com"
	emailEndpoint    = "/auth/userinfo.email"
	profileEndpoint  = "/auth/userinfo.profile"
	userInfoEndpoint = "/oauth2/v2/userinfo"
)

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Picture   string `json:"picture"`
	Locale    string `json:"locale"`
}

type GoogleOauth interface {
	GetSigninURL() string
	ExchangeCodeWithUserInfo(ctx context.Context, code string) (*User, error)
}

type GoogleOauthConfig struct {
	State        string
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type oauth struct {
	config       *GoogleOauthConfig
	oauth2Config *oauth2.Config
}

func NewGoogleOauth(c *GoogleOauthConfig) GoogleOauth {
	return &oauth{
		config: c,
		oauth2Config: &oauth2.Config{
			ClientID:     c.ClientID,
			ClientSecret: c.ClientSecret,
			RedirectURL:  c.RedirectURL,
			Endpoint:     google.Endpoint,
			Scopes: []string{
				googleAPIURL + emailEndpoint,
				googleAPIURL + profileEndpoint,
			},
		},
	}
}

func (o *oauth) GetSigninURL() string {
	return o.oauth2Config.AuthCodeURL(o.config.State, oauth2.SetAuthURLParam("select_user", googleAPIURL))
}

func (o *oauth) ExchangeCodeWithUserInfo(ctx context.Context, code string) (user *User, err error) {
	token, err := o.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return
	}

	client := o.oauth2Config.Client(context.Background(), token)
	resp, err := client.Get(googleAPIURL + userInfoEndpoint)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.NewServiceErr(errors.Unauthorized, errors.Unauthorized)
		return
	}

	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return
	}
	return
}
