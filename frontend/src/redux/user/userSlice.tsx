import { configureStore, createSlice , PayloadAction} from "@reduxjs/toolkit";

export interface UserState{
    idnumber:string;
    email:string;
    username:string;
    role:string;
}

const initialState: UserState = {
    idnumber:"",
    email:"",
    username:"",
    role:"",
  };

export const userSlice = createSlice({
    name:"user",
    initialState,
    reducers:{
        login:(state,action: PayloadAction<UserState>)=>{
            console.log("from userSlice",action.payload);
            const {
                idnumber,email,username,role
            } = action.payload; //action.payload will pass to
            //all here at user
            state.idnumber = idnumber;
            state.email = email;
            state.username = username;
            state.role = role;
        },
        //not include action because we are not receive any value for the logout
        logout:(state)=>{
            state.idnumber = "";
            state.email = "";
            state.username = "";
            state.role = "";
        }
    }
});

export const { login, logout } = userSlice.actions;
export default userSlice.reducer;