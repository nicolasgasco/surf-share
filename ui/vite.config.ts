import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'
import {tanstackRouter} from '@tanstack/router-plugin/vite'

// https://vite.dev/config/
export default defineConfig({
    plugins: [
        react({
            babel: {
                plugins: [['babel-plugin-react-compiler']],
            },
        }),
        tanstackRouter({
            target: 'react',
            autoCodeSplitting: true,
        }),
        tailwindcss()
    ],
    server: {
        host: '0.0.0.0',
        port: 5173,
        strictPort: true,
    }
})
