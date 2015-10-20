// Copyright 2015 Yury Kozyrev. License MIT.
// Proprietary license.
package emoji

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"strconv"
)

func TestInsert(t *testing.T) {
	parser := NewEmojiParser()

	var text = "a #💩 #and #🍦 #😳"
	var i = -1
	replased := parser.ReplaceAllStringFunc(text, func(s string) string {
		i++
		return strconv.Itoa(i)
	})
	assert.Equal(t, replased, "a #0 #and #1 #2")

	htmlEnt := parser.ToHtmlEntities(text)

	assert.Equal(t, htmlEnt, "a #&#x1F4A9; #and #&#x1F366; #&#x1F633;")

	htmlImg := parser.ToHtmlImages(text)

	assertion := `a #<img
class="emoji"
draggable="false"
alt="💩"
src="https://twemoji.maxcdn.com/36x36/1f4a9.png"> #and #<img
class="emoji"
draggable="false"
alt="🍦"
src="https://twemoji.maxcdn.com/36x36/1f366.png"> #<img
class="emoji"
draggable="false"
alt="😳"
src="https://twemoji.maxcdn.com/36x36/1f633.png">`
	assert.Equal(t, htmlImg, assertion)
}
