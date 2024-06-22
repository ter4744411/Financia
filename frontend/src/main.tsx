import React from 'react'
import ReactDOM from 'react-dom/client'
import Home from './pages/Home/Home.tsx'
import Register from './pages/Register/Register.tsx'
import App from './App.tsx'
import './index.css'
import Login from './pages/Login/Login.tsx'
import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouterProvider,
} from "react-router-dom";
import Layout from './Layout'
import UserDashboard from './pages/UserDashboard/UserDashboard.tsx'
import AdminDashboard from './pages/AdminDashboard/AdminDashboard.tsx'
import {Provider} from "react-redux";
import store, { persistor } from './redux/store';
import { PersistGate } from 'redux-persist/integration/react';

const router = createBrowserRouter([
  {
    element: <Layout />,
    children: [
      {
        path:"/",
        element:<Home/>,
      },
      {
        path:"/register",
        element:<Register/>,
      },
      {
        path:"/login",
        element:<Login/>,
      },
      {
        path:"/userdashboard",
        element:<UserDashboard/>,
      },
      {
        path:"/admindashboard",
        element:<AdminDashboard/>,
      },
    ],
  },
]);
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <RouterProvider router={router}/>
      </PersistGate>
    </Provider>
  </React.StrictMode>
)
