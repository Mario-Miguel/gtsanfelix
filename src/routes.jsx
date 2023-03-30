import {AiFillCalendar, AiFillHome, AiFillPhone, AiFillRead} from "react-icons/ai";
import {CalendarView, ContactView, GaleryView, HistoryView, HomeView, RolesView} from "./views/index.jsx";

const routes = [
    {
        path: "/",
        name: "Iniciu",
        navbar: true,
        icon: <AiFillHome color="inherit"/>,
        element: <HomeView/>,
    },
    {
        path: "/obres",
        name: "Obres en Cartel",
        navbar: true,
        icon: <AiFillRead color="inherit"/>,
        element: <RolesView/>,
    },
    {
        path: "/calendariu",
        name: "Calendariu",
        navbar: true,
        icon: <AiFillCalendar color="inherit"/>,
        element: <CalendarView/>,
    },
    {
        path: "/contautu",
        name: "Contautu",
        navbar: true,
        icon: <AiFillPhone color="inherit"/>,
        element: <ContactView/>,
    },
    {
        path: "/galeria",
        name: "Galería",
        navbar: true,
        element: <GaleryView/>
    },
    {
        path: "/historia",
        name: "Compare Offers",
        navbar: true,
        element: <HistoryView/>
    }

];

export default routes;
