const LandingPage = () => {
    const arr = [];
    for (let i = 0; i < 20; i++) {
        arr.push(i);
    }

    return (
        <div className="w-full max-w-4xl h-full flex flex-col gap-2 justify-center mx-auto text-center overflow-hidden no-scrollbar">
            <h1 className="text-6xl font-extrabold">Find your Dream Job</h1>
            <h2 className="text-5xl font-bold">and get Hired</h2>
            <h3 className="text-lg m-8">Explore thousands of job listings for perfect candidate</h3>

            <div className="flex flex-row justify-center gap-8">
                <button className="bg-blue-600 rounded-xl px-8 py-4">Find Jobs</button>
                <button className="bg-red-600 rounded-xl px-8 py-4">Post a Job</button>
            </div>

            {/* Corousal */}
            <div className="h-32 mt-12 flex flex-row gap-8 marquee-content animate-marquee">
                {
                    arr.map(e => (
                        <img key={e} src="/images/companyLogos/image-no-bg.png" className="h-32" />
                    ))
                }
            </div>

        </div>
    )
}

export default LandingPage