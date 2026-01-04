import {createLazyFileRoute, useParams} from '@tanstack/react-router'

export const Route = createLazyFileRoute('/breaks/$breakSlug')({
    component: RouteComponent,
})

function RouteComponent() {
    const {breakSlug} = useParams({
        from: '/breaks/$breakSlug',
    })

    return <div>
        <h1>Your break info for {breakSlug}</h1>
    </div>
}
