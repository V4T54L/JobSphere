import { useEffect, useState } from "react"
import { serverBaseUrl } from "../utils/constants"
import { useParams } from "react-router-dom"


const ViewJob = () => {
  const { id } = useParams();

  const [job, setJob] = useState(null)
  const [error, setError] = useState("")

  useEffect(() => {
    ; (
      async () => {
        try {
          const resp = await fetch(`${serverBaseUrl}/jobs/${id}`)
          const data = await resp.json()
          if (data.error) {
            setError(data.error)
          } else {
            setJob(data)
          }
        } catch (error) {
          if (error.name === "TypeError" && error.message.includes("Failed to fetch")) {
            setError("Connection refused. Please check your network connection.")
          } else {
            setError(`(${error}`)
          }
        }
      }
    )()
  }, [id])

  if (job == null && error == "") {
    return <div>Loading...</div>
  }

  return (
    <div className="container mx-auto p-6 sha">
      {
        error&& <div>Error {error}</div>
      }
      {
        job &&
        <div className="bg-indigo-950 shadow-lg shadow-indigo-900 rounded-lg p-6">
          <h1 className="text-2xl font-bold mb-4">{job.title}</h1>
          <div className="text-sm text-gray-200 mb-2">
            <span className="mr-2">{job.company}</span>
            <span className="mr-2">|</span>
            <span>{job.location}</span>
          </div>
          <div className="text-lg text-gray-200 mb-2">
            Salary: <span className="font-bold">{job.salary}</span>
          </div>
          <p className="text-gray-400 mb-4">{job.description}</p>
          <div className="mt-4 border-t pt-4">
            <div className="flex items-center justify-between">
              <span className="text-gray-400">Posted on: {new Date(job.createdAt).toLocaleDateString()}</span>
              <button className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600 transition duration-200">
                Apply Now
              </button>
            </div>
          </div>
        </div>
      }
    </div>

  )
}

export default ViewJob