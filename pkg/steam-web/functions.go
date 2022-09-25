package steamweb

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL           = "https://steamcommunity.com/profiles/"
	avatarUrlSelector = "#responsive_page_template_content > div.no_header.profile_page > div.profile_header_bg > div > div > div > div.playerAvatar > div > img"
	nickNameSelector  = "#responsive_page_template_content > div.no_header.profile_page > div.profile_header_bg > div > div > div > div.profile_header_centered_persona > div.persona_name > span.actual_persona_name"
)

type ProfilePage struct {
	Url       string
	SteamID   string
	NickName  string
	AvatarUrl string
}

func Parse(steamID string) (ProfilePage, error) {
	pp := ProfilePage{
		Url: baseURL + steamID,
		SteamID: steamID,
	}

	resp, err := http.Get(pp.Url)
	if err != nil {
		return pp, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return pp, fmt.Errorf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return pp, err
	}

	pp.NickName = doc.Find(nickNameSelector).Text()
	pp.AvatarUrl = doc.Find(avatarUrlSelector).AttrOr("src", "")

	return pp, nil
}
