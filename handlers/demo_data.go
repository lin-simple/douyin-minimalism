package handlers

var DemoVideo = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		CommentCount:  0,
		FavoriteCount: 0,
		IsFavorite:    false,
		Title:         "bear",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestJack",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

var DemoUserLoginInfo = map[string]User{
	"jack12345678": {
		Id:            1,
		Name:          "jack",
		FollowerCount: 20,
		FollowCount:   5,
		IsFollow:      true,
	},
	"mike12345678910": {
		Id:            2,
		Name:          "mike",
		FollowerCount: 30,
		FollowCount:   5,
		IsFollow:      true,
	},
}
