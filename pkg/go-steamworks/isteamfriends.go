package steamworks

import (
	"runtime"
	"unsafe"
)

type ISteamFriends interface {
	ActivateGameOverlayToWebPage(url string, eMod EActivateGameOverlayToWebPageMode)
	GetPersonaName() string
	GetPersonaState() EPersonaState
	GetFriendPersonaName(steamIDFriend CSteamID) string
	GetPlayerNickname(steamIDPlayer CSteamID) string
	GetFriendPersonaState(steamIDFriend CSteamID) EPersonaState
}

const (
	flatAPI_SteamFriends                               = "SteamAPI_SteamFriends_v017"
	flatAPI_ISteamFriends_ActivateGameOverlayToWebPage = "SteamAPI_ISteamFriends_ActivateGameOverlayToWebPage"
	flatAPI_ISteamFriends_GetPersonaName               = "SteamAPI_ISteamFriends_GetPersonaName"
	flatAPI_ISteamFriends_GetPersonaState              = "SteamAPI_ISteamFriends_GetPersonaState"
	flatAPI_ISteamFriends_GetFriendPersonaName         = "SteamAPI_ISteamFriends_GetFriendPersonaName"
	flatAPI_ISteamFriends_GetFriendPersonaState        = "SteamAPI_ISteamFriends_GetFriendPersonaState"
	flatAPI_ISteamFriends_GetPlayerNickname            = "SteamAPI_ISteamFriends_GetPlayerNickname"
)

type steamFriends uintptr

func SteamFriends() ISteamFriends {
	v, err := theDLL.call(flatAPI_SteamFriends)
	if err != nil {
		panic(err)
	}
	return steamFriends(v)
}

func (s steamFriends) ActivateGameOverlayToWebPage(url string, eMod EActivateGameOverlayToWebPageMode) {
	cUrl := append([]byte(url), 0)
	defer runtime.KeepAlive(cUrl)

	_, err := theDLL.call(flatAPI_ISteamFriends_ActivateGameOverlayToWebPage,
		uintptr(s), uintptr(unsafe.Pointer(&cUrl[0])), uintptr(eMod))
	if err != nil {
		panic(err)
	}
}

func (s steamFriends) GetPersonaName() string {
	v, err := theDLL.call(flatAPI_ISteamFriends_GetPersonaName, uintptr(s))
	if err != nil {
		panic(err)
	}

	bs := make([]byte, 0, 256)
	for i := int32(0); ; i++ {
		b := *(*byte)(unsafe.Pointer(v))
		v += unsafe.Sizeof(byte(0))
		if b == 0 {
			break
		}
		bs = append(bs, b)
	}
	return string(bs)
}

func (s steamFriends) GetPersonaState() EPersonaState {
	v, err := theDLL.call(flatAPI_ISteamFriends_GetPersonaState, uintptr(s))
	if err != nil {
		panic(err)
	}

	return EPersonaState(v)
}

func (s steamFriends) GetFriendPersonaName(steamIDFriend CSteamID) string {
	v, err := theDLL.call(flatAPI_ISteamFriends_GetFriendPersonaName, uintptr(s), uintptr(steamIDFriend))
	if err != nil {
		panic(err)
	}

	bs := make([]byte, 0, 256)
	for i := int32(0); ; i++ {
		b := *(*byte)(unsafe.Pointer(v))
		v += unsafe.Sizeof(byte(0))
		if b == 0 {
			break
		}
		bs = append(bs, b)
	}
	return string(bs)
}

func (s steamFriends) GetPlayerNickname(steamIDPlayer CSteamID) string {
	v, err := theDLL.call(flatAPI_ISteamFriends_GetPlayerNickname, uintptr(s), uintptr(steamIDPlayer))
	if err != nil {
		panic(err)
	}

	if v == 0x0 {
		return ""
	}

	bs := make([]byte, 0, 256)
	for i := int32(0); ; i++ {
		b := *(*byte)(unsafe.Pointer(v))
		v += unsafe.Sizeof(byte(0))
		if b == 0 {
			break
		}
		bs = append(bs, b)
	}
	return string(bs)
}

func (s steamFriends) GetFriendPersonaState(steamIDFriend CSteamID) EPersonaState {
	v, err := theDLL.call(flatAPI_ISteamFriends_GetFriendPersonaState, uintptr(s), uintptr(steamIDFriend))
	if err != nil {
		panic(err)
	}

	return EPersonaState(v)
}
