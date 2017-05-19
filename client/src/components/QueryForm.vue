<template>
  <div class="hello">
  <div v-for="query in previousqueries">
    <button v-on:click="replacequerystring(query.query, $event)">query is: {{ query.query }} </button>
  </div>
    <h1>{{ msg }}</h1>
    <div class="form">
      <div class="tab-content">
        <div id="signup">   
          <form>
            <div class="top-row">
              <div class="field-wrap">
                <label>
                  Database <span class="req" >*</span>
                </label>
                <input type="text" v-model="dbname" placeholder="ae1"/>
              </div>
              <div class="field-wrap">
                <label>
                  Query String<span class="req">*</span>
                </label>
                <textarea type="multiline" cols="40" rows="5" v-model="querystring" size="200" placeholder="Describe Users"/>
            <button v-on:click="makequery">make query to db</button>
              </div>
            </div>
          </form>
        </div>
      </div><!-- tab-content -->
    </div> <!-- /form -->
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
</template>

<script>
const config = {
  endpoint: 'http://localhost:9099',
};
export default {
  name: 'QueryForm',
  data() {
    return {
      querystring: '',
      msg: 'Query Form at your service',
      dbname: '',
      columns: [],
      rowdata: [], 
      dbs: ['ae1','New','Test'],
      previousqueries: []
    };
  },
  methods: {
    addToQueryString: function (newVal,event) {
      this.querystring += " " + newVal
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
        return response.json();
      })
        .then((data) => {
          this.rowdata = data
          return data;
      })
        .then((data) => {
          console.log(data)
          this.columns = Object.keys(data[0])
          var parsedArray = data.map(function(row){
          return Object.values(row);
          }) 
          this.rowdata = parsedArray
        })
        .catch(function(error){
        console.log("error", error)
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
height: 30px;
  float: left;
}
form{
margin: auto;
width: 50%;
}
label{
font-size: 1.0em;
  float: left;
}

.field-checklist .wrapper {
	width: 100%;
}
table{
width: 90%;
max-height: 400px;
  overflow: auto;
}
th {
font-weight: bold;
}
tr:nth-child(even){
background:#cdcccc;
}
.hello{
  font-size:0.8em;
}
</style>
