import {HomePage} from "./pages/HomePage.tsx";

function App() {
    return (
        <div className="w-full flex flex-col items-center justify-start min-h-screen">
            <div className="w-full py-4 px-8">
                <p className="text-2xl font-bold">SurfShare</p>
            </div>
            <div className="h-full flex flex-col items-center justify-start flex-1 py-24 px-8">
                <HomePage/>
            </div>
        </div>
    )
}

export default App
