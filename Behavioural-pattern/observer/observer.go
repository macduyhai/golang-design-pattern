package observer

import "log"

// Tạo mối liên hệ one to many giữa các subject và các Observer với nhau(Vd 1 subject sẽ có thuộc
// tính là mọt mảng nhiều observer) nên khi trạng thái của object thay đổi, tất cả các observer
// liên kết với subject này sẽ được thông báo và tự động cập nhật

type Observer interface {
	update(interface{})
}
type Subject interface {
	RegisterObserver(obs Observer)
	RemoveObserver(obs Observer)
	NotifyObservers()
}
type Video struct {
	Title string
}

// YoutubeChannel is a concrete implementation of Subject interface
type YoutubeChannel struct {
	Observers []Observer
	NewVideo  *Video
}

func (yt *YoutubeChannel) RegisterObserver(obs Observer) {
	yt.Observers = append(yt.Observers, obs)
}

func (yt *YoutubeChannel) RemoveObserver(obs Observer) {
	index := 0
	listObs := yt.Observers
	for _, o := range listObs {
		if o != obs {
			listObs[index] = o
			index++
		}
	}
	yt.Observers = listObs[:index]
}

// notify to all observers when a new video is released
func (yt *YoutubeChannel) NotifyObservers() {
	for _, obs := range yt.Observers {
		obs.update(yt.NewVideo)
	}
}

func (yt *YoutubeChannel) ReleaseNewvideo(video *Video) {
	yt.NewVideo = video
	yt.NotifyObservers()
}

// UserInterface is a concrete implementation of Observer interface
type UserInterface struct {
	UserName string
	Videos   []*Video
}

func (ui *UserInterface) update(video interface{}) {
	ui.Videos = append(ui.Videos, video.(*Video))
	log.Printf("UI %s - Video: '%s' has just been released\n", ui.UserName, video.(*Video).Title)
}
func NewUserInterface(username string) Observer {
	return &UserInterface{UserName: username, Videos: make([]*Video, 0)}
}
