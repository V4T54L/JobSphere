import { useState } from "react"
import LoginModal from "./LoginModal"
import SignupModal from "./SignupModal"
import { useDispatch, useSelector } from "react-redux"
import { setAccessToken } from "../features/userSlice"

const Layout = ({ children }) => {
  const accessToken = useSelector(state => state.accessToken)
  const dispatch = useDispatch()

  const [modalState, setModalState] = useState(0)
  const closeModal = () => {
    setModalState(0)
  }
  const openLoginModal = () => {
    setModalState(1)
  }
  const openSignupModal = () => {
    setModalState(2)
  }

  const [userInfoDropdown, setUserInfoDropdown] = useState(false)
  const closeDropdown = () => setUserInfoDropdown(false)
  const toggleDropdown = () => setUserInfoDropdown(prev => !prev)

  const signout = ()=>{
    dispatch(setAccessToken(""))
  }

  return (
    <div className="flex flex-col w-screen h-screen overflow-x-hidden no-scrollbar bg-blue-950 text-blue-100">
      {/* Navbar */}
      <div className="flex flex-row justify-between px-8 py-4 w-full">
        {/* LOGO */}
        <h1 className="text-4xl font-extrabold">JobSphere</h1>
        {
          accessToken?.length == 0 ?
            (
              <div className="flex flex-row gap-4">
                <button className="px-2 py-1 border-2 rounded-lg" onClick={openLoginModal}>Login</button>
                <button className="px-2 py-1 border-2 rounded-lg" onClick={openSignupModal}>Signup</button>
              </div>
            ) : (
              <div
                onClick={toggleDropdown}
                className="flex flex-row gap-2 items-center mr-8 hover:bg-blue-700 py-1 px-2 transition-all duration-100 rounded-xl">
                <img src={"https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg"} className="h-12 rounded-full" />
                <h6>@username</h6>
              </div>
            )
        }
      </div>
      {
        userInfoDropdown &&
        <div id="dropdownInformation" onBlur={closeDropdown} className="absolute right-10 top-20 z-10 bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700 dark:divide-gray-600 left-50">
          <div className="px-4 py-3 text-sm text-gray-900 dark:text-white">
            <div>Bonnie Green</div>
            <div className="font-medium truncate">name@flowbite.com</div>
          </div>
          <ul className="py-2 text-sm text-gray-700 dark:text-gray-200" aria-labelledby="dropdownInformationButton">
            <li>
              <a href="#" className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Dashboard</a>
            </li>
            <li>
              <a href="#" className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Settings</a>
            </li>
            <li>
              <a href="#" className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Earnings</a>
            </li>
          </ul>
          <div className="py-2">
            <a onClick={signout} href="#" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white">Sign out</a>
          </div>
        </div>
      }

      {/* Content */}
      <main className="flex-1">
        {children}
      </main>

      <footer className="pt-4">
        <div className="container mx-auto pt-2 text-xs flex flex-col md:flex-row justify-around">
          <ul className="flex justify-center mb-2">
            <li className="mr-6">
              <a href="#" className="">
                About
              </a>
            </li>
            <li className="mr-6">
              <a href="#" className="">
                Contact
              </a>
            </li>
            <li>
              <a href="#" className="">
                Terms & Conditions
              </a>
            </li>
          </ul>
          <p className="text-center">
            Copyright &copy; {new Date().getFullYear()} JobSphere
          </p>
        </div>
      </footer>

      {
        accessToken == "" &&
          modalState == 1 ? (
          <LoginModal closeModal={closeModal} openSignupModal={openSignupModal} />
        ) : (

          modalState == 2 &&
          <SignupModal closeModal={closeModal} openLoginModal={openLoginModal} />
        )
      }
    </div>
  )
}

export default Layout