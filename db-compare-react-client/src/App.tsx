import React from "react";
/* import TextArea from "./TextArea"; */
import "./App.css";

export default function App(props: any) {
  /* * const [state, setState] = useState([]); */
  /* const [query, setQuery] = useState(); */

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

  /* const makeQuery = (e: any) => { */
  /*   e.preventDefault(); */
  /*   /1* return QueryForm(payload); *1/ */
  /* }; */

  /* const QueryForm = (payload: Payload) => { */
  /*   let rowdata = []; */
  /*   let columns = []; */

  /*   fetch(config.endpoint, { */
  /*     method: "POST", */
  /*     body: JSON.stringify(payload), */
  /*   }) */
  /*     .then(function (response) { */
  /*       if (!response.ok) { */
  /*         throw Error(response.statusText); */
  /*       } */
  /*       return response; */
  /*     }) */
  /*     .then(function (response) { */
  /*       return response.json(); */
  /*     }) */
  /*     .then((data) => { */
  /*       rowdata = data; */
  /*       return data; */
  /*     }) */
  /*     .then((data) => { */
  /*       columns = Object.keys(data.Data[0]); */
  /*       var parsedArray = data.Data.map(function (row: object) { */
  /*         return Object.values(row); */
  /*       }); */
  /*       rowdata = parsedArray; */
  /*     }) */
  /*     .catch(function (error) {}); */
  /* }; */

  const uploadForm = () => {
    return (
      <form
        encType="multipart/form-data"
        action="http://localhost:9099/upload"
        method="post"
      >
        <input type="file" />
        {/* token */}
        <input type="hidden" />
        <input type="submit" value="upload" />
      </form>
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
