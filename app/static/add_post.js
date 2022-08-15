function submit() {
   var data_send = document.getElementById("POST_SEND").value;

   // Пример отправки POST запроса:
   async function postData(url = '', data = {}) {
      // Default options are marked with *
      const response = await fetch(url, {
         method: 'POST', // *GET, POST, PUT, DELETE, etc.
         mode: 'cors', // no-cors, *cors, same-origin
         cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
         credentials: 'same-origin', // include, *same-origin, omit
         headers: {
            'Content-Type': 'application/json'
            // 'Content-Type': 'application/x-www-form-urlencoded',
         },
         redirect: 'follow', // manual, *follow, error
         referrerPolicy: 'no-referrer', // no-referrer, *client
         body: JSON.stringify(data) // body data type must match "Content-Type" header
      });
      return await response.json(); // parses JSON response into native JavaScript objects
   }

   postData('/api/v1/posts?post=' + data_send)
      .then((data) => {
         console.log(data); // JSON data parsed by `response.json()` call
      });

   x.send(null);
   var Myelement = document.querySelector('textarea[name="POST_SEND"]');
   Myelement.value = "";
   location.reload();
}
