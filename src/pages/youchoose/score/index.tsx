import Item from "./list/item";
import StyleScore from "./Score.module.scss";
import movies from "assets/movies.json";

export default function Score() {
    
    return (
        <>
            <div className={StyleScore.box_score}>
                <div className={StyleScore.score}>
                    <Item id={movies[0].id} poster={movies[0].poster} title={movies[0].title} score={movies[0].score} />
                    <Item id={movies[1].id} poster={movies[1].poster} title={movies[1].title} score={movies[1].score} />
                    <Item id={movies[2].id} poster={movies[2].poster} title={movies[2].title} score={movies[2].score} />
                </div>
            </div>
        </>

    )
}