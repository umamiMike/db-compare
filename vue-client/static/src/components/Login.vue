<template>
  <div class="queryBox">
    <h1>{{ msg }}</h1>
    <h2>The queryBox</h2>
    <div class="form">
      <ul class="tab-group">
        <li class="tab active"><a href="#signup">Sign Up</a></li>
        <li class="tab"><a href="#login">Log In</a></li>
      </ul>
      <div class="tab-content">
        <div id="signup">   
          <h1>Sign Up</h1>
          <form action="/" method="post">
            <div class="top-row">
              <div class="field-wrap">
                <label>
                  First Name<span class="req">*</span>
                </label>
                <input type="text" required autocomplete="off" />
              </div>
              <div class="field-wrap">
                <label>
                  Last Name<span class="req">*</span>
                </label>
                <input type="text"required autocomplete="off"/>
              </div>
            </div>
            <div class="field-wrap">
              <label>
                Email Address<span class="req">*</span>
              </label>
              <input type="email"required autocomplete="off"/>
            </div>

            <div class="field-wrap">
              <label>
                Set A Password<span class="req">*</span>
              </label>
              <input type="password"required autocomplete="off"/>
            </div>
          </form>
        </div>
      </div><!-- tab-content -->
    </div> <!-- /form -->
    <ul>
    </ul>
    <ul>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'queryForm',
  data() {
    return {
      username: {},
      msg: 'I would like to query a db',
    };
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
  margin: 0 10px;
}

a {
  color: #42b983;
}

fieldset {
	border: 0;
}
button {
 width: 50%;
 height: 30px;
 margin: 8px
}
input {
width: 100%;
height: 30px;
}
form{
margin: auto;
width: 50%;
}
label{
font-size: 1.5em;
}
.panel-heading {
	color: #333;
	background-color: #f5f5f5;
	border-color: #ddd;

	padding: 10px 15px;
	border-bottom: 1px solid transparent;
	border-top-left-radius: 3px;
	border-top-right-radius: 3px;        
}

.panel-body {
	padding: 15px;
}				

.field-checklist .wrapper {
	width: 100%;
}
</style>
