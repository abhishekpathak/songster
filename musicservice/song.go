package musicservice

type Song struct {
	Id        int      `json:"id"`
	Videoid   string   `json:"videoid"`
	Name      string   `json:"name"`
	Length    int      `json:"length"`
	Seek      int      `json:"seek"`
	AddedBy   string   `json:"added_by"`
	Thumbnail string   `json:"thumbnail"`
	Details   SongInfo `json:"details"`
}

type LibSong struct {
	Videoid string `json:"videoid"`
	Track   string `json:"track"`
	Fav     bool   `json:"fav"`
	Source  string `json:"source"`
}

type User struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func (s *Song) Score() float64 {
	var score float64
	defer func() {
		if err := recover(); err != nil {
			score = -100.00
		}
	}()
	likes := float64(s.Details.Likes)
	dislikes := float64(s.Details.Dislikes)
	views := float64(s.Details.Views)
	score = (likes * 100.00 / views) - (dislikes / likes)
	return score
}
