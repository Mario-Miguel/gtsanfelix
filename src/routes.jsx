import { AiFillCalendar, AiFillHome, AiFillPhone, AiFillRead } from "react-icons/ai";
import { CalendarView, ContactView, GaleryView, HistoryView, HomeView, RolesView } from "./views/index.jsx";

const routes = [
	{
		path: "/",
		name: "Iniciu",
		label: "nav-home",
		navbar: true,
		icon: <AiFillHome color="inherit" />,
		element: <HomeView />,
	},
	{
		path: "/historia",
		name: "Hestoria",
		label: "nav-history",
		navbar: true,
		element: <HistoryView />,
	},
	{
		path: "/obres",
		name: "Obres en Cartel",
		label: "nav-roles",
		navbar: true,
		icon: <AiFillRead color="inherit" />,
		element: <RolesView />,
	},
	{
		path: "/calendariu",
		name: "Calendariu",
		label: "nav-calendar",
		navbar: true,
		icon: <AiFillCalendar color="inherit" />,
		element: <CalendarView />,
	},

	{
		path: "/galeria",
		name: "Galería",
		label: "nav-gallery",
		navbar: true,
		element: <GaleryView />,
	},

	{
		path: "/contautu",
		name: "Contautu",
		label: "nav-contact",
		navbar: true,
		icon: <AiFillPhone color="inherit" />,
		element: <ContactView />,
	},
];

export default routes;
