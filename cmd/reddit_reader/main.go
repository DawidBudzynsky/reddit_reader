package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/DawidBudzynsky/reddit_reader/pkg/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env file")
	}

	// textToRead := `"Anna, can you room the next patient?" I asked, glancing up from my notes. "Sure, she'll be in Room 5," she replied, heading down the hallway. I quickly reviewed the patient's history that Anna had written up. The lady had been complaining of a haze that seemed to come and go and occasional stabbing pains in her left eye that began several days ago. Nothing way out of the ordinary, really. It was probably just a combination of floaters and dry eye. In my mind, the diagnosis was already locked and loaded, but just out of precaution, the patient was dilated to make sure there wasn’t anything going on with her retina. Satisfied with my guess, I took off my readers, picked up my laptop, and entered Room 5. "Mrs. Myers, it’s good to meet you. I’m Dr. Ronald." "It’s nice to meet you too," she replied, her voice tense. Mrs. Myers was a young lady, her hands clutching her handbag in the exam chair, her eyes hidden behind a pair of sunglasses far too large for her face. "Sorry for the sunglasses, doc. Since yesterday, my eyes have been really sensitive to light." Faint alarm bells started to ring in my head, but I pushed the thought aside for now. "Oh, it’s no problem. Why don’t you tell me a little more about what’s been going on with your eyes?" I asked, trying to sound reassuring. "Well, it’s just my left eye, really. Since Saturday, I feel like my vision's been getting worse, like a fog's been building up. And I get these shooting pains from time to time. It almost feels like someone is poking my eye from the back. I had lunch with a friend on Friday who had similar symptoms at the time. I was wondering if it’s a weird form of pink eye or something?" "Could be. How about let’s take those sunglasses off, and we'll get a look at what’s going on with your eye." "Oh, of course," she said, removing the oversized sunglasses. Under the slit lamp, her eyes seemed normal enough, with the exception of her left eye being somewhat bloodshot. With no clear explanation for her symptoms, I took out my fundus lens and peered into the back of her eye. I had caught no more than a glimpse of her retina before she shrieked and slapped the lens, knocking it rattling onto the floor. But even without that abrupt interruption, I would have dropped the lens in shock. In that brief, horrifying instant, I saw something that defied all logic—her optic nerve was a pitch-black void, an abyss that seemed to swallow light itself. The optic nerve usually looks like a donut, with a dark orange outer ring surrounding a brighter inner circle. The inside of Mrs. Myers' optic nerve was pure darkness. A perfect little black circle, almost like a miniature black hole had manifested in the back of her eye, absorbing all light passing into it. As my mind spun in a whirlwind of confusion, Mrs. Myers tumbled from her chair, collapsing onto the floor with a shriek that pierced the air. She clutched her hand over her eye, her voice rising in a crescendo of panic. Just then, Anna burst into the room, snapping me out of my stupor. We rushed to her side, dropping to our knees to help Mrs. Myers. She was inconsolable, moaning and rocking back and forth, her distress palpable. "Oh god, oh god, oh god," she wailed, her voice cracking with terror. "There's SOMETHING MOVING IN MY EYE!" As her screaming reached an unbearable climax, Mrs. Myers suddenly looked up from her prostrated position. Her hand was still cupped firmly over her eye, but now we could see the blood trickling down between her fingers, a crimson stream staining her pale cheek. My heart pounded in my chest as I watched in horror. Something was moving beneath her skin, forcing her fingers apart from the inside. The sight was grotesque, a nightmarish struggle playing out just beneath her flesh. With a final, blood-curdling scream, Mrs. Myers collapsed to the floor, her body convulsing violently. Her arms fell limply to her sides, revealing the true extent of the horror. Jet-black, spider-like appendages erupted from her eye, each one glistening with a wet, otherworldly sheen. They waved and writhed, searching for a grip on the slick linoleum tiles. Each movement left behind bloody scratches on Mrs. Myers' face, the lines of red contrasting starkly with her ashen skin`
	//
	// textToSpeach := texttospeach.NewTikTokTTS(texttospeach.VoiceEnglishUK2, "output.mp3")
	// textToSpeach.TextToMp3(textToRead)

	session, err := discordgo.New(os.Getenv("DISCORD_BOT_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	commandHandler := commands.NewCommandHandler("!", session)
	commandHandler.AddCommand("!hello", &commands.HelloWorldHandler{})
	commandHandler.AddCommand("!read", &commands.ReadHandler{})

	session.AddHandler(commandHandler.HandleCommand)

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGABRT, os.Interrupt)
	<-sc
}
