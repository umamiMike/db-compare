<template>
  <div class="hello">
  <div id="replacequery_cont" v-for="query in previousqueries">
    <button class="replacequery" v-on:click="replacequerystring(query.query, $event)">query is: {{ query.query }} </button>
  </div>
    <h1>{{ msg }}</h1>
    <div id="someform" class="form">
      <div class="tab-content">
        <div id="signup">   
          <form>
            <div class="top-row">
              <div class="field-wrap">
                <label>
                  Database <span class="req" >*</span>
                </label>
                <input type="text" v-model="dbname" placeholder="ae1">
              </div>
              <div class="field-wrap">
                <label>
                  Query String<span class="req">*</span>
                </label>
                <textarea class="queryinput" type="multiline" cols="40" rows="5" v-on:click="saveCursor" v-on:keyup="saveCursor" v-model="querystring" size="200" placeholder="Describe Users">show tables </textarea>
            <button v-on:click="makequery">make query to db</button>
              </div>
            </div>
          </form>
        </div>
      </div><!-- tab-content -->
    </div> <!-- /form -->
    <div class ="tablecontainer">
    <table>
      <tr class="header">
        <th v-for="header in columns">
           <h2 v-on:click="addToQueryString(header, $event)"> {{ header }} </h2>
        </th>  
      </tr>
      <tr v-for="row in rowdata">
        <td v-for="fielddata in row">
    <h2 v-on:click="addToQueryString(fielddata, $event)"> {{ fielddata }} </h2>
        </td></tr>
    </table>
    </div>
  </div>
</template>

<script>
  const config = {
    endpoint: 'http://localhost:9099',
  };
export default {
  name: 'QueryForm',
  data() {
    return {
      querystring: 'show tables',
      querypos: 0,
      msg: 'Query Form at your service',
      dbname: 'new',
      columns: [],
      rowdata: [], 
      dbs: ['ae1','New','Test'],
      previousqueries: []
    };
  },

  methods: {
    saveCursor: function (event) {
      this.querypos =  event.target.selectionStart
      let se = event.target.selectionEnd

    },
    addToQueryString: function (newVal,event) {
      let astr = this.querystring.substr(0,this.querypos)
      let cstr = this.querystring.substr(this.querypos)
      this.querystring = astr + newVal + cstr
    },
    replacequerystring: function (newVal,event) {
      this.querystring =  newVal
    },
    makequery: function (event) {
      this.rowdata = [];
      this.previousqueries.push({db:this.dbname,query:this.querystring})
      var payload = {
        query: this.querystring,
        db: this.dbname,
      };
      fetch(config.endpoint,
        {
          method: 'POST',
          body: JSON.stringify(payload)
        })
        .then(function (response) {
          if (!response.ok){ 
            throw Error(response.statusText);
          }
          return response
        })
        .then(function (response) {
          console.log(response.ok);
          return response.json();
        })
        .then((data) => {
          this.rowdata = data
          return data;
        })
        .then((data) => {
          this.columns = Object.keys(data.Dat[0])
          var parsedArray = data.Dat.map(function(row){
            return Object.values(row);
          }) 
          this.rowdata = parsedArray
        })
        .catch(function(error){
        });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
  text-align:center;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
}

a {
  color: #42b983;
}

button {
}
input {
width: 25%;
  float: left;
}
form{
margin: auto;
width: 50%;
}
label{
  float: left;
}
#replacequery_cont{
        float:none;

}
.replacequery{
float: left;
padding: 4px;
margin: 4px;
}

.queryinput {
       background: rgb(240,240,240); 
        font-family:inherit;
        margin: 16px;
        padding: 16px;
}
.header {
  padding: 16px;
  margin: 0px;
  border: 2px;
  border-color: black;
}
.field-checklist .wrapper {
}
.tablecontainer {
height: 400px;
  overflow: scroll;
}
table{
}
th {
}
tr:nth-child(even){
background:#cdcccc;
}
</style>
