import { useEffect, useState } from "react"
import { mockJobs } from "../utils/mock"
import JobInfoCard from "./JobInfoCard"

const ListJobs = () => {
  const [jobs] = useState(mockJobs)
  const [filteredJobs, setFilteredJobs] = useState(mockJobs)

  useEffect(() => {
    setFilteredJobs(jobs)
  }, [jobs])

  const searchJobs = (query) => {
    query = query.toLowerCase()
    setFilteredJobs(jobs.filter(job => job.title.toLowerCase().includes(query)))
  }

  return (
    <div className="w-full max-w-7xl h-full mx-auto text-center overflow-x-hidden no-scrollbar p-8">
      <h1 className="text-6xl font-extrabold mb-8">Latest Jobs</h1>
      <form className="flex flex-row gap-3 my-2">
        <div className="px-2 py-1 border-2 border-blue-600 rounded-md basis-11/12 flex-grow">
          <input type="text" name="searchByTitle" id="searchByTitle"
            onChange={(e) => searchJobs(e.target.value)}
            placeholder="Search Jobs by Title..."
            className="bg-inherit outline-none w-full" />
        </div>
        <button type="submit" className="bg-blue-600 px-2 py-1 rounded-md basis-1/12 flex-grow">Search</button>
      </form>

      <div className="flex flex-row gap-3 my-2">
        <div className="px-2 py-1 border-2 border-blue-600 rounded-md basis-5/12 flex-grow">
          <select className="w-full bg-inherit outline-none">
            <option>Filter by Location</option>
          </select>
        </div>
        <div className="px-2 py-1 border-2 border-blue-600 rounded-md basis-5/12 flex-grow">
          <select className="w-full bg-inherit outline-none">
            <option>Filter by Company</option>
          </select>
        </div>
        <button type="submit" className="bg-red-600 px-2 py-1 rounded-md basis-2/12 flex-grow">Clear filters</button>
      </div>

      <div className="grid grid-cols-3 gap-4 mt-8">
        {
          filteredJobs.map(job => (
            <JobInfoCard key={job.id} job={job} />
          ))
        }
      </div>
    </div>
  )
}

export default ListJobs