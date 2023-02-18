package usecases

type TagDto struct {
	ID        string `json:"tag_id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type InputCreateTagDto struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputCreateTagDto struct {
	Tag TagDto `json:"tag"`
}

type InputDeleteTagDto struct {
	TagId string `json:"tag_id"`
}

type OutputDeleteTagDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputFindTagDto struct {
	TagId string `json:"tag_id"`
}

type OutputFindTagDto struct {
	Tag TagDto `json:"tag"`
}

type InputFindTagByNameDto struct {
	TagName string `json:"tag_name"`
}

type OutputFindTagByNameDto struct {
	Tag TagDto `json:"tag"`
}

type InputUpdateTagDto struct {
	TagId   string `json:"tag_id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputUpdateTagDto struct {
	Tag TagDto `json:"tag"`
}

type InputIsDeletedTagDto struct {
	TagId string `json:"tag_id"`
}

type OutputIsDeletedTagDto struct {
	IsDeleted bool `json:"is_tag_deleted"`
}

type OutputFindAllTagDto struct {
	Tags []TagDto `json:"tags"`
}

type InputFindAllTagMovieListsDto struct {
	TagId string `json:"tag_id"`
}

type OutputFindAllTagMovieListsDto struct {
	Tag        TagDto         `json:"tag"`
	MovieLists []MovieListDto `json:"movies"`
}
