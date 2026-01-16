import {createLazyFileRoute} from '@tanstack/react-router'
import {type MouseEvent, useState} from "react";

export const Route = createLazyFileRoute('/signin')({
    component: RouteComponent,
})

function RouteComponent() {
    const [showSignUp, setShowSignUp] = useState(false);

    const handleToggleSignUp = (e: MouseEvent) => {
        e.preventDefault();
        setShowSignUp(prev => !prev);
    }

    return <>
        <h1 className="mb-8">Welcome back!</h1>

        {showSignUp ? (
            <>
                <form className="flex flex-col gap-4 w-full max-w-sm mb-4">
                    <input type="text" placeholder="Username" className="p-2 border border-gray-300 rounded"/>
                    <input type="email" placeholder="Email" className="p-2 border border-gray-300 rounded"/>
                    <input type="password" placeholder="Password" className="p-2 border border-gray-300 rounded"/>

                    <button type="submit"
                            className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600 transition">
                        Create account
                    </button>
                </form>
                <p>Already have an account? <a onClick={handleToggleSignUp}>Log in</a></p>
            </>
        ) : (
            <>
                <form className="flex flex-col gap-4 w-full max-w-sm mb-4">
                    <input type="email" placeholder="Email" className="p-2 border border-gray-300 rounded"/>
                    <input type="password" placeholder="Password" className="p-2 border border-gray-300 rounded"/>
                    <button type="submit"
                            className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600 transition">
                        Log in
                    </button>
                </form>
                <p>Don't have an account yet? <a onClick={handleToggleSignUp}>Create a new account</a></p>
            </>
        )}
    </>
}
