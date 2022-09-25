package tekken

import (
	"fmt"

	"golang.org/x/sys/windows"

	"github.com/nullentrypoint/tekken7-web-displayer/pkg/go-steamworks"
	"github.com/nullentrypoint/tekken7-web-displayer/pkg/irma"
)

type Api struct {
	handle       windows.Handle
	steamApiInfo *irma.ModuleInfo
	hooks        map[Hook]Hook
	Data         Data
	OldData      Data
}

type Data struct {
	UserSteamID     steamworks.CSteamID
	OpponentSteamID steamworks.CSteamID
}

type PlayerInfo struct {
	Name        string
	AvatarUrl   string
	SteamID     string
	IP          string
	Location    string
	CountryCode string
}

func (r PlayerInfo) String() string {
	return fmt.Sprintf("Name: %v SteamID: %v IP: %v Location: %v", r.Name, r.SteamID, r.IP, r.Location)
}
