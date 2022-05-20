import StyleItem from "./Item.module.scss";

interface Props {
    poster: string;
    title: string;
    id: number;
    score: number;
}

export default function Item(props: Props) {

    const { id, poster, title, score } = props;

    const myStyle = {
        backgroundImage: `url(${poster})`,
        backgroundSize: "cover",
        backgroundRepeat: "no-repeat",
        backgroundPosition: "center"
    };

    return (
        <>
            <div key={id} className={StyleItem.item}>
                <img src={poster} alt={title}  style={myStyle} className={StyleItem.item_poster} />
                <div className={StyleItem.item_title}>{title}</div>
                <div className={StyleItem.item_score}>{score}</div>
            </div>
        </>
    )
}