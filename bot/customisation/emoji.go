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
	EmojiId         = NewCustomEmoji("id", 1341144840204718191, false)
	EmojiOpen       = NewCustomEmoji("open", 1341144830704746629, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1341144822110490635, false)
	EmojiClose      = NewCustomEmoji("close", 1341144882173050921, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1341144867408969728, false)
	EmojiReason     = NewCustomEmoji("reason", 1341144769052541029, false)
	EmojiSubject    = NewCustomEmoji("subject", 1341144753118515231, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1341144731526107228, false)
	EmojiClaim      = NewCustomEmoji("claim", 1341144891022905466, false)
	EmojiPanel      = NewCustomEmoji("panel", 1341144810379018250, false)
	EmojiRating     = NewCustomEmoji("rating", 1341144779190173799, false)
	EmojiStaff      = NewCustomEmoji("staff", 1341144761163321354, false)
	EmojiThread     = NewCustomEmoji("thread", 1341144744348352587, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1341144898681962538, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1341144801604534302, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1341144850304729138, false)
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
