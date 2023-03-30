import './App.css'
import UserLayout from "./layouts/User.jsx";
import {createBrowserRouter, redirect, RouterProvider} from "react-router-dom";
import routes from "./routes.jsx";


const router = createBrowserRouter(
    [
        {
            path: "/",
            element: <UserLayout routeOptions={routes}/>,
            children: routes,
        },
        {
            path: "*",
            element: <div/>,
            loader: async () => {
                return redirect("/");
            },
        },
    ]
)


function App() {

  return (
     <RouterProvider router={router}>
        </RouterProvider>
  )
}

export default App
