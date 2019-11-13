import React from 'react';
import './App.css';

function App() {
  return (
    <div className="App">
      <div className="title-header">
        <h1>the <br></br> javascript <br></br> quiz<span className="dot">.</span></h1>
      </div>
      <div className="content">
        <br></br>
        <p>Just a gamified version of the
          javascript questions by <a href="https://github.com/lydiahallie/javascript-questions">Lydia Halie</a> <br></br>
          A project by Junaid Rahim
        </p>
        <br></br>
      </div>
      <div className="button-container">
        <button className="btn btn-danger btn-lg">Start Playing</button>
      </div>
    </div>
  );
}

export default App;
