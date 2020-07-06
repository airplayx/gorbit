package gorbit

import "testing"

func TestEmojiEncode(t *testing.T) {
	t.Parallel()
	t.Log(EmojiEncode(""))
	t.Log(EmojiEncode("❤"))
	t.Log(EmojiEncode("💝😄🛅"))
}

func TestEmojiDecode(t *testing.T) {
	t.Parallel()
	t.Log(EmojiDecode(""))
	t.Log(EmojiDecode("❤"))
	t.Log(EmojiDecode(EmojiEncode("💝😄🛅")))
}
