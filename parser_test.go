// Copyright 2015 Yury Kozyrev. License MIT.
// Proprietary license.
package emoji

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"strconv"
	"regexp"
)

func TestInsert(t *testing.T) {
	container  := make(map[string]string)
	parser := NewEmojiParser()

	var text = "a #💩 #and #🍦 #😳"
	var i = -1
	replased := parser.ReplaceAllStringFunc(text, func(s string) string {
		i++
		key := "_$"+strconv.Itoa(i)+"_"
		container[key] = s
		return key
		return strconv.Itoa(i)
	})
	assert.Equal(t, replased, "a #_$0_ #and #_$1_ #_$2_")

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

	recovered := regexp.MustCompile(`\_\$\d+\_`).ReplaceAllStringFunc(replased, func(s string) string {
		return container[s]
	})
	assert.Equal(t, recovered, text)
}
