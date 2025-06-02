import { TabMenu } from 'primereact/tabmenu';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function Navbar() {

	const [activeIndex, setActiveIndex] = useState(0);
	const navigate = useNavigate();

	const itemsMenu = [
		{ label: 'Home', icon: 'pi pi-fw pi-home', command: () => { navigate("/home") } },
	];

	return (
		<TabMenu model={itemsMenu} activeIndex={activeIndex} onTabChange={(e) => setActiveIndex(e.index)} />
	)
}