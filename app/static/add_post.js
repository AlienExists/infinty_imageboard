function submit() {
   var url = "/api/v1/posts";
   var data = new FormData();
   data.append('PostData', document.getElementById("POST_SEND").value);
   var xhr = new XMLHttpRequest();
   xhr.open("POST", url);

   xhr.setRequestHeader("Content-Type", "application/json");

   xhr.onreadystatechange = function () {
      if (xhr.readyState === 4) {
         var Myelement = document.querySelector('textarea[name="POST_SEND"]');
         Myelement.value = "";
         console.log(xhr.status);
         console.log(xhr.responseText);
         location.reload();
      }
   };

   xhr.send(data);
}
