package main

import (
	"os"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"encoding/binary"
	"io"
	"strings"
	"time"
	"math/rand"
)

var JohnCena = make([][]byte, 0);
var Elevator = make([][]byte, 0);
var Rickroll = make([][]byte, 0);
var Letter = make([][]byte, 0);
var Cri = make([][]byte, 0);
var NumberHat = make([][]byte, 0);
var ExoticButters = make([][]byte, 0);
var DamnSon = make([][]byte, 0);
var Jeff = make([][]byte, 0);
var Nigga = make([][]byte, 0);
var RussianSinger = make([][]byte, 0);
var SadViolin = make([][]byte, 0);
var ShutUp = make([][]byte, 0);
var Triple = make([][]byte, 0);
var TurTur = make([][]byte, 0);
var Weed = make([][]byte, 0);
var XFiles = make([][]byte, 0);

var statuses = []string{
	"hidden object games",
	"Oh... Sir!",
	"Minecraft 1.0 ALPHA",
	"with your mother",
	"something",
	"something else",
	"bored",
	"dead"}

type Settings struct{
	vc *discordgo.VoiceConnection
	playing bool
	commander string
}
var settings = make(map[string]*Settings);

func main(){
	args := os.Args[1:];

	if(len(args) < 1){
		fmt.Println("No token provided!");
		return;
	}
	token := args[0];

	load("John Cena", &JohnCena);
	load("Elevator", &Elevator);
	load("Rickroll", &Rickroll);
	load("Cri", &Cri);
	load("Letter", &Letter);
	load("NumberHat", &NumberHat);
	load("ExoticButters", &ExoticButters);
	load("damnson", &DamnSon);
	load("jeff", &Jeff);
	load("nigga", &Nigga);
	load("russianSinger", &RussianSinger);
	load("sadviolin", &SadViolin);
	load("shutup", &ShutUp);
	load("triple", &Triple);
	load("turtur", &TurTur);
	load("weed", &Weed);
	load("xfiles", &XFiles);

	d, _ := discordgo.New("Bot " + token);
	d.AddHandler(ready);
	d.AddHandler(messageCreate);
	d.Open();

	<-make(chan struct{});
}

func load(file string, buffer *[][]byte){
	f, err := os.Open("Dank/" + file + ".dca");
	if err != nil { fmt.Println("File not found: " + file); return; }

	var length int16;
	for {
		err := binary.Read(f, binary.LittleEndian, &length);

		if err == io.EOF || err == io.ErrUnexpectedEOF{
			return;
		}

		buf := make([]byte, length);
		binary.Read(f, binary.LittleEndian, &buf);

		*buffer = append(*buffer, buf);
	}
}

func play(buffer [][]byte, session *discordgo.Session, guild, channel string, s *Settings){
	s.playing = true;
	s.vc, _ = session.ChannelVoiceJoin(guild, channel, false, true);

	s.vc.Speaking(true);

	for _, buf := range buffer{
		if s.vc == nil { return; }
		s.vc.OpusSend <- buf;
	}

	s.vc.Speaking(false);
	s.vc.Disconnect();
	s.playing = false;
}

func messageCreate(session *discordgo.Session, event *discordgo.MessageCreate){
	msg := strings.ToLower(strings.TrimSpace(event.Content));
	author := event.Author;

	channel, _ := session.State.Channel(event.ChannelID);
	guild, _ := session.State.Guild(channel.GuildID);
	//member, _ := session.State.Member(guild.ID, author.ID);

	s := settings[guild.ID];
	if s == nil{
		s = &Settings{};
		settings[guild.ID] = s;
	}
	fmt.Println(*s);

	if(s.commander != "" && s.commander != author.ID){
		return;
	}

	var buffer [][]byte = nil;
	switch msg {
		case "john cena":
			buffer = JohnCena;
		case "waiting":
			buffer = Elevator;
		case "rickroll":
			buffer = Rickroll;
		case "cri":
			buffer = Cri;
		case "letter":
			buffer = Letter;
		case "numbr hat":
			buffer = NumberHat;
		case "exotic butters":
			buffer = ExoticButters;
		case "damn son":
			buffer = DamnSon;
		case "jeff":
			buffer = Jeff;
		case "nigga":
			buffer = Nigga;
		case "russian singer":
			buffer = RussianSinger;
		case "sad violin":
			buffer = SadViolin;
		case "shut up":
			buffer = ShutUp;
		case "triple":
			buffer = Triple;
		case "turtur":
			buffer = TurTur;
		case "weed":
			buffer = Weed;
		case "illuminati":
			buffer = XFiles;
		case "thx":
			if s.vc != nil{
				s.vc.Speaking(false);
				s.vc.Disconnect();
				s.playing = false;
			}
		case "listen only to me plz":
			s.commander = author.ID;
		case "every1 owns u stopad robot":
			s.commander = "";
		case "cler da chat plz":
			messages, _ := session.ChannelMessages(event.ChannelID, 100, "", "");
			ids := make([]string, 0);
			for _, message := range messages{
				ids = append(ids, message.ID);
			}
			session.ChannelMessagesBulkDelete(event.ChannelID, ids);
	}

	if buffer != nil && !s.playing{
		for _, state := range guild.VoiceStates{
			if state.UserID == event.Author.ID{
				play(buffer, session, guild.ID, state.ChannelID, s);
			}
		}
	}
}

func ready(session *discordgo.Session, event *discordgo.Ready){
	ticker := time.NewTicker(time.Second * 5);
	go func(){
		for{
			<- ticker.C;
			session.UpdateStatus(0, statuses[rand.Intn(len(statuses))]);
		}
	}();
}
