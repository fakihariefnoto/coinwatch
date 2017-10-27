package notifier

const (
	inisial int = iota
	SMILE
	SAD
	CRY
	HAPPY
	INFO
	COOL
	CRITICAL
	ANGRY
	ANGEL
	DEVIL
	SHY
	SMART
	SICK
	SMIRK
	WINK
	WORRIED
	LAUGH
	SURPRISE
)

var (
	messageIcon map[int]string
)

func init() {
	messageIcon = map[int]string{
		SMILE:    "face-smile",
		ANGRY:    "face-angry",
		ANGEL:    "face-angel",
		COOL:     "face-cool",
		SAD:      "face-sad",
		CRY:      "face-crying",
		HAPPY:    "face-kiss",
		INFO:     "face-uncertain",
		CRITICAL: "face-tired",
		DEVIL:    "face-devilish",
		SHY:      "face-embarrassed",
		SMART:    "face-glasses",
		SICK:     "face-sick",
		SMIRK:    "face-smirk",
		WINK:     "face-wink",
		WORRIED:  "face-worried",
		LAUGH:    "face-laugh",
		SURPRISE: "face-surprise",
	}
}
