import React, {FunctionComponent, useEffect, useState} from 'react';
import baily from './baily.svg';
import './App.css';

const App: FunctionComponent = () => {

  const [error, setError] = useState(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [status, setStatus] = useState(null);

  useEffect(() => {
    console.log(process.env.REACT_APP_BFLOBOX_API_URL)
    fetch(`${process.env.REACT_APP_BFLOBOX_API_URL}/health`)
      .then(res => res.json())
      .then(
        (result) => {
          setIsLoaded(true);
          setStatus(result.status);
        },
        (error) => {
          console.log("ERROR: ", error)
          setIsLoaded(true);
          setError(error);
        }
      )
  }, [])

  if (error) {
    return <div>Error :(</div>;
  } else if (!isLoaded) {
    return <div>Loading...</div>;
  } else {
    return (
      <div className="App">
        <header className="App-header">
          <img src={baily} className="App-logo" alt="logo" />
          <h1>{status}</h1>
        </header>
      </div>
    );
  }
}

export default App;
