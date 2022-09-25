import {api } from './api-client'
export const handleQuery = () => {
  console.log(query);

  fetch(api + '/queries', {
    method: 'POST',
    body: JSON.stringify({
      data: {
        type: 'query',
        attributes: {
          query_string: query,
        },
      },
    }),
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
    .then(data => {
      console.log(data);
      return data;
    })
    .catch(error => {
      console.log(error);
    });
};
