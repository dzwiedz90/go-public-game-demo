package core

import (
	"github.com/dzwiedz90/go-public-game-demo/models/character"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Character = character.Character{}
)

func LoadMusic(stream string) rl.Music {
	rl.InitAudioDevice()
	musicStream := rl.LoadMusicStream(stream)
	rl.PlayMusicStream(musicStream)

	return musicStream
}

func UnloadMusic(stream rl.Music) {
	rl.UnloadMusicStream(stream)
	rl.CloseAudioDevice()
}
