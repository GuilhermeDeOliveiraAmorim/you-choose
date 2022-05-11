import { Link } from "react-router-dom";
import StyleMenu from "./Menu.module.scss";

export default function Menu() {
    function myFunction() {
        var x: HTMLElement | any = document.getElementById("myTopnav");
        if (x.className === "topnav") {
            x.className += " responsive";
        } else {
            x.className = "topnav";
        }
    }

    const rotas = [
		{
			label: 'Início',
			to: '/'
		},
		{
			label: 'Home',
			to: '/home'
		},
		{
			label: 'Lists',
			to: '/lists'
		},
		{
			label: 'Score',
			to: '/score'
		}
	];

    return (
        <div id="myTopnav">
            <ul className={StyleMenu.topnav}>
                <li>Home</li>
                <li>Lists</li>
                <li>Score</li>
                <li>You Choose</li>
            </ul>
        </div>
    );
}
