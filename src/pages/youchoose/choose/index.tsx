import StyleChoose from "./Choose.module.scss";
import movies from "assets/movies.json";
import Movie from "../movie";
import { useState } from "react";

export default function Choose() {
    const initialIds = arrayNumberRandom();

    const [movieIds, setMovieIds] = useState({
        movieIdA: initialIds[0],
        movieIdB: initialIds[1],
    });

    function handleIds(idClick: number) {

        idClick = idClick;
        let idA = movieIds.movieIdA;
        let idB = movieIds.movieIdB;

        if (idA === idClick) {
            do {
                var aux = numberRandom();
                idB = aux;
            } while (idA === aux);
        } else {
            do {
                var aux = numberRandom();
                idA = aux;
            } while (idB === aux);
        }

        //console.log(movieIdA, movieIdB);
        /*Os ids não podem ser iguais*/

        let movieIdA = idA;
        let movieIdB = idB;

        setMovieIds({ movieIdA, movieIdB });
    }

    function numberRandom() {
        let num = 0;
        while (num === 0) {
            num = Math.floor(Math.random() * Object.keys(movies).length);
        }
        return num;
    }

    function arrayNumberRandom() {
        let x = 0;
        let y = 0;

        while (x === y) {
            x = numberRandom();
            y = numberRandom();
        }
        return [x, y];
    }

    return (
        <>
            <div className={StyleChoose.box_choose}>
                <div onClick={() => handleIds(movieIds.movieIdA)}>
                    <Movie
                        title={movies[movieIds.movieIdA].title}
                        id={movies[movieIds.movieIdA].id}
                        poster={movies[movieIds.movieIdA].poster}
                        score={movies[movieIds.movieIdA].score}
                        votes={movies[movieIds.movieIdA].votes}
                    />
                </div>
                <div onClick={() => handleIds(movieIds.movieIdB)}>
                    <Movie
                        title={movies[movieIds.movieIdB].title}
                        id={movies[movieIds.movieIdB].id}
                        poster={movies[movieIds.movieIdB].poster}
                        score={movies[movieIds.movieIdB].score}
                        votes={movies[movieIds.movieIdB].votes}
                    />
                </div>
            </div>
        </>
    );
}
