import { useSelector } from "react-redux"
import LandingPage from "./components/LandingPage"
import Layout from "./components/Layout"
import ListJobs from "./components/ListJobs"
import ViewJob from "./components/ViewJob"
import { BrowserRouter, Routes, Route } from "react-router-dom"
import PageNotFound from "./components/PageNotFound"
import ProtectedRoute from "./components/ProtectedRoute"

function App() {
  const accessToken = useSelector(state => state.accessToken)

  const Page = accessToken == "" ? LandingPage : ListJobs;
  // const Page = accessToken == "" ? LandingPage : ViewJob;oiui

  return (
    <BrowserRouter>
      <Layout>
        <Routes>
          <Route path="/" element={<Page />} />
          <Route path="/job/:id" element={<ProtectedRoute><ViewJob /></ProtectedRoute>} />
          <Route path="/*" element={<PageNotFound />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  )
}

export default App
