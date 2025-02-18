package customisation

import (
	"fmt"

	"github.com/rxdn/gdl/objects"
	"github.com/rxdn/gdl/objects/guild/emoji"
)

type CustomEmoji struct {
	Name     string
	Id       uint64
	Animated bool
}

func NewCustomEmoji(name string, id uint64, animated bool) CustomEmoji {
	return CustomEmoji{
		Name: name,
		Id:   id,
	}
}

func (e CustomEmoji) String() string {
	if e.Animated {
		return fmt.Sprintf("<a:%s:%d>", e.Name, e.Id)
	} else {
		return fmt.Sprintf("<:%s:%d>", e.Name, e.Id)
	}
}

func (e CustomEmoji) BuildEmoji() *emoji.Emoji {
	return &emoji.Emoji{
		Id:       objects.NewNullableSnowflake(e.Id),
		Name:     e.Name,
		Animated: e.Animated,
	}
}

var (
	EmojiId         = NewCustomEmoji("id", 1341456261601951764, false)
	EmojiOpen       = NewCustomEmoji("open", 1341456248603934803, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1341456236570349620, false)
	EmojiClose      = NewCustomEmoji("close", 1341456291490435112, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1341456282409762898, false)
	EmojiReason     = NewCustomEmoji("reason", 1341456165896192001, false)
	EmojiSubject    = NewCustomEmoji("subject", 1341456147391189065, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1341456125471494275, false)
	EmojiClaim      = NewCustomEmoji("claim", 1341456291490435112, false)
	EmojiPanel      = NewCustomEmoji("panel", 1341456200113328198, false)
	EmojiRating     = NewCustomEmoji("rating", 1341456177791373362, false)
	EmojiStaff      = NewCustomEmoji("staff", 1341456155972730922, false)
	EmojiThread     = NewCustomEmoji("thread", 1341456136682868938, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1341456309995962390, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1341456188755410974, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1341456271013838898, false)
	//EmojiTime       = NewCustomEmoji("time", 974006684622159952, false)
)

// PrefixWithEmoji Useful for whitelabel bots
func PrefixWithEmoji(s string, emoji CustomEmoji, includeEmoji bool) string {
	if includeEmoji {
		return fmt.Sprintf("%s %s", emoji, s)
	} else {
		return s
	}
}
