// 适配器（游戏机PlaySound()可以像mp3 PlayMusic()一样连接音乐播放器play(player Player)，需要适配器组合，然后构建PlayMusic()方法）
// 适配器模式是在类型不匹配的时候使用,
// 目的是将一种类型伪装成另一种类型以便我们的代码可以正常使用;而装饰者模式装饰者和被装饰者拥有相同的类型(相同的超类),目的是为了增强功能或者方便使用.
package main

import "fmt"

type Player interface {
	PlayMusic()
}

type MusicPlayer struct {
	Src string
}

func (p MusicPlayer) PlayMusic() {
	fmt.Println("play music: " + p.Src)
}

type GameSoundPlayer struct {
	Src string
}

func (p GameSoundPlayer) PlaySound() {
	fmt.Println("play sound: " + p.Src)
}

type GameSoundAdapter struct {
	SoundPlayer GameSoundPlayer
}

func (p GameSoundAdapter) PlayMusic() {
	p.SoundPlayer.PlaySound()
}

func main() {
	gameSound := GameSoundPlayer{Src: "game.mid"}
	gameAdapter := GameSoundAdapter{SoundPlayer: gameSound}
	play(gameAdapter)
}

func play(player Player) {
	player.PlayMusic()
}
