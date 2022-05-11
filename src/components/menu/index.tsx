import { Link } from "react-router-dom";
import StyleMenu from "./Menu.module.scss";

export default function Menu() {
    const rotas = [
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
		},
		{
			label: 'You Choose',
			to: '/you-choose'
		}
	];

    return (
        <nav>
			<ul className={StyleMenu.menu__list}>
				{rotas.map((rota, index) => (
					<li key={index} className={StyleMenu.menu__link}>
						<Link to={rota.to}>
							{rota.label}
						</Link>
					</li>
				))}
			</ul>
        </nav>
    );
}
