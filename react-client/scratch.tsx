import React, { useState } from "react";
/* import TextArea from "./TextArea"; */
import "./App.css";




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





export default function App(props: any) {
  const [state, setState] = useState([]);

  /* const config = { */
  /*   endpoint: "http://localhost:9099", */
  /* }; */

  /* // TODO: 
     define 
     convert this data from hard coded vars to interface, and props 
    do I care about  intial working or just doing a redesign based on functionality?
  */

  interface Payload {
    query: string;
    db: string;
  }


  /* TODO: convert to react form upload; */
  const uploadForm = (count: any) => {
    return (
      <div>
        <form
          encType="multipart/form-data"
          action="http://localhost:9099/upload"
          method="post"
        >
          <input type="file" />
          <input type="hidden" />
          <input type="submit" value="upload" />
        </form>
        <h1>{count}</h1>

        <button onClick={() => setCount(count + 1)}>click me</button>
        <h1>{barf.barf} barfs</h1>
        <h1>{barf.snarf} snarfs</h1>
      </div>
    );
  };

  return (
    <div className="App">
      {uploadForm()}
      <div className="Hello">
        <div id="queryForm" className="form">
          <div className="tab-content">
            {/* <div id="signup"> */}
            {/*   <form> */}
            {/*     <div className="top-row"> */}
            {/*       <div className="field-wrap"> */}
            {/*         <label> */}
            {/*           Database <span className="req">*</span> */}
            {/*         </label> */}
            {/*         <input */}
            {/*           type="text" */}
            {/*           v-model="dbname" */}
            {/*           placeholder="ae1" */}
            {/*         ></input> */}
            {/*       </div> */}
            {/*       <div className="field-wrap"> */}
            {/*         <label> */}
            {/*           Query String<span className="req">*</span> */}
            {/*           :q */}
            {/*         </label> */}
            {/*         <textarea */}
            {/*           className="queryinput" */}
            {/*           cols={40} */}
            {/*           rows={5} */}
            {/*           placeholder="Describe Users" */}
            {/*         > */}
            {/*           TextArea.tsx show tables */}
            {/*         </textarea> */}

            {/*         <button type="button" onClick={makeQuery}></button> */}
            {/*       </div> */}
            {/*     </div> */}
            {/*   </form> */}
            {/* </div> */}
          </div>
        </div>
      </div>

      {/* <TextArea name="testo" onChange={makeQuery} /> */}
      {/* TODO: create tabular data
          headers are the keys 
          <TabularData/>

        */}
      {/* <div className="tablecontainer"> */}
      {/*   <table> */}
      {/*     <tr className="header"> */}
      {/*       {/1* <th v-for="header in columns"> *1/} */}
      {/*       <th>header a</th> */}
      {/*       <th>header b</th> */}
      {/*       <th>header c</th> */}
      {/*     </tr> */}
      {/*     <tr> */}
      {/*       {/1* <td v-for="fielddata in row"> *1/} */}
      {/*       {/1* TODO: DataPointRows list *1/} */}
      {/*       <td> */}
      {/*         datapoint 1{/1* <h2 onClick={insertIntoQuery}> fielddata </h2> *1/} */}
      {/*       </td> */}
      {/*       <td> */}
      {/*         datapoint 2{/1* <h2 onClick={insertIntoQuery}> fielddata </h2> *1/} */}
      {/*       </td> */}

      {/*       <td> */}
      {/*         datapoint 3{/1* <h2 onClick={insertIntoQuery}> fielddata </h2> *1/} */}
      {/*       </td> */}
      {/*     </tr> */}
      {/*   </table> */}
      {/* </div> */}
    </div>
  );
}

```js

  let querystring = " ";
  let querypos = 0;
  let dbname = "new";
  let columns = [];
  let rowdata = [];
  // list of dbs, of type string, when sent to the api will dictate the db to read from
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
  const replaceQueryString = (newVal: any, event: Event) => {
    let querystring = newVal;
    return querystring;
  };
