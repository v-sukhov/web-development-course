import React, {useState} from 'react';

//import './App.css';

import datahub from './ajax/datahub';

import Main from './main/Main';
import LoginForm from './LoginForm';

function App(props) {

  const [isAuthenticated, setIsAuthenticated] = useState(datahub.getAuthenticaion);

  function handleTokenExpired() {
    datahub.clearAuthentication();
    setIsAuthenticated(false);
  }

  function handleSuccessAuthentication() {
    setIsAuthenticated(true);
  }

  return (
    isAuthenticated ? 
      <Main
        onTokenExpired = {handleTokenExpired}
      /> 
    : 
      <LoginForm
        onSuccessAuthentication = {handleSuccessAuthentication}
      />
  )
}

export default App;


// {
//   isAuthenticated ? <Main/> : <LoginForm/>
// }
