<template>
  <div id="app">
    <h3>Querying multiple DBS and showing response on a table</h3>
    
    <div class="container">
      <QueryForm></QueryForm>
      <QueryForm></QueryForm>
    </div>
  </div>
</template>

<script>
import QueryForm from './components/QueryForm';

export default {
  name: 'app',
  components: {
    QueryForm,
  },
};
</script>
<style>
@import url('https://fonts.googleapis.com/css?family=Nova+Mono');
#app {
  font-family: 'Nova Mono', monospace;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;

}
input {font-family: inherit;}
</style>
