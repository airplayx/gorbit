package emoji

import "testing"

func TestEmojiEncode(t *testing.T) {
	t.Parallel()
	t.Log(Encode(""))
	t.Log(Encode("❤"))
	t.Log(Encode("💝😄🛅"))
}

func TestEmojiDecode(t *testing.T) {
	t.Parallel()
	t.Log(Decode(""))
	t.Log(Decode("❤"))
	t.Log(Decode(Encode("💝😄🛅")))
}
