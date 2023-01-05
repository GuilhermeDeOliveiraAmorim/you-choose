package addchoosertolist

func AddChooserToList(input *InputAddChooserToMovieListDto) *OutputAddChooserToMovieListDto {
	chooser := input.Chooser
	movieList := input.MovieList

	movieList.AddChooser(chooser)

	output := &OutputAddChooserToMovieListDto{
		IDMovieList: movieList.ID,
		IDChooser:   chooser.ID,
	}

	return output
}
