import {createLazyFileRoute, Outlet} from '@tanstack/react-router'

export const Route = createLazyFileRoute('/breaks')({
    component: RouteComponent,
})

function RouteComponent() {
    return <Outlet/>
}
