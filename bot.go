package main

import (
	"crypto/tls"
	"github.com/thoj/go-ircevent"
)

func main() {

	conn := irc.IRC("Interjectbot", "inflik")
	conn.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	conn.UseTLS = true
	conn.Connect("irc.swehack.org:6697")
	conn.AddCallback("001", func(e *irc.Event) {
		conn.Join("#swehack")
	})

	conn.AddCallback("PRIVMSG", func(e *irc.Event) {
		if e.Message() == ".inflik" {
			conn.Privmsg(e.Arguments[0], "Jag skulle bara vilja inflika för ett ögonblick.")
			conn.Privmsg(e.Arguments[0], "Vad du kallar Linux är faktiskt GNU/Linux, eller som jag själv kallar det, GNU+Linux. Linux är inte ett operativsystem i sig själv, utan snarare ännu en del utav ett funktionellt GNU-system, som görs användbart av GNU corelibs, shell-utils, och andra nödvändiga delar, som tillsammans definerar ett OS enligt POSIX.")
			conn.Privmsg(e.Arguments[0], "Många datorer kör ett modifierat GNU-system varje dag, utan att inse det. Genom en lustig härva av händelser kallas det GNU som används ofta Linux, och många av dess användare inser inte att de använder GNU-systemet, som utvecklats utav GNU-projektet.")
			conn.Privmsg(e.Arguments[0], "Det finns ett Linux, och dessa människor använder det, men det är bara en del av systemet de använder. Linux är kerneln, programmet i systemet som allokerar maskinens resurser till de andra programmen du kör. Kerneln är en viktig del utav ett operativsystem, men helt oanvändbart i sig själv; den kan bara fungera i samband med ett helt operativsystem.")
			conn.Privmsg(e.Arguments[0], "Linux används oftast i samband med GNU-operativsystemet: hela systemet är bara GNU med Linux tillagt, eller GNU/Linux. Alla så kallade Linux-distrubitioner är egentligen distrubutioner utav GNU/Linux!")
		}
		if e.Message() == ".interject" {
			conn.Privmsg(e.Arguments[0], "I'd just like to interject for moment. What you're refering to as Linux, is in fact, GNU/Linux, or as I've recently taken to calling it, GNU plus Linux. Linux is not an operating system unto itself, but rather another free component of a fully functioning GNU system made useful by the GNU corelibs, shell utilities and vital system components comprising a full OS as defined by POSIX.")
			conn.Privmsg(e.Arguments[0], "Many computer users run a modified version of the GNU system every day, without realizing it. Through a peculiar turn of events, the version of GNU which is widely used today is often called Linux, and many of its users are not aware that it is basically the GNU system, developed by the GNU Project. ")
			conn.Privmsg(e.Arguments[0], "There really is a Linux, and these people are using it, but it is just a part of the system they use. Linux is the kernel: the program in the system that allocates the machine's resources to the other programs that you run. ")
			conn.Privmsg(e.Arguments[0], "The kernel is an essential part of an operating system, but useless by itself; it can only function in the context of a complete operating system. Linux is normally used in combination with the GNU operating system: the whole system is basically GNU with Linux added, or GNU/Linux. All the so-called Linux distributions are really distributions of GNU/Linux! ")
		}
	})
	conn.Loop()
}
