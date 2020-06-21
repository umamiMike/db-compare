
<template>
  <div id="app">
    <form enctype="multipart/form-data" action="http://localhost:9099/upload" method="post">
    <input type="file" name="uploadfile" />
    <input type="hidden" name="token" :token="value"/>
    <input type="submit" value="upload" />
    </form>
    <h3>Querying multiple DBS and showing response on a table</h3>

    <div class="container">
      <QueryForm></QueryForm>
    </div>
  </div>
</template>

<script>
import QueryForm from './QueryForm';

export default {
  name: 'app',
  components: {
    QueryForm,
  },
};
</script>
<style>
</style>
