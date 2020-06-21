import React from "react";
import "./App.css";

const App = (props: any) => {
  const config = {
    endpoint: "http://localhost:9099",
  };

  let querystring = " ";
  let querypos = 0;
  let dbname = "new";
  let columns = [];
  let rowdata = [];
  let dbs = ["ae1", "New", "Test"];
  let previousqueries = [];

  const QueryForm = (props: any) => {
    const saveCursor = (event: any) => {
      let querypos = event.target.selectionStart;
      let se = event.target.selectionEnd;
    };

    const addToQueryString = (newVal: any, event: any) => {
      let astr = querystring.substr(0, querypos);
      let cstr = querystring.substr(querypos);
      querystring = astr + " " + newVal + cstr;
    };
    const replacequerystring = (newVal: any, event: any) => {
      querystring = newVal;
    };

    const makequery = (event: any) => {
      let rowdata = [];
      /* let previousqueries.push({db: dbname, query: querystring}) */
      let payload = {
        query: querystring,
        db: dbname,
      };

      fetch(config.endpoint, {
        method: "POST",
        body: JSON.stringify(payload),
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
          rowdata = data;
          return data;
        })
        .then((data) => {
          columns = Object.keys(data.Data[0]);
          var parsedArray = data.Data.map(function (row: object) {
            return Object.values(row);
          });
          rowdata = parsedArray;
        })
        .catch(function (error) {});
    };
  };

  return <div className="App"></div>;
};

export default App;
