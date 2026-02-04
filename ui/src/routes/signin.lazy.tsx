import {createLazyFileRoute, useRouter} from '@tanstack/react-router'
import {type FormEvent, type MouseEvent, useState} from "react";
import {useAuth} from '../contexts/AuthContext';

export const Route = createLazyFileRoute('/signin')({
    component: RouteComponent,
})

function RouteComponent() {
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const [showSignUp, setShowSignUp] = useState(false);
    const {login} = useAuth();
    const {navigate} = useRouter()

    const handleToggleSignUp = (e: MouseEvent) => {
        e.preventDefault();
        setShowSignUp(prev => !prev);
    }

    const handleLogin = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const res = await fetch('http://localhost:8080/auth/login', {
            method: 'POST',
            body: new FormData(e.currentTarget),
        });


        if (!res.ok) {
            setErrorMessage("Invalid email or password");
            return;
        }

        const {user, token} = await res.json();
        login(user, token);
        setErrorMessage(null)
        navigate({to: '/'});
    }

    const handleSignUp = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const res = await fetch('http://localhost:8080/auth/register', {
            method: 'POST',
            body: new FormData(e.currentTarget),
        });

        if (!res.ok) {
            setErrorMessage("Failed to create account. Please try again.");
            return;
        }

        const {user, token} = await res.json();
        login(user, token);
        setErrorMessage(null)
        navigate({to: '/'});
    }

    const handleOnChange = () => {
        if (errorMessage) {
            setErrorMessage(null);
        }
    }

    return <>
        <h1 className="mb-8 title-1">Welcome back!</h1>

        {showSignUp ? (
                <>
                    <form className="flex flex-col gap-4 w-full max-w-sm mb-4" onSubmit={handleSignUp}
                          onChange={handleOnChange}>
                        <input type="text" placeholder="Username" className="p-2 border border-gray-300 rounded"
                               name="username" required/>
                        <input type="email" placeholder="Email" className="p-2 border border-gray-300 rounded"
                               name="email" required/>
                        <input type="password" placeholder="Password" className="p-2 border border-gray-300 rounded"
                               name="password" required/>

                        {errorMessage && <p className="text-red-500">{errorMessage}</p>}
                        <button type="submit"
                                className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600 transition">
                            Create account
                        </button>
                    </form>

                    <p>Already have an account? <a onClick={handleToggleSignUp}>Log in</a>
                    </p>
                </>
            ) :
            (
                <>
                    <form className="flex flex-col gap-4 w-full max-w-sm mb-4" onSubmit={handleLogin}
                          onChange={handleOnChange}>
                        <input type="email" name="email" placeholder="Email"
                               className="p-2 border border-gray-300 rounded" required/>
                        <input type="password" name="password" placeholder="Password"
                               className="p-2 border border-gray-300 rounded" required/>

                        {errorMessage && <p className="text-red-500">{errorMessage}</p>}
                        <button type="submit"
                                className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600 transition">
                            Log in
                        </button>
                    </form>
                    <p>Don't have an account yet? <a onClick={handleToggleSignUp}>Create a new account</a>
                    </p>
                </>
            )
        }

    </>
}
