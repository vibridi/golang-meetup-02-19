package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	varcharColumnLen = 500

	// len = 571
	longTextASCII = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
`
	// len = 624
	longTextUnicode = `
竹販東初強念元浅注校働牙核丸下。因復速億展百中販感験東金学万像呼。調聞急年性新室作河供応点質食済重摘。済高市球掲機山就批員五宅暮社禁泰受袋。会五太同護脱政先動済百喪載学世試億岡広。自発囲違立判水斎内計栄面回声反田。怪本氏楽比題性次光計見強。含写東備市九術覧位将王支属球阪勝祝。却者宣調明物伊学帯馬度分使娘度作変迎前意。
動宅況学元瀬病名能興投自止。消軽上肩柱場別視自入職計所。算用変限書究武売淀年制体立議。増広故上童年絵校転基不救房喜宮逃当。対合迎口式梨真新住容特心巡注況法情旅代。台会操前管断編費先葬訓置。合必番打券示供止性就報次覚。陸民司車更対誌告明民瀬全。座歌先心主万論載変江関家賞在。三立著省治民取受済村公目験。
将顧目道化住正集平寄東特。共好極処朝然界京石得連報連。容著集所事文平汁暮真理奪氷職壊年取者正。著内席税情麻重碁護伊塚展政重権。案京毎亡奏省以誘継漂校図禁各功気加。権事区献感商住右辺通各質共敗意小。応訴条志拳済争輪公空漂女。藤大博観立示命最所騎烈十語席旧。毎市別恵観約稿融訃拍月転知城銀。掲降選多能勉最忘体大司板団能生前。
文動利作将輩東福駒男面罪大速不。会平民知決軍要朝知査局戦記堀朝平。学少中代案言真年掲生込場政試。経高語権番育車刊均身写言。若必音併今長択待談時止託央田惑果人樫氷。自弾島給索同姿所界臣能示読輔続。嘉高雪崩欧運暮村指棋移通特出格切国観野。問連出江現真夜味基覚人藤徒方寄東断文婦。覚葉覧備有限択河真育障号。
`
	// len = 501
	longTextWithFinalUnicode = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Lorem ipsum dolor sit amet, consectetur adipis 强大
`
)

func TestElide(t *testing.T) {
	t.Run("elides long ASCII text", func(t *testing.T) {
		assert.True(t, len(longTextASCII) > FreeTextMaxLength)

		elided := Elide(longTextASCII)
		assert.Len(t, elided, varcharColumnLen)
		assert.True(t, strings.HasSuffix(elided, "..."))
		assert.Equal(t, elided[FreeTextMaxLength:varcharColumnLen], "...")
	})

	t.Run("elides long unicode text", func(t *testing.T) {
		// 624 runes, 1862 bytes
		assert.True(t, len(longTextUnicode) > FreeTextMaxLength)

		elided := Elide(longTextUnicode)
		assert.Len(t, elided, varcharColumnLen)
		assert.True(t, strings.HasSuffix(elided, "..."))
		assert.Equal(t, elided[FreeTextMaxLength:varcharColumnLen], "...")
	})

	t.Run("does not corrupt unicode points", func(t *testing.T) {
		// 497 runes, 501 bytes
		assert.True(t, len(longTextWithFinalUnicode) > FreeTextMaxLength)

		elided := Elide(longTextWithFinalUnicode)
		assert.Len(t, elided, varcharColumnLen)
		assert.True(t, strings.HasSuffix(elided, "强..."))
		assert.Equal(t, elided[FreeTextMaxLength:varcharColumnLen], "...")
	})
}