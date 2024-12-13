import { createSlice } from "@reduxjs/toolkit"


const initialState = {
    accessToken: "",
    info: {
        username: "",
        email: "",
        profilePicture: "",
    },
}

export const userSlice = createSlice({
    name: "user data",
    initialState,
    reducers: {
        setAccessToken: (state, action) => {
            state.accessToken = action.payload
        },
    }
})

export const { setAccessToken } = userSlice.actions

export default userSlice.reducer