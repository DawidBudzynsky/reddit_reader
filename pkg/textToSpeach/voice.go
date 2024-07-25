package texttospeach

type Voice string

const (
	// DISNEY VOICES
	VoiceGhostFace    Voice = "en_us_ghostface"    // Ghost Face
	VoiceChewbacca    Voice = "en_us_chewbacca"    // Chewbacca
	VoiceC3PO         Voice = "en_us_c3po"         // C3PO
	VoiceStitch       Voice = "en_us_stitch"       // Stitch
	VoiceStormtrooper Voice = "en_us_stormtrooper" // Stormtrooper
	VoiceRocket       Voice = "en_us_rocket"       // Rocket
	// ENGLISH VOICES
	VoiceEnglishAU1 Voice = "en_au_001" // English AU - Female
	VoiceEnglishAU2 Voice = "en_au_002" // English AU - Male
	VoiceEnglishUK1 Voice = "en_uk_001" // English UK - Male 1
	VoiceEnglishUK2 Voice = "en_uk_003" // English UK - Male 2
	VoiceEnglishUS1 Voice = "en_us_001" // English US - Female (Int. 1)
	VoiceEnglishUS2 Voice = "en_us_002" // English US - Female (Int. 2)
	VoiceEnglishUS3 Voice = "en_us_006" // English US - Male 1
	VoiceEnglishUS4 Voice = "en_us_007" // English US - Male 2
	VoiceEnglishUS5 Voice = "en_us_009" // English US - Male 3
	VoiceEnglishUS6 Voice = "en_us_010" // English US - Male 4
	// EUROPE VOICES
	VoiceFrench1 Voice = "fr_001" // French - Male 1
	VoiceFrench2 Voice = "fr_002" // French - Male 2
	VoiceGermanF Voice = "de_001" // German - Female
	VoiceGermanM Voice = "de_002" // German - Male
	VoiceSpanish Voice = "es_002" // Spanish - Male
	// AMERICA VOICES
	VoiceSpanishMX    Voice = "es_mx_002" // Spanish MX - Male
	VoicePortugueseF1 Voice = "br_001"    // Portuguese BR - Female 1
	VoicePortugueseF2 Voice = "br_003"    // Portuguese BR - Female 2
	VoicePortugueseF3 Voice = "br_004"    // Portuguese BR - Female 3
	VoicePortugueseM  Voice = "br_005"    // Portuguese BR - Male
	// ASIA VOICES
	VoiceIndonesianF Voice = "id_001" // Indonesian - Female
	VoiceJapaneseF1  Voice = "jp_001" // Japanese - Female 1
	VoiceJapaneseF2  Voice = "jp_003" // Japanese - Female 2
	VoiceJapaneseF3  Voice = "jp_005" // Japanese - Female 3
	VoiceJapaneseM   Voice = "jp_006" // Japanese - Male
	VoiceKoreanM1    Voice = "kr_002" // Korean - Male 1
	VoiceKoreanF     Voice = "kr_003" // Korean - Female
	VoiceKoreanM2    Voice = "kr_004" // Korean - Male 2
	// SINGING VOICES
	VoiceAlto         Voice = "en_female_f08_salut_damour" // Alto
	VoiceTenor        Voice = "en_male_m03_lobby"          // Tenor
	VoiceWarmyBreeze  Voice = "en_female_f08_warmy_breeze" // Warmy Breeze
	VoiceSunshineSoon Voice = "en_male_m03_sunshine_soon"  // Sunshine Soon
	// OTHER
	VoiceNarration Voice = "en_male_narration"   // Narrator
	VoiceWacky     Voice = "en_male_funny"       // Wacky
	VoicePeaceful  Voice = "en_female_emotional" // Peaceful
)
