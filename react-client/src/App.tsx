import React, { useState } from "react";
import "./App.css";
import DatasourceCredentialsForm from "./DatasourceCredentialsForm";
import TextArea from "./TextArea";

export default function App() {
  const api = "http://localhost:9099";
  const [query, setQuery] = useState("");

  const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    /* console.log(e.target.value); */
  };
  //TODO: refactor duplicate in DatasourceCredsForm
  const initCreds = {
    username: "",
    host: "",
    password: "",
    dbname: "",
  };

  /* TODO: set expected response headers */
  const handleSubmit = (formdata: any) => {
    fetch(api + "datasources", {
      method: "POST",
      body: JSON.stringify(formdata),
    })
      .then(function (response) {
        if (!response.ok) {
          throw Error(response.statusText);
        }
        return response;
      })
      .then(function (response) {
        return response.json();
      })
      .then((data) => {
        return data;
      })
      .catch((error) => {});
  };

  const handleQuery = (q: string) => {
    console.log(q);
  };

  return (
    <div className="App">
      <DatasourceCredentialsForm creds={initCreds} onSubmit={handleSubmit} />
      <TextArea name="query" onChange={(e) => setQuery(e.target.value)} />
      <button
        value={query}
        onClick={() => {
          handleQuery(query);
        }}
      >
        {" "}
        query the db
      </button>
    </div>
  );
}
