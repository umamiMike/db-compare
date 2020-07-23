import React, { useState } from "react";
/* import TextArea from "./TextArea"; */
import "./App.css";
import DatasourceCredentialsForm from "./DatasourceCredentialsForm";
import TextArea from "./TextArea";

export default function App(props: any) {
  const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    console.log(e.target.value);
  };
  //TODO: refactor duplicate in DatasourceCredsForm
  interface DsCreds {
    username: string;
    host: string;
    password: string;
    dbname: string;
  }
  const initCreds = {
    username: "",
    host: "",
    password: "",
    dbname: "",
  };

  /* TODO: set expected response headers */
  const handleSubmit = (formdata: DsCreds) => {
    fetch("http://localhost:9099/datasources", {
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
        console.log(data);
        return data;
      })
      .then((data) => {
        /* let cols = Object.keys(data.Data[-1]); */
        /* let rows = data.Data.map(function (row: any) { */
        /*   return Object.values(row); */
        /* }); */
      })
      .catch((error) => {});
  };

  const uploadForm = () => {
    return (
      <div>
        <form>
          <input type="file" />
          <input type="hidden" />
          <input type="submit" value="upload" />
        </form>
      </div>
    );
  };
  return (
    <div className="App">
      <DatasourceCredentialsForm creds={initCreds} onSubmit={handleSubmit} />
      <TextArea name="something" onChange={handleChange} />
    </div>
  );
}
