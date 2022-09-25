package steamworks

import "unsafe"

type ISteamNetworking interface {
	GetP2PSessionState(steamIDRemote CSteamID, pConnectionState *P2PSessionState_t) bool
}

const (
	flatAPI_SteamNetworking                     = "SteamAPI_SteamNetworking_v006"
	flatAPI_ISteamNetworking_GetP2PSessionState = "SteamAPI_ISteamNetworking_GetP2PSessionState"
)

type steamNetworking uintptr

func SteamNetworking() ISteamNetworking {
	v, err := theDLL.call(flatAPI_SteamNetworking)
	if err != nil {
		panic(err)
	}
	return steamNetworking(v)
}

func (s steamNetworking) GetP2PSessionState(steamIDRemote CSteamID, pConnectionState *P2PSessionState_t) bool {
	v, err := theDLL.call(flatAPI_ISteamNetworking_GetP2PSessionState,
		uintptr(s), uintptr(steamIDRemote), uintptr(unsafe.Pointer(pConnectionState)))
	if err != nil {
		panic(err)
	}

	return byte(v) != 0
}
