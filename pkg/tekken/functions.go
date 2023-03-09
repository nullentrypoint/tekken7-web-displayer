package tekken

import (
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows"

	"github.com/nullentrypoint/tekken7-web-displayer/pkg/go-steamworks"
	"github.com/nullentrypoint/tekken7-web-displayer/pkg/ip"
	"github.com/nullentrypoint/tekken7-web-displayer/pkg/irma"
	steamweb "github.com/nullentrypoint/tekken7-web-displayer/pkg/steam-web"
)

func NewApi() *Api {
	return &Api{
		hooks: make(map[Hook]Hook),
	}
}

func (r *Api) Run() {
	for {
		handle, err := irma.FindProcessByName(TEKKEN_EXE_NAME)
		if err == nil && handle != windows.InvalidHandle {
			log.Println("[INFO] Tekken process found!")
			time.Sleep(time.Second * 10)
			err := r.processHandler(handle)
			if err != nil {
				log.Println("[ERROR]", err)
			}
			break
		}
		log.Println("[WARNING] Tekken not found. Try again...")
		time.Sleep(time.Second * 2)
	}
}

func (r *Api) Subscribe(hook Hook) {
	r.hooks[hook] = hook
}

func (r *Api) Unsubscribe(hook Hook) {
	delete(r.hooks, hook)
}

func (r *Api) fireEventSteamWorksInit() {
	for _, hook := range r.hooks {
		hook.EventSteamWorksInit()
	}
}

func (r *Api) fireEventNewChallenger(data *Data) {
	if len(r.hooks) == 0 {
		return
	}

	info := GetPlayerInfo(data.OpponentSteamID)

	for _, hook := range r.hooks {
		hook.EventNewChallenger(info)
	}
}

func GetPlayerInfo(steamID steamworks.CSteamID) PlayerInfo {
	var sessionState steamworks.P2PSessionState_t
	info := PlayerInfo{
		//Name: steamworks.SteamFriends().GetFriendPersonaName(steamID),
		SteamID: fmt.Sprint(steamID),
		Time: time.Now().Format(timeLayout),
	}

	pp, err := steamweb.Parse(info.SteamID)
	if err != nil {
		log.Println("[ERROR] steamweb.Parse", err)
	}

	info.Name = pp.NickName
	info.AvatarUrl = pp.AvatarUrl

	if steamworks.SteamNetworking().GetP2PSessionState(steamID, &sessionState) && sessionState.M_nRemoteIP != 0 {
		info.IP = ip.Itoa(int64(sessionState.M_nRemoteIP))
		ipInfo, err := ip.GetInfo(info.IP)
		if err != nil {
			log.Println("[ERROR] ip.GetInfo", err)
		} else {
			info.Location = ipInfo.String()
			info.CountryCode = strings.ToLower(ipInfo.CountryCode)
		}

	}
	log.Println("[INFO]", info)

	return info
}

func (r *Api) initApi(handle windows.Handle) error {
	r.handle = handle
	hSteamApi, info, err := r.findModuleByName(STEAM_API_MODULE_EDITED_NAME)
	if err != nil || hSteamApi == syscall.InvalidHandle {
		hSteamApi, info, err = r.findModuleByName(STEAM_API_MODULE_NAME)
		if err != nil || hSteamApi == syscall.InvalidHandle {
			return err
		}
	}

	r.steamApiInfo = info

	if err := initSteamworks(); err != nil {
		return err
	}

	return nil
}

func (r *Api) processHandler(handle windows.Handle) error {

	if err := initSteamworks(); err != nil {
		return err
	}
	r.fireEventSteamWorksInit()
	//defer steamworks.Shutdown()

	err := r.initApi(handle)
	if err != nil {
		return err
	}

	for {
		newChallenger, err := r.UpdateData()
		if err != nil {
			return err
		}
		if newChallenger {
			r.fireEventNewChallenger(&r.Data)
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func initSteamworks() error {

	_, err := strconv.Atoi(TEKKEN_STEAM_APP_ID)
	if err != nil {
		return fmt.Errorf("invalid app_id: %v", TEKKEN_STEAM_APP_ID)
	}

	// if !steamworks.RestartAppIfNecessary(uint32(appID)) {
	// 	return
	// }
	if !steamworks.Init() {
		return fmt.Errorf("steamworks.Init failed")
	}
	//log.Println("GetCurrentGameLanguage:", steamworks.SteamApps().GetCurrentGameLanguage())
	//steamworks.SteamFriends().ActivateGameOverlayToWebPage("http://localhost:8080/", steamworks.K_EActivateGameOverlayToWebPageMode_Default)

	return nil
}

func (r Api) getUserSteamID() (steamworks.CSteamID, error) {

	data, err := r.readProcessMemory(r.steamApiInfo.BaseOfDll+STEAM_ID_USER_STATIC_POINTER, 8, nil)
	if err != nil {
		return 0, err
	}
	steamID := binary.LittleEndian.Uint64(data)

	return steamworks.CSteamID(steamID), nil
}

func (r Api) getOpponentSteamID() (steamworks.CSteamID, error) {

	data, err := r.readProcessMemory(r.steamApiInfo.BaseOfDll+STEAM_ID_BETTER_STATIC_POINTER, 8, STEAM_ID_BETTER_POINTER_OFFSETS[:])
	if err != nil {
		return 0, err
	}
	steamID := binary.LittleEndian.Uint64(data)

	return steamworks.CSteamID(steamID), nil
}

func (r *Api) UpdateData() (bool, error) {

	r.OldData = r.Data

	var err error
	r.Data.UserSteamID, err = r.getUserSteamID()
	if err != nil {
		return false, err
	}

	r.Data.OpponentSteamID, err = r.getOpponentSteamID()
	if err != nil {
		return false, err
	}

	if r.Data.OpponentSteamID == r.Data.UserSteamID {
		return false, nil
	}

	if strings.Index(fmt.Sprint(r.Data.OpponentSteamID), "765") != 0 {
		return false, nil
	}

	return !reflect.DeepEqual(r.Data, r.OldData), nil
}

func (r Api) OverlayOpenUrl(url string) {
	steamworks.SteamFriends().ActivateGameOverlayToWebPage(url, steamworks.K_EActivateGameOverlayToWebPageMode_Default)
}

func (r Api) GetLanguage() string {
	return steamworks.SteamApps().GetCurrentGameLanguage()
}
