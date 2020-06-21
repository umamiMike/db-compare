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

  const saveCursor = (event: any) => {
    event.preventDefault();
    let querypos = event.target.selectionStart;
    let se = event.target.selectionEnd;
    return "";
  };

  const insertIntoQuery = (newVal: any, event: any) => {
    let astr = querystring.substr(0, querypos);
    let cstr = querystring.substr(querypos);
    querystring = astr + " " + newVal + cstr;
  };
  interface Payload {
    query: string;
    db: string;
  }

  const makeQuery = (e: Event, payload: Payload) => {
    let rowdata = [];
    QueryForm(payload);
  };

  const QueryForm = (payload: Payload) => {
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

  const replaceQueryString = (newVal: any, event: Event) => {
    let querystring = newVal;
    return querystring;
  };

  return (
    <div className="App">
      <div className="Hello">
        {/* <div id="replacequery_cont"> */}
        {/*   <button className="replacequery"></button> */}
        {/* </div> */}
        <h1> put message here</h1>
        <div id="someform" className="form">
          <div className="tab-content">
            <div id="signup">
              <form>
                <div className="top-row">
                  <div className="field-wrap">
                    <label>
                      Database <span className="req">*</span>
                    </label>
                    <input
                      type="text"
                      v-model="dbname"
                      placeholder="ae1"
                    ></input>
                  </div>
                  <div className="field-wrap">
                    <label>
                      Query String<span className="req">*</span>
                    </label>
                    <textarea
                      className="queryinput"
                      cols={40}
                      rows={5}
                      onClick={() => saveCursor}
                      placeholder="Describe Users"
                    >
                      show tables
                    </textarea>
                    {/* <button type="button" onClick={makeQuery}> */}
                    {/*   load data */}
                    {/* </button> */}
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
      <div className="tablecontainer">
        <table>
          <tr className="header">
            //display headers
            {/* <th v-for="header in columns"> */}
            <th>{/* <h2 onClick={insertIntoQuery}> header info</h2> */}</th>
          </tr>
          <tr v-for="row in rowdata">
            {/* <td v-for="fielddata in row"> */}
            <td>
              td
              {/* <h2 onClick={insertIntoQuery}> fielddata </h2> */}
            </td>
            <td>
              td
              {/* <h2 onClick={insertIntoQuery}> fielddata </h2> */}
            </td>

            <td>
              td
              {/* <h2 onClick={insertIntoQuery}> fielddata </h2> */}
            </td>
          </tr>
        </table>
      </div>
    </div>
  );
};

export default App;
