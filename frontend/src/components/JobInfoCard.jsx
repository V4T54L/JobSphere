import { FaLocationDot } from "react-icons/fa6";
import { FaRegHeart } from "react-icons/fa";
import { useNavigate } from "react-router-dom";


const JobInfoCard = ({job}) => {
    const navigate = useNavigate()
    const navigateTo = ()=>{
        // console.log(`endpoint: `)
        navigate("/job/1")
    }
    return (
        <div className="p-4 rounded-lg shadow bg-indigo-900 border-indigo-800 text-left" onClick={navigateTo}>
            <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{job?.title}</h5>
            <div className="flex flex-row justify-between items-center">
                {/* <img src="/images/companyLogos/image-no-bg.png" className="h-12 bg-gradient-to-b from-stone-950 via-stone-700 to-stone-400 rounded-lg" /> */}
                <img src="/images/companyLogos/image-does-not-exists.png" alt={job?.company} className="h-12 bg-gradient-to-b from-stone-950 via-stone-700 to-stone-400 rounded-lg" />
                <h6 className="flex flex-row gap-2 items-center mr-8">
                    <FaLocationDot />
                    {job?.location}
                </h6>
            </div>
            <hr className="my-2 opacity-30"/>
            <p className="m-3 font-normal text-gray-700 dark:text-gray-400">{job?.description}</p>
            <div className="flex flex-row gap-2">
                <button className="px-2 py-1 basis-11/12 flex-grow bg-blue-600 rounded-md">Read more</button>
                <button className="px-2 py-1 rounded-md border basis-1/12 flex-grow mx-auto"><FaRegHeart /></button>
            </div>
        </div>
    )
}

export default JobInfoCard