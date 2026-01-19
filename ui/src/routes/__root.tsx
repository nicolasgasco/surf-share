import {createRootRoute, Link, Outlet} from "@tanstack/react-router";
import {useAuth} from "../contexts/AuthContext.tsx";


export const Route = createRootRoute({
    component: RouteComponent,

});

function RouteComponent() {
    const {user} = useAuth()

    return <div className="w-full flex flex-col items-center justify-start min-h-screen">
        <div className="w-full flex items-center justify-between py-4 px-8">
            <Link to="/" className="text-2xl font-bold">SurfShare</Link>

            {user ? <p>Hello, {user.username}!</p> : <Link to="/signin" className="font-bold">Sign in</Link>}
        </div>
        <div className="h-full flex flex-col items-center justify-start flex-1 py-24 px-8">
            <Outlet/>
        </div>
    </div>
}