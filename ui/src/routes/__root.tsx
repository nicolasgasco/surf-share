import {createRootRoute, Link, Outlet} from "@tanstack/react-router";


export const Route = createRootRoute({
    component: () => (
        <div className="w-full flex flex-col items-center justify-start min-h-screen">
            <div className="w-full flex items-center justify-between py-4 px-8">
                <Link to="/" className="text-2xl font-bold">SurfShare</Link>

                <Link to="/signin" className="font-bold">Sign in</Link>
            </div>
            <div className="h-full flex flex-col items-center justify-start flex-1 py-24 px-8">
                <Outlet/>
            </div>
        </div>
    ),
});