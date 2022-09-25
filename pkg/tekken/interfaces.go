package tekken

type Hook interface {
	EventSteamWorksInit()
	EventNewChallenger(data PlayerInfo)
}