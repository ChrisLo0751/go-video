import React from 'react';
import Dropzone from './components/dropzone';
import './App.css'

function App() {
  const handleDrop = (acceptedFiles: File[]) => {
    // Handle the dropped files here
    console.log(acceptedFiles)
  };

  return (
    <div className="App">
      <Dropzone handlerUploadedFiles={handleDrop}></Dropzone>
    </div>
  );
}

export default App;
