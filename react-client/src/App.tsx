import React, { useState } from "react";
import "./App.css";
import DatasourceCredentialsForm from "./DatasourceCredentialsForm";
import TextArea from "./TextArea";

export default function App() {
  const api = "http://localhost:9099";
  const [query, setQuery] = useState("");

  const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => { console.log(e.target.value); */ };
  //TODO: refactor duplicate in DatasourceCredsForm
  const initCreds = {
    username: "",
    host: "",
    password: "",
    dbname: "",
  };

  /* TODO: set expected response headers */
  const handleSubmit = (formdata: any) => {
    fetch(api + "/datasources", {
      method: "POST",
      body: JSON.stringify({
        data: {
          type: "datasource",
          properties: { ...formdata },
        },
      }),
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

  const handleQuery = () => {
    console.log(query);

    fetch(api + "/queries", {
      method: "POST",
      body: JSON.stringify({
        data: {
          type: "query",
          properties: {
            query_string: query,
          },
        },
      }),
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

  return (
    <div className="App">
      <DatasourceCredentialsForm creds={initCreds} onSubmit={handleSubmit} />
      <TextArea name="query" onChange={(e) => setQuery(e.target.value)} />
      <button
        onClick={() => {
          handleQuery();
        }}
      >
        {" "}
        query the db
      </button>
    </div>
  );
}
