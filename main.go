package main

import (
	"crypto/tls"
	"log"
	"regexp"

	"github.com/thoj/go-ircevent"
)

func main() {
	CHAN := "#swehack"
	SERVER := "irc.swehack.org:6697"
	NICK := "placebot"

	i := irc.IRC(NICK, NICK)
	i.Version = "666"
	i.UseTLS = true
	i.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := i.Connect(SERVER)
	if err != nil {
		log.Fatal(err)
	}

	i.AddCallback("PRIVMSG", func(event *irc.Event) {
		if event.Nick == i.GetNick() {
			return
		}

		re := regexp.MustCompile("!gif ([A-z0-9 \\.\\,\\-]+)")
		gifMatch := re.FindStringSubmatch(event.Message())

		if len(gifMatch) != 0 {
			r, err := GetGif(gifMatch[1])
			if err != nil {
				return
			}
			if r != "" {
				i.Privmsg(CHAN, r)
			}
			return
		}

		urlreg := regexp.MustCompile("(?i)(https?:\\/\\/)([a-z\\-0-9\\._]+\\.[a-z]{2,3}[a-z0-9%\\-\\._?#&\\/=]*)")
		urlMatch := urlreg.FindStringSubmatch(event.Message())
		if len(urlMatch) != 0 {
			clientPtr := prepareProxyClient()
			title, err := httpGetTitle(clientPtr, urlMatch[0])
			if err != nil || title == "" {
				return
			}

			i.Privmsg(CHAN, "[Title] "+title)
			return
		}
		if event.Message() == ".inflik" {
			i.Privmsg(event.Arguments[0], "Jag skulle bara vilja inflika för ett ögonblick.")
			i.Privmsg(event.Arguments[0], "Vad du kallar Linux är faktiskt GNU/Linux, eller som jag själv kallar det, GNU+Linux. Linux är inte ett operativsystem i sig själv, utan snarare ännu en del utav ett funktionellt GNU-system, som görs användbart av GNU corelibs, shell-utils, och andra nödvändiga delar, som tillsammans definerar ett OS enligt POSIX.")
			i.Privmsg(event.Arguments[0], "Många datorer kör ett modifierat GNU-system varje dag, utan att inse det. Genom en lustig härva av händelser kallas det GNU som används ofta Linux, och många av dess användare inser inte att de använder GNU-systemet, som utvecklats utav GNU-projektet.")
			i.Privmsg(event.Arguments[0], "Det finns ett Linux, och dessa människor använder det, men det är bara en del av systemet de använder. Linux är kerneln, programmet i systemet som allokerar maskinens resurser till de andra programmen du kör. Kerneln är en viktig del utav ett operativsystem, men helt oanvändbart i sig själv; den kan bara fungera i samband med ett helt operativsystem.")
			i.Privmsg(event.Arguments[0], "Linux används oftast i samband med GNU-operativsystemet: hela systemet är bara GNU med Linux tillagt, eller GNU/Linux. Alla så kallade Linux-distrubitioner är egentligen distrubutioner utav GNU/Linux!")
		}
		if event.Message() == ".interject" {
			i.Privmsg(event.Arguments[0], "I'd just like to interject for moment. What you're refering to as Linux, is in fact, GNU/Linux, or as I've recently taken to calling it, GNU plus Linux. Linux is not an operating system unto itself, but rather another free component of a fully functioning GNU system made useful by the GNU corelibs, shell utilities and vital system components comprising a full OS as defined by POSIX.")
			i.Privmsg(event.Arguments[0], "Many computer users run a modified version of the GNU system every day, without realizing it. Through a peculiar turn of events, the version of GNU which is widely used today is often called Linux, and many of its users are not aware that it is basically the GNU system, developed by the GNU Project. ")
			i.Privmsg(event.Arguments[0], "There really is a Linux, and these people are using it, but it is just a part of the system they usevent. Linux is the kernel: the program in the system that allocates the machine's resources to the other programs that you run. ")
			i.Privmsg(event.Arguments[0], "The kernel is an essential part of an operating system, but useless by itself; it can only function in the context of a complete operating system. Linux is normally used in combination with the GNU operating system: the whole system is basically GNU with Linux added, or GNU/Linux. All the so-called Linux distributions are really distributions of GNU/Linux! ")
		}
		if event.Message() == "wow #rude" {
			i.Privmsg(event.Arguments[0], "I feel offended by your recent action(s). Please read http://stop-irc-bullying.eu/stop")
		}
	})

	i.Debug = true
	i.Loop()
}
