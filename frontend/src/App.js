import React, {useContext, useState} from 'react';
import {BrowserRouter} from "react-router-dom";
import AppRouter from "./components/AppRouter";
import Footer from "./components/UI/footer/Footer";
import { AuthContext } from "./context";

function App() {

  const [reg, setReg] = useState(false)
  const [userLogin, setUserLogin] = useState("")
  const [userStatus, setUserStatus] = useState(0)
  const [userID, setUserID] = useState(0)
  const [ignoreThat, setIgnoreThat] = useState(0)

  return (
    <AuthContext.Provider value = {{reg, setReg, userLogin, setUserLogin, userStatus, setUserStatus, userID, setUserID, ignoreThat, setIgnoreThat}}>
      <BrowserRouter>
          <AppRouter/>
      </BrowserRouter>
      </AuthContext.Provider>
  )
}

export default App;
