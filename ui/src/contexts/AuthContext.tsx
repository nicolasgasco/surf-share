import React, {createContext, useContext, useEffect, useState} from 'react';
import {jwtDecode} from 'jwt-decode';
import Cookies from 'js-cookie';
import type {User} from '../types/user';

interface AuthContextType {
    user: User | null;
    token: string | null;
    login: (user: User, token: string) => void;
    logout: () => void;
    isLoading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({children}: { children: React.ReactNode }) {
    const [user, setUser] = useState<User | null>(null);
    const [token, setToken] = useState<string | null>(null);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const storedToken = Cookies.get('auth_token');
        if (storedToken) {
            try {
                const decoded = jwtDecode<{ id: string; username: string; email: string }>(storedToken);
                setToken(storedToken);
                setUser({
                    id: decoded.id,
                    username: decoded.username,
                    email: decoded.email,
                });
            } catch (error) {
                console.error('Invalid token:', error);
                Cookies.remove('auth_token');
            }
        }
        setIsLoading(false);
    }, []);

    const login = (newUser: User, newToken: string) => {
        setUser(newUser);
        setToken(newToken);

        const decoded = jwtDecode<{ exp: number }>(newToken);
        const expiryDate = new Date(decoded.exp * 1000);
        Cookies.set('auth_token', newToken, {
            expires: expiryDate,
            secure: true,
            sameSite: 'strict',
        });
    };

    const logout = () => {
        setUser(null);
        setToken(null);
        Cookies.remove('auth_token');
    };

    return (
        <AuthContext.Provider value={{user, token, login, logout, isLoading}}>
            {children}
        </AuthContext.Provider>
    );
}

export function useAuth() {
    const context = useContext(AuthContext);

    if (!context) {
        throw new Error('useAuth must be used within AuthProvider');
    }

    return context;
}
