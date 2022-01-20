package Services

import (
	"web-service-gin/Dao"
	"web-service-gin/Model"
)

func GetAlbum() []Model.Album {
	return Dao.QueryAlbums()
}
