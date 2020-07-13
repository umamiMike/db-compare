import React, { useState } from "react";
/* import TextArea from "./TextArea"; */
import "./App.css";

export default function App(props: any) {
  const [state, setState] = useState([]);
  const [count, setCount] = useState(1);

  const uploadForm = () => {
    return (
      <div>
        <form>
          <input type="file" />
          <input type="hidden" />
          <input type="submit" value="upload" />
        </form>
        <button onClick={() => setCount(1)}>click me</button>
      </div>
    );
  };

  return <div className="App"> {uploadForm()} </div>;
}
